package main

func fillSpace(str string, needLen int) string {
	res := str
	strLen := len(str)
	for strLen < needLen {
		res = string(append([]byte(res), ' '))
		strLen++
	}
	return res
}
