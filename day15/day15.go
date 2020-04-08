package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
	"github.com/pbavinck/AofCode2019/math2"
	"github.com/pbavinck/AofCode2019/system"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day15/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "D15"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
}

const (
	wall int = iota
	empty
	oxygen
	start
)
const (
	up int = iota + 1
	down
	left
	right
)

const (
	emptyPixel = " "
	wallPixel  = "\u2588"
	oxyPixel   = "O"
	startPixel = "S"
)

var pixelTable = map[int]string{
	wall:   wallPixel,
	empty:  emptyPixel,
	oxygen: oxyPixel,
	start:  startPixel,
}

type coord struct {
	x int
	y int
}

func (c *coord) up() coord    { return coord{x: c.x, y: c.y - 1} }
func (c *coord) right() coord { return coord{x: c.x + 1, y: c.y} }
func (c *coord) down() coord  { return coord{x: c.x, y: c.y + 1} }
func (c *coord) left() coord  { return coord{x: c.x - 1, y: c.y} }

type location struct {
	position coord
	visited  bool
	is       int
	distance int
}

type robot struct {
	position  coord
	cpu       *machines.Computer
	floorplan map[coord]*location
	oxyAt     coord
	xMax      int
	yMax      int
	xMin      int
	yMin      int
}

func newRobot(data []string) *robot {
	r := robot{}
	r.position = coord{x: 0, y: 0}
	r.oxyAt = coord{x: 0, y: 0} // not there == starting position
	r.cpu = machines.NewComputer("O", data, 10000)
	r.floorplan = make(map[coord]*location)
	r.floorplan[r.position] = &location{
		position: r.position,
		visited:  true,
		is:       start,
		distance: 0,
	}
	return &r
}

func (r *robot) print() {
	screen := make([][]string, r.yMax-r.yMin+1)
	for y := range screen {
		screen[y] = make([]string, r.xMax-r.xMin+1)
	}
	for c := range r.floorplan {
		screen[-r.yMin+c.y][-r.xMin+c.x] = pixelTable[r.floorplan[c].is]
	}
	system.CallClear()
	started := false
	for y := range screen {
		s := ""
		for x := range screen[y] {
			if screen[y][x] == string("") {
				s += fmt.Sprintf("%+1v", emptyPixel)
			} else {
				s += fmt.Sprintf("%+1v", screen[y][x])
				started = true
			}
		}
		if started {
			fmt.Printf("%v\n", s)
		}
	}
	fmt.Println(strings.Repeat("\u2014", 1*(r.xMax-r.xMin+1)))
	return
}

func reverse(d int) int {
	if d == 1 || d == 2 {
		return 1 + (d % 2)
	}
	return 3 + (d % 2)
}

func (r *robot) checkDirection(pos coord, d int) {
	var newpos coord
	switch d {
	case up:
		newpos = pos.up()
	case down:
		newpos = pos.down()
	case left:
		newpos = pos.left()
	case right:
		newpos = pos.right()
	}

	distance := r.floorplan[pos].distance + 1
	_, ok := r.floorplan[newpos]
	if ok && distance >= r.floorplan[newpos].distance {
		// if we have been here before and it is not a shorter route
		return
	}
	r.xMin = math2.IntMin(r.xMin, newpos.x)
	r.xMax = math2.IntMax(r.xMax, newpos.x)
	r.yMin = math2.IntMin(r.yMin, newpos.y)
	r.yMax = math2.IntMax(r.yMax, newpos.y)

	// move to position
	r.cpu.Input <- strconv.Itoa(d)
	result, _ := strconv.Atoi(<-r.cpu.Output)
	if ok && distance < r.floorplan[newpos].distance && result == empty {
		lg.Print(LogTagInfo, "Found shorter route at %+v (%v < %v)", newpos, distance, r.floorplan[newpos].distance)
	}
	r.floorplan[newpos] = &location{
		position: newpos,
		visited:  true,
		is:       result,
		distance: distance,
	}
	lg.Print(LogTagInfo, "%v at %+v (%v steps)", pixelTable[result], newpos, distance)
	// r.print()
	switch result {
	case wall:
		// no need to back paddle, so leave now
		return
	case oxygen:
		r.oxyAt = newpos
	case empty:
		// find possible next positions
		r.checkDirection(newpos, left)
		r.checkDirection(newpos, down)
		r.checkDirection(newpos, right)
		r.checkDirection(newpos, up)
	}
	r.cpu.Input <- strconv.Itoa(reverse(d))
	_, _ = strconv.Atoi(<-r.cpu.Output)
}

func (r *robot) oxygenFlow() (minutes int) {
	todo1 := make(map[coord]int) // use two todo maps in order to swap between them
	todo2 := make(map[coord]int)
	todo1[r.oxyAt] = oxygen
	todonow, todonext, t := &todo1, &todo2, &todo2
	for {
		if len(*todonow) == 0 {
			return minutes - 1 // we don't count the initial one
		}
		for pos := range *todonow {
			r.floorplan[pos].is = oxygen
			delete((*todonow), pos)
			if r.floorplan[pos.up()].is == empty {
				(*todonext)[pos.up()] = r.floorplan[pos.up()].is
			}
			if r.floorplan[pos.down()].is == empty {
				(*todonext)[pos.down()] = r.floorplan[pos.down()].is
			}
			if r.floorplan[pos.left()].is == empty {
				(*todonext)[pos.left()] = r.floorplan[pos.left()].is
			}
			if r.floorplan[pos.right()].is == empty {
				(*todonext)[pos.right()] = r.floorplan[pos.right()].is
			}
		}
		minutes++
		t = todonow // swap map to use
		todonow = todonext
		todonext = t
		// r.print()
	}
}

//SolvePart1 solves part 1 of day 15
func SolvePart1(data []string) {
	r := newRobot(data)
	go r.cpu.Run(nil)
	// explore the four directions
	r.checkDirection(r.position, left)
	r.checkDirection(r.position, down)
	r.checkDirection(r.position, right)
	r.checkDirection(r.position, up)
	// r.print()
	fmt.Printf("Part 1 - Answer: %v (Oxygen at %+v)\n", r.floorplan[r.oxyAt].distance, r.oxyAt)
}

//SolvePart2 solves part 2 of day 15
func SolvePart2(data []string) {
	r := newRobot(data)
	go r.cpu.Run(nil)
	r.checkDirection(r.position, left)
	r.checkDirection(r.position, down)
	r.checkDirection(r.position, right)
	r.checkDirection(r.position, up)
	r.floorplan[r.position].is = empty
	fmt.Println("Part 2 - Answer:", r.oxygenFlow())
}

// Solve runs day 15 assignment
func Solve() {
	fmt.Printf("\n*** DAY 15 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
