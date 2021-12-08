package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day08/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeLCDDisplay(os.Args[1])
	lcd := s.GetLCDDisplay()
	lcd.Count1478()
	lcd.Decode()
}
