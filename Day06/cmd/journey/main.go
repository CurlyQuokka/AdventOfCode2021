package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2021/Day06/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeFishObservatory(os.Args[1])

	fo := s.GetFishObservatory()
	fo.ObserveInitialPopulation()

	v, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fo.SimulatePopulation(v)
	fo.HowMuchIsTheFish()
}
