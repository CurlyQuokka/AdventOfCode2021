package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day13/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeManualDecoder(os.Args[1])
	md := s.GetManualDecoder()

	// md.PrintData()

	instr := md.GetInstruction(0)
	md.FoldPage(instr)

	md.CountDots()

	// md.PrintData()
	md.FoldRemaining()
	md.CountDots()

	md.PrintData()
}
