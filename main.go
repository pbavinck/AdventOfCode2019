package main

import (
	"io/ioutil"
	"log"

	"github.com/pbavinck/AofCod2019/day7"
)

func main() {
	log.SetOutput(ioutil.Discard)
	// log.SetOutput(os.Stderr)

	// day1.Solve()
	// day2.Solve()
	// day2a.Solve()
	// day3.Solve()
	// day4.Solve()
	// day5.Solve()
	// day6.Solve()
	day7.Solve()

}
