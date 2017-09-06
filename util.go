package main

import (
	"fmt"
	"strconv"
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

// merge two byte to one
func mergeBytes(input []byte) []byte {
	var res = []byte{}
	// the last is change line ?
	for i := 0; i < len(input)-1; i += 2 {
		num, _ := strconv.ParseInt(string(input[i:i+2]), 16, 32)
		res = append(res, byte(num))
	}
	fmt.Println(res)
	return res
}
