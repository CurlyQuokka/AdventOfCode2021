package hydrothermal

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	lineSeparator  = " -> "
	PointSeparator = ","
)

type Point struct {
	X, Y int
}

type line struct {
	a, b Point
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

func CretePoint(p string) Point {
	v := strings.Split(p, PointSeparator)

	return Point{
		X: convertPoint(v[0]),
		Y: convertPoint(v[1]),
	}
}

func convertData(data []string) lines {
	var allLines lines
	for _, d := range data {
		p := strings.Split(d, lineSeparator)
		l := line{
			a: CretePoint(p[0]),
			b: CretePoint(p[1]),
		}
		allLines = append(allLines, l)
	}
	return allLines
}

func (l *lines) findMaxValues() (int, int) {
	maxX, maxY := 0, 0
	for _, item := range *l {
		utils.GetHigher(&item.a.X, &maxX)
		utils.GetHigher(&item.b.X, &maxX)
		utils.GetHigher(&item.a.Y, &maxY)
		utils.GetHigher(&item.b.Y, &maxY)
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
		if v.a.X == v.b.X || v.a.Y == v.b.Y {
			if v.a.X == v.b.X {
				if v.a.Y > v.b.Y {
					tmp := v.a
					v.a = v.b
					v.b = tmp
				}
				straight = append(straight, v)
			}
			if v.a.Y == v.b.Y {
				if v.a.X > v.b.X {
					tmp := v.a
					v.a = v.b
					v.b = tmp
				}
				straight = append(straight, v)
			}
		} else {
			// filter diagonal lines
			if v.a.X > v.b.X {
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
		for i := l.a.Y; i <= l.b.Y; i++ {
			for j := l.a.X; j <= l.b.X; j++ {
				ha.board[i][j]++
			}
		}
	}
	if countDiagonals {
		for _, l := range ha.diagonalLines {
			if l.a.Y <= l.b.Y {
				for i, j := l.a.Y, l.a.X; i <= l.b.Y; i, j = i+1, j+1 {
					ha.board[i][j]++
				}
			} else {
				for i, j := l.a.Y, l.a.X; i >= l.b.Y; i, j = i-1, j+1 {
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
