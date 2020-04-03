package day12

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"

	"github.com/pbavinck/AofCode2019/loader"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day12/input.txt"

type coord struct {
	x int
	y int
	z int
}

type vector struct {
	x int
	y int
	z int
}
type aMoon struct {
	id       int
	position coord
	velocity vector
}

type jupiterMoons struct {
	ticks int // number of passed time ticks
	moons []aMoon
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func newMoons(data []string) *jupiterMoons {
	j := jupiterMoons{}
	j.moons = make([]aMoon, 4)
	r, _ := regexp.Compile(`<x=(\-?\d*)\s*,\s*y=(\-?\d*)\s*,\s*z=(\-?\d*)>`)
	for i, s := range data {
		matches := r.FindStringSubmatch(s)
		j.moons[i].position = coord{}
		j.moons[i].position.x, _ = strconv.Atoi(matches[1])
		j.moons[i].position.y, _ = strconv.Atoi(matches[2])
		j.moons[i].position.z, _ = strconv.Atoi(matches[3])
		j.moons[i].id = i

		// fmt.Printf("%+3v: %+v\n", j.moons[i].id, j.moons[i].position)
	}
	return &j
}

func (m *aMoon) energy() int {
	p := m.position
	v := m.velocity
	pot := intAbs(p.x) + intAbs(p.y) + intAbs(p.z)
	kin := intAbs(v.x) + intAbs(v.y) + intAbs(v.z)
	return pot * kin
}

func (m *aMoon) gravityFrom(o aMoon) {
	// calculates the gravity m feels from m1
	if m.position.x > o.position.x {
		m.velocity.x--
	} else if m.position.x < o.position.x {
		m.velocity.x++
	}
	if m.position.y > o.position.y {
		m.velocity.y--
	} else if m.position.y < o.position.y {
		m.velocity.y++
	}
	if m.position.z > o.position.z {
		m.velocity.z--
	} else if m.position.z < o.position.z {
		m.velocity.z++
	}
}

func (j *jupiterMoons) print() {
	for i := 0; i < len(j.moons); i++ {
		fmt.Printf("%v - pos: %+v, vel: %+v\n", i, j.moons[i].position, j.moons[i].velocity)
	}
}

func (j *jupiterMoons) doGravity() {
	for i1 := range j.moons {
		for i2 := range j.moons {
			if j.moons[i1].id != j.moons[i2].id {
				j.moons[i1].gravityFrom(j.moons[i2])
			}
		}
	}
}

func (j *jupiterMoons) addVelocity() {
	for i := range j.moons {
		j.moons[i].position.x += j.moons[i].velocity.x
		j.moons[i].position.y += j.moons[i].velocity.y
		j.moons[i].position.z += j.moons[i].velocity.z
	}
}

func (j *jupiterMoons) cycles(count int) {
	for i := 0; i < count; i++ {
		j.doGravity()
		j.addVelocity()
	}
}

func (j *jupiterMoons) energy() int {
	e := 0
	for _, m := range j.moons {
		e += m.energy()
	}
	return e
}

func (j *jupiterMoons) startCalculation(cycles int) int {
	j.cycles(cycles)
	e := j.energy()
	return e
}

func (j *jupiterMoons) findStep(dim string) (step int64) {
	// find steps until patterm repeats for only 1 dimension at a time
	pos1 := make([]int, len(j.moons))
	vel1 := make([]int, len(j.moons))
	pos2 := make([]int, len(j.moons))
	vel2 := make([]int, len(j.moons))
	for i, m := range j.moons {
		switch dim {
		case "x":
			pos1[i] = m.position.x
			vel1[i] = m.velocity.x
			pos2[i] = m.position.x
			vel2[i] = m.velocity.x
		case "y":
			pos1[i] = m.position.y
			vel1[i] = m.velocity.y
			pos2[i] = m.position.y
			vel2[i] = m.velocity.y
		case "z":
			pos1[i] = m.position.z
			vel1[i] = m.velocity.z
			pos2[i] = m.position.z
			vel2[i] = m.velocity.z
		}
	}

	step = 0
	for {
		step++
		// calculate velocity
		for i1 := 0; i1 < len(pos1); i1++ {
			for i2 := 0; i2 < len(pos1); i2++ {
				if pos2[i1] > pos2[i2] {
					vel2[i1]--
					vel2[i2]++
				} else if pos2[i1] > pos2[i2] {
					vel2[i1]++
					vel2[i2]--
				}
			}
		}
		// calculate new positiom
		for i1 := 0; i1 < len(pos1); i1++ {
			pos2[i1] += vel2[i1]
		}
		if reflect.DeepEqual(pos1, pos2) && reflect.DeepEqual(vel1, vel2) {
			break
		}
	}
	return
}

// GCD greatest common divisor via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple via GCD
func LCM(a, b int64, integers ...int64) int64 {
	var result, i int64
	result = a * b / GCD(a, b)

	for i = 0; i < int64(len(integers)); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

//SolvePart1 solves part 1 of day 12
func SolvePart1(data []string) {
	j := newMoons(data)
	j.cycles(1000)
	fmt.Println("Part 1 - Answer:", j.energy())
}

//SolvePart2 solves part 2 of day 12
func SolvePart2(data []string) {
	j := newMoons(data)
	xSteps := j.findStep("x")
	ySteps := j.findStep("y")
	zSteps := j.findStep("z")

	fmt.Println("Part 2 - Answer:", LCM(xSteps, ySteps, zSteps))
}

// Solve runs day 12 assignment
func Solve() {
	fmt.Printf("\n*** DAY 12 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)

}
