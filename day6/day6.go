package day6

import (
	"fmt"
	"log"

	"github.com/pbavinck/AofCode2019/loader"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day6/input.txt"

type objectInfo struct {
	name         string
	orbits       string
	jumpsToSanta int
}

type aUniverse map[string]*objectInfo

// method of the aUniverse map make(map[...]... returns a pointer to the map header: *hmap)
// therefore we can simply use aUniverse as receiver instead of *aUniverse, like with slice methods
//https://stackoverflow.com/questions/49176090/map-as-a-method-receiver
func (u aUniverse) load(data []string) {
	for i := 0; i < len(data); i++ {
		name := data[i][4:]
		object := objectInfo{
			name:         name,
			orbits:       data[i][0:3],
			jumpsToSanta: -1,
		}
		u[name] = &object
	}
}

func (u aUniverse) calculateJumpsToSanta() {
	for item := range u {
		object := u[item]
		jumpsToSanta := -1
		for {
			if object.jumpsToSanta != -1 { // already processed
				break
			}

			if object.name == "SAN" {
				object.jumpsToSanta = 0
				jumpsToSanta = 0
			} else {
				if jumpsToSanta >= 0 {
					jumpsToSanta++
					object.jumpsToSanta = jumpsToSanta
					log.Printf("%3v orbits %3v %+6v jumps to Santa", object.name, object.orbits, object.jumpsToSanta)
				}
			}

			if object.orbits == "COM" {
				break
			}
			object = u[object.orbits]
		}
	}
}

func (u aUniverse) countOrbits() int {
	count := 0
	for item := range u {
		o := u[item]
		for {
			count++
			if o.orbits == "COM" {
				break
			}
			o = u[o.orbits]
		}
	}
	return count
}

func (u aUniverse) countJumpsToSanta() int {
	u.calculateJumpsToSanta()
	you := u["YOU"]
	o := u[you.orbits]
	jumps := 1 // for consistency with Santa to Planet jump being counted as 1
	for {
		if o.jumpsToSanta != -1 {
			log.Printf("Jump %v, arrived at %v and found %v jumps to Santa", jumps, o.name, o.jumpsToSanta)
			return o.jumpsToSanta + jumps
		}
		jumps++
		o = u[o.orbits]
	}
}

//SolvePart1 solves part 1 of day 6
func SolvePart1(u aUniverse) {
	fmt.Println("Part 2 - Answer:", u.countOrbits())
}

//SolvePart2 solves part 2 of day 6
func SolvePart2(u aUniverse) {
	// we counted both the "YOU to planet" jump and the "Santa to planet" jump, so subtract 2
	fmt.Println("Part 2 - Answer:", u.countJumpsToSanta()-2)
}

// Solve runs day 6 assignment
func Solve() {
	fmt.Printf("\n*** DAY 6 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	universe := make(aUniverse)
	universe.load(data)

	SolvePart1(universe)

	data = loader.ReadStringsFromFile(inputFile, false)
	SolvePart2(universe)
}
