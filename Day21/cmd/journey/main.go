package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day21/pkg/diracdicegame"
	"github.com/CurlyQuokka/AdventOfCode2021/Day21/pkg/submarine"
)

const (
	fields = 10
	score1 = 1000
	score2 = 21
	sides  = 100
)

func main() {
	s := submarine.NewSubmarine()
	die := diracdicegame.NewDeterministicDie(sides)
	s.InitializeDiracDicegame(fields, die, os.Args[1])
	ddg := s.GetDiracDiceGame()
	ddg.PlayFirstGame(score1)
	ddg.ResetPlayers(os.Args[1])
	ddg.PlaySecondGame(score2, fields)
}
