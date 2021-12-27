package crabalignment

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	inputSeparator = ","
)

type CrabAlignment struct {
	crabPosition    []int
	fuelConsumption int
}

func InitializeCrabAlignment(path string) *CrabAlignment {
	ca := &CrabAlignment{}
	data := utils.LoadData(path)
	splitData := strings.Split(data[0], inputSeparator)

	for _, s := range splitData {
		v, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(41)
		}
		ca.crabPosition = append(ca.crabPosition, v)
	}

	sort.Ints(ca.crabPosition)

	return ca
}

func (ca *CrabAlignment) CalculateFuelConsumption(notConstantConsumption bool) {
	var fc int

	if notConstantConsumption {
		avg := ca.calculateMean()
		for _, crabPos := range ca.crabPosition {
			v := int(math.Round(math.Abs(float64(avg) - float64(crabPos))))
			fc += int(math.Round(float64(v * (v + 1) / 2)))
		}
	} else {
		median := ca.calculateMedian()
		for _, crabPos := range ca.crabPosition {
			fc += int(math.Round(math.Abs(float64(median) - float64(crabPos))))
		}
	}

	ca.fuelConsumption = int(fc)
}

func (ca *CrabAlignment) calculateMedian() int {
	length := len(ca.crabPosition)
	pos := length / 2
	if length%2 == 0 {
		return int(math.Round(float64(ca.crabPosition[pos])+float64(ca.crabPosition[pos-1])) / 2.0)
	}

	return ca.crabPosition[pos]
}

func (ca *CrabAlignment) calculateMean() int {
	var sum int
	for _, crabPos := range ca.crabPosition {
		sum += crabPos
	}

	avg := float64(sum) / float64(len(ca.crabPosition))

	return int(math.Floor(avg))
}

func (ca *CrabAlignment) FuelLevelGauge() {
	fmt.Printf("Fuel consumption: %d\n", ca.fuelConsumption)
}
