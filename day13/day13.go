package day13

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/pbavinck/AofCode2019/loader"
	"github.com/pbavinck/AofCode2019/machines"
	"github.com/pbavinck/AofCode2019/math2"
	"github.com/pbavinck/AofCode2019/system"
	"github.com/pbavinck/lg"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day13/input.txt"

var ballInit = coord{x: -1, y: -1}
var paddleInit = coord{x: -1, y: -1}

// LogGroup The default log group this packages logs to
var LogGroup = "D13"

// LogTagInfo Used to prefix info log items
var LogTagInfo int

// LogTagDebug Used to prefix debug log items
var LogTagDebug int

const (
	emptyTile  = iota // 0
	wallTile          // 1
	blockTile         // 2
	paddleTile        // 3
	ballTile          // 4
)
const (
	emptyPixel  = " "
	wallPixel   = "\u2588"
	blockPixel  = "\u25A1"
	paddlePixel = "\u2582"
	ballPixel   = "\u2299"
)

var pixelTable = map[int]string{
	emptyTile:  emptyPixel,
	wallTile:   wallPixel,
	blockTile:  blockPixel,
	paddleTile: paddlePixel,
	ballTile:   ballPixel,
}

type coord struct {
	x int
	y int
}
type aDisplay struct {
	xMin       int
	xMax       int
	yMin       int
	yMax       int
	items      map[coord]int
	tileCounts map[int]int
}

type anArcade struct {
	cpu            *machines.Computer
	display        *aDisplay
	ballPosition   coord
	paddlePosition coord
	score          int
	blocks         int
}

func init() {
	LogTagInfo, _ = lg.CreateTag("", LogGroup, lg.InfoLevel)
	LogTagDebug, _ = lg.CreateTag("", LogGroup, lg.DebugLevel)
}

func (d *aDisplay) print() (onscreen int) {
	screen := make([][]string, d.yMax-d.yMin+1)
	for y := range screen {
		screen[y] = make([]string, d.xMax-d.xMin+1)
	}
	for c := range d.items {
		screen[-d.yMin+c.y][-d.xMin+c.x] = pixelTable[d.items[c]]
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
				onscreen++
				started = true
			}
		}
		if started {
			fmt.Printf("%v\n", s)
		}
	}
	fmt.Println(strings.Repeat("\u2014", 1*(d.xMax-d.xMin+1)))
	return
}
func (d *aDisplay) setPixel(position coord, tileID int) {
	d.items[position] = tileID
	d.xMin = math2.IntMin(d.xMin, position.x)
	d.xMax = math2.IntMax(d.xMax, position.x)
	d.yMin = math2.IntMin(d.yMin, position.y)
	d.yMax = math2.IntMax(d.yMax, position.y)
}

func newArcade(data []string) *anArcade {
	a := anArcade{}
	a.cpu = machines.NewComputer("A", data, 10000)
	a.ballPosition = coord{x: -1, y: -1}
	a.paddlePosition = coord{x: -1, y: -1}
	a.attachDisplay()
	return &a
}

func (a *anArcade) attachDisplay() {
	a.display = &aDisplay{}
	a.display.items = make(map[coord]int)
	a.display.tileCounts = make(map[int]int)
}

func (a *anArcade) print() {
	a.display.print()
	fmt.Println("Score:", a.score)
}

func (a *anArcade) play(wg *sync.WaitGroup) {
	defer wg.Done()
	var x, y, p int
	for {
		t, ok := <-a.cpu.Output
		if !ok { // if channel has closed
			break
		}
		x, _ = strconv.Atoi(t)
		y, _ = strconv.Atoi(<-a.cpu.Output)
		p, _ = strconv.Atoi(<-a.cpu.Output)
		if x == -1 && y == 0 {
			a.score = p
			lg.Print(LogTagInfo, "Score: %v", a.score)
		} else {
			a.display.tileCounts[p]++
			c := coord{x: x, y: y}
			lg.Print(LogTagDebug, "Read output: %+v = %+2v\n", c, p)
			a.display.setPixel(c, p)
			switch p {
			case ballTile:
				lg.Print(LogTagDebug, "ball at %+v\n", c)
				a.ballPosition = c
				lg.Print(LogTagDebug, "ball X:%v, paddle X:%v\n", a.ballPosition.x, a.paddlePosition.x)
				if a.ballPosition.x > a.paddlePosition.x {
					a.cpu.Input <- "1"
					lg.Print(LogTagInfo, "Input sent: 1 (move right)\n")
				} else if a.ballPosition.x < a.paddlePosition.x {
					a.cpu.Input <- "-1"
					lg.Print(LogTagInfo, "Input sent: -1 (move left)\n")
				} else {
					a.cpu.Input <- "0"
					lg.Print(LogTagInfo, "Input sent: 0 (no move)\n")
				}
				// a.print()
			case paddleTile:
				lg.Print(LogTagInfo, "paddle at %+v\n", c)
				a.paddlePosition = c
			}
		}
	}
}

//SolvePart1 solves part 1 of day 13
func SolvePart1(data []string) {
	var wg sync.WaitGroup
	a := newArcade(data)
	wg.Add(2)
	go a.cpu.Run(&wg)
	go a.play(&wg)
	wg.Wait()
	// a.display.print()
	fmt.Println("Part 1 - Answer:", a.display.tileCounts[blockTile])
}

//SolvePart2 solves part 2 of day 13
func SolvePart2(data []string) {
	var wg sync.WaitGroup
	a := newArcade(data)
	a.cpu.SetLineValue(0, "2")
	wg.Add(2)
	go a.cpu.Run(&wg)
	go a.play(&wg)
	wg.Wait()
	// a.print()
	fmt.Println("Part 2 - Answer:", a.score)
}

// Solve runs day 13 assignment
func Solve() {
	fmt.Printf("\n*** DAY 13 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
