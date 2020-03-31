package day8

import (
	"fmt"
	"log"

	"github.com/pbavinck/AofCode2019/loader"
)

const inputFile = "/Users/pbavinck/Automation/golang/src/github.com/pbavinck/AofCode2019/day8/input.txt"

type anImage struct {
	width      int
	height     int
	layerCount int
	pixels     [][][]int
}

func newImage(data string, width int, height int) *anImage {
	img := anImage{
		width:      width,
		height:     height,
		layerCount: len(data) / (width * height),
		pixels:     make([][][]int, len(data)/(width*height)),
	}
	for l := 0; l < img.layerCount; l++ {
		img.pixels[l] = make([][]int, img.height)
		for r := 0; r < img.height; r++ {
			img.pixels[l][r] = make([]int, img.width)
		}
	}

	for l := 0; l < img.layerCount; l++ {
		for r := 0; r < img.height; r++ {
			for c := 0; c < img.width; c++ {
				i := l*img.width*img.height + r*img.width + c
				img.pixels[l][r][c] = int(data[i] - '0')
			}
		}
	}

	log.Printf("%v layers loaded\n", img.layerCount)
	return &img
}

func (img *anImage) fewest0Digits() (fewestZeros, bestLayer int) {
	fewestZeros = img.width*img.height + 1
	bestLayer = -1
	for l := 0; l < img.layerCount; l++ {
		zeros := 0
		for r := 0; r < img.height; r++ {
			for c := 0; c < img.width; c++ {
				if img.pixels[l][r][c] == 0 {
					zeros++
				}
			}
		}
		if zeros < fewestZeros {
			fewestZeros = zeros
			bestLayer = l
		}
	}
	log.Printf("Layer %v has the fewest zeros: %v\n", bestLayer, fewestZeros)
	return
}

func (img *anImage) multiply1x2(l int) int {
	ones, twos := 0, 0
	for r := 0; r < img.height; r++ {
		for c := 0; c < img.width; c++ {
			switch {
			case img.pixels[l][r][c] == 1:
				ones++
			case img.pixels[l][r][c] == 2:
				twos++
			}
		}
	}
	log.Printf("Layer %v has %v 1's and %v 2's -> 1's X 2's = %v\n", l, ones, twos, ones*twos)
	return ones * twos
}

func (img *anImage) pixelColor(r, c int) int {
	for l := 0; l < img.layerCount; l++ {
		switch {
		case img.pixels[l][r][c] == 0:
			return 0
		case img.pixels[l][r][c] == 1:
			return 1
		}
	}
	return 2
}

//SolvePart1 solves part 1 of day 8
func SolvePart1(data []string) {
	img := newImage(data[0], 25, 6)
	_, bestLayer := img.fewest0Digits()
	onesTimesZeros := img.multiply1x2(bestLayer)
	fmt.Println("Part 1 - Answer:", onesTimesZeros)
}

func (img *anImage) decode() {
	for r := 0; r < img.height; r++ {
		for c := 0; c < img.width; c++ {
			switch img.pixelColor(r, c) {
			case 0:
				fmt.Printf(" ")
			case 1:
				fmt.Printf("\u2588")
			}
		}
		fmt.Printf("\n")
	}
}

//SolvePart2 solves part 2 of day 8
func SolvePart2(data []string) {
	img := newImage(data[0], 25, 6)
	fmt.Println("Part 2 - Answer:")
	img.decode()
}

// Solve runs day 8 assignment
func Solve() {
	fmt.Printf("\n*** DAY 8 ***\n")
	data := loader.ReadStringsFromFile(inputFile, false)

	SolvePart1(data)

	SolvePart2(data)
}
