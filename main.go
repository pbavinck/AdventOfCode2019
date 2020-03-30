package main

import (
	"io/ioutil"
	"log"

	"github.com/pbavinck/AofCode2019/day1"
	"github.com/pbavinck/AofCode2019/day2"
	"github.com/pbavinck/AofCode2019/day2a"
	"github.com/pbavinck/AofCode2019/day3"
	"github.com/pbavinck/AofCode2019/day4"
	"github.com/pbavinck/AofCode2019/day5"
	"github.com/pbavinck/AofCode2019/day6"
	"github.com/pbavinck/AofCode2019/day7"
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
	day6.Solve()
	day7.Solve()

}
