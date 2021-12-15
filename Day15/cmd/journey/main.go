package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day15/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeChitonsAvoider(os.Args[1])
	ca := s.GetChitonsAvoider()
	ca.CalculateRisk(false)
	ca.CalculateRisk(true)
}
