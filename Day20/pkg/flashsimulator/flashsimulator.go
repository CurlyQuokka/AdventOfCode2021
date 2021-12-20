package flashsimulator

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CurlyQuokka/AdventOfCode2021/Day20/pkg/utils"
)

const (
	flashValue = 9
)

type FlashSimulator struct {
	data             []string
	energyLevels     [][]int
	rowSize, colSize int
	flashCounter     int
	synchronizedStep int
	currentStep      int
}

func NewFlashSimulator(path string) *FlashSimulator {
	fs := &FlashSimulator{
		data: utils.LoadData(path),
	}
	fs.Reset()
	return fs
}

func (fs *FlashSimulator) Reset() {
	fs.energyLevels = [][]int{}
	fs.prepareData()
	fs.rowSize = len(fs.energyLevels)
	fs.colSize = len(fs.energyLevels[0])
	fs.flashCounter = 0
	fs.synchronizedStep = 0
	fs.currentStep = 0
}

func (fs *FlashSimulator) RunSimulation(steps int) {
	if steps > 0 {
		for i := 0; i < steps; i++ {
			fs.runStep()
		}
	} else {
		for ok, isSynced := false, false; !ok; ok = isSynced {
			fs.currentStep++
			fs.runStep()

			if isSynced = fs.isSynchonizedFlash(); isSynced {
				if fs.synchronizedStep == 0 {
					fs.synchronizedStep = fs.currentStep
				}
			}
		}
	}
}

func (fs *FlashSimulator) PrintNumOfFlashes() {
	fmt.Printf("Number of flashes: %d\n", fs.flashCounter)
}

func (fs *FlashSimulator) PrintSynchronizedStep() {
	fmt.Printf("Synchornized step: %d\n", fs.synchronizedStep)
}

func (fs *FlashSimulator) prepareData() {
	for _, line := range fs.data {
		var row []int
		for _, v := range line {
			iv, err := strconv.Atoi(string(v))
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(99)
			}
			row = append(row, iv)
		}
		fs.energyLevels = append(fs.energyLevels, row)
	}
}

func (fs *FlashSimulator) runStep() {
	fs.increaseEnergyLevels()
	initialFlashMap := fs.createFlashMap()
	fs.simulateFlashes(initialFlashMap)
	fs.zeroizeFlashed()
}

func (fs *FlashSimulator) increaseEnergyLevels() {
	for i, row := range fs.energyLevels {
		for j := range row {
			fs.energyLevels[i][j]++
		}
	}
}

func (fs *FlashSimulator) createFlashMap() *[][]int {
	fm := make([][]int, fs.rowSize)
	for i := 0; i < fs.rowSize; i++ {
		fm[i] = make([]int, fs.colSize)
	}
	return &fm
}

func (fs *FlashSimulator) findOctopusesToFlash() *[][]int {
	fm := fs.createFlashMap()
	for row, d := range fs.energyLevels {
		for col, value := range d {
			if value > flashValue {
				(*fm)[row][col] = 1
			}
		}
	}
	return fm
}

func (fs *FlashSimulator) getFlashMapDifference(oldFlashMap, newFlashMap *[][]int) (*[][]int, int) {
	counter := 0
	result := fs.createFlashMap()
	for row := range *oldFlashMap {
		for col := range (*oldFlashMap)[0] {
			if (*oldFlashMap)[row][col] != (*newFlashMap)[row][col] {
				(*result)[row][col]++
				counter++
			}
		}
	}
	return result, counter
}

func (fs *FlashSimulator) simulateFlashes(oldFlashMap *[][]int) {
	newFlashMap := fs.findOctopusesToFlash()
	difFlashMap, counter := fs.getFlashMapDifference(oldFlashMap, newFlashMap)
	if counter > 0 {
		fs.flashCounter += counter
		fs.doFlashes(difFlashMap)
		fs.simulateFlashes(newFlashMap)
	}
}

func (fs *FlashSimulator) doFlashes(flashMap *[][]int) {
	for row := range *flashMap {
		for col := range (*flashMap)[0] {
			if (*flashMap)[row][col] > 0 {
				fs.flash(row, col)
			}
		}
	}
}

func (fs *FlashSimulator) zeroizeFlashed() {
	for i, row := range fs.energyLevels {
		for j := range row {
			if fs.energyLevels[i][j] > flashValue {
				fs.energyLevels[i][j] = 0
			}
		}
	}
}

func (fs *FlashSimulator) flash(row, col int) {
	if row > 0 && col > 0 {
		fs.energyLevels[row-1][col-1]++
	}
	if row > 0 {
		fs.energyLevels[row-1][col]++
	}
	if col > 0 {
		fs.energyLevels[row][col-1]++
	}
	if row < fs.rowSize-1 && col < fs.colSize-1 {
		fs.energyLevels[row+1][col+1]++
	}
	if row < fs.rowSize-1 {
		fs.energyLevels[row+1][col]++
	}
	if col < fs.colSize-1 {
		fs.energyLevels[row][col+1]++
	}
	if row > 0 && col < fs.colSize-1 {
		fs.energyLevels[row-1][col+1]++
	}
	if row < fs.rowSize-1 && col > 0 {
		fs.energyLevels[row+1][col-1]++
	}
}

func (fs *FlashSimulator) isSynchonizedFlash() bool {
	for i, row := range fs.energyLevels {
		for j := range row {
			if fs.energyLevels[i][j] > 0 {
				return false
			}
		}
	}
	return true
}
