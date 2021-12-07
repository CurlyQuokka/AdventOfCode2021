package submarine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/bingo"
	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/crabalignment"
	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/fishobservatory"
	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/hydrothermal"
	"github.com/CurlyQuokka/AdventOfCode2021/Day07/pkg/utils"
)

type commandData struct {
	cmd   string
	value int
}

type Submarine struct {
	currentDepth      int
	currentHorizontal int
	aim               int
	gammaRate         int
	epsilonRate       int
	oxygenGenRating   int
	co2ScrubberRating int
	journeyData       []commandData
	powerReport       []string

	bingoSys        *bingo.BingoSubsystem
	hydroSys        *hydrothermal.HydrothermalAvoidance
	fishobservatory *fishobservatory.Fishobservatory
	crabSubsystem   *crabalignment.CrabAlignment
}

func (s *Submarine) LoadPowerReport(path string) {
	s.powerReport = utils.LoadData(path)
	s.caluclatePowerCoefficients()
	s.oxygenGenRating = getLifeSupportRating(s.powerReport, 0, false)
	s.co2ScrubberRating = getLifeSupportRating(s.powerReport, 0, true)
}

func (s *Submarine) LoadJourney(path string) {
	data := utils.LoadData(path)

	s.journeyData = []commandData{}

	for _, d := range data {
		c := strings.Split(d, " ")
		val, _ := strconv.Atoi(c[1])
		cmd := commandData{
			cmd:   c[0],
			value: val,
		}
		s.journeyData = append(s.journeyData, cmd)
	}
}

func (s *Submarine) caluclatePowerCoefficients() {
	var gamma, epislon string
	for i := 0; i < len(s.powerReport[0]); i++ {
		dom, least := utils.GetDomLeastValues(s.powerReport, i)
		gamma += dom
		epislon += least
	}
	s.gammaRate = utils.ConvertBinToDec(gamma)
	s.epsilonRate = utils.ConvertBinToDec(epislon)
}

func getLifeSupportRating(data []string, index int, keepLeast bool) int {
	if len(data) == 1 {
		return utils.ConvertBinToDec(data[0])
	}

	dom, least := utils.GetDomLeastValues(data, index)

	if keepLeast {
		return getLifeSupportRating(utils.TrimData(data, least, index), index+1, keepLeast)
	}
	return getLifeSupportRating(utils.TrimData(data, dom, index), index+1, keepLeast)
}

func NewSubmarine() Submarine {
	s := Submarine{
		currentDepth:      0,
		currentHorizontal: 0,
		aim:               0,
		gammaRate:         0,
		epsilonRate:       0,
	}
	return s
}

func (s *Submarine) ProcessJourney() (int, int) {
	for _, cmd := range s.journeyData {
		switch c := cmd.cmd; c {
		case "up":
			s.currentDepth -= cmd.value
		case "down":
			s.currentDepth += cmd.value
		case "forward":
			s.currentHorizontal += cmd.value
		}
	}

	return s.currentDepth, s.currentHorizontal
}

func (s *Submarine) ProcessJourneyWithAim() (int, int) {
	for _, cmd := range s.journeyData {
		switch c := cmd.cmd; c {
		case "up":
			s.aim -= cmd.value
		case "down":
			s.aim += cmd.value
		case "forward":
			s.currentHorizontal += cmd.value
			s.currentDepth += s.aim * cmd.value
		}
	}

	return s.currentDepth, s.currentHorizontal
}

func (s *Submarine) PrintDestination() {
	fmt.Printf("Destination: %d\n", s.currentDepth*s.currentHorizontal)
}

func (s *Submarine) PrintPowerUsage() {
	fmt.Printf("Power usage: %d\n", s.gammaRate*s.epsilonRate)
}

func (s *Submarine) PrintLifeSupportRating() {
	fmt.Printf("Life support rating: %d\n", s.oxygenGenRating*s.co2ScrubberRating)
}

func (s *Submarine) ResetSubmarineToFactoryDefault() {
	s.currentDepth = 0
	s.currentHorizontal = 0
	s.aim = 0
	s.gammaRate = 0
	s.epsilonRate = 0
	s.journeyData = []commandData{}
	s.powerReport = []string{}
}

func (s *Submarine) InitializeBingoSubsystem(path string) {
	s.bingoSys = bingo.NewBingoSubsystem(path)
}

func (s *Submarine) GetBingoSubsystem() *bingo.BingoSubsystem {
	return s.bingoSys
}

func (s *Submarine) InitializeHydrothermalSybsystem(path string) {
	s.hydroSys = hydrothermal.NewHydrothermalAvoidance(path)
}

func (s *Submarine) GetHydrothermalSubsystem() *hydrothermal.HydrothermalAvoidance {
	return s.hydroSys
}

func (s *Submarine) InitializeFishObservatory(path string) {
	s.fishobservatory = fishobservatory.InitilizeFishobservatory(path)
}

func (s *Submarine) GetFishObservatory() *fishobservatory.Fishobservatory {
	return s.fishobservatory
}

func (s *Submarine) InitializeCrabArmy(path string) {
	s.crabSubsystem = crabalignment.InitializeCrabAlignment(path)
}

func (s *Submarine) GetCrabArmy() *crabalignment.CrabAlignment {
	return s.crabSubsystem
}
