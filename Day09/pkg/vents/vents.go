package vents

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2021/Day09/pkg/utils"
)

type VentsAnalyzer struct {
	data  []string
	vents [][]int
}

func NewVentsAnalyzer(path string) *VentsAnalyzer {
	va := &VentsAnalyzer{
		data: utils.LoadData(path),
	}
	va.prepareMap()
	return va
}

func (va *VentsAnalyzer) prepareMap() {
	for _, line := range va.data {
		var row []int
		for _, v := range line {
			iv, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(99)
			}
			row = append(row, iv)
		}
		va.vents = append(va.vents, row)
	}
}

func (va *VentsAnalyzer) isBottomHigher(row, col int) bool {
	if row == (len(va.vents)-1) || va.vents[row+1][col] > va.vents[row][col] {
		return true
	}
	return false
}

func (va *VentsAnalyzer) isUpperHigher(row, col int) bool {
	if row == 0 || va.vents[row-1][col] > va.vents[row][col] {
		return true
	}
	return false
}

func (va *VentsAnalyzer) isLeftHigher(row, col int) bool {
	if col == 0 || va.vents[row][col-1] > va.vents[row][col] {
		return true
	}
	return false
}

func (va *VentsAnalyzer) isRighttHigher(row, col int) bool {
	if col == (len(va.vents[0])-1) || va.vents[row][col+1] > va.vents[row][col] {
		return true
	}
	return false
}

func (va *VentsAnalyzer) checkVent(row, col int) (bool, int) {
	if va.isBottomHigher(row, col) && va.isLeftHigher(row, col) && va.isRighttHigher(row, col) && va.isUpperHigher(row, col) {
		return true, va.vents[row][col] + 1
	}
	return false, 0
}

func (va *VentsAnalyzer) CalculateRiskLevel() {
	rowSize := len(va.vents)
	colSize := len(va.vents[0])
	riskLevel := 0
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			_, v := va.checkVent(row, col)
			riskLevel += v
		}
	}
	fmt.Printf("Risk level: %d\n", riskLevel)
}

type pair struct {
	row, col int
}

func prepareBasinMap(row, col int) [][]int {
	b := make([][]int, row)
	for i := range b {
		b[i] = make([]int, col)
	}
	return b
}

func appendPoint(points *[]pair, basinMap *[][]int, row, col int) {
	if (*basinMap)[row][col] == 0 {
		p := pair{
			row: row,
			col: col,
		}

		*points = append(*points, p)
	}
}

func (va *VentsAnalyzer) exploreBasin(basinMap *[][]int, points []pair) {
	rowSize := len(va.vents)
	colSize := len(va.vents[0])
	newPoints := []pair{}
	for _, p := range points {
		(*basinMap)[p.row][p.col] = 1
		if p.row > 0 {
			if va.vents[p.row-1][p.col] < 9 {
				appendPoint(&newPoints, basinMap, p.row-1, p.col)
			}
		}
		if p.row < (rowSize - 1) {
			if va.vents[p.row+1][p.col] < 9 {
				appendPoint(&newPoints, basinMap, p.row+1, p.col)
			}
		}
		if p.col > 0 {
			if va.vents[p.row][p.col-1] < 9 {
				appendPoint(&newPoints, basinMap, p.row, p.col-1)
			}
		}
		if p.col < (colSize - 1) {
			if va.vents[p.row][p.col+1] < 9 {
				appendPoint(&newPoints, basinMap, p.row, p.col+1)
			}
		}
		va.exploreBasin(basinMap, newPoints)
	}
}

func calculateBasinSize(basinMap *[][]int) int {
	sum := 0
	for _, row := range *basinMap {
		for _, col := range row {
			sum += col
		}
	}
	return sum
}

func (va *VentsAnalyzer) ScanBasins() {
	rowSize := len(va.vents)
	colSize := len(va.vents[0])
	basins := []int{}
	for row := 0; row < rowSize; row++ {
		for col := 0; col < colSize; col++ {
			if isMin, _ := va.checkVent(row, col); isMin {
				tmp := prepareBasinMap(rowSize, colSize)
				p := pair{
					row: row,
					col: col,
				}
				points := []pair{}
				points = append(points, p)
				va.exploreBasin(&tmp, points)
				basins = append(basins, calculateBasinSize(&tmp))
			}
		}
	}
	sort.Ints(basins)
	numOfBasins := len(basins)
	basinsVal := basins[numOfBasins-1] * basins[numOfBasins-2] * basins[numOfBasins-3]
	fmt.Printf("3 largest basins multiplied: %d\n", basinsVal)
}
