package bestroute

import (
	"fmt"
	"sort"

	"github.com/CurlyQuokka/AdventOfCode2021/Day15/pkg/utils"
)

const (
	autoScoreMultiplier = 5
)

type symbolStack []rune

func (ss *symbolStack) push(r rune) {
	*ss = append(*ss, r)
}

func (ss *symbolStack) pop() rune {
	numOfElements := len(*ss)
	if numOfElements == 0 {
		return rune(0)
	}
	r := (*ss)[numOfElements-1]
	if numOfElements <= 1 {
		ss.clearStack()
	} else {
		ss.removeLastElement()
	}
	return r
}

func (ss *symbolStack) clearStack() {
	*ss = []rune{}
}

func (ss *symbolStack) removeLastElement() {
	numOfElements := len(*ss)
	*ss = (*ss)[0 : numOfElements-1]
}

func (ss *symbolStack) getLastElement() rune {
	numOfElements := len(*ss)
	if numOfElements == 0 {
		return rune(0)
	}
	return (*ss)[numOfElements-1]
}

type BestRoute struct {
	data             []string
	symbols          symbolStack
	charactersMap    map[rune]rune
	scoreMap         map[rune]int
	autoScoreMap     map[rune]int
	corruptedSymbols []rune
	autoScores       []int
}

func NewBesRoute(path string) *BestRoute {
	br := &BestRoute{
		data:          utils.LoadData(path),
		charactersMap: make(map[rune]rune),
		scoreMap:      make(map[rune]int),
		autoScoreMap:  make(map[rune]int),
	}

	br.charactersMap[')'] = '('
	br.charactersMap[']'] = '['
	br.charactersMap['}'] = '{'
	br.charactersMap['>'] = '<'

	br.scoreMap[')'] = 3
	br.scoreMap[']'] = 57
	br.scoreMap['}'] = 1197
	br.scoreMap['>'] = 25137

	br.autoScoreMap['('] = 1
	br.autoScoreMap['['] = 2
	br.autoScoreMap['{'] = 3
	br.autoScoreMap['<'] = 4

	return br
}

func (br *BestRoute) PrintData() {
	for _, l := range br.data {
		fmt.Println(l)
	}
}

func (br *BestRoute) ProcessData() {
	for _, line := range br.data {
		isCorrupted := false
		for _, s := range line {

			if val, closingSymbol := br.charactersMap[s]; closingSymbol {
				if val != br.symbols.getLastElement() {
					br.corruptedSymbols = append(br.corruptedSymbols, s)
					isCorrupted = true
					break
				} else {
					br.symbols.pop()
				}
			} else {
				br.symbols.push(s)
			}
		}
		if !isCorrupted {
			autoScore := 0
			for r := br.symbols.pop(); r != rune(0); r = br.symbols.pop() {
				autoScore *= autoScoreMultiplier
				autoScore += br.autoScoreMap[r]
			}
			br.autoScores = append(br.autoScores, autoScore)
		}
		br.symbols.clearStack()
	}

	score := 0
	for _, r := range br.corruptedSymbols {
		score += br.scoreMap[r]
	}

	fmt.Printf("Corrupted score: %d\n", score)

	sort.Ints(br.autoScores)

	fmt.Printf("Autocomplete score: %v\n", br.autoScores[len(br.autoScores)/2])
}
