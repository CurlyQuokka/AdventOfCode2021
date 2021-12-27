package snailfishmath

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	splitValue  = 10
	reduceLevel = 4
)

type SnailfishMath struct {
	data []string
	root *snailfishNumber
}

type snailfishNumber struct {
	value      int
	parent     *snailfishNumber
	leftChild  *snailfishNumber
	rightChild *snailfishNumber
	magnitude  int
}

func NewSnailfishNumber(v int, p *snailfishNumber) *snailfishNumber {
	return &snailfishNumber{
		value:  v,
		parent: p,
	}
}

func NewDefaultSnailfishNumber() *snailfishNumber {
	return &snailfishNumber{
		value:  -1,
		parent: nil,
	}
}

func (sn *snailfishNumber) Add(snToAdd *snailfishNumber) *snailfishNumber {
	newSn := NewDefaultSnailfishNumber()
	sn.parent = newSn
	snToAdd.parent = newSn
	newSn.leftChild = sn
	newSn.rightChild = snToAdd
	return newSn
}

func (sn *snailfishNumber) Print() {
	if sn.value > -1 {
		fmt.Print(sn.value)
	} else {
		fmt.Print("[")
		sn.leftChild.Print()
		fmt.Print(",")
		sn.rightChild.Print()
		fmt.Print("]")
	}
}

func NewSnailfishMath(path string) *SnailfishMath {
	sm := &SnailfishMath{
		data: utils.LoadData(path),
	}
	sm.root = processNumber(sm.data[0], nil)
	return sm
}

func (sm *SnailfishMath) Print() {
	utils.PrintStringSlice(sm.data)
	sm.root.Print()
	fmt.Println()
}

func processNumber(data string, parent *snailfishNumber) *snailfishNumber {
	sn := NewSnailfishNumber(-1, parent)

	separatorIndex := -1

	if !strings.Contains(data, "[") {
		val, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(5252)
		}
		sn.value = val
	} else {
		data = data[1 : len(data)-1]

		if data[0] != '[' {
			separatorIndex = strings.Index(data, ",")
		} else {
			openCounter := 0
			closeCounter := 0
			openIndex := -1

			for i, c := range data {
				if c == '[' {
					openCounter++
					if openIndex == -1 {
						openIndex = i
					}
				}
				if c == ']' {
					closeCounter++
				}
				if openCounter == closeCounter && i < len(data)-1 {
					separatorIndex = i + 1
					break
				}
			}
		}

		left := data[0:separatorIndex]
		right := data[separatorIndex+1:]

		sn.leftChild = processNumber(left, sn)
		sn.rightChild = processNumber(right, sn)
	}

	return sn
}

func (sm *SnailfishMath) ReplaceData(newData []string) {
	sm.data = newData
}

func (sm *SnailfishMath) GetData() []string {
	return sm.data
}

func (sm *SnailfishMath) DoMath() int {
	sm.root = processNumber(sm.data[0], nil)

	for i := 1; i < len(sm.data); i++ {
		toAdd := processNumber(sm.data[i], nil)
		sm.root = sm.root.Add(toAdd)

		j := 0
		for {
			j++
			numbers := []*snailfishNumber{}
			sm.root.traverse(&numbers)
			exploded := false
			splitted := false
			for k, number := range numbers {
				if k > 0 {
					exploded = number.explode()
				}
				if exploded {
					break
				}
			}
			if exploded {
				continue
			}

			for _, number := range numbers {
				splitted = number.split()
				if splitted {
					break
				}
			}
			if splitted {
				continue
			}

			if !exploded && !splitted {
				break
			}
		}
	}
	sm.root.calculateMagnitude()
	return sm.root.magnitude
}

func (sn *snailfishNumber) split() bool {
	if sn.value >= splitValue {
		l := sn.value / 2
		r := int(math.Ceil(float64(sn.value) / 2.0))

		newData := fmt.Sprintf("[%d,%d]", l, r)
		newSn := processNumber(newData, sn.parent)

		if sn == sn.parent.leftChild {
			sn.parent.leftChild = newSn
		} else {
			sn.parent.rightChild = newSn
		}
		sn.parent = nil

		return true
	}
	return false
}

func (sn *snailfishNumber) checkLevel() int {
	if sn == sn.getRoot() {
		return 0
	}
	p := sn.parent
	level := 1
	for {
		if p.parent == nil {
			return level
		} else {
			p = p.parent
			level++
		}
	}
}

func (sn *snailfishNumber) traverse(numbers *[]*snailfishNumber) {
	// sn.Print()
	// fmt.Println()
	*numbers = append(*numbers, sn)
	if sn.leftChild != nil {
		sn.leftChild.traverse(numbers)
	}
	if sn.rightChild != nil {
		sn.rightChild.traverse(numbers)
	}
}

func (sn *snailfishNumber) explode() bool {
	left := sn.leftChild
	right := sn.rightChild

	if sn.value == -1 && sn.checkLevel() >= reduceLevel && left.value > -1 && right.value > -1 {
		numbers := []*snailfishNumber{}
		root := sn.getRoot()
		root.traverse(&numbers)

		leafs := []*snailfishNumber{}

		for _, number := range numbers {
			if number.value > -1 {
				leafs = append(leafs, number)

			}
		}

		leftIndex := -1
		rightIndex := -1
		for i := 0; i < len(leafs); i++ {
			if leafs[i] == sn.leftChild {
				leftIndex = i
			}
			if leafs[i] == sn.rightChild {
				rightIndex = i
			}
			if leftIndex > -1 && rightIndex > -1 {
				break
			}
		}
		if leftIndex > 0 {
			leafs[leftIndex-1].value += sn.leftChild.value
		}
		if rightIndex < len(leafs)-1 {
			leafs[rightIndex+1].value += sn.rightChild.value
		}
		sn.leftChild = nil
		sn.rightChild = nil
		sn.value = 0
		return true
	}
	return false
}

func (sn *snailfishNumber) getRoot() *snailfishNumber {
	p := sn.parent
	for {
		if p.parent == nil {
			return p
		}
		p = p.parent
	}
}

func (sn *snailfishNumber) calculateMagnitude() {
	if sn.value == -1 {
		sn.leftChild.calculateMagnitude()
		sn.rightChild.calculateMagnitude()
		sn.magnitude = sn.leftChild.magnitude*3 + sn.rightChild.magnitude*2
	} else {
		sn.magnitude = sn.value
	}
}

func (sm *SnailfishMath) Reset() {
	sm.root = nil
}
