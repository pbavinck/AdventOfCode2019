package main

import (
	"io/ioutil"
	"log"

	"github.com/pbavinck/AofCode2019/day1"
	"github.com/pbavinck/AofCode2019/day10"
	"github.com/pbavinck/AofCode2019/day11"
	"github.com/pbavinck/AofCode2019/day2"
	"github.com/pbavinck/AofCode2019/day2a"
	"github.com/pbavinck/AofCode2019/day3"
	"github.com/pbavinck/AofCode2019/day4"
	"github.com/pbavinck/AofCode2019/day5"
	"github.com/pbavinck/AofCode2019/day5a"
	"github.com/pbavinck/AofCode2019/day6"
	"github.com/pbavinck/AofCode2019/day7"
	"github.com/pbavinck/AofCode2019/day8"
	"github.com/pbavinck/AofCode2019/day9"
)

func main() {
	log.SetOutput(ioutil.Discard)
	// log.SetOutput(os.Stderr)

	day1.Solve()
	day2.Solve()
	day2a.Solve()
	day3.Solve()
	day4.Solve()
	day5.Solve()
	day5a.Solve()
	day6.Solve()
	day7.Solve()
	day8.Solve()
	day9.Solve()
	day10.Solve()
	day11.Solve()

	// Compute days
	// day2a.Solve()
	// day5a.Solve()
	// day7.Solve()
	// day9.Solve()
}
