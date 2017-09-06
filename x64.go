package main

import (
	"fmt"
	"io"
	"os/exec"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX64 struct{}

var sortedX64RegNames = []string{
	"rax", "rbx", "rcx", "rdx", "end",
	"rsi", "rdi", "r8", "r9", "end",
	"r10", "r11", "r12", "r13", "end",
	"r14", "r15", "end",
	"rip", "rbp", "rsp", "end",
	"cs", "ss", "ds", "es", "end",
	"fs", "gs", "end",
	"flags", "end",
}

var regMapX64 = map[string]int{
	// 通用寄存器
	"rax": uc.X86_REG_RAX,
	"rbx": uc.X86_REG_RBX,
	"rcx": uc.X86_REG_RCX,
	"rdx": uc.X86_REG_RDX,
	"rsi": uc.X86_REG_RSI,
	"rdi": uc.X86_REG_RDI,
	"r8":  uc.X86_REG_R8,
	"r9":  uc.X86_REG_R9,
	"r10": uc.X86_REG_R10,
	"r11": uc.X86_REG_R11,
	"r12": uc.X86_REG_R12,
	"r13": uc.X86_REG_R13,
	"r14": uc.X86_REG_R14,
	"r15": uc.X86_REG_R15,

	"rip":   uc.X86_REG_RIP,
	"rbp":   uc.X86_REG_RBP,
	"rsp":   uc.X86_REG_RSP,
	"flags": uc.X86_REG_EFLAGS,
	"cs":    uc.X86_REG_CS,
	"ss":    uc.X86_REG_SS,
	"ds":    uc.X86_REG_DS,
	"es":    uc.X86_REG_ES,
	"fs":    uc.X86_REG_FS,
	"gs":    uc.X86_REG_GS,
}
var muX64 uc.Unicorn

func initX64() {
	m := machineX64{}
	machineMap[keyX64] = m
	muX64, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
}

func (m machineX64) setOutput(out io.Writer) {
}

func (m machineX64) displayRegisters() {
	startLine := "----------------- cpu context -----------------"
	fmt.Println(cyan(startLine))
	for _, regName := range sortedX64RegNames {
		if regName == "end" {
			fmt.Println()
			continue
		}

		reg := regMapX64[regName]
		res, _ := muX64.RegRead(reg)
		//res = 1002300000
		resStr := fmt.Sprintf("%0#[1]*[2]x", 16, res)
		regName = fillSpace(regName, 3)
		fmt.Printf("%v : %v ", purple(regName), resStr)
	}

}

func (m machineX64) displayStack() {
	startLine := "----------------- stack context -----------------"
	fmt.Println(yellow(startLine))
}

func (m machineX64) execute(cmd string) error {
	var args = []string{
		"-a", "x86", cmd,
	}
	res, err := exec.Command("rasm2", args...).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))

	// FIXME
	code := mergeBytes(res)
	//code := []byte{184, 210, 4, 0, 0}
	muX64.MemMap(0x1000, 0x1000)
	muX64.MemWrite(0x1000, code)
	if err := muX64.Start(0x1000, 0x1000+uint64(len(code))); err != nil {
		panic(err)
	}
	return err
}
