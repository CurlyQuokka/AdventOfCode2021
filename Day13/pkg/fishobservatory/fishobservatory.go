package fishobservatory

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day13/pkg/utils"
)

const (
	inputSeparator   = ","
	youngTimeToBreed = 8
	oldTimeToBreed   = 6
)

type Fishobservatory struct {
	input    []int
	observed [youngTimeToBreed + 1]int
}

func InitilizeFishobservatory(path string) *Fishobservatory {
	bs := &Fishobservatory{}
	data := utils.LoadData(path)
	splitData := strings.Split(data[0], inputSeparator)

	for _, s := range splitData {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(41)
		}
		bs.input = append(bs.input, v)
	}

	return bs
}

func (fb *Fishobservatory) ObserveInitialPopulation() {
	fb.observed = [youngTimeToBreed + 1]int{}
	for _, v := range fb.input {
		fb.observed[v]++
	}
}

func (fb *Fishobservatory) SimulatePopulation(days int) {
	for day := 0; day <= days; day++ {
		breeded := fb.observed[0]

		for i := 1; i < len(fb.observed); i++ {
			fb.observed[i-1] = fb.observed[i]
		}

		fb.observed[oldTimeToBreed] += breeded

		if day < days {
			fb.observed[youngTimeToBreed] = breeded
		} else {
			fb.observed[youngTimeToBreed] = 0 // exclude the lanterfishes that won't be born until next day
		}
	}
}

func (fb *Fishobservatory) HowMuchIsTheFish() {
	sum := 0
	for _, v := range fb.observed {
		sum += v
	}

	fmt.Printf("Nuber of fishes in the school: %d\n", sum)
}
