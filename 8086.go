package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os/exec"
	"strings"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machine8086 struct{}

var sorted8086RegNames = []string{
	"ax", "bx", "cx", "dx", "end",
	"ip", "bp", "sp", "end",
	"cs", "ss", "ds", "es", "fs", "gs", "end",
	"flags", "end",
}

var regMap8086 = map[string]int{
	"ax": uc.X86_REG_AX,
	"bx": uc.X86_REG_AX,
	"cx": uc.X86_REG_AX,
	"dx": uc.X86_REG_AX,

	"ip":    uc.X86_REG_IP,
	"bp":    uc.X86_REG_BP,
	"sp":    uc.X86_REG_SP,
	"flags": uc.X86_REG_EFLAGS,
	"cs":    uc.X86_REG_CS,
	"ss":    uc.X86_REG_SS,
	"ds":    uc.X86_REG_DS,
	"es":    uc.X86_REG_ES,
	"fs":    uc.X86_REG_FS,
	"gs":    uc.X86_REG_GS,
}

var mu8086 uc.Unicorn

func init8086() {
	m := machine8086{}
	machineMap[key8086] = m
	mu8086, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_16)
	mu8086.RegWrite(uc.X86_REG_SP, 0x01300000)
	mu8086.RegWrite(uc.X86_REG_BP, 0x10000000)
	mu8086.MemMap(0x0000, 0x20000000)
}

func (m machine8086) setOutput(out io.Writer) {
}

func (m machine8086) displayRegisters() {
	startLine := "----------------- cpu context -----------------"
	fmt.Println(cyan(startLine))
	for _, regName := range sorted8086RegNames {
		if regName == "end" {
			fmt.Println()
			continue
		}

		reg := regMap8086[regName]
		res, _ := mu8086.RegRead(reg)

		resStr := fmt.Sprintf("%0#[1]*[2]x", 4, res)
		beforeVal, ok := beforeExecRegVals[regName]
		paddedRegName := fillSpace(regName, 3)
		if ok && beforeVal != res {
			fmt.Printf("%v : %v ", purple(paddedRegName), red(resStr))
		} else {
			fmt.Printf("%v : %v ", purple(paddedRegName), resStr)
		}
		beforeExecRegVals[regName] = res
	}
}

func (m machine8086) execute(cmd string) error {
	var args = []string{
		"-a", "x86", cmd,
	}
	res, err := exec.Command("rasm2", args...).Output()
	if err != nil {
		fmt.Println(err)
	}

	resStr := strings.Trim(string(res), "\n")
	fmt.Printf("%v: %v\t%v: %v\n", purple("opcode"), cmd, purple("hex"), resStr)
	helperInfo()
	code, _ := hex.DecodeString(resStr)

	// NOTICE
	// push/pop rax commands must ensure that
	// the rsp point into the range of memmap
	mu8086.MemWrite(0x0000, code)
	if err := mu8086.Start(0x0000, 0x0000+uint64(len(code))); err != nil {
		fmt.Println(red(fmt.Sprintf("err : %v", err)))
	}
	return err
}

func (m machine8086) displayStack() {

	startLine := "----------------- stack context -----------------"
	fmt.Println(yellow(startLine))
	// need 4 column, 5 line
	readStartAddr := uint64(0x012ffff0)
	byteSize := 2
	readOffset := uint64(byteSize * 4 * 5)
	bytesData, err := mu8086.MemRead(readStartAddr, readOffset)
	if err != nil {
	}
	spVal, _ := mu8086.RegRead(uc.X86_REG_SP)

	for i := 0; i < len(bytesData); i += byteSize {
		if i%(4*byteSize) == 0 && i != 0 {
			fmt.Println()
		}

		// the first line didn't change line
		if i%(4*byteSize) == 0 {
			fmt.Printf("%0#[1]*[2]x : ", 2*byteSize, readStartAddr+uint64(i))
		}

		var reversedBytes = bytesData[i : i+byteSize]
		for i := 0; i <= byteSize/2-1; i++ {
			j := byteSize - 1 - i
			reversedBytes[i], reversedBytes[j] = reversedBytes[j], reversedBytes[i]
		}

		currentAddr := readStartAddr + uint64(i)
		if spVal == currentAddr {
			fmt.Printf("%s ", red(hex.EncodeToString(reversedBytes[0:])))
		} else {
			fmt.Printf("%s ", hex.EncodeToString(reversedBytes[0:]))
		}
	}
	fmt.Println()
}
