package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day04/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeBingoSubsystem(os.Args[1])
	bs := s.GetBingoSubsystem()
	bs.PlayGame()
	bs.PrintScore()
	bs.WreckThisCasino()
	bs.FindLoosingBoard()
	bs.PrintScore()
}
