package day1

import (
	"fmt"
	"strconv"

	"github.com/pbavinck/AofCod2019/loader"
)

// https://adventofcode.com/2019/day/1

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCod2019/day1/input.txt"

//SolvePart1 solves part 1 of the the day1 puzzle
func SolvePart1(data []int) int {
	totalFuel := 0
	for i := range data {
		mass := data[i]
		fuel := int(mass/3) - 2

		totalFuel += fuel
	}

	fmt.Println("Total (Part 1): ", strconv.Itoa(totalFuel))
	return totalFuel
}

//SolvePart2 solves part 1 of the the day1 puzzle
func SolvePart2(data []int) int {
	totalFuel := 0
	for i := range data {
		mass := data[i]
		fuel := int(mass/3) - 2
		extra := fuel
		for {
			extra = int(extra/3) - 2
			if extra <= 0 {
				break
			}
			fuel += extra
		}

		totalFuel += fuel
	}

	fmt.Println("Total (Part 2): ", strconv.Itoa(totalFuel))
	return totalFuel

}

// Solve runs day 1 assignment
func Solve() {
	fmt.Printf("\n*** DAY 1 : Part 1 ***")
	data := loader.ReadIntsFromFile(inputFile, false)

	fmt.Printf("%v line(s) read from input\n", len(data))
	SolvePart1(data)

	fmt.Println("\n*** DAY 1 : Part 2 ***")
	SolvePart2(data)
}
