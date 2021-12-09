package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day09/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeVentsSubsystem(os.Args[1])
	vs := s.GetVentsSubsystem()
	vs.CalculateRiskLevel()
	vs.ScanBasins()
}
