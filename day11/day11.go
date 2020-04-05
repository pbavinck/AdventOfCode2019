package day11

import (
	"fmt"
	"strings"
	"sync"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day11/input.txt"

const dimension = 150

// const white = "\u2588"
const (
	blackHull  = "."
	whitePaint = "#"
	blackPaint = " "
	blackCode  = "0"
	whiteCode  = "1"

	turnLeft  = "0"
	turnRight = "1"
)

type robot struct {
	computer  *machines.Computer
	position  coord
	direction int
	surface   map[coord]string
	xMin      int
	xMax      int
	yMin      int
	yMax      int
}
type coord struct {
	x int
	y int
}

func intMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func intMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func newRobot() *robot {
	r := robot{}
	r.surface = make(map[coord]string)
	r.position = coord{x: 0, y: 0}
	return &r
}

func (r *robot) print() {
	fmt.Println(strings.Repeat("-", r.xMax-r.xMin+1))
	canvas := make([][]string, r.yMax-r.yMin+1)
	for y := range canvas {
		canvas[y] = make([]string, r.xMax-r.xMin+1)
	}
	for c := range r.surface {
		canvas[-r.yMin+c.y][-r.xMin+c.x] = r.surface[c]
	}

	started := false				
	for y := range canvas {
		s := ""
		for x := range canvas[y] {
			if canvas[y][x] == string("") {
				s += fmt.Sprintf("%v", blackHull)
			} else {
				s += fmt.Sprintf("%v", canvas[y][x])
				started = true
			}
		}
		if started {
			fmt.Printf("%v\n", s)
		}
	}
	fmt.Println(strings.Repeat("-", r.xMax-r.xMin+1))
}

func (r *robot) getLocation(c coord) string {
	if color, ok := r.surface[c]; ok {
		return color
	}
	return blackHull
}

func (r *robot) setLocation(c coord, color string) {
	r.surface[c] = color
}

func (r *robot) paint(colorCode string) {
	if colorCode == blackCode {
		r.surface[r.position] = blackPaint
	} else {
		r.surface[r.position] = whitePaint
	}
	r.xMin = intMin(r.xMin, r.position.x)
	r.xMax = intMax(r.xMax, r.position.x)
	r.yMin = intMin(r.yMin, r.position.y)
	r.yMax = intMax(r.yMax, r.position.y)
}

func (r *robot) turn(direction string) {
	// direction stored as values 0 to 3 (0 top, 1 right, 2 down, 3 left).
	if direction == turnLeft {
		r.direction = (r.direction - 1) % 4
	} else {
		r.direction = (r.direction + 1) % 4
	}
	// correct for negative
	if r.direction < 0 {
		r.direction = r.direction + 4
	}
}

func (r *robot) move() {
	switch r.direction {
	case 0:
		r.position = coord{x: r.position.x, y: r.position.y - 1} // up
	case 1:
		r.position = coord{x: r.position.x + 1, y: r.position.y} // right
	case 2:
		r.position = coord{x: r.position.x, y: r.position.y + 1} // down
	case 3:
		r.position = coord{x: r.position.x - 1, y: r.position.y} // left
	}
}

func (r *robot) look() (colorCode string, alreadyPainted bool) {
	p := r.getLocation(coord{x: r.position.x, y: r.position.y})
	switch p {
	case whitePaint:
		colorCode = whiteCode
		alreadyPainted = true
	case blackHull:
		colorCode = blackCode
		alreadyPainted = false
	case blackPaint:
		colorCode = blackCode
		alreadyPainted = true
	}

	return
}

func (r *robot) start(data []string) int {
	var wg sync.WaitGroup
	r.computer = machines.NewComputer("R", data, 10000)
	painted := 0

	colorCode, _ := r.look()
	r.computer.Input <- colorCode
	wg.Add(2)
	go r.computer.Run(&wg)
	go func(*machines.Computer, *sync.WaitGroup) {
		defer wg.Done()

		first := true
		for {
			output, ok := <-r.computer.Output
			if !ok {
				return
			}

			if first {
				_, alreadyPainted := r.look()
				if !alreadyPainted {
					painted++
				}

				r.paint(output)
				first = false
			} else {
				r.turn(output)
				r.move()
				first = true
				colorCode, _ := r.look()
				r.computer.Input <- colorCode
			}
		}
	}(r.computer, &wg)

	wg.Wait()
	return painted
}

//SolvePart1 solves part 1 of day 11
func SolvePart1(data []string) {
	r := newRobot()
	fmt.Println("Part 1 - Answer:", r.start(data))
}

//SolvePart2 solves part 2 of day 11
func SolvePart2(data []string) {
	r := newRobot()
	r.paint("1") // paint white tile under robot
	fmt.Println("Part 2 - Answer:")
	r.start(data)
	r.print()
}

// Solve runs day 11 assignment
func Solve() {
	fmt.Printf("\n*** DAY 11 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
