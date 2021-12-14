package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day14/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializePolymerArmour(os.Args[1])
	pa := s.GetPolymerArmour()
	pa.GenerateArmour(10)
	pa.GenerateArmour(40)
}
