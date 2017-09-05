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

func lightGrey(str interface{}) string {
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

/*

# change output color
def lightPurple(s,e="\n") : print("\033[94m%v\033[00m".format(s), end=e)
def cyan(s,e="\n"): print("\033[96m%v\033[00m".format(s), end=e)
def lightGray(s,e="\n"): print("\033[97m%v\033[00m".format(s), end=e)
def black(s,e="\n"): print("\033[98m%v\033[00m".format(s), end=e)
def white(s,e="\n"): print("\033[00m", end=e)
*/
