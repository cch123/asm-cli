package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os/exec"
	"strings"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX86 struct{}

var sortedX86RegNames = []string{
	"eax", "ebx", "ecx", "edx", "end",
	"esi", "edi", "end",
	"eip", "ebp", "esp", "end",
	"eflags", "end",
	"cs", "ss", "ds", "es", "end",
	"fs", "gs", "end",
}

var regMapX86 = map[string]int{
	// 通用寄存器
	"eax": uc.X86_REG_EAX,
	"ebx": uc.X86_REG_EBX,
	"ecx": uc.X86_REG_ECX,
	"edx": uc.X86_REG_EDX,
	"esi": uc.X86_REG_ESI,
	"edi": uc.X86_REG_EDI,

	"eip":    uc.X86_REG_EIP,
	"ebp":    uc.X86_REG_EBP,
	"esp":    uc.X86_REG_ESP,
	"eflags": uc.X86_REG_EFLAGS,
	"cs":     uc.X86_REG_CS,
	"ss":     uc.X86_REG_SS,
	"ds":     uc.X86_REG_DS,
	"es":     uc.X86_REG_ES,
	"fs":     uc.X86_REG_FS,
	"gs":     uc.X86_REG_GS,
}
var muX86 uc.Unicorn

func initX86() {
	machine := machineX86{}
	machineMap[keyX86] = machine
	muX86, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_32)
	muX86.RegWrite(uc.X86_REG_ESP, 0x01300000)
	muX86.RegWrite(uc.X86_REG_EBP, 0x10000000)
	muX86.MemMap(0x0000, 0x20000000)
}

func (m machineX86) setOutput(out io.Writer) {
}

func (m machineX86) displayRegisters() {
	startLine := "----------------- cpu context -----------------"
	fmt.Println(cyan(startLine))
	for _, regName := range sortedX86RegNames {
		if regName == "end" {
			fmt.Println()
			continue
		}

		reg := regMapX86[regName]
		res, _ := muX86.RegRead(reg)

		resStr := fmt.Sprintf("%0#[1]*[2]x", 8, res)
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

func (m machineX86) execute(cmd string) error {
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
	muX86.MemWrite(0x0000, code)
	if err := muX86.Start(0x0000, 0x0000+uint64(len(code))); err != nil {
		fmt.Println(red(fmt.Sprintf("err : %v", err)))
	}
	return err
}

func (m machineX86) displayStack() {
	startLine := "----------------- stack context -----------------"
	fmt.Println(yellow(startLine))
	// 64 bit = 8 bytes
	// need 4 column, 5 line
	readStartAddr := uint64(0x012fffe0)
	byteSize := 4
	readOffset := uint64(byteSize * 4 * 5)
	bytesData, err := muX86.MemRead(readStartAddr, readOffset)
	if err != nil {
	}
	espVal, _ := muX86.RegRead(uc.X86_REG_ESP)

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
		if espVal == currentAddr {
			fmt.Printf("%s ", red(hex.EncodeToString(reversedBytes[0:])))
		} else {
			fmt.Printf("%s ", hex.EncodeToString(reversedBytes[0:]))
		}
	}
	fmt.Println()
}
