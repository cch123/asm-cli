package main

import (
	"fmt"

	prompt "github.com/c-bata/go-prompt"
)

func fillSpace(str string, needLen int) string {
	res := str
	strLen := len(str)
	for strLen < needLen {
		res = string(append([]byte(res), ' '))
		strLen++
	}
	return res
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "MOV", Description: "MOV dest, src"},
		{Text: "PUSH", Description: "PUSH src"},
		{Text: "POP", Description: "POP dest"},
		{Text: "ADD", Description: "ADD dest,src"},
		{Text: "SUB", Description: "SUB dest,src"},
		{Text: "INC", Description: "INC dest"},
		{Text: "DEC", Description: "DEC dest"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func helperInfo() {
	fmt.Println("detailed info can be referred at http://ref.x86asm.net/")
}
