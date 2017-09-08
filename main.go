package main

import (
	"flag"
	"fmt"
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
	execute(string) error
}

func init() {
	initX86()
	initX64()
	init8086()
}

var ma machine

func main() {
	//machineName := keyX64
	//machineName := keyX86
	var arch = flag.String("a", "x64", "x86/x64/8086")
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
	var ok bool
	ma, ok = machineMap[machineName]
	if !ok {
		fmt.Println("wrong key")
		os.Exit(1)
	}

	ma.displayRegisters()
	ma.displayStack()

	p := prompt.New(myExecutor, completer, prompt.OptionPrefix(machineName+">> "))
	p.Run()

}

func myExecutor(cmd string) {
	if cmd == "quit" || cmd == "exit" || cmd == "q" {
		os.Exit(0)
	}
	ma.execute(cmd)
	ma.displayRegisters()
	ma.displayStack()
}
