package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day10/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeBesRouteSubsystem(os.Args[1])
	brs := s.GetBestRouteSubsystem()
	// brs.PrintData()
	brs.ProcessData()
}
