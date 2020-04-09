package day16

import (
	"fmt"
	"strconv"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/math2"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day16/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "D16"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
}

func convert(data string) []int {
	result := []int{}
	for i := 0; i < len(data); i++ {
		t, _ := strconv.Atoi(string(data[i]))
		result = append(result, t)
	}
	return result
}

// calcDigit calculates a single of the signal
func calcDigit(signal []int, digitIndex int) (resultDigit int) {
	basePattern := []int{0, 1, 0, -1}
	resultDigit = 0
	patternLength := len(basePattern)
	repeatCount := digitIndex + 1 // number of times the digit of the base pattern needs to be repeated

	// we can start  from digit index, because all values before are multiplied by 0
	for i := digitIndex; i < len(signal); i++ {
		// factor is the multiply factor to be used: 0, 1, 0 or -1
		// i+1 because the pattern used is skipping first value
		// example if we are calculating the value for digit 5, then reapeatCount == 6
		// skipping first zero, gives us 4 x 0, then 5 x 1, 5 x 0 and 5 x -1
		// if i == 87 we get 87 / 6 = 14(.5), and with MOD (14 % 4) = 2
		// this is the index to be used in basePattern, in this case for i this is 0
		factor := basePattern[((i+1)/repeatCount)%patternLength]
		resultDigit += factor * signal[i]
	}
	resultDigit = math2.IntAbs(resultDigit) % 10
	return
}

func calcPhase(signal []int) (result []int) {
	result = make([]int, len(signal))
	for digitIndex := 0; digitIndex < len(signal); digitIndex++ {
		result[digitIndex] = calcDigit(signal, digitIndex)
	}
	return
}

//SolvePart1 solves part 1 of day 16
func SolvePart1(data []string) {
	signal := convert(data[0])
	for i := 0; i < 100; i++ {
		signal = calcPhase(signal)
	}
	fmt.Printf("Part 1 - Answer: ")
	for i := 0; i < 8; i++ {
		fmt.Printf("%v", signal[i])
	}
	fmt.Printf("\n")
}

//SolvePart2 solves part 2 of day 16
func SolvePart2(data []string) {
	signalMultiplier := 10_000
	messageSize := 8
	offsetSize := 7

	signal := convert(data[0])
	offset, _ := strconv.Atoi(data[0][0:offsetSize])
	signal2 := make([]int, signalMultiplier*len(signal))
	for i := 0; i < signalMultiplier*len(signal); i++ {
		signal2[i] = signal[i%len(signal)]
	}

	// looking at the input data, the offset is 5.9M, while total signal is 6.5M long, so first
	// 5.9M multiply factors are 0, then the next 5.9M multiply factors are 1, so we just
	// need for eacht digit to sum the remaining digits. It is easier to go from the back to the front
	for i := 0; i < 100; i++ {
		sum := 0
		for i := len(signal2) - 1; i >= offset; i-- {
			sum += signal2[i]
			signal2[i] = math2.IntAbs(sum) % 10
		}
	}
	fmt.Printf("Part 2 - Answer: ")
	for i := offset; i < offset+messageSize; i++ {
		fmt.Printf("%v", signal2[i])
	}
	fmt.Printf("\n")

}

// Solve runs day 16 assignment
func Solve() {
	fmt.Printf("\n*** DAY 16 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
