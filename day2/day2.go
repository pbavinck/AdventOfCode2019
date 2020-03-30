package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pbavinck/AofCode2019/loader"
)

// https://adventofcode.com/2019/day/2

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day2/input.txt"

func testProgram(lines []string, in1, in2 int) string {
	s := strings.Split(lines[0], ",")
	if in1 != -1 {
		s[1] = strconv.Itoa(in1)
		s[2] = strconv.Itoa(in2)
	}
	i := 0
	for {
		opcode, _ := strconv.Atoi(s[i])
		switch {
		case opcode == 1:
			{
				// add
				i1, _ := strconv.Atoi(s[i+1])
				a, _ := strconv.Atoi(s[i1])
				i2, _ := strconv.Atoi(s[i+2])
				b, _ := strconv.Atoi(s[i2])
				c, _ := strconv.Atoi(s[i+3])
				s[c] = strconv.Itoa(a + b)
				i += 4
			}
		case opcode == 2:
			{
				// multiply
				i1, _ := strconv.Atoi(s[i+1])
				a, _ := strconv.Atoi(s[i1])
				i2, _ := strconv.Atoi(s[i+2])
				b, _ := strconv.Atoi(s[i2])
				c, _ := strconv.Atoi(s[i+3])
				s[c] = strconv.Itoa(a * b)
				i += 4
			}
		case opcode == 99:
			{
				// halt
				// printState(s)
				return s[0]
			}
		default:
			{
				i++
			}
		}

	}
}

// SolvePart1 solves part 1 of Day 2 assignment
func SolvePart1() {
	s := loader.ReadStringsFromFile(inputFile, false)
	fmt.Printf("Part 1 - Answer: %v\n", testProgram(s, 12, 2))
}

// SolvePart2 solves part 2 of Day 2 assignment
func SolvePart2() {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			s := loader.ReadStringsFromFile(inputFile, false)
			if output := testProgram(s, noun, verb); output == "19690720" {
				fmt.Printf("Part 2 - Answer: %v (noun: %v, verb: %v)\n", 100*noun+verb, noun, verb)
				return
			}
		}
	}
}

//Solve solves the the day2 puzzle
func Solve() {
	fmt.Println("\n*** DAY 2 ***")
	SolvePart1()
	SolvePart2()
}
