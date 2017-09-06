package main

import "fmt"

func red(str interface{}) string {
	return fmt.Sprintf("\033[91m%v\033[00m", str)
}

func green(str interface{}) string {
	return fmt.Sprintf("\033[92m%v\033[00m", str)
}

func yellow(str interface{}) string {
	return fmt.Sprintf("\033[93m%v\033[00m", str)
}

func lightPurple(str interface{}) string {
	return fmt.Sprintf("\033[94m%v\033[00m", str)
}

func purple(str interface{}) string {
	return fmt.Sprintf("\033[95m%v\033[00m", str)
}

func cyan(str interface{}) string {
	return fmt.Sprintf("\033[96m%v\033[00m", str)
}

func lightGray(str interface{}) string {
	return fmt.Sprintf("\033[97m%v\033[00m", str)
}

func boldCyan(str interface{}) string {
	return fmt.Sprintf("\033[96m\033[1m%v\033[00m", str)
}

func boldGreen(str interface{}) string {
	return fmt.Sprintf("\033[92m\033[1m%v\033[00m", str)
}

func boldYellow(str interface{}) string {
	return fmt.Sprintf("\033[93m\033[1m%v\033[00m", str)
}

func black(str interface{}) string {
	return fmt.Sprintf("\033[98m%v\033[00m", str)
}
