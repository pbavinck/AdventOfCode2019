package dayXXX

import (
	"fmt"
	"log"
	"os"

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
	// log.SetOutput(ioutil.Discard)
	log.SetOutput(os.Stderr)

	fmt.Printf("\n*** DAY XXX ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)
	fmt.Printf("%v line(s) in input\n", len(data))

	SolvePart1()

	data := loader.ReadStringsFromFile(inputFile, false)
	SolvePart2()
}
