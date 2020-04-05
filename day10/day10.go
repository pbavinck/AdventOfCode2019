package day10

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day10/input.txt"

// LogGroup The default log group this packages logs to
var LogGroup = "D10"

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
	xd int
	yd int
}

type aTarget struct {
	position  coord
	direction float32
}

type asteroidField struct {
	xSize      int
	ySize      int
	aSize      int               // number of asteroids
	field      [][]string        // the asteroid map itself
	directions map[coord]float32 // directions from sensor to asteroid
	sensor     coord             // (possible) location of the placed sensor
	targets    []coord           //aTarget  // list of targets
}

func newField(data []string) *asteroidField {
	a := asteroidField{}
	a.ySize = len(data)
	a.xSize = len(data[0])
	a.field = make([][]string, a.ySize)
	a.directions = make(map[coord]float32, a.ySize)
	for y := 0; y < a.ySize; y++ {
		a.field[y] = make([]string, a.xSize)
		for x := 0; x < a.xSize; x++ {
			a.field[y][x] = string(data[y][x : x+1])
			if a.field[y][x] == "#" {
				a.aSize++
			}
		}
	}
	a.targets = make([]coord, a.aSize-1) // sensor will be placed on one
	return &a
}

func (a *asteroidField) log(title string) {
	lg.Print(LogTagDebug, "%v", title)
	lg.Print(LogTagDebug, "%v", strings.Repeat("-", 4*a.xSize))
	for y := 0; y < a.ySize; y++ {
		s := ""
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			switch {
			case c == a.sensor:
				s = s + fmt.Sprintf("%+4v", "\u2588")
			case a.getLocation(c) == ".":
				s = s + fmt.Sprintf("%+4v", " ")
			default:
				s = s + fmt.Sprintf("%+4v", a.getLocation(c))
			}
		}
		lg.Print(LogTagDebug, "%v", s)
	}
}

func (a *asteroidField) contains(c coord) bool {
	return c.x >= 0 && c.x < a.xSize &&
		c.y >= 0 && c.y < a.ySize
}

func (a *asteroidField) getLocation(c coord) string {
	return a.field[c.y][c.x]
}

func (a *asteroidField) setLocation(c coord, s string) {
	a.field[c.y][c.x] = s
}

func getVector(c1, c2 coord) vector {
	return vector{
		xd: c2.x - c1.x,
		yd: c2.y - c1.y,
	}
}

func (c *coord) addVecor(v vector) {
	c.x += v.xd
	c.y += v.yd
}

func (a *asteroidField) printDirections() {
	fmt.Println("ySize:", a.ySize)
	a.placeSensor(coord{x: 20, y: 21})
	fmt.Printf("sensor: y %+v\n        ", a.sensor)
	for x := 0; x < a.xSize; x++ {
		fmt.Printf("%+8v", x)
	}
	fmt.Printf("\n")
	for y := 0; y < a.ySize; y++ {
		fmt.Printf("%+8v", y)
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			if c == a.sensor {
				fmt.Printf("%+8v", "\u2588")
				continue
			}
			if a.getLocation(c) == "#" {
				s := fmt.Sprintf("%.2f", a.directions[c])
				a.setLocation(c, s)
				fmt.Printf("%+8v", s)
			} else {
				fmt.Printf("%+8v", a.getLocation(c))

			}
		}
		fmt.Printf("\n")
	}
}

func getDirection(v vector, ySize int) float32 {
	// Keep in mind that the Y axis is reversed. Top Left == x:0, y:0, bottom right x:a.xSize-1,y:ySize-1
	// Directions are calculated by the quotient of x & y (not degrees). We add a Qx factor to these values
	// to make sure Q1 values > Q2 values > Q3 values > Q4 values
	// The origin of the vector is Q1 (lowest x, highest y)
	// Q4 | Q1
	// -------
	// Q3 | Q2
	xd := float32(v.xd)
	yd := float32(v.yd)
	var max float32 = float32(ySize + 1)
	switch {
	case xd >= 0 && -yd >= 0:
		// Q1 uses 3x max to prioritize Q1 as #1 (will use sort later)
		if xd == 0 {
			return 3.0*max + max // represents infinity
		}
		return 3.0*max + (-yd / xd)
	case xd >= 0 && -yd < 0:
		// Q2 uses 2x max
		if xd == 0 {
			return 2.0 * max // represents infinity
		}
		return 2.0*max + (xd / yd)
	case xd < 0 && -yd < 0:
		// Q3 uses 1x max
		return 1.0*max + (yd / -xd)
	case xd < 0 && -yd >= 0:
		// Q4
		if yd == 0 {
			return 1.0 * max // represents infinity
		}
		return -xd / -yd
	}
	return 0
}

func (a *asteroidField) calculateDirections() {
	i := 0
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			if c == a.sensor {
				continue
			}
			if a.getLocation(c) == "#" {
				v := getVector(a.sensor, c)
				a.directions[c] = getDirection(v, a.ySize)
				i++
			}
		}
	}
}

