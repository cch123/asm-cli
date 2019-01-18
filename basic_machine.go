package main

import (
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/keystone-engine/keystone/bindings/go/keystone"
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"
)

var beforeExecRegVals = map[string]uint64{}

type basicMachine struct {
	sp             int
	sortedRegNames []string
	regMap         map[string]int
	mu             uc.Unicorn
	byteSize       int
}

func readFlagVals(flags uint64) map[string]int {
	res := make(map[string]int)
	// cf:0 zf:0 of:0 sf:0 pf:0 af:0 df:0
	flagNames := []string{"cf", "zf", "of", "sf", "pf", "af", "df"}
	var nameToBitMap = map[string]uint{
		"cf": 0,
		"pf": 2,
		"af": 4,
		"zf": 6,
		"sf": 7,
		"df": 10,
		"of": 11,
	}
	for _, flagName := range flagNames {
		bitPos := nameToBitMap[flagName]

		res[flagName] = 0
		if flags>>bitPos&1 > 0 {
			res[flagName] = 1
		}
	}
	return res
}

func (m basicMachine) displayRegisters() {
	startLine := "----------------- cpu context -----------------"
	fmt.Println(cyan(startLine))
	for _, regName := range m.sortedRegNames {
		if regName == "end" {
			fmt.Println()
			continue
		}

		reg := m.regMap[regName]
		res, _ := m.mu.RegRead(reg)

		resStr := fmt.Sprintf("%0#[1]*[2]x", m.byteSize*2, res)
		beforeVal, ok := beforeExecRegVals[regName]
		// pad the reg name to 3 bytes
		paddedRegName := fillSpace(regName, 3)
		if ok && beforeVal != res {
			fmt.Printf("%v : %v ", purple(paddedRegName), red(resStr))
		} else {
			fmt.Printf("%v : %v ", purple(paddedRegName), resStr)
		}
		beforeExecRegVals[regName] = res

		// flag register detail
		if regName == "flags" {
			flagValMap := readFlagVals(res)
			fmt.Printf("(cf:%v zf:%v of:%v sf:%v pf:%v af:%v df:%v)",
				flagValMap["cf"], flagValMap["zf"], flagValMap["of"], flagValMap["sf"], flagValMap["pf"], flagValMap["af"], flagValMap["df"])
		}
	}

}

func (m basicMachine) displayStack() {
	startLine := "----------------- stack context -----------------"
	fmt.Println(yellow(startLine))

	readStartAddr := uint64(0x1300000 - m.byteSize*8)
	readOffset := uint64(m.byteSize * 4 * 5)
	bytesData, err := m.mu.MemRead(readStartAddr, readOffset)
	if err != nil {
		// TODO do some thing
	}
	spVal, _ := m.mu.RegRead(m.sp)

	for i := 0; i < len(bytesData); i += m.byteSize {
		if i%(m.byteSize*4) == 0 && i != 0 {
			fmt.Println()
		}

		// the first line didn't change line
		if i%(m.byteSize*4) == 0 {
			fmt.Printf("%0#[1]*[2]x : ", m.byteSize*2, readStartAddr+uint64(i))
		}

		var reversedBytes = bytesData[i : i+m.byteSize]
		for i := 0; i <= m.byteSize/2-1; i++ {
			j := m.byteSize - 1 - i
			reversedBytes[i], reversedBytes[j] = reversedBytes[j], reversedBytes[i]
		}

		currentAddr := readStartAddr + uint64(i)
		if spVal == currentAddr {
			fmt.Printf("%s ", red(hex.EncodeToString(reversedBytes[0:])))
		} else {
			fmt.Printf("%s ", hex.EncodeToString(reversedBytes[0:]))
		}
	}
	fmt.Println()

}

/*
func assemble(mnemonic string) ([]byte, error) {
	// TODO: What is the effect of the second argument(address) of Assemble
	code, cnt, ok := keystone.Assemble(mnemonic, 0)
	if !ok || cnt == 0 {
		return nil, fmt.Errorf("Error: assemble instruction(%s)", mnemonic)
	}
	return code, nil
}
*/

func (m basicMachine) execute(cmd string) error {
	/*
		var args = []string{
			"-a", "x86", cmd,
		}
		res, err := exec.Command("rasm2", args...).Output()
		if err != nil {
			fmt.Println(err)
		}
	*/

	ks, _ := keystone.New(keystone.ARCH_X86, keystone.MODE_64)
	ks.Option(keystone.OPT_SYNTAX, keystone.OPT_SYNTAX_INTEL)
	res, cnt, ok := ks.Assemble(cmd, 0)

	if !ok || cnt == 0 {
		return errors.New("assemble failed")
	}

	//resStr := strings.Trim(string(res), "\n")
	fmt.Printf("%v: %v\t%v: %v\n", purple("opcode"), cmd, purple("hex"), res)
	helperInfo()
	//code, _ := hex.DecodeString(resStr)

	// NOTICE
	// push/pop rax commands must ensure that
	// the rsp point into the range of memmap
	m.mu.MemWrite(0x0000, res)
	if err := m.mu.Start(0x0000, 0x0000+uint64(len(res))); err != nil {
		fmt.Println(red(fmt.Sprintf("err : %v", err)))
	}
	return nil
}
