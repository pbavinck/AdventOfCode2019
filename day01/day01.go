package day01

import (
	"fmt"
	"strconv"

	"github.com/pbavinck/AofCode2019/loader"
)

// https://adventofcode.com/2019/day/1

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day01/input.txt"

//SolvePart1 solves part 1 of the the day01 puzzle
func SolvePart1(data []int) int {
	totalFuel := 0
	for i := range data {
		mass := data[i]
		fuel := int(mass/3) - 2

		totalFuel += fuel
	}

	fmt.Println("Part 1 - Total fuel:", strconv.Itoa(totalFuel))
	return totalFuel
}

//SolvePart2 solves part 1 of the the day01 puzzle
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

	fmt.Println("Part 2 - Total fuel:", strconv.Itoa(totalFuel))
	return totalFuel

}

// Solve runs day 1 assignment
func Solve() {
	fmt.Printf("\n*** DAY 1 ***\n")
	data := loader.ReadIntsFromFile(inputFile, false)
	SolvePart1(data)
	SolvePart2(data)
}
