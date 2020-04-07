package day14

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day14/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "D14"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

// LogTagLoad Used to prefix log items of data load
var LogTagLoad int

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
	LogTagLoad, _ = lg.CreateTag("load", LogGroup, lg.DebugLevel)
}

type elementAmount struct {
	name   string // name of the element
	amount int    // the amount associated
}

type aReaction struct {
	input  []elementAmount // the amount of input ingredients (elements) required by the reaction
	output elementAmount   // the amount of element a reaction produces
}

type nanoFactory struct {
	reactions map[string]*aReaction // the reaction formula with input and output
	// available map[string]int
	todo map[string]int
}

func newReaction(input string, output string) *aReaction {
	r := aReaction{}
	for _, chem := range strings.Split(input, ",") {
		split := strings.Split(strings.TrimSpace(chem), " ")
		amount, _ := strconv.Atoi(split[0])
		name := split[1]
		r.input = append(r.input, elementAmount{name: name, amount: amount})
	}
	split := strings.Split(strings.TrimSpace(output), " ")
	amount, _ := strconv.Atoi(split[0])
	name := split[1]
	r.output = elementAmount{name: name, amount: amount}
	return &r
}

func newNanoFactory(data []string) *nanoFactory {
	n := nanoFactory{}
	n.reactions = make(map[string]*aReaction)
	n.todo = make(map[string]int)
	for _, s := range data {
		split := strings.Split(s, " => ")
		input := split[0]
		output := split[1]
		split = strings.Split(strings.TrimSpace(output), " ")
		name := split[1]

		n.reactions[name] = newReaction(input, output)
	}
	lg.Print(LogTagInfo, "%v reactions loaded", len(n.reactions))
	n.printReactions()
	return &n
}

func (n *nanoFactory) printReactions() {
	for _, r := range n.reactions {
		lg.Print(LogTagLoad, "%+v => %+v", r.input, r.output)
	}
}

func (e *elementAmount) formatted() string {
	return fmt.Sprintf("%v %v", e.amount, e.name)
}
func formattedElems(elems []elementAmount) string {
	s := ""
	for _, e := range elems {
		s += e.formatted() + ", "
	}
	if len(s) > 0 {
		s = s[:len(s)-2]
	}
	return s
}

func (n *nanoFactory) doTodo() {
	for {
		stillBusy := false
		for element, amount := range n.todo {
			if element == "ORE" {
				continue // no reaction to get ORE
			}
			if amount > 0 { // still something to do
				stillBusy = true
				times := 1 + (amount-1)/n.reactions[element].output.amount
				n.todo[element] -= times * n.reactions[element].output.amount
				for _, input := range n.reactions[element].input {
					n.todo[input.name] += times * input.amount
				}
			}
		}

		if !stillBusy {
			return
		}
	}
}

//SolvePart1 solves part 1 of day 14
func SolvePart1(data []string) {
	n := newNanoFactory(data)
	n.todo["FUEL"] = 1
	n.doTodo()
	fmt.Println("Part 1 - Answer:", n.todo["ORE"])
}

func calcMaxFuel(data []string) int {
	n := newNanoFactory(data)
	n.todo["FUEL"] = 1
	n.doTodo()
	availableOre := 1_000_000_000_000
	oreFor1Fuel := n.todo["ORE"]
	fuel, step := 0, availableOre/oreFor1Fuel
	backup := make(map[string]int) // to restore after failed attempts
	for {
		// reset
		n.todo = make(map[string]int)
		for name, amount := range backup {
			n.todo[name] = amount
		}

		n.todo["FUEL"] += step
		n.doTodo()

		if n.todo["ORE"] <= availableOre {
			// it fitted, let's try another one
			fuel += step
			backup = n.todo
			continue
		}

		if step > 1 {
			// it did not fit, try half the step size
			step /= 2
			continue
		}

		// Now even step size of 1 failed
		break
	}
	return fuel
}

//SolvePart2 solves part 2 of day 14
func SolvePart2(data []string) {

	fmt.Println("Part 2 - Answer:", calcMaxFuel(data))
}

// Solve runs day 14 assignment
func Solve() {
	lg.DisableTag(LogTagLoad)
	fmt.Printf("\n*** DAY 14 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
