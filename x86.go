package main

import (
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machineX86 struct {
	basicMachine
}

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

func initX86() machine {
	m := machineX86{}
	m.regMap = regMapX86
	m.byteSize = 4
	m.sp = uc.X86_REG_ESP
	m.sortedRegNames = sortedX86RegNames
	m.mu, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_32)
	m.mu.RegWrite(uc.X86_REG_ESP, 0x01300000)
	m.mu.RegWrite(uc.X86_REG_EBP, 0x10000000)
	m.mu.MemMap(0x0000, 0x20000000)
	return m
}
