package trenchmapper

import (
	"fmt"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	windowSize = 3
	dark       = '.'
	lit        = '#'
)

type TrenchMapper struct {
	data      []string
	algorithm string
	trenchMap [][]rune
}

func NewTrenchMapper(path string) *TrenchMapper {
	tm := &TrenchMapper{
		data: utils.LoadData(path),
	}
	tm.processData()
	return tm
}

func (tm *TrenchMapper) processData() {
	tm.algorithm = tm.data[0]
	tm.data = tm.data[2:]
	for _, l := range tm.data {
		line := []rune{}
		for _, c := range l {
			line = append(line, c)
		}
		tm.trenchMap = append(tm.trenchMap, line)
	}
}

func extendMap(m *[][]rune, n int, background rune) *[][]rune {
	newMap := [][]rune{}
	for i := 0; i < len(*m)+n*2; i++ {
		line := []rune{}
		for j := 0; j < len((*m)[0])+n*2; j++ {
			line = append(line, background)
		}
		newMap = append(newMap, line)
	}

	for i, r := n, 0; i < len(newMap)-n; i, r = i+1, r+1 {
		for j, c := n, 0; j < len(newMap[0])-n; j, c = j+1, c+1 {
			newMap[i][j] = (*m)[r][c]
		}
	}

	return &newMap
}

func returnBin(v rune) string {
	if v == dark {
		return "0"
	}
	return "1"
}

func (tm *TrenchMapper) processMapPoint(m *[][]rune, row, col int) rune {
	val := ""
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			val += returnBin((*m)[row+i][col+j])
		}
	}
	index := utils.BinStringToDec(val)
	return rune(tm.algorithm[index])
}

func (tm *TrenchMapper) enchanceMap(m *[][]rune) *[][]rune {
	newMap := [][]rune{}
	for r := 1; r < len(*m)-1; r++ {
		line := []rune{}
		for c := 1; c < len((*m)[0])-1; c++ {
			line = append(line, tm.processMapPoint(m, r, c))
		}
		newMap = append(newMap, line)
	}
	return &newMap
}

func countLit(m *[][]rune) int {
	counter := 0
	for _, r := range *m {
		for _, c := range r {
			if c == lit {
				counter++
			}
		}
	}
	return counter
}

func (tm *TrenchMapper) MapCave(steps int) {
	currentBackground := dark
	currentMap := &tm.trenchMap
	invertBackground := tm.algorithm[0] == lit
	for i := 0; i < steps; i++ {
		currentMap = extendMap(currentMap, windowSize, currentBackground)
		currentMap = tm.enchanceMap(currentMap)
		if invertBackground {
			if currentBackground == lit {
				currentBackground = dark
			} else {
				currentBackground = lit
			}
		}
	}
	counter := countLit(currentMap)
	fmt.Printf("Lit elements: %d\n", counter)
}
