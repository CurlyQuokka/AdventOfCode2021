package main

import (
	"fmt"
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day18/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeSnailfishMath(os.Args[1])
	sm := s.GetSnailfishMath()
	magnitude := sm.DoMath()
	fmt.Printf("Magnitude: %d\n", magnitude)

	allData := sm.GetData()
	magnitudes := []int{}
	for i := 0; i < len(allData); i++ {
		for j := 0; j < len(allData); j++ {
			if i != j {
				newData := []string{}
				newData = append(newData, allData[i])
				newData = append(newData, allData[j])
				sm.Reset()
				sm.ReplaceData(newData)
				result := sm.DoMath()
				magnitudes = append(magnitudes, result)
			}
		}
	}
	max := 0
	for _, v := range magnitudes {
		if v > max {
			max = v
		}
	}
	fmt.Printf("Max of 2: %d\n", max)
}
