package machines

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/pbavinck/lg"
)

// LogGroup The default group this package logs to
const LogGroup = "CPU"

// LogInfoTag info
var LogInfoTag int

// LogOpcodeTag opcode logs
var LogOpcodeTag int

// LogDebugTag debug
var LogDebugTag int

// LogParamTag parameter logs
var LogParamTag int

// LogIOTag action logs
var LogIOTag int

func init() {
	LogInfoTag, _ = lg.CreateTag("info", "CPU", lg.InfoLevel)
	LogOpcodeTag, _ = lg.CreateTag("opcode", "CPU", lg.InfoLevel)
	LogDebugTag, _ = lg.CreateTag("debug", "CPU", lg.DebugLevel)
	LogParamTag, _ = lg.CreateTag("param", "CPU", lg.DebugLevel)
	LogIOTag, _ = lg.CreateTag("IO", "CPU", lg.DebugLevel)
	lg.DisableGroup(LogGroup)
}

// Computer The computer to run the code
type Computer struct {
	name         string
	program      []string
	relativeBase int
	Input        chan string
	Output       chan string
}

func (c *Computer) getParamValue(line int, paramIndex int) int {
	opcode := c.program[line]
	if paramMode(opcode, paramIndex) == "0" {
		//postion mode
		index, _ := strconv.Atoi(c.program[line+paramIndex+1])
		r, _ := strconv.Atoi(c.program[index])
		lg.Print(LogParamTag, "%v       Param %v == %-6v (position ). line[%v] has value: %v\n", c.name, paramIndex, index, index, r)
		return r
	}

	if paramMode(opcode, paramIndex) == "2" {
		//relative mode
		index, _ := strconv.Atoi(c.program[line+paramIndex+1])
		r, _ := strconv.Atoi(c.program[c.relativeBase+index])
		lg.Print(LogParamTag, "%v       Param %v == %-6v (relative ). Base at %v, line[%v + %v] has value: %v\n", c.name, paramIndex, index, c.relativeBase, c.relativeBase, index, r)
		return r
	}

	// Immediate mode
	r, _ := strconv.Atoi(c.program[line+paramIndex+1])
	lg.Print(LogParamTag, "%v       Param %v == %-6v (immediate).\n", c.name, paramIndex, r)
	return r
}

func (c *Computer) getTargetIndex(line int, paramIndex int) int {
	opcode := c.program[line]
	if paramMode(opcode, paramIndex) == "0" {
		//postion mode
		index, _ := strconv.Atoi(c.program[line+paramIndex+1])
		lg.Print(LogParamTag, "%v       Param %v == %-6v (position ). Target line: %v\n", c.name, paramIndex, index, index)
		return index
	}

	if paramMode(opcode, paramIndex) == "2" {
		//relative mode
		index, _ := strconv.Atoi(c.program[line+paramIndex+1])
		lg.Print(LogParamTag, "%v       Param %v == %-6v (relative ). Base at %v, target line: %v + %v = %v\n", c.name, paramIndex, index, c.relativeBase, c.relativeBase, index, c.relativeBase+index)
		index = c.relativeBase + index
		return index
	}

	log.Fatal("Invalid param mode for terget index")
	return -1
}

func (c *Computer) add(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (ADD)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex := c.getTargetIndex(line, 2)
	c.program[targetIndex] = strconv.Itoa(a + b)
	lg.Print(LogParamTag, "%v       Param 2 == %-6v (output   ). line[%v] set to value: %v (%v + %v)\n", c.name, targetIndex, targetIndex, strconv.Itoa(a+b), a, b)
	nextLine = line + 4
	return
}

func (c *Computer) multiply(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (MULTIPLY)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex := c.getTargetIndex(line, 2)
	c.program[targetIndex] = strconv.Itoa(a * b)
	lg.Print(LogParamTag, "%v       Param 2 == %-6v (output). line[%v] set to value: %v (%v * %v)\n", c.name, targetIndex, targetIndex, strconv.Itoa(a*b), a, b)
	nextLine = line + 4
	return
}

