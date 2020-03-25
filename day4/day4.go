package day4

import (
	"fmt"
	"strconv"
)

const rangeStart = 254032
const rangeEnd = 789860

func hasAPairPart1(a int) bool {
	s := strconv.Itoa(a)
	var p byte
	for i := 0; i < len(s); i++ {
		if s[i] == p {
			return true
		}
		p = s[i]
	}
	return false
}

func hasAPairPart2(a int) bool {
	s := strconv.Itoa(a)
	count := 1
	var p byte
	for i := 0; i < len(s); i++ {
		if s[i] == p {
			count++
		} else {
			if count == 2 {
				return true
			}
			count = 1
		}
		p = s[i]
	}
	return (count == 2)
}

func forceIncrease(a int) int {
	s := strconv.Itoa(a)
	var result string = ""
	var p byte = 0
	for i := 0; i < len(s); i++ {
		if s[i] < p {
			result = result + string(p)
		} else {
			result = result + string(s[i])
			p = s[i]
		}
	}

	r, _ := strconv.Atoi(result)
	return r
}

func generatePasswords(start, end int, isPart2 bool) int {
	result := 0
	b := forceIncrease(start)
	fmt.Printf("From: %6v, To:%6v\n", start, end)

	for {
		if b > end {
			break
		}

		if (!isPart2 && hasAPairPart1(b)) ||
			(isPart2 && hasAPairPart2(b)) {
			result++
		}
		b = forceIncrease(b + 1)

	}
	return result
}

//SolvePart1 solves part 1 of day 4
func SolvePart1() int {
	result := generatePasswords(rangeStart, rangeEnd, false)
	fmt.Println("Answer (Part 1): ", result)
	return result
}

//SolvePart2 solves part 2 of day 4
func SolvePart2() int {
	result := generatePasswords(rangeStart, rangeEnd, true)
	fmt.Println("Answer (Part 2): ", result)
	return result
}

// Solve runs day 4 assignment
func Solve() {
	fmt.Printf("\n*** DAY 4 : Part 1 ***\n")

	SolvePart1()

	fmt.Printf("\n*** DAY 4 : Part 2 ***\n")
	SolvePart2()
}
