package main

/*
  rasm2 seems doesn't support 8086 asm
  so this feature will be delayed
*/

import (
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machine8086 struct {
	basicMachine
}

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

func init8086() machine {
	m := machine8086{}
	machineMap[key8086] = m
	mu8086, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_16)
	mu8086.RegWrite(uc.X86_REG_SP, 0x1000)
	mu8086.RegWrite(uc.X86_REG_BP, 0xffff)
	mu8086.MemMap(0x0000, 0xffff)
	return m
}
