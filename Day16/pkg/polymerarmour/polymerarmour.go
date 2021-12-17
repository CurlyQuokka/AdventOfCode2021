package polymerarmour

import (
	"fmt"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day16/pkg/utils"
)

const (
	pairsSeparator = " -> "
)

type PolymerArmour struct {
	data      []string
	poly      string
	pairs     map[string]string
	generated string
}

func getPairs(data []string) map[string]string {
	m := make(map[string]string)
	for _, d := range data {
		ds := strings.Split(d, pairsSeparator)
		m[ds[0]] = ds[1]
	}
	return m
}

func NewPolymerArmour(path string) *PolymerArmour {
	pa := &PolymerArmour{
		data:      utils.LoadData(path),
		generated: "",
	}
	pa.poly = pa.data[0]
	pa.pairs = getPairs(pa.data[2:])
	return pa
}

func divPoly(p *string) *map[string]int {
	m := make(map[string]int)
	for i := 1; i < len(*p); i++ {
		key := (*p)[i-1 : i+1]
		m[key]++
	}
	return &m
}

func (pa *PolymerArmour) GenerateArmour(steps int) {
	val := divPoly(&pa.poly)
	for i := 0; i < steps; i++ {
		val = pa.processPoly(val)
	}

	count := getCountMap(val, pa.poly)
	min, max := findMinMax(&count)

	fmt.Printf("Substracted: %d\n", max-min)
}

func (pa *PolymerArmour) processPoly(polyCount *map[string]int) *map[string]int {
	polyCountCopy := copyPolyCount(polyCount)
	for key, value := range *polyCount {
		newVal := pa.pairs[key]
		newKey1 := string(key[0]) + newVal
		newKey2 := newVal + string(key[1])

		(*polyCountCopy)[key] -= value
		(*polyCountCopy)[newKey1] += value
		(*polyCountCopy)[newKey2] += value
	}
	return polyCountCopy
}

func copyPolyCount(polyCount *map[string]int) *map[string]int {
	m := make(map[string]int)
	for key, value := range *polyCount {
		m[key] = value
	}
	return &m
}

func getCountMap(data *map[string]int, poly string) map[string]int {
	result := make(map[string]int)
	for key, value := range *data {
		if _, exists := result[string(key[0])]; exists {
			result[string(key[0])] += value
		} else {
			result[string(key[0])] = value
		}

		if _, exists := result[string(key[1])]; exists {
			result[string(key[1])] += value
		} else {
			result[string(key[1])] = value
		}
	}

	for key, value := range result {
		result[key] = value / 2
	}

	result[string(poly[0])]++
	result[string(poly[len(poly)-1])]++

	return result
}

func findMinMax(m *map[string]int) (int, int) {
	max := 0
	min := 0

	i := 0
	for _, value := range *m {
		if i == 0 {
			max = value
			min = value
			i++
		} else {
			if value > max {
				max = value
			}
			if value < min {
				min = value
			}
		}
	}

	return min, max
}
