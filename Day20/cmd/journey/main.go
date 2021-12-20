package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2021/Day20/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeTrenchMapper(os.Args[1])
	tm := s.GetTrenchMapper()
	steps, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(252)
	}
	tm.MapCave(steps)
}
