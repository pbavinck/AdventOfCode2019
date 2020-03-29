package machines

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

// Computer The computer to run the code
type Computer struct {
	name    string
	program []string
	Input   chan string
	Output  chan string
}

func (c *Computer) getParamValue(line int, paramIndex int) int {
	opcode := c.program[line]
	if paramMode(opcode, paramIndex) == "0" {
		//postion mode
		index, _ := strconv.Atoi(c.program[line+paramIndex+1])
		r, _ := strconv.Atoi(c.program[index])
		log.Printf("%v       Param %v == %-6v (position ). line[%v] has value: %v\n", c.name, paramIndex, index, index, r)
		return r
	}

	// Immediate
	r, _ := strconv.Atoi(c.program[line+paramIndex+1])
	log.Printf("%v       Param %v == %-6v (immediate).\n", c.name, paramIndex, r)
	return r
}

func (c *Computer) add(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (ADD)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex, _ := strconv.Atoi(c.program[line+3])
	c.program[targetIndex] = strconv.Itoa(a + b)
	log.Printf("%v       Param 2 == %-6v (output   ). line[%v] set to value: %v (%v + %v)\n", c.name, targetIndex, targetIndex, strconv.Itoa(a+b), a, b)
	return line + 4
}

func (c *Computer) multiply(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (MULTIPLY)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex, _ := strconv.Atoi(c.program[line+3])
	c.program[targetIndex] = strconv.Itoa(a * b)
	log.Printf("%v       Param 2 == %-6v (output). line[%v] set to value: %v (%v * %v)\n", c.name, targetIndex, targetIndex, strconv.Itoa(a*b), a, b)
	return line + 4
}

func (c *Computer) in(line int, input string) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (INPUT)\n", string(c.name[0]), line, c.program[line])
	targetIndex, _ := strconv.Atoi(c.program[line+1])
	c.program[targetIndex] = input
	log.Printf("%v       Input == %v\n", c.name, input)
	log.Printf("%v       Param 0 == %-6v (output). line[%v] set to value: %v\n", c.name, targetIndex, targetIndex, input)
	return line + 2
}

func (c *Computer) out(line int) (nextline int, output int) {
	log.Printf("%v%+5v: Opcode: %v (OUTPUT)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	log.Printf("%v -> Output: %v\n", c.name, a)
	return line + 2, a
}

func (c *Computer) jumpIfTrue(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (JUMP_IF_TRUE)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	if a != 0 {
		nextline = b
		log.Printf("%v       JUMP to %-6v\n", c.name, nextline)
	} else {
		nextline = line + 3
		log.Printf("%v       NO JUMP\n", c.name)
	}
	return nextline
}

func (c *Computer) jumpIfFalse(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (JUMP_IF_FALSE)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	if a == 0 {
		nextline = b
		log.Printf("%v       JUMP to %+6v\n", c.name, nextline)
	} else {
		nextline = line + 3
		log.Printf("%v       NO JUMP\n", c.name)
	}
	return nextline
}

func (c *Computer) lessThan(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (LESS_THAN)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex, _ := strconv.Atoi(c.program[line+3])
	if a < b {
		c.program[targetIndex] = string("1")
	} else {
		c.program[targetIndex] = string("0")
	}
	log.Printf("        Param 3 == %-6v (output). line[%v] set to %v\n", targetIndex, targetIndex, c.program[targetIndex])
	return line + 4
}

func (c *Computer) equal(line int) (nextline int) {
	log.Printf("%v%+5v: Opcode: %v (EQUAL_TO)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex, _ := strconv.Atoi(c.program[line+3])
	if a == b {
		c.program[targetIndex] = string("1")
	} else {
		c.program[targetIndex] = string("0")
	}
	log.Printf("        Param 3 == %-6v (output). line[%v] set to %v\n", targetIndex, targetIndex, c.program[targetIndex])
	return line + 4
}

// SetLineValue enables setting the value of a single line
func (c *Computer) SetLineValue(line int, value string) {
	c.program[line] = value
}

// GetLineValue returns the value at the requested line
func (c *Computer) GetLineValue(line int) string {
	return c.program[line]
}

// Run Instructs the machine to execute the laoded program
func (c *Computer) Run(wg *sync.WaitGroup) string {
	// var inputIndex int
	output := -1
	line := 0
	for {
		c.program[line] = padWithZeros(c.program[line])
		opcode := c.program[line]
		operation := opcode[len(opcode)-2:]
		switch {
		case operation == "01":
			line = c.add(line)
		case operation == "02":
			line = c.multiply(line)
		case operation == "03":
			// input := c.input[inputIndex]
			// inputIndex++
			// line = c.in(line, input)
			line = c.in(line, <-c.Input)
		case operation == "04":
			// line, output = c.out(line)
			line, output = c.out(line)
			c.Output <- strconv.Itoa(output)
		case operation == "05":
			line = c.jumpIfTrue(line)
		case operation == "06":
			line = c.jumpIfFalse(line)
		case operation == "07":
			line = c.lessThan(line)
		case operation == "08":
			line = c.equal(line)
		// case operation == "99":
		default:
			if wg != nil {
				wg.Done()
			}
			log.Printf("END OF program %v with output (%v)\n\n", c.name, output)
			return strconv.Itoa(output)
		}
	}
}

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

// NewComputer creates a new computer and loads the provided program
func NewComputer(name string, data []string) *Computer {
	c := Computer{
		name:    name,
		program: strings.Split(data[0], ","),
		Input:   make(chan string, 2),
		Output:  make(chan string, 2),
	}
	return &c
}
