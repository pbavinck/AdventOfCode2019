package day17

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day17/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "D17"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
}

type coord struct {
	x int
	y int
}
type vector struct {
	x int
	y int
}

func (c *coord) add(v vector) coord { return coord{x: c.x + v.x, y: c.y + v.y} }

type location struct {
	position coord
}
type robot struct {
	position      coord
	direction     vector
	path          []string
	cpu           *machines.Computer
	intersections map[coord]*location
	floorplan     [][]string
	xMax          int
	yMax          int
	xMin          int
	yMin          int
}

func newRobot(data []string) *robot {
	r := robot{}
	r.position = coord{x: 0, y: 0}
	r.cpu = machines.NewComputer("O", data, 10000)
	r.intersections = make(map[coord]*location)
	// create y array of zero length
	r.floorplan = make([][]string, 0)
	return &r
}

func (r *robot) print() {
	for y := 0; y <= r.yMax; y++ {
		for x := 0; x <= r.xMax; x++ {
			fmt.Printf("%+2v", r.floorplan[y][x])

		}
		fmt.Printf("\n")

	}
	fmt.Printf("\n")
}

func (r *robot) inBounds(c coord) bool {
	if c.x < r.xMin || c.x > r.xMax ||
		c.y < r.yMin || c.y > r.yMax {
		return false
	}
	return true
}

func (r *robot) scanSkaffold() {
	go r.cpu.Run(nil)
	x, y := 0, 0
	// add first y line and prep x array
	r.floorplan = append(r.floorplan, make([]string, 0))
	r.floorplan[0] = make([]string, 0)

	for pixel := range r.cpu.Output {
		// fmt.Printf("%+3v", pixel)
		tmp, _ := strconv.Atoi(pixel)
		switch tmp {
		case 10:
			if x == 0 {
				continue
			}
			x = 0
			y++
			r.floorplan = append(r.floorplan, make([]string, 0))
			r.floorplan[y] = make([]string, 0)
		default:
			switch tmp {
			case 60: // <
				r.position = coord{x: x, y: y}
				r.direction = vector{x: -1, y: 0}
			case 62: // >
				r.position = coord{x: x, y: y}
				r.direction = vector{x: 1, y: 0}
			case 94: // ^
				r.position = coord{x: x, y: y}
				r.direction = vector{x: 0, y: -1}
			case 118: // v
				r.position = coord{x: x, y: y}
				r.direction = vector{x: 0, y: 1}
			}
			r.floorplan[y] = append(r.floorplan[y], fmt.Sprintf("%c", tmp))
			x++
		}
	}
	r.floorplan = r.floorplan[:len(r.floorplan)-1]
	r.yMax = len(r.floorplan) - 1
	r.xMax = len(r.floorplan[0]) - 1
	// r.print()
}

func pathChar(v vector) rune {
	switch {
	case v.x == 0 && v.y == -1:
		return '^'
	case v.x == 0 && v.y == 1:
		return 'v'
	case v.x == -1 && v.y == 0:
		return '<'
	case v.x == 1 && v.y == 0:
		return '>'
	}
	return ' '
}

func nextDirection(direction vector, clockwise bool) (vector, error) {
	if clockwise {
		switch direction {
		case vector{x: 0, y: -1}:
			return vector{x: 1, y: 0}, nil
		case vector{x: 1, y: 0}:
			return vector{x: 0, y: +1}, nil
		case vector{x: 0, y: +1}:
			return vector{x: -1, y: 0}, nil
		case vector{x: -1, y: 0}:
			return vector{x: 0, y: -1}, nil
		}
	} else {
		switch direction {
		case vector{x: 0, y: -1}:
			return vector{x: -1, y: 0}, nil
		case vector{x: -1, y: 0}:
			return vector{x: 0, y: +1}, nil
		case vector{x: 0, y: +1}:
			return vector{x: 1, y: 0}, nil
		case vector{x: 1, y: 0}:
			return vector{x: 0, y: -1}, nil
		}
	}
	return vector{x: 0, y: 0}, fmt.Errorf("invalid turn")
}

func (r *robot) findIntersections() {
	for y := 0; y <= r.yMax; y++ {
		for x := 0; x <= r.xMax; x++ {
			if r.floorplan[y][x] == "." {
				continue
			}

			// check right
			if x+1 <= r.xMax && r.floorplan[y][x+1] == "." {
				continue
			}
			// check left
			if x-1 >= r.xMin && r.floorplan[y][x-1] == "." {
				continue
			}
			// check down
			if y+1 <= r.yMax && r.floorplan[y+1][x] == "." {
				continue
			}
			// check left
			if y-1 >= r.yMin && r.floorplan[y-1][x] == "." {
				continue
			}
			c := coord{x: x, y: y}
			r.intersections[c] = &location{position: c}
			r.floorplan[y][x] = "O"

		}

	}
	// r.print()
}