func (a *asteroidField) placeSensor(sensor coord) {
	a.sensor = sensor
	a.calculateDirections()
}

func (a *asteroidField) unmarkBlockedAsteroids() {
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			if a.getLocation(c) == "-" {
				a.setLocation(c, "#")
			}
		}
	}
}

func (a *asteroidField) markBlockedAsteroids() {
	// given a sensor position
	// find all astroid blocked by another astroid
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			v := getVector(a.sensor, c)
			firstObserved := true
			for {
				if !a.contains(c) || c == a.sensor {
					break
				}
				if a.getLocation(c) == "#" {
					if firstObserved {
						// This astroid is not blocked as it is the first one on the line from the sensor to this position
						// All following astroids on this line (multiples of same vector) are blocked
						firstObserved = false
					} else {
						// a.setLocation(c, "X")
						a.setLocation(c, "-")
					}
				}
				c.addVecor(v)
			}
		}
	}
	return
}

func (a *asteroidField) observable() (observed int) {
	// Marks unobservable astroids and counts the ones still observable
	a.markBlockedAsteroids()
	observed = 0
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			if a.getLocation(c) == "#" {
				observed++
			}
		}
	}
	return
}

func findPosition(data []string) (bestLocation coord, bestObserved int) {
	// check all possible locations for the sensor
	a := newField(data)
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			a = newField(data)

			c := coord{x: x, y: y}
			if a.getLocation(c) != "#" {
				continue
			}

			a.setLocation(c, "\u2588")
			a.sensor = c
			seen := a.observable()
			if seen > bestObserved {
				bestObserved = seen
				bestLocation = c
			}
			a.setLocation(c, "#")
		}
	}
	return
}

func (a *asteroidField) printTargets(title string) {
	lg.Print(LogTagDebug, "\nPRINT TARGETS: %v\n", title)
	a.log("at printTargets")
	i := 0
	for j, t := range a.targets {
		if a.directions[t] != 0.0 {
			if a.getLocation(t) == "#" {
				lg.Print(LogTagDebug, "%+3v [%v] - %+v, direction: %.3f\n", j, a.getLocation(t), t, a.directions[t])
			} else {
				lg.Print(LogTagDebug, "%v [%v] - %+v, direction: %.3f (blocked)\n", j, a.getLocation(t), t, a.directions[t])
			}
			i++
		}
	}
	lg.Print(LogTagDebug, "Target count: %v\n", i)
}

func (a *asteroidField) findTargets() {
	i, j := 0, 0
	a.targets = []coord{}
	for y := 0; y < a.ySize; y++ {
		for x := 0; x < a.xSize; x++ {
			c := coord{x: x, y: y}
			if a.getLocation(c) == "#" && c != a.sensor {
				a.targets = append(a.targets, c)
				i++
			} else if a.getLocation(c) == "-" {
				j++
			}
		}
	}

	// sort the targets from high to low
	sort.Slice(a.targets[:], func(i, j int) bool {
		return a.directions[a.targets[i]] > a.directions[a.targets[j]]
	})

	// fmt.Printf("Found targets: %v (%v blocked)\n", i, j)
}

func (a *asteroidField) fire360() int {
	// firing one full circle
	hit := 0
	a.markBlockedAsteroids()
	a.findTargets()
	for _, t := range a.targets {
		a.setLocation(t, ".")
		hit++
		if hit == 200 {
			fmt.Printf("Part 2 - 200th destroyed asteroid at %+v. Answer: %v\n", t, 100*t.x+t.y)
		}
	}
	a.unmarkBlockedAsteroids()
	return hit
}

func (a *asteroidField) keepFiring() {
	hit := 0
	round := 1
	lg.Print(LogTagInfo, "Asteroids to clear:  %v", a.aSize-1)
	for {
		if hit >= a.aSize-1 { // we don't shoot the sensor asteroid
			break
		}
		lg.Print(LogTagInfo, "Round: %v", round)
		hit += a.fire360()
		a.log("")
		lg.Print(LogTagInfo, "Hits so far: %v", hit)
		round++
	}
}

//SolvePart1 solves part 1 of day 10
func SolvePart1(data []string) {
	bestLocation, bestSeen := findPosition(data)

	// show the best solution
	a := newField(data)
	a.placeSensor(bestLocation)
	a.log("After sensor placed")
	fmt.Println("Part 1 - Answer:", bestSeen)
	fmt.Printf("Best location: %+v\n\n", bestLocation)
}

//SolvePart2 solves part 2 of day 10
func SolvePart2(data []string) {
	a := newField(data)
	a.placeSensor(coord{x: 20, y: 21})
	a.keepFiring()
}

// Solve runs day 10 assignment
func Solve() {
	fmt.Printf("\n*** DAY 10 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
