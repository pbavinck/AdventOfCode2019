package day5

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pbavinck/AofCod2019/loader"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCod2019/day5/input.txt"

func padWithZeros(s string) string {
	// Adds zeros to opcode
	a, _ := strconv.Atoi(s)
	return fmt.Sprintf("%05d", a)
}

func paramMode(opcode string, index int) string {
	// returns the mode of param index of the opcode
	opcode = padWithZeros(opcode)
	return string(opcode[2-index])
}

func valueToUse(code []string, opcodeIndex int, paramIndex int) int {
	opcode := code[opcodeIndex]
	if paramMode(opcode, paramIndex) == "0" {
		//postion mode
		index, _ := strconv.Atoi(code[opcodeIndex+paramIndex+1])
		r, _ := strconv.Atoi(code[index])
		log.Printf("        Param %v == %-6v (position ). code[%v] has value: %v\n", paramIndex, index, index, r)
		return r
	}

	// Immediate
	r, _ := strconv.Atoi(code[opcodeIndex+paramIndex+1])
	log.Printf("        Param %v == %-6v (immediate).\n", paramIndex, r)
	return r
}

func doInstr1(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (ADD)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	targetIndex, _ := strconv.Atoi(code[opcodeIndex+3])
	code[targetIndex] = strconv.Itoa(a + b)
	log.Printf("        Param 2 == %-6v (output   ). code[%v] set to value: %v (%v + %v)\n", targetIndex, targetIndex, strconv.Itoa(a+b), a, b)
	return opcodeIndex + 4
}

func doInstr2(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (MULTIPLY)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	targetIndex, _ := strconv.Atoi(code[opcodeIndex+3])
	code[targetIndex] = strconv.Itoa(a * b)
	log.Printf("        Param 2 == %-6v (output). code[%v] set to value: %v (%v * %v)\n", targetIndex, targetIndex, strconv.Itoa(a*b), a, b)
	return opcodeIndex + 4
}

func doInstr3(code []string, opcodeIndex int, input string) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (INPUT)\n", opcodeIndex, code[opcodeIndex])
	targetIndex, _ := strconv.Atoi(code[opcodeIndex+1])
	code[targetIndex] = input
	log.Printf("        Input == %v\n", input)
	log.Printf("        Param 0 == %-6v (output). code[%v] set to value: %v\n", targetIndex, targetIndex, input)

	return opcodeIndex + 2
}

func doInstr4(code []string, opcodeIndex int) (nextOpcodeIndex int, output int) {
	log.Printf("%+6v: Opcode: %v (OUTPUT)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	log.Printf(" -> Output: %v\n", a)
	return opcodeIndex + 2, a
}

func doInstr5(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (JUMP_IF_TRUE)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	if a != 0 {
		nextOpcodeIndex = b
		log.Printf("        JUMP to %-6v\n", nextOpcodeIndex)
	} else {
		nextOpcodeIndex = opcodeIndex + 3
		log.Printf("        NO JUMP\n")

	}
	return nextOpcodeIndex
}

func doInstr6(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (JUMP_IF_FALSE)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	if a == 0 {
		nextOpcodeIndex = b
		log.Printf("        JUMP to %+6v\n", nextOpcodeIndex)
	} else {
		nextOpcodeIndex = opcodeIndex + 3
		log.Printf("        NO JUMP\n")
	}
	return nextOpcodeIndex
}

func doInstr7(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (LESS_THAN)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	targetIndex, _ := strconv.Atoi(code[opcodeIndex+3])
	if a < b {
		code[targetIndex] = string("1")
	} else {
		code[targetIndex] = string("0")
	}
	log.Printf("        Param 3 == %-6v (output). code[%v] set to %v\n", targetIndex, targetIndex, code[targetIndex])
	return opcodeIndex + 4
}

func doInstr8(code []string, opcodeIndex int) (nextOpcodeIndex int) {
	log.Printf("%+6v: Opcode: %v (EQUAL_TO)\n", opcodeIndex, code[opcodeIndex])
	a := valueToUse(code, opcodeIndex, 0)
	b := valueToUse(code, opcodeIndex, 1)
	targetIndex, _ := strconv.Atoi(code[opcodeIndex+3])
	if a == b {
		code[targetIndex] = string("1")
	} else {
		code[targetIndex] = string("0")
	}
	log.Printf("        Param 3 == %-6v (output). code[%v] set to %v\n", targetIndex, targetIndex, code[targetIndex])
	return opcodeIndex + 4
}

func run(code []string, input string) int {
	output := -1
	opcodeIndex := 0
	for {
		code[opcodeIndex] = padWithZeros(code[opcodeIndex])
		opcode := code[opcodeIndex]
		operation := opcode[len(opcode)-2:]
		switch {
		case operation == "01":
			opcodeIndex = doInstr1(code, opcodeIndex)
		case operation == "02":
			opcodeIndex = doInstr2(code, opcodeIndex)
		case operation == "03":
			opcodeIndex = doInstr3(code, opcodeIndex, input)
		case operation == "04":
			opcodeIndex, output = doInstr4(code, opcodeIndex)
		case operation == "05":
			opcodeIndex = doInstr5(code, opcodeIndex)
		case operation == "06":
			opcodeIndex = doInstr6(code, opcodeIndex)
		case operation == "07":
			opcodeIndex = doInstr7(code, opcodeIndex)
		case operation == "08":
			opcodeIndex = doInstr8(code, opcodeIndex)
		case operation == "99":
			log.Printf("END OF PROGRAM\n\n")
			return output
		}
	}
}

//SolvePart1 solves part 1 of day 5
func SolvePart1(code []string) {
	fmt.Println("Part 1 - Program output:", run(code, string("1")))
}

//SolvePart2 solves part 2 of day 5
func SolvePart2(code []string) {
	fmt.Println("Part 2 - Program output:", run(code, string("5")))

}

// Solve runs day 5 assignment
func Solve() {
	fmt.Printf("\n*** DAY 5 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	code := strings.Split(data[0], ",")
	SolvePart1(code)

	code = strings.Split(data[0], ",")
	SolvePart2(code)
}
