package main

import (
	"io"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX64 struct{}

var regMap = map[string]int{
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

func initX64() {
	m := machineX64{}
	machineMap[keyX64] = m
}

func (m machineX64) setOutput(out io.Writer) {
}

func (m machineX64) displayRegisters() {
}

func (m machineX64) execute() {
}
