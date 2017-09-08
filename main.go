package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/c-bata/go-prompt"
)

const (
	keyX86  = "x86"
	keyX64  = "x64"
	key8086 = "8086"
)

var machineMap = map[string]machine{}

type machine interface {
	displayRegisters()
	displayStack()
	setOutput(io.Writer)
	execute(string) error
}

func init() {
	initX86()
	initX64()
	init8086()
}

func main() {
	//machineName := keyX64
	//machineName := keyX86
	var arch = flag.String("a", "x86", "x86/x64/8086")
	flag.Parse()
	var machineName string
	switch *arch {
	case keyX86:
		fallthrough
	case keyX64:
		fallthrough
	case key8086:
		machineName = *arch

	default:
		fmt.Println("invalid arch type")
	}
	ma, ok := machineMap[machineName]
	if !ok {
		fmt.Println("wrong key")
		os.Exit(1)
	}

	ma.displayRegisters()
	ma.displayStack()

	for {
		fmt.Println("Input q to quit.")
		// FIXME when input ctrl+c/ctrl+z/ctrl+d
		// prompt will become very slow
		t := prompt.Input(machineName+"> ", completer)
		if t == "q" || t == "quit" || t == "exit" {
			break
		}
		ma.execute(t)
		ma.displayRegisters()
		ma.displayStack()
	}

}
