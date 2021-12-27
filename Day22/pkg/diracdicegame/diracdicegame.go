package diracdicegame

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day22/pkg/utils"
)

const (
	positionSeprator = ": "
)

type player struct {
	position uint
	score    uint
}

func newPlayer(initalPosition uint) *player {
	return &player{
		position: initalPosition,
		score:    0,
	}
}

func (p *player) move(d Die, fields uint) {
	dieResult := uint(0)
	for i := 0; i < 3; i++ {
		dieResult += d.Roll()
	}

	p.position += dieResult
	mod := p.position % fields
	if mod == 0 {
		p.position = fields
	} else {
		p.position = mod
	}
	p.score += p.position
}

type Die interface {
	Roll() uint
	GetNumberOfRolls() uint
}

type DeterministicDie struct {
	sides         uint
	numberOfRolls uint
	fullRolls     uint
}

func NewDeterministicDie(s uint) *DeterministicDie {
	return &DeterministicDie{
		sides:         s,
		numberOfRolls: 0,
	}
}

func (dd *DeterministicDie) Roll() uint {
	dd.numberOfRolls++
	result := dd.numberOfRolls
	if dd.numberOfRolls >= dd.sides {
		dd.numberOfRolls = 0
		dd.fullRolls++
	}
	return result
}

func (dd *DeterministicDie) GetNumberOfRolls() uint {
	return dd.fullRolls*dd.sides + dd.numberOfRolls
}

type DiracDiceGame struct {
	numberOfFields uint
	wininngScore   uint
	gameDie        Die
	players        []*player
}

func NewDiracDiceGame(fields uint, gameDie Die, path string) *DiracDiceGame {
	return &DiracDiceGame{
		numberOfFields: fields,
		gameDie:        gameDie,
		players:        createPlayers(path),
	}
}

func (ddg *DiracDiceGame) PlayFirstGame(winningScore uint) {
	winning := false
	loosingPlayer := -1
	for {
		for i, p := range ddg.players {
			p.move(ddg.gameDie, ddg.numberOfFields)
			if p.score >= winningScore {
				winning = true
				loosingPlayer = 1 - i
				break
			}
		}
		if winning {
			break
		}
	}
	rolls := ddg.gameDie.GetNumberOfRolls()
	score := ddg.players[loosingPlayer].score
	fmt.Printf("Game result: %d\n", rolls*score)
}

func createPlayers(path string) []*player {
	data := utils.LoadData(path)
	players := []*player{}
	for _, l := range data {
		input := strings.Split(l, positionSeprator)
		initial, err := strconv.Atoi(input[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(49)
		}
		players = append(players, newPlayer(uint(initial)))
	}
	return players
}

func (ddg *DiracDiceGame) ResetPlayers(path string) {
	ddg.players = createPlayers(path)
}

func (ddg *DiracDiceGame) copyPlayers() []player {
	copy := []player{}
	for _, p := range ddg.players {
		copy = append(copy, *p)
	}
	return copy
}

func (ddg *DiracDiceGame) PlaySecondGame(winningScore, fields uint) {
	numOfWins := uint(0)
	numOfUniverses := uint(0)

	numOfRollValues := map[uint]uint{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}

	for key, value := range numOfRollValues {
		numOfUniversesTmp, numOfWinsTmp := rollQuantumDie(ddg.players[0].position, ddg.players[1].position, ddg.players[0].score,
			ddg.players[1].score, key, fields, winningScore, true, &numOfRollValues)
		numOfWins += numOfWinsTmp * value
		numOfUniverses += numOfUniversesTmp * value
	}

	fmt.Printf("Number of universes: %d\n", numOfUniverses)
	fmt.Printf("Number of Player 1 wins: %d\n", numOfWins)
	fmt.Printf("Number of Player 2 wins: %d\n", numOfUniverses-numOfWins)
}

func rollQuantumDie(p1Position, p2Position, p1Score, p2Score, value, fields, winningScore uint, isPlayerOne bool, numOfRollValues *map[uint]uint) (uint, uint) {
	numOfWins := uint(0)
	numOfUniverses := uint(0)

	if isPlayerOne {
		p1Position = ((p1Position + value - 1) % fields) + 1
		p1Score += p1Position
	} else {
		p2Position = ((p2Position + value - 1) % fields) + 1
		p2Score += p2Position
	}

	if p1Score >= winningScore || p2Score >= winningScore {
		if p1Score >= winningScore {
			return 1, 1
		} else {
			return 1, 0
		}
	} else {
		isPlayerOne = !isPlayerOne
		for key, value := range *numOfRollValues {
			numOfUniversesTmp, numOfWinsTmp := rollQuantumDie(p1Position, p2Position, p1Score, p2Score, key, fields, winningScore, isPlayerOne, numOfRollValues)
			numOfWins += numOfWinsTmp * value
			numOfUniverses += numOfUniversesTmp * value
		}
	}

	return numOfUniverses, numOfWins
}
