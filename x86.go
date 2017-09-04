package main

import uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"

var machineX86 machine

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

func registerX86() {
	machineMap[keyX86] = machineX86
}
