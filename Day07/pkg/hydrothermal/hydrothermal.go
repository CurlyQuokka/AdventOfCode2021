package hydrothermal

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/utils"
)

const (
	lineSeparator  = " -> "
	pointSeparator = ","
)

type point struct {
	x, y int
}

type line struct {
	a, b point
}
type lines []line

type HydrothermalAvoidance struct {
	rawData          []string
	allLines         lines
	straightLines    lines
	diagonalLines    lines
	maxX, maxY       int
	board            [][]int
	obstaclesToAvoid int
}

func convertPoint(p string) int {
	xv, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(888)
	}
	return xv
}

func cretePoint(p string) point {
	v := strings.Split(p, pointSeparator)

	return point{
		x: convertPoint(v[0]),
		y: convertPoint(v[1]),
	}
}

func convertData(data []string) lines {
	var allLines lines
	for _, d := range data {
		p := strings.Split(d, lineSeparator)
		l := line{
			a: cretePoint(p[0]),
			b: cretePoint(p[1]),
		}
		allLines = append(allLines, l)
	}
	return allLines
}

func (l *lines) findMaxValues() (int, int) {
	maxX, maxY := 0, 0
	for _, item := range *l {
		utils.GetHigher(&item.a.x, &maxX)
		utils.GetHigher(&item.b.x, &maxX)
		utils.GetHigher(&item.a.y, &maxY)
		utils.GetHigher(&item.b.y, &maxY)
	}
	return maxX, maxY
}

func NewHydrothermalAvoidance(path string) *HydrothermalAvoidance {
	ha := &HydrothermalAvoidance{
		rawData:  utils.LoadData(path),
		allLines: lines{},
		maxX:     0,
		maxY:     0,
	}
	ha.allLines = convertData(ha.rawData)
	ha.maxX, ha.maxY = ha.allLines.findMaxValues()
	ha.straightLines, ha.diagonalLines = filterLines(ha.allLines)
	return ha
}

func (ha *HydrothermalAvoidance) Print() {
	fmt.Printf("%v\n", ha.allLines)
	fmt.Printf("%v\n", ha.straightLines)
	fmt.Printf("%v\n", ha.diagonalLines)
	fmt.Printf("Max x: %d, max y: %d\n", ha.maxX, ha.maxY)
	for _, row := range ha.board {
		for _, col := range row {
			fmt.Printf("%d ", col)
		}
		fmt.Println()
	}
}

func (ha *HydrothermalAvoidance) prepareBoard() {
	b := make([][]int, ha.maxY+1)
	for i := range b {
		b[i] = make([]int, ha.maxX+1)
	}
	ha.board = b
}

func filterLines(l lines) (lines, lines) {
	var straight, diagonal lines
	for _, v := range l {
		// filter straight lines
		if v.a.x == v.b.x || v.a.y == v.b.y {
			if v.a.x == v.b.x {
				if v.a.y > v.b.y {
					tmp := v.a
					v.a = v.b
					v.b = tmp
				}
				straight = append(straight, v)
			}
			if v.a.y == v.b.y {
				if v.a.x > v.b.x {
					tmp := v.a
					v.a = v.b
					v.b = tmp
				}
				straight = append(straight, v)
			}
		} else {
			// filter diagonal lines
			if v.a.x > v.b.x {
				tmp := v.a
				v.a = v.b
				v.b = tmp
			}
			diagonal = append(diagonal, v)
		}
	}
	return straight, diagonal
}

func (ha *HydrothermalAvoidance) MarkLines(countDiagonals bool) {
	ha.prepareBoard()
	for _, l := range ha.straightLines {
		for i := l.a.y; i <= l.b.y; i++ {
			for j := l.a.x; j <= l.b.x; j++ {
				ha.board[i][j]++
			}
		}
	}
	if countDiagonals {
		for _, l := range ha.diagonalLines {
			if l.a.y <= l.b.y {
				for i, j := l.a.y, l.a.x; i <= l.b.y; i, j = i+1, j+1 {
					ha.board[i][j]++
				}
			} else {
				for i, j := l.a.y, l.a.x; i >= l.b.y; i, j = i-1, j+1 {
					ha.board[i][j]++
				}
			}

		}
	}
	ha.countObstacles()
}

func (ha *HydrothermalAvoidance) countObstacles() {
	ha.obstaclesToAvoid = 0
	for _, row := range ha.board {
		for _, col := range row {
			if col > 1 {
				ha.obstaclesToAvoid++
			}
		}
	}
}

func (ha *HydrothermalAvoidance) PrintObstaclesCount() {
	fmt.Printf("Obstacle count: %d\n", ha.obstaclesToAvoid)
}
