package submarine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type commandData struct {
	cmd   string
	value int
}

type Submarine struct {
	currentDepth      int
	currentHorizontal int
	aim               int
	journeyData       []commandData
}

func (s *Submarine) loadJourney(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var commands []commandData

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(c[1])
		cmd := commandData{
			cmd:   c[0],
			value: val,
		}
		commands = append(commands, cmd)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	s.journeyData = commands
}

func NewSubmarine(path string) Submarine {
	s := Submarine{
		currentDepth:      0,
		currentHorizontal: 0,
	}
	s.loadJourney(path)
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

func (s *Submarine) ResetSubmarineToFactoryDefault() {
	s.currentDepth = 0
	s.currentHorizontal = 0
	s.aim = 0
}
