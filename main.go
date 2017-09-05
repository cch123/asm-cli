package main

import (
	"fmt"
	"io"
	"os"
)

const (
	keyX86  = "x86"
	keyX64  = "x64"
	key8086 = "8086"
)

var machineMap = map[string]machine{}

type machine interface {
	displayRegisters()
	setOutput(io.Writer)
	execute()
}

func init() {
	initX86()
	initX64()
	init8086()
}

func main() {
	machineName := keyX86
	ma, ok := machineMap[machineName]
	if !ok {
		fmt.Println("wrong key")
		os.Exit(1)
	}
	ma.displayRegisters()

	fmt.Println(ma)

}
