package main

import uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"

var machine8086 machine

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

func register8086() {
	machineMap[key8086] = machine8086
}
