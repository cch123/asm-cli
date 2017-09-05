package main

import (
	"fmt"
	"io"

	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

type machine8086 struct{}

var sorted8086RegNames = []string{
	"ax", "bx", "cx", "dx",
	"ip", "bp", "sp",
	"flags",
	"cs", "ss", "ds", "es", "fs", "gs",
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

func init8086() {
	m := machine8086{}
	machineMap[key8086] = m
	mu8086, _ = uc.NewUnicorn(uc.ARCH_X86, uc.MODE_16)
}

func (m machine8086) setOutput(out io.Writer) {
}

func (m machine8086) displayRegisters() {
	for _, regName := range sortedX64RegNames {
		reg := regMap8086[regName]
		res, _ := mu8086.RegRead(reg)
		fmt.Printf("read from %v : %v\n", regName, res)
	}
}

func (m machine8086) execute() {
}
