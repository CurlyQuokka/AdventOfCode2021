package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day05/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeHydrothermalSybsystem(os.Args[1])
	hs := s.GetHydrothermalSubsystem()
	hs.MarkLines(false)
	hs.PrintObstaclesCount()
	hs.MarkLines(true)
	hs.PrintObstaclesCount()
}
