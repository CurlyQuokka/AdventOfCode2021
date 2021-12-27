package chitonsavoider

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	inputExtension = 5
	maxRisk        = 9
)

type node struct {
	id         int
	risk       uint
	tentative  uint
	neighbours []int
	visited    bool
}

func (ca *ChitonsAvoider) newNode(row, col int) node {
	n := node{
		id:        ca.calculateNodeId(row, col),
		risk:      uint(ca.input[row][col]),
		tentative: math.MaxInt,
		visited:   false,
	}

	if row > 0 {
		n.neighbours = append(n.neighbours, ca.calculateNodeId(row-1, col))
	}

	if row < len(ca.input)-1 {
		n.neighbours = append(n.neighbours, ca.calculateNodeId(row+1, col))
	}

	if col > 0 {
		n.neighbours = append(n.neighbours, ca.calculateNodeId(row, col-1))
	}

	if col < len(ca.input)-1 {
		n.neighbours = append(n.neighbours, ca.calculateNodeId(row, col+1))
	}

	if n.id == 0 {
		n.tentative = 0
	}

	return n
}

func (ca *ChitonsAvoider) calculateNodeId(row, col int) int {
	return row*len(ca.input[0]) + col
}

type ChitonsAvoider struct {
	data           []string
	input          [][]int
	nodes          map[int]node
	unvisitedNodes map[int]*node
	stopNodeId     int
	smallestRisk   uint
}

func NewChitonsAvoider(path string) *ChitonsAvoider {
	ca := &ChitonsAvoider{
		data:           utils.LoadData(path),
		nodes:          make(map[int]node),
		unvisitedNodes: make(map[int]*node),
	}
	ca.convertData()
	return ca
}

func (ca *ChitonsAvoider) convertData() {
	for _, l := range ca.data {
		line := []int{}
		for _, c := range l {
			v, err := strconv.Atoi(string(c))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(646)
			}
			line = append(line, v)
		}
		ca.input = append(ca.input, line)
	}
}

func (ca *ChitonsAvoider) Print() {
	fmt.Printf("%v\n", ca.input)
}

func (ca *ChitonsAvoider) createNodes() {
	ca.nodes = make(map[int]node)
	ca.unvisitedNodes = make(map[int]*node)
	for row := 0; row < len(ca.input); row++ {
		for col := 0; col < len(ca.input[row]); col++ {
			n := ca.newNode(row, col)
			ca.nodes[n.id] = n
			ca.unvisitedNodes[n.id] = &n
		}
	}
	initial := ca.nodes[0]
	initial.tentative = 0
	initial.risk = 0
}

func (ca *ChitonsAvoider) GetInitialNode() *node {
	n := ca.nodes[0]
	return &n
}

func (ca *ChitonsAvoider) FindPath(n *node) {
	for _, id := range n.neighbours {
		if neigh, exists := ca.unvisitedNodes[id]; exists {
			newTen := neigh.risk + n.tentative
			// fmt.Printf("%d + %d\n", neigh.risk, n.tentative)
			if newTen < neigh.tentative {
				neigh.tentative = newTen
			}
		}
	}
	n.visited = true
	delete(ca.unvisitedNodes, n.id)

	toCheck := len(ca.unvisitedNodes)
	if toCheck%1000 == 0 {
		fmt.Printf("%d\n", toCheck)
	}

	if n.id == ca.stopNodeId {
		ca.smallestRisk = n.tentative
		fmt.Printf("Risk: %d\n", ca.smallestRisk)
		return
	}
	sk := ca.FindSmallestUnvisited()
	ca.FindPath(ca.unvisitedNodes[sk])
}

func (ca *ChitonsAvoider) FindSmallestUnvisited() int {
	minKey := 0
	var minValue uint
	minValue = math.MaxUint
	for key, value := range ca.unvisitedNodes {
		if value.tentative < minValue {
			minValue = value.tentative
			minKey = key
		}
	}
	return minKey
}

func (ca *ChitonsAvoider) ExtendInput() {
	rows := len(ca.input) * inputExtension
	cols := len(ca.input[0]) * inputExtension

	extendedInput := make([][]int, rows)
	for i := 0; i < rows; i++ {
		extendedInput[i] = make([]int, cols)
	}

	for i := 0; i < inputExtension; i++ {
		for r := 0; r < len(ca.input); r++ {
			for c := 0; c < len(ca.input[r]); c++ {
				newRisk := ca.input[r][c] + i
				if newRisk > maxRisk {
					newRisk -= maxRisk
				}

				newRow := r
				newCol := c + len(ca.input[r])*i

				extendedInput[newRow][newCol] = newRisk
			}
		}
	}

	for i := 0; i < inputExtension; i++ {
		for r := 0; r < len(ca.input); r++ {
			for c := 0; c < len(extendedInput[r]); c++ {
				newRisk := extendedInput[r][c] + i
				if newRisk > maxRisk {
					newRisk -= maxRisk
				}

				newRow := r + len(ca.input)*i
				newCol := c

				extendedInput[newRow][newCol] = newRisk
			}
		}
	}

	ca.input = extendedInput
}

func (ca *ChitonsAvoider) CalculateRisk(extended bool) {
	if extended {
		ca.ExtendInput()
	}

	ca.stopNodeId = ca.calculateNodeId(len(ca.input)-1, len(ca.input[0])-1)
	ca.createNodes()
	in := ca.GetInitialNode()
	ca.FindPath(in)
}
