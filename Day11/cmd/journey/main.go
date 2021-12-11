package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day11/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeFlashSimulator(os.Args[1])
	fs := s.GetFlashSImulator()
	fs.RunSimulation(100)
	fs.PrintNumOfFlashes()
	fs.Reset()
	fs.RunSimulation(0)
	fs.PrintSynchronizedStep()
}
