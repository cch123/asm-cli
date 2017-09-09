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

var ma machine

func main() {
	//machineName := keyX64
	//machineName := keyX86
	var arch = flag.String("a", "x64", "x86/x64/8086")
	flag.Parse()
	ma = getMachine(*arch)
	if ma == nil {
		fmt.Println("unsupported arch, please use x86 or x64")
		os.Exit(0)
	}

	ma.displayRegisters()
	ma.displayStack()

	p := prompt.New(myExecutor, completer, prompt.OptionPrefix(*arch+">> "))
	p.Run()

}

func getMachine(arch string) machine {
	var ma machine
	switch arch {
	case keyX86:
		ma = initX86()
	case keyX64:
		ma = initX64()
	case key8086:
		fallthrough
	default:
		fmt.Println("invalid arch type")
		return nil
	}
	return ma
}

func myExecutor(cmd string) {
	if cmd == "quit" || cmd == "exit" || cmd == "q" {
		os.Exit(0)
	}
	ma.execute(cmd)
	ma.displayRegisters()
	ma.displayStack()
}
