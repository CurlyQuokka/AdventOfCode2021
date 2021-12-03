package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day03/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.LoadPowerReport(os.Args[1])
	s.PrintPowerUsage()
	s.PrintLifeSupportRating()
}
