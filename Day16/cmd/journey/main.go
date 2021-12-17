package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day16/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeBitsDecoder(os.Args[1])
	bd := s.GetBitsDecoder()
	bd.Decode()
	bd.SumVersions()
	bd.EvaluateBits()
}
