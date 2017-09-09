package main

import (
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX64 struct {
	basicMachine
}

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

func initX64() machine {
	m := machineX64{}
	m.regMap = regMapX64
	m.byteSize = 8
	m.sp = uc.X86_REG_RSP
	m.sortedRegNames = sortedX64RegNames
	m.mu, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_64)
	m.mu.RegWrite(uc.X86_REG_RSP, 0x01300000)
	m.mu.RegWrite(uc.X86_REG_RBP, 0x10000000)
	m.mu.MemMap(0x0000, 0x20000000)
	return m
}
