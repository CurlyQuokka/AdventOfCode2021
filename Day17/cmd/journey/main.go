package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day17/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeProbeLauncher(os.Args[1])
	pl := s.GetProbeLauncher()
	pl.CalculatePossibleVelocities()
}