func (c *Computer) in(line int, input string) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (INPUT)\n", string(c.name[0]), line, c.program[line])
	targetIndex := c.getTargetIndex(line, 0)
	c.program[targetIndex] = input
	lg.Print(LogIOTag, "%v    Input: <- %v\n", c.name, input)
	lg.Print(LogParamTag, "%v       Param 0 == %-6v (output). line[%v] set to value: %v\n", c.name, targetIndex, targetIndex, input)
	nextLine = line + 2
	return
}

func (c *Computer) out(line int) (nextLine int, output int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (OUTPUT)\n", string(c.name[0]), line, c.program[line])
	output = c.getParamValue(line, 0)
	lg.Print(LogIOTag, "%v -> Output: %v\n", c.name, output)
	nextLine = line + 2
	return
}

func (c *Computer) jumpIfTrue(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (JUMP_IF_TRUE)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	if a != 0 {
		nextLine = b
		lg.Print(LogDebugTag, "%v       JUMP to %-6v\n", c.name, nextLine)
	} else {
		nextLine = line + 3
		lg.Print(LogDebugTag, "%v       NO JUMP\n", c.name)
	}
	return
}

func (c *Computer) jumpIfFalse(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (JUMP_IF_FALSE)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	if a == 0 {
		nextLine = b
		lg.Print(LogDebugTag, "%v       JUMP to %+6v\n", c.name, nextLine)
	} else {
		nextLine = line + 3
		lg.Print(LogDebugTag, "%v       NO JUMP\n", c.name)
	}
	return
}

func (c *Computer) lessThan(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (LESS_THAN)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex := c.getTargetIndex(line, 2)
	if a < b {
		c.program[targetIndex] = string("1")
	} else {
		c.program[targetIndex] = string("0")
	}
	lg.Print(LogParamTag, "%v       Param 3 == %-6v (output). line[%v] set to %v\n", c.name, targetIndex, targetIndex, c.program[targetIndex])
	nextLine = line + 4
	return
}

func (c *Computer) equal(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (EQUAL_TO)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	b := c.getParamValue(line, 1)
	targetIndex := c.getTargetIndex(line, 2)
	if a == b {
		c.program[targetIndex] = string("1")
	} else {
		c.program[targetIndex] = string("0")
	}
	lg.Print(LogParamTag, "%v       Param 3 == %-6v (output). line[%v] set to %v\n", c.name, targetIndex, targetIndex, c.program[targetIndex])
	nextLine = line + 4
	return
}

func (c *Computer) base(line int) (nextLine int) {
	lg.Print(LogOpcodeTag, "%v%+5v: Opcode: %v (SET_BASE)\n", string(c.name[0]), line, c.program[line])
	a := c.getParamValue(line, 0)
	c.relativeBase = c.relativeBase + a
	lg.Print(LogDebugTag, "%v       BASE now %v\n", c.name, c.relativeBase)
	nextLine = line + 2
	return
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
			line = c.in(line, <-c.Input)
		case operation == "04":
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
		case operation == "09":
			line = c.base(line)
		default:
			if wg != nil {
				wg.Done()
			}
			close(c.Output)
			lg.Print(LogInfoTag, "END OF program %v with output (%v)\n\n", c.name, output)
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
func NewComputer(name string, data []string, memorySize int) *Computer {
	c := Computer{
		name:         name,
		program:      strings.Split(data[0], ","),
		relativeBase: 0,
		// make the channels have 2 capacity, because day 7 requires, phase and input signal to be passed
		Input: make(chan string, 2),
		// make the channels have 2 capacity, because day 7 requires it
		Output: make(chan string, 2),
	}
	if memorySize > len(c.program) {
		c.program = append(c.program, make([]string, memorySize-len(c.program))...)
	}
	return &c
}
