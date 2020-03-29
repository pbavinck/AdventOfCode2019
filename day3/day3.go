package day3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pbavinck/AofCod2019/loader"
)

type coord struct {
	x int
	y int
}

type segmentType struct {
	start     coord
	end       coord
	xd        int
	yd        int
	direction string
	length    int
}

type wireType struct {
	segments []segmentType
	steps    int
	index    int
}

type intersectionType struct {
	pos       coord
	wireSteps [2]int
}
type intersectionsType struct {
	list       []intersectionType
	coordIndex map[coord]int
}

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCod2019/day3/input.txt"

func intAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
func distance(a coord, b coord) int {
	return intAbs(a.x, b.x) + intAbs(a.y, b.y)
}
func sortStartEnd(a, b *int) {
	if *a > *b {
		t := *a
		*a = *b
		*b = t
	}
}
func intersect(a, b segmentType) (coord, error) {
	sortStartEnd(&a.start.x, &a.end.x)
	sortStartEnd(&b.start.x, &b.end.x)
	sortStartEnd(&a.start.y, &a.end.y)
	sortStartEnd(&b.start.y, &b.end.y)
	origin := coord{x: 0, y: 0}
	if a.start == origin || b.start == origin {
		return coord{x: 0, y: 0}, fmt.Errorf("No intersection")
	}

	if (a.direction == "U" || a.direction == "D") &&
		(b.direction == "L" || b.direction == "R") {
		if (b.start.x <= a.start.x && b.end.x >= a.start.x) &&
			(a.start.y <= b.start.y && a.end.y >= b.start.y) {
			return coord{a.start.x, b.start.y}, nil
		}
	}

	if (a.direction == "L" || a.direction == "R") &&
		(b.direction == "U" || b.direction == "D") {
		if (a.start.x <= b.start.x && a.end.x >= b.start.x) &&
			(b.start.y <= a.start.y && b.end.y >= a.start.y) {
			return coord{b.start.x, a.start.y}, nil
		}
	}
	return coord{x: 0, y: 0}, fmt.Errorf("No intersection")
}

func createWire(s string, index int) wireType {
	instructions := strings.Split(s, ",")

	wire := wireType{}
	wire.index = index
	wire.segments = make([]segmentType, len(instructions))
	position := coord{x: 0, y: 0}
	for i, instr := range instructions {
		segm := &wire.segments[i]
		segm.direction = string(instr[0])
		segm.length, _ = strconv.Atoi(instr[1:])

		segm.start = position
		switch {
		case segm.direction == "U":
			{
				segm.xd = 0
				segm.yd = 1
				segm.end = coord{x: position.x, y: position.y + segm.length}
			}
		case segm.direction == "R":
			{
				segm.xd = 1
				segm.yd = 0
				segm.end = coord{x: position.x + segm.length, y: position.y}
			}
		case segm.direction == "D":
			{
				segm.xd = 0
				segm.yd = -1
				segm.end = coord{x: position.x, y: position.y - segm.length}
			}
		case segm.direction == "L":
			{
				segm.xd = -1
				segm.yd = 0
				segm.end = coord{x: position.x - segm.length, y: position.y}
			}
		}
		position = segm.end
	}
	return wire
}

func findIntersections(wire0 wireType, wire1 wireType) intersectionsType {
	origin := coord{x: 0, y: 0}
	var intersections intersectionsType
	intersections.coordIndex = make(map[coord]int)
	for _, segm0 := range wire0.segments {
		// for each segment in the first wire
		// check if there is an intersection with the second wire
		for _, segm1 := range wire1.segments {
			if cross, err := intersect(segm0, segm1); err == nil {
				if cross != origin {
					intersections.list = append(intersections.list, intersectionType{pos: cross})
					intersections.coordIndex[cross] = len(intersections.list) - 1

				}
			}
		}
	}
	return intersections
}

func findClosestIntersection(intersections intersectionsType) (coord, int) {
	var bestPosition coord
	var bestDistance int = -1
	for i := 0; i < len(intersections.list); i++ {
		d := distance(coord{x: 0, y: 0}, intersections.list[i].pos)
		if bestDistance == -1 || d < bestDistance {
			bestPosition = intersections.list[i].pos
			bestDistance = d
		}
	}
	return bestPosition, bestDistance
}

func traverseWire(wire wireType, intersections *intersectionsType) {
	posList := make(map[coord]int)
	position := coord{x: 0, y: 0}
	steps := 0
	for _, segm := range wire.segments {
		for i := 0; i < segm.length; i++ {
			steps++
			position = coord{x: position.x + segm.xd, y: position.y + segm.yd}
			// remember lowest number of steps to get here
			if _, ok := posList[position]; !ok {
				posList[position] = steps
				if index, ok := intersections.coordIndex[position]; ok {
					intersections.list[index].wireSteps[wire.index] = posList[position]
				}
			}
		}
	}
}

func getLowestSteps(intersections intersectionsType) (coord, int) {
	var steps int = -1
	var pos coord
	for i := 0; i < len(intersections.list); i++ {
		wireSteps := intersections.list[i].wireSteps
		if steps == -1 || wireSteps[0]+wireSteps[1] < steps {
			steps = wireSteps[0] + wireSteps[1]
			pos = intersections.list[i].pos
		}
	}
	return pos, steps
}

func solvePart1(s []string) int {
	wire0 := createWire(s[0], 0)
	wire1 := createWire(s[1], 1)
	intersections := findIntersections(wire0, wire1)
	bestPosition, bestDistance := findClosestIntersection(intersections)
	fmt.Printf("Part 1 - Closest intersection: (%v, %v) with distance: %v\n", bestPosition.x, bestPosition.y, bestDistance)
	return bestDistance
}

func solvePart2(s []string) int {
	wire0 := createWire(s[0], 0)
	wire1 := createWire(s[1], 1)
	intersections := findIntersections(wire0, wire1)
	traverseWire(wire0, &intersections)
	traverseWire(wire1, &intersections)

	var bestPosition coord
	var bestSteps int
	bestPosition, bestSteps = getLowestSteps(intersections)
	fmt.Printf("Part 2 - Lowest step count: (%v, %v) with steps: %v\n", bestPosition.x, bestPosition.y, bestSteps)
	return bestSteps
}

// Solve solves the Day 3 assignments
func Solve() {
	fmt.Println("\n*** DAY 3 ***")
	data := loader.ReadStringsFromFile(inputFile, false)
	solvePart1(data)
	solvePart2(data)
}
