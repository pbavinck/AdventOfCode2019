package day9

import (
	"fmt"
	"sync"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day9/input.txt"

// While the computer produces output, this function captures and processes them
// Returns the last output
func processOutput(c *machines.Computer, wg *sync.WaitGroup, output *string) {
	defer wg.Done()
	for *output = range c.Output {
		// read all outputs from the output channel
		// fmt.Printf("out: %v\n", *output)
	}
}

func execute(data []string, input string) string {
	var wg sync.WaitGroup
	output := ""
	c := machines.NewComputer("A", data, 10000)
	if input != "-1" {
		c.Input <- input
	}
	wg.Add(2)
	go c.Run(&wg)
	go processOutput(c, &wg, &output)
	wg.Wait()
	return output
}

//SolvePart1 solves part 1 of day 9
func SolvePart1(data []string) {
	fmt.Println("Part 1 - Program output:", execute(data, "1"))
}

//SolvePart2 solves part 2 of day 9
func SolvePart2(data []string) {
	fmt.Println("Part 2 - Program output:", execute(data, "2"))
}

// Solve runs day 9 assignment
func Solve() {
	fmt.Printf("\n*** DAY 9 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
