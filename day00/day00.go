package dayXXX

import (
	"fmt"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/dayXXX/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "DXXX"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
}

//SolvePart1 solves part 1 of day XXX
func SolvePart1(data []string) {
	result := 0
	fmt.Println("Part 1 - Answer:", result)
}

//SolvePart2 solves part 2 of day XXX
func SolvePart2(data []string) {
	result := 0
	fmt.Println("Part 2 - Answer:", result)
}

// Solve runs day XXX assignment
func Solve() {
	fmt.Printf("\n*** DAY XXX ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	// uncomment if you need to reinitialize based on data
	// data = loader.ReadStringsFromFile(inputFile, false)
	SolvePart2(data)
}
