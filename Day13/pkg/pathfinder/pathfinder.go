package pathfinder

import (
	"fmt"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day13/pkg/utils"
)

const (
	caveSeparator = "-"
	startCaveKey  = "start"
	endCaveKey    = "end"
)

type cave struct {
	name        string
	connections []string
}

type PathFinder struct {
	data  []string
	caves map[string]*cave
	paths []*[]string
}

func NewPathFInder(path string) *PathFinder {
	pf := &PathFinder{
		data:  utils.LoadData(path),
		caves: make(map[string]*cave),
	}
	pf.FindAllCaves()
	return pf
}

func (pf *PathFinder) FindAllCaves() {
	for _, l := range pf.data {
		caves := strings.Split(l, caveSeparator)
		pf.AddCave(caves[0], caves[1])
		pf.AddCave(caves[1], caves[0])
	}
}

func (pf *PathFinder) AddCave(first, second string) {
	if exCave, exists := pf.caves[first]; !exists {
		c := cave{
			name: first,
		}
		c.connections = append(c.connections, second)
		pf.caves[first] = &c
	} else {
		caveAlreadyAdded := false
		for _, con := range exCave.connections {
			if con == second {
				caveAlreadyAdded = true
			}
		}
		if !caveAlreadyAdded {
			exCave.connections = append(exCave.connections, second)
		}
	}
}

func (pf *PathFinder) FindPaths(extendedSearch bool) {
	path := []string{}
	startCave := pf.caves[startCaveKey]
	pf.checkCave(startCave, path, extendedSearch)
}

func checkIfCaveInPath(caveName string, path []string) bool {
	for _, cave := range path {
		if cave == caveName {
			return true
		}
	}
	return false
}

func checkIfSmallCave(caveName string) bool {
	return caveName[0] >= 97
}

func checkIfVisitedSmallCaveTwice(path []string) bool {
	for i := 0; i < len(path); i++ {
		if checkIfSmallCave(path[i]) {
			for j := 0; j < len(path); j++ {
				if i != j {
					if path[i] == path[j] {
						return true
					}
				}
			}
		}
	}
	return false
}

func (pf *PathFinder) checkCave(c *cave, path []string, extendedSearch bool) {
	path = append(path, c.name)
	if c.name == endCaveKey {
		pf.paths = append(pf.paths, &path)
		return
	}
	for _, con := range c.connections {
		if checkIfSmallCave(con) {
			if extendedSearch {
				if checkIfVisitedSmallCaveTwice(path) {
					if !checkIfCaveInPath(con, path) && con != startCaveKey {
						pf.checkCave(pf.caves[con], path, extendedSearch)
					}
				} else if con != startCaveKey {
					pf.checkCave(pf.caves[con], path, extendedSearch)
				}
			} else {
				if !checkIfCaveInPath(con, path) && con != startCaveKey {
					pf.checkCave(pf.caves[con], path, extendedSearch)
				}
			}
		} else if con != startCaveKey {
			pf.checkCave(pf.caves[con], path, extendedSearch)
		}
	}
}

func (pf *PathFinder) PrintData() {
	for _, l := range pf.data {
		fmt.Println(l)
	}
}

func (pf *PathFinder) PrintCaves() {
	for _, cave := range pf.caves {
		cave.print()
	}
}

func (pf *PathFinder) PrintPaths() {
	for _, path := range pf.paths {
		for _, c := range *path {
			fmt.Printf("%s ", c)
		}
		fmt.Println()
	}
}

func (c *cave) print() {
	fmt.Printf("%s connected to: %v\n", c.name, c.connections)
}

func (pf *PathFinder) GetNumberOfPaths() {
	fmt.Printf("Number of paths: %d\n", len(pf.paths))
}

func (pf *PathFinder) Reset() {
	pf.paths = []*[]string{}
}
