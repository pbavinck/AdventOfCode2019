package loader

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/pbavinck/lg"
)

// ReadStringsFromFile reads the file and returns a []string
func ReadStringsFromFile(filepath string, showLoaded bool) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not read file %s", filepath)
	}
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if showLoaded {
		lg.Debug("%v line(s) in input\n", len(input))
		lg.Debug("Data loaded:")
		for i := range input {
			lg.Debug("%v %v", i, input[i])
		}
	}

	lg.Info("%v line(s) in input\n", len(input))
	return input
}

// ReadIntsFromFile reads the file and returns a []int
func ReadIntsFromFile(filepath string, showLoaded bool) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not read file %s", filepath)
	}
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if showLoaded {
		lg.Debug("Data loaded:")
		for i := range input {
			lg.Debug("%v %v", i, input[i])
		}
	}

	lg.Info("%v line(s) in input\n", len(input))
	return input
}
