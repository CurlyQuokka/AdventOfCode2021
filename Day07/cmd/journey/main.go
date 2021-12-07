package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeCrabArmy(os.Args[1])

	ca := s.GetCrabArmy()
	ca.CalculateFuelConsumption(false)
	ca.FuelLevelGauge()
	ca.CalculateFuelConsumption(true)
	ca.FuelLevelGauge()
}
