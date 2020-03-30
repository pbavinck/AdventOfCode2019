// Package day2a Day 2 revisited
// This solution uses the new Computer from the machines package, which was created at day 7
package day2a

import (
	"fmt"
	"strconv"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day2/input.txt"

//SolvePart1 solves part 1 of day 2
func SolvePart1(data []string) {
	c := machines.NewComputer("A", data)
	c.SetLineValue(1, "12")
	c.SetLineValue(2, "2")
	c.Run(nil)

	result := c.GetLineValue(0)
	fmt.Println("Part 1 - Answer:", result)
}

//SolvePart2 solves part 2 of day 2
func SolvePart2(data []string) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			c := machines.NewComputer("B", data)
			c.SetLineValue(1, strconv.Itoa(noun))
			c.SetLineValue(2, strconv.Itoa(verb))
			c.Run(nil)

			result := c.GetLineValue(0)
			if result == "19690720" {
				fmt.Printf("Part 2 - Answer: %v (noun: %v, verb: %v)\n", 100*noun+verb, noun, verb)
				return
			}
		}
	}

}

// Solve runs day 2 assignment
func Solve() {
	fmt.Printf("\n*** DAY 2a ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
