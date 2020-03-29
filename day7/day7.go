package day7

import (
	"fmt"
	"log"
	"strconv"
	"sync"

	"github.com/pbavinck/AofCod2019/day7/permutation"
	"github.com/pbavinck/AofCod2019/loader"
	"github.com/pbavinck/AofCod2019/machines"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCod2019/day7/input.txt"

func tryPhase(data []string, phase string, part1 bool) int {
	var amplifiers [5]*machines.Computer
	for i := 0; i < len(amplifiers); i++ {
		// Create and start computers
		amplifiers[i] = machines.NewComputer(string(65+i), data)

		// Connect amps
		if i > 0 {
			amplifiers[i].Input = amplifiers[i-1].Output
		}
	}

	if !part1 {
		// Add feedback loop
		amplifiers[0].Input = amplifiers[4].Output
	}

	// !! First make sure all computers and have been created before sending info through channels!
	var wg sync.WaitGroup
	for i := 0; i < len(amplifiers); i++ {

		// Set phases
		amplifiers[i].Input <- string(phase[i])

		if i == 0 {
			// Send first signal to first amp
			amplifiers[0].Input <- "0"
		}

		// Only interested in the last amplifier
		if i == 4 {
			wg.Add(1)
			go amplifiers[i].Run(&wg)
		} else {
			go amplifiers[i].Run(nil)
		}
	}
	wg.Wait()

	r, _ := strconv.Atoi(<-amplifiers[4].Output)
	return r
}

//SolvePart1 solves part 1 of day 7
func SolvePart1(data []string) {
	r := new(permutation.Request)
	phases := r.GenerateFor("01234")
	maxSignal := 0

	log.Printf("Number of phases to try: %v\n", len(phases))

	for _, phase := range phases {
		output := tryPhase(data, phase, true)
		if output > maxSignal {
			maxSignal = output
		}
	}
	fmt.Println("Part 1 - Highest signal output:", maxSignal)
}

//SolvePart2 solves part 2 of day 7
func SolvePart2(data []string) {
	r := new(permutation.Request)
	phases := r.GenerateFor("56789")
	maxSignal := 0

	log.Printf("Number of phases to try: %v\n", len(phases))

	for _, phase := range phases {
		output := tryPhase(data, phase, false)
		if output > maxSignal {
			maxSignal = output
		}
	}
	fmt.Println("Part 2 - Highest signal output:", maxSignal)
}

// Solve runs day 7 assignment
func Solve() {
	fmt.Printf("\n*** DAY 7 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)
	SolvePart1(data)
	SolvePart2(data)
}