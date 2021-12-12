package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day12/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializePathFInderSubsystem(os.Args[1])
	pf := s.GetPathFinderSubsystem()
	pf.FindPaths(false)
	pf.GetNumberOfPaths()
	pf.Reset()
	pf.FindPaths(true)
	pf.GetNumberOfPaths()
}
