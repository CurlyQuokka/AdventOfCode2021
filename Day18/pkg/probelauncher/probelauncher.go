package probelauncher

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day18/pkg/utils"
)

const (
	inputTag        = "target area: "
	xySeparator     = ", "
	minMaxSeparator = ".."

	yMutiplier = 1
)

type ProbeLauncher struct {
	target *targetArea
	paths  []*[]*point
}

func NewProbeLauncher(path string) *ProbeLauncher {
	pl := &ProbeLauncher{}
	data := utils.LoadData(path)
	pl.target = NewTargetArea(data[0])
	return pl
}

func (pl *ProbeLauncher) CalculatePossibleVelocities() {
	for x := 0; x <= pl.target.maxX; x++ {
		// have no idea how big should be the y to be checked, but this surprisingly does the job, so...
		for y := pl.target.minY; y < int(math.Abs(float64(pl.target.minY)))*yMutiplier; y++ {
			v := &velocity{x: x, y: y}
			pl.Launch(v)
		}
	}

	maxY := pl.target.minY
	for _, path := range pl.paths {
		for _, p := range *path {
			if p.y > maxY {
				maxY = p.y
			}
		}
	}

	fmt.Printf("Max Y: %d\n", maxY)
	fmt.Printf("Number of velocities: %d\n", len(pl.paths))
}

func (pl *ProbeLauncher) Launch(v *velocity) {
	x := 0
	y := 0
	step := 0

	inTarget := false
	path := []*point{}
	for {
		x += v.x
		y += v.y

		path = append(path, &point{x, y})

		inTarget = isInTarget(x, y, pl.target)

		if inTarget {
			pl.paths = append(pl.paths, &path)
			break
		} else if y < pl.target.minY {
			break
		}
		step++
		if v.x > 0 {
			v.x = v.x - 1
		}
		v.y--
	}
}

func NewTargetArea(data string) *targetArea {
	ta := &targetArea{}
	data = strings.ReplaceAll(data, inputTag, "")
	splitted := strings.Split(data, xySeparator)
	ta.minX, ta.maxX = processTag(splitted[0])
	ta.minY, ta.maxY = processTag(splitted[1])
	return ta
}

type targetArea struct {
	minX, maxX, minY, maxY int
}

type velocity struct {
	x, y int
}

// duplicated for readibility
type point struct {
	x, y int
}

func isInTarget(x, y int, t *targetArea) bool {
	return x <= t.maxX && x >= t.minX && y >= t.minY && y <= t.maxY
}

func processTag(tag string) (int, int) {
	if strings.Contains(tag, "x") {
		tag = strings.ReplaceAll(tag, "x=", "")
	} else {
		tag = strings.ReplaceAll(tag, "y=", "")
	}
	splitted := strings.Split(tag, minMaxSeparator)
	min, err := strconv.Atoi(splitted[0])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(575)
	}

	max, err := strconv.Atoi(splitted[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(576)
	}

	return min, max
}