func (r *robot) alignment() int {
	sum := 0
	for c := range r.intersections {
		sum += c.x * c.y
	}
	return sum
}

func (r *robot) move(steps int, path *[]string) int {
	found := false
	turn := ' '
	c := r.position
	direction := r.direction
	for i := 0; i < 3; i++ {
		switch i {
		case 0:
			// try straight
			c = r.position.add(direction)
			if r.inBounds(c) &&
				(r.floorplan[c.y][c.x] == "^" || r.floorplan[c.y][c.x] == "v" ||
					r.floorplan[c.y][c.x] == "<" || r.floorplan[c.y][c.x] == ">") {
				// jump over intersection
				c = c.add(direction)
				steps++
			}
		case 1: // turn right
			direction, _ = nextDirection(r.direction, true)
			c = r.position.add(direction)
			turn = 'R'
		case 2: // turn left
			direction, _ = nextDirection(r.direction, false)
			c = r.position.add(direction)
			turn = 'L'
		}

		if r.inBounds(c) && r.floorplan[c.y][c.x] == "#" {
			found = true
			if turn != ' ' {
				if steps > 0 {
					*path = append(*path, strconv.Itoa(steps))
				}
				*path = append(*path, string(turn))
				steps = 1
			} else {
				steps++
			}
			break
		}
	}

	if found {
		r.position = c
		r.direction = direction
		r.floorplan[c.y][c.x] = string(pathChar(r.direction))
		return steps
	}
	*path = append(*path, strconv.Itoa(steps))

	return steps
}

func (r *robot) buildPath() {
	r.scanSkaffold()
	steps, newSteps := 0, 0
	path := make([]string, 0)
	for {
		newSteps = r.move(steps, &path)
		if newSteps == steps {
			break
		}
		steps = newSteps
	}
	for i := 0; i < len(path); i = i + 2 {
		r.path = append(r.path, path[i]+path[i+1])
	}
}

// countOccurences counts the number of times s exists within path
func countOccurences(path, s []string) (int, []int) {
	count := 0
	indeces := []int{}
	i := 0
	for {
		if i > len(path)-len(s) {
			break
		}

		j := 0
		for j = 0; j < len(s); j++ {
			if s[j] != path[i+j] {
				break
			}
		}
		if j == len(s) {
			// we made it to the end, so it is a match
			indeces = append(indeces, i)
			count++
			i += len(s)
		} else {
			i++
		}
	}
	return count, indeces
}

// removeFunction clears the slice items covered by the function
// length = length of the functions
// indeces = starting points of the functions
func removeFunction(path []string, indeces []int, length int) []string {
	result := make([]string, len(path))
	k := 0 // index of active function block to look at (indeces are ordered)
	for i := 0; i < len(path); i++ {
		if k < len(indeces) && i >= indeces[k]+length {
			k++ // start looking at next block
		}
		if k >= len(indeces) || i < indeces[k] {
			result[i] = path[i]
		}
	}
	return result
}

type solutionType struct {
	main  []string
	funcs [3]struct {
		value   []string
		indeces []int
	}
	remaining int
}

// findFunctions finds all possible path breakdowns in to functions
// path = the path that needs to broken down into functions
// solutions = keep strack of all the found solutions
// draftSolution = the possible solution we are now working on
// fIndex = recursive depth -> each function used increases depth
// returns a possibly updated list of solutions
func findFunctions(path []string, solutions []solutionType, draftSolution solutionType, fIndex int) []solutionType {
	if draftSolution.remaining == 0 {
		solutions = append(solutions, draftSolution)
		return solutions
	}
	if fIndex > 2 {
		// not done, but we have run out of functions
		return solutions
	}
	// find start and end points (boundaries) of possible next functions
	start, end := 0, 0
	for start = 0; start < len(path); start++ {
		if path[start] != "" {
			break
		}
	}
	end = 0
	for end = start; end < len(path) && end-start < 7; end++ { // 7 because L1,L1,L1,L1,L1,L1,L1 == already > 20
		if path[end] == "" {
			break
		}
	}
	end = end - 1                       // we went 1 too far
	for l := 2; l <= end-start+1; l++ { // more than one otherwise no benefit
		f := path[start : start+l]

		// create a new function in the draft solution
		draftSolution.funcs[fIndex].value = make([]string, l)
		copy(draftSolution.funcs[fIndex].value, f)
		// keep track of where the new function was used in path
		count, indeces := countOccurences(path, f)
		draftSolution.funcs[fIndex].indeces = make([]int, count)
		copy(draftSolution.funcs[fIndex].indeces, indeces)

		draftSolution.remaining -= (len(indeces) * len(f))
		solutions = findFunctions(removeFunction(path, indeces, l), solutions, draftSolution, fIndex+1)
		draftSolution.remaining += (len(indeces) * len(f))
	}
	return solutions
}

