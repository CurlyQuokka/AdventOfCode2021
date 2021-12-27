package reactorrebooter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	commandSeparator = " "
	coordSeparator   = ","
	valueSeparator   = "="
	minMaxSeparator  = ".."

	onCmd  = "on"
	offCmd = "off"

	maxCoord = 50
	minCoord = -50
)

type coord struct {
	min, max int
}

func newCoord(data string) coord {
	vv := strings.Split(data, valueSeparator)
	values := strings.Split(vv[1], minMaxSeparator)
	mm := []int{}
	for _, v := range values {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(252)
		}
		mm = append(mm, val)
	}
	return coord{
		min: mm[0],
		max: mm[1],
	}
}

func getCommon(first, second *coord) (*coord, bool) {
	if second.min > first.max || second.max < first.min {
		return nil, false
	}
	if second.min <= first.min && second.max >= first.max {
		return &coord{first.min, first.max}, true
	}
	if second.min >= first.min && second.max <= first.max {
		return &coord{second.min, second.max}, true
	}
	if second.min <= first.min && second.max <= first.max {
		return &coord{first.min, second.max}, true
	}
	if second.min >= first.min && second.max >= first.max {
		return &coord{second.min, first.max}, true
	}
	return nil, false
}

type cube struct {
	x, y, z coord
	value   int
}

func newCube(data string) cube {
	coords := strings.Split(data, coordSeparator)
	crds := []coord{}
	for _, c := range coords {
		crds = append(crds, newCoord(c))
	}

	return newCubeMinMax(crds[0], crds[1], crds[2])
}

func newCubeMinMax(x, y, z coord) cube {
	return cube{
		x:     x,
		y:     y,
		z:     z,
		value: 0,
	}
}

func (c cube) copy() cube {
	return newCubeMinMax(c.x, c.y, c.z)
}

func (c cube) count() int {
	return (c.x.max - c.x.min + 1) * (c.y.max - c.y.min + 1) * (c.z.max - c.z.min + 1) * c.value
}

func getIntersection(first, second cube) *cube {
	commonX, xExists := getCommon(&first.x, &second.x)
	if !xExists {
		return nil
	}
	commonY, yExists := getCommon(&first.y, &second.y)
	if !yExists {
		return nil
	}
	commonZ, zExists := getCommon(&first.z, &second.z)
	if !zExists {
		return nil
	}

	c := newCubeMinMax(*commonX, *commonY, *commonZ)
	return &c
}

type command struct {
	cmd string
	cb  cube
}

func newCommand(data string) command {
	splitLine := strings.Split(data, commandSeparator)
	cmd := splitLine[0]
	cb := newCube(splitLine[1])
	return command{
		cmd: cmd,
		cb:  cb,
	}
}

type ReactorRebooter struct {
	commands    []command
	cubes       []cube
	rangeSecure bool
}

func NewReactorRebooter(path string) *ReactorRebooter {
	data := utils.LoadData(path)
	rr := &ReactorRebooter{}
	for _, line := range data {
		rr.commands = append(rr.commands, newCommand(line))
	}
	rr.rangeSecure = true
	return rr
}

func (rr *ReactorRebooter) Print() {
	for _, cmd := range rr.commands {
		fmt.Printf("%v\n", cmd)
	}
}

func (rr *ReactorRebooter) SetRangeSecure(val bool) {
	rr.rangeSecure = val
}

func (rr *ReactorRebooter) RebootReactor() {
	start := 0
	rr.cubes = []cube{}
	for i, cmd := range rr.commands {
		if cmd.cmd == onCmd {
			c := cmd.cb.copy()
			c.value = 1
			rr.cubes = append(rr.cubes, c)
			start = i
			break
		}
	}
	for i := start + 1; i < len(rr.commands); i++ {
		if rr.rangeSecure {
			if rr.commands[i].cb.x.min < minCoord || rr.commands[i].cb.x.max > maxCoord ||
				rr.commands[i].cb.y.min < minCoord || rr.commands[i].cb.y.max > maxCoord ||
				rr.commands[i].cb.z.min < minCoord || rr.commands[i].cb.z.max > maxCoord {
				continue
			}
		}
		for j := len(rr.cubes) - 1; j >= 0; j-- {
			inter := getIntersection(rr.cubes[j], rr.commands[i].cb)
			if inter != nil {
				if rr.cubes[j].value == 1 {
					inter.value = -1
				} else {
					inter.value = 1
				}
				rr.cubes = append(rr.cubes, *inter)
			}
		}
		if rr.commands[i].cmd == onCmd {
			cpyc := rr.commands[i].cb.copy()
			cpyc.value = 1
			rr.cubes = append(rr.cubes, cpyc)
		}
	}
	counter := 0
	for _, c := range rr.cubes {
		v := c.count()
		counter += v
	}
	fmt.Printf("Counted: %d\n", counter)
}
