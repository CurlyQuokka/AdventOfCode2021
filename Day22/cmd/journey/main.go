package main

import (
	"os"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/submarine"
)

func main() {
	s := submarine.NewSubmarine()
	s.InitializeReactorRebooter(os.Args[1])
	rr := s.GetReactorRebooter()
	rr.RebootReactor()
	rr.SetRangeSecure(false)
	rr.RebootReactor()
}
