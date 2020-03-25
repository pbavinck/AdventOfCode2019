package dayXXX

import (
	"fmt"

	"github.com/pbavinck/AofCod2019/loader"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCod2019/dayXXX/input.txt"

//SolvePart1 solves part 1 of day XXX
func SolvePart1() int {
	result := 0
	fmt.Println("Answer (Part 1): ", result)
	return result
}

//SolvePart2 solves part 2 of day XXX
func SolvePart2() int {
	result := 0
	fmt.Println("Answer (Part 1): ", result)
	return result
}

// Solve runs day XXX assignment
func Solve() {
	fmt.Printf("\n*** DAY XXX : Part 1 ***\n")
	data := loader.ReadIntsFromFile(inputFile, false)
	fmt.Printf("%v line(s) read from input\n", len(data))

	SolvePart1()

	fmt.Printf("\n*** DAY XXX : Part 2 ***\n")
	SolvePart2()
}