// filterSolutions filters out solutions with too high memory footrpint
func filterSolutions(solutions []solutionType) []solutionType {
	filtered := make([]solutionType, 0)

solutions:
	for solIndex := range solutions {
		l := 2*len(solutions[solIndex].main) - 1
		if l > 20 {
			break
		}
		for _, f := range solutions[solIndex].funcs {
			l := 4*len(f.value) - 1 // L10 L5 (2 itmes) -> L,10,L,5, which becomes 7, don't ask me why 10 is counted as 1!! (4 * 2 -1)
			if l > 20 {
				continue solutions
			}
		}
		filtered = append(filtered, solutions[solIndex])
	}
	return filtered
}

func pickSolution(path []string, solutions []solutionType) solutionType {
	for solIndex := range solutions {
		// calculate main functions for the different solutions
		sol := &solutions[solIndex]
		f0, f1, f2, i := 0, 0, 0, 0
		remaining := len(path)
		for {
			if f0 < len(sol.funcs[0].indeces) && i == sol.funcs[0].indeces[f0] {
				f0++
				sol.main = append(sol.main, "A")
				remaining -= len(sol.funcs[0].value)
			} else if f1 < len(sol.funcs[1].indeces) && i == sol.funcs[1].indeces[f1] {
				f1++
				sol.main = append(sol.main, "B")
				remaining -= len(sol.funcs[1].value)
			} else if f2 < len(sol.funcs[2].indeces) && i == sol.funcs[2].indeces[f2] {
				f2++
				sol.main = append(sol.main, "C")
				remaining -= len(sol.funcs[2].value)
			}
			if remaining == 0 {
				break
			}
			i++
		}
	}

	solutions = filterSolutions(solutions)

	// pick a random solution
	rand.Seed(time.Now().UnixNano())
	solIndex := rand.Intn(len(solutions))
	return solutions[solIndex]
}

// commaMain creates comma delimited string of items in the main routine
func commaMain(in []string) string {
	s := ""
	for i := 0; i < len(in); i++ {
		s = s + string(in[i][0]) + ","
	}
	s = s[:len(s)-1] + "\n"
	return s
}

// commaFunc creates comma delimited string of items in one of the functions
func commaFunc(solution solutionType, index int) string {
	s := ""
	v := solution.funcs[index].value
	for i := 0; i < len(v); i++ {
		s = s + string(v[i][0]) + ","
		s = s + v[i][1:] + ","
	}
	s = s[:len(s)-1] + "\n"
	return s
}

//SolvePart1 solves part 1 of day 17
func SolvePart1(data []string) {
	r := newRobot(data)
	r.scanSkaffold()
	r.findIntersections()

	fmt.Println("Part 1 - Answer:", r.alignment())
}

//SolvePart2 solves part 2 of day 17
func SolvePart2(data []string) {
	var wg sync.WaitGroup
	r := newRobot(data)
	r.buildPath()
	path := make([]string, len(r.path))
	copy(path, r.path)

	solutions := make([]solutionType, 0)
	solutions = findFunctions(path, solutions, solutionType{remaining: len(path)}, 0)
	solution := pickSolution(path, solutions)

	r = newRobot(data)
	r.cpu.SetLineValue(0, "2")

	wg.Add(1)
	go r.cpu.Run(&wg)
	input := commaMain(solution.main) +
		commaFunc(solution, 0) +
		commaFunc(solution, 1) +
		commaFunc(solution, 2) +
		"n\n"

	for _, c := range input {
		r.cpu.Input <- strconv.Itoa(int(c))
	}

	dust := ""
	for dust = range r.cpu.Output {
		// drain output, remember last
	}
	wg.Wait()
	fmt.Println("Part 2 - Answer:", dust)

}

// Solve runs day 17 assignment
func Solve() {
	fmt.Printf("\n*** DAY 17 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
