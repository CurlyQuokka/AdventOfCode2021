package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day02/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine(os.Args[1])
	s.ProcessJourney()
	s.PrintDestination()

	s.ResetSubmarineToFactoryDefault()
	s.ProcessJourneyWithAim()
	s.PrintDestination()
}
