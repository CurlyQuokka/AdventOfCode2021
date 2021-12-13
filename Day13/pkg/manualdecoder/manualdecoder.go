package manualdecoder

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day13/pkg/hydrothermal"
	"github.com/CurlyQuokka/AdventOfCode2021/Day13/pkg/utils"
)

const (
	positionSeparator         = ","
	instructionSeparator      = " "
	instructionValueSeparator = "="

	up   direction = "up"
	left direction = "left"
)

type direction string

type instruction struct {
	dir  direction
	line int
}

func convertInstruction(l string) instruction {
	ls := strings.Split(l, instructionSeparator)
	ils := strings.Split(ls[len(ls)-1], instructionValueSeparator)

	val, err := strconv.Atoi(ils[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(524)
	}

	instr := instruction{
		dir:  up,
		line: val,
	}
	fmt.Println(ils)

	if ils[0] == "x" {
		instr.dir = left
	}

	return instr
}

type ManualDecoder struct {
	data         []string
	points       []*hydrothermal.Point
	instructions []instruction
	maxX, maxY   int

	page *[][]int
}

func NewManualDecoder(path string) *ManualDecoder {
	md := &ManualDecoder{
		data: utils.LoadData(path),
	}
	md.convertPoints()
	md.convertInstructions()
	md.findMaxes()
	md.preparePage()
	return md
}

func (md *ManualDecoder) convertPoints() {
	for _, l := range md.data {
		if l == "" {
			break
		}
		p := hydrothermal.CretePoint(l)
		md.points = append(md.points, &p)
	}
}

func (md *ManualDecoder) convertInstructions() {
	blank := 0
	for _, v := range md.data {
		if v == "" {
			break
		}
		blank++
	}

	for i := blank + 1; i < len(md.data); i++ {
		md.instructions = append(md.instructions, convertInstruction(md.data[i]))
	}
}

func (md *ManualDecoder) findMaxes() {
	for _, p := range md.points {
		if p.X > md.maxX {
			md.maxX = p.X
		}
		if p.Y > md.maxY {
			md.maxY = p.Y
		}
	}
}

func (md *ManualDecoder) PrintData() {
	utils.PrintStringSlice(md.data)

	fmt.Println()
	for _, p := range md.points {
		fmt.Printf("%v\n", p)
	}

	fmt.Println()
	for _, i := range md.instructions {
		fmt.Printf("%v\n", i)
	}

	fmt.Printf("\nMax X: %d, Max Y: %d\n", md.maxX, md.maxY)

	fmt.Println()
	for _, l := range *md.page {
		for _, v := range l {
			if v > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func (md *ManualDecoder) preparePage() {
	md.page = utils.Prepare2DInt(md.maxX+1, md.maxY+1)
	for _, p := range md.points {
		(*md.page)[p.Y][p.X]++
	}
}

func (md *ManualDecoder) GetInstruction(index int) instruction {
	return md.instructions[index]
}

func (md *ManualDecoder) FoldPage(instr instruction) {
	if instr.dir == up {
		md.foldUp(instr.line)
	} else {
		md.foldLeft(instr.line)
	}
}

func (md *ManualDecoder) foldUp(line int) {
	newPage := utils.Prepare2DInt(md.maxX+1, line)

	for y := 0; y < line; y++ {
		for x := 0; x < md.maxX+1; x++ {
			(*newPage)[y][x] = (*md.page)[y][x]
		}
	}

	for y1, y2 := line+1, line-1; y1 < md.maxY+1; y1, y2 = y1+1, y2-1 {
		for x := 0; x < md.maxX+1; x++ {
			(*newPage)[y2][x] += (*md.page)[y1][x]
		}
	}

	md.page = newPage
	md.maxY = line - 1
}

func (md *ManualDecoder) foldLeft(line int) {
	newPage := utils.Prepare2DInt(line, md.maxY+1)

	for y := 0; y < md.maxY+1; y++ {
		for x := 0; x < line; x++ {
			(*newPage)[y][x] = (*md.page)[y][x]
		}
	}

	for y := 0; y < md.maxY+1; y++ {
		for x1, x2 := line+1, line-1; x1 < md.maxX+1; x1, x2 = x1+1, x2-1 {
			(*newPage)[y][x2] += (*md.page)[y][x1]
		}
	}

	md.page = newPage
	md.maxX = line - 1
}

func (md *ManualDecoder) CountDots() {
	counter := 0
	for _, l := range *md.page {
		for _, v := range l {
			if v > 0 {
				counter++
			}
		}
	}
	fmt.Printf("Visible dots: %d\n", counter)
}

func (md *ManualDecoder) FoldRemaining() {
	for i := 1; i < len(md.instructions); i++ {
		fmt.Println(md.instructions[i])
		md.FoldPage(md.instructions[i])
	}
}
