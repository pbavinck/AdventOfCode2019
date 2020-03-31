// Package day5a Day 5 revisited
// This solution uses the new Computer from the machines package, which was created at day 7
package day5a

import (
	"fmt"
	"sync"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day5a/input.txt"

//SolvePart1 solves part 1 of day 5
func SolvePart1(data []string) {
	var wg sync.WaitGroup
	output := ""
	c := machines.NewComputer("A", data, 0)
	c.Input <- "1"
	wg.Add(2)
	go c.Run(&wg)
	go func(*machines.Computer, *sync.WaitGroup) {
		defer wg.Done()
		// Clear output channel (using output through closure, output will contain the last one)
		for output = range c.Output {
			// fmt.Printf("out: %v\n", output)
		}
	}(c, &wg)
	wg.Wait()
	fmt.Println("Part 1 - Program output:", output)
}

//SolvePart2 solves part 2 of day 5
func SolvePart2(data []string) {
	var wg sync.WaitGroup
	output := ""
	c := machines.NewComputer("A", data, 0)
	c.Input <- "5"
	wg.Add(2)
	go c.Run(&wg)
	go func(*machines.Computer, *sync.WaitGroup) {
		defer wg.Done()
		// Clear output channel (using output through closure, output will contain the last one)
		for output = range c.Output {
			// fmt.Printf("out: %v\n", output)
		}
	}(c, &wg)
	wg.Wait()
	fmt.Println("Part 2 - Program output:", output)
}

// Solve runs day 5 assignment
func Solve() {
	fmt.Printf("\n*** DAY 5a ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
