package main

import (
	"fmt"
	"io"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX86 struct{}

var sortedX86RegNames = []string{
	"eax", "ebx", "ecx", "edx",
	"esi", "edi",
	"eip", "ebp", "esp",
	"eflags",
	"cs", "ss", "ds", "es", "fs", "gs",
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
}

func (m machineX86) setOutput(out io.Writer) {
}

func (m machineX86) displayRegisters() {
	for _, regName := range sortedX86RegNames {
		reg := regMapX86[regName]
		res, _ := muX86.RegRead(reg)
		fmt.Printf("read from %v : %v\n", regName, res)
	}
}

func (m machineX86) execute() {
}
