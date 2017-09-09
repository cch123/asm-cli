package main

import (
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

var beforeExecRegVals = map[string]uint64{}

type basicMachine struct {
	sp             int
	sortedRegNames []string
	regMap         map[string]int
	mu             uc.Unicorn
	byteSize       int
}

func (m basicMachine) displayRegisters() {
	startLine := "----------------- cpu context -----------------"
	fmt.Println(cyan(startLine))
	for _, regName := range m.sortedRegNames {
		if regName == "end" {
			fmt.Println()
			continue
		}

		reg := m.regMap[regName]
		res, _ := m.mu.RegRead(reg)

		resStr := fmt.Sprintf("%0#[1]*[2]x", m.byteSize*2, res)
		beforeVal, ok := beforeExecRegVals[regName]
		// pad the reg name to 3 bytes
		paddedRegName := fillSpace(regName, 3)
		if ok && beforeVal != res {
			fmt.Printf("%v : %v ", purple(paddedRegName), red(resStr))
		} else {
			fmt.Printf("%v : %v ", purple(paddedRegName), resStr)
		}
		beforeExecRegVals[regName] = res
	}

}

func (m basicMachine) displayStack() {
	startLine := "----------------- stack context -----------------"
	fmt.Println(yellow(startLine))

	readStartAddr := uint64(0x1300000 - m.byteSize*8)
	readOffset := uint64(m.byteSize * 4 * 5)
	bytesData, err := m.mu.MemRead(readStartAddr, readOffset)
	if err != nil {
		// TODO do some thing
	}
	spVal, _ := m.mu.RegRead(m.sp)

	for i := 0; i < len(bytesData); i += m.byteSize {
		if i%(m.byteSize*4) == 0 && i != 0 {
			fmt.Println()
		}

		// the first line didn't change line
		if i%(m.byteSize*4) == 0 {
			fmt.Printf("%0#[1]*[2]x : ", m.byteSize*2, readStartAddr+uint64(i))
		}

		var reversedBytes = bytesData[i : i+m.byteSize]
		for i := 0; i <= m.byteSize/2-1; i++ {
			j := m.byteSize - 1 - i
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

func (m basicMachine) execute(cmd string) error {
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
	m.mu.MemWrite(0x0000, code)
	if err := m.mu.Start(0x0000, 0x0000+uint64(len(code))); err != nil {
		fmt.Println(red(fmt.Sprintf("err : %v", err)))
	}
	return err
}
