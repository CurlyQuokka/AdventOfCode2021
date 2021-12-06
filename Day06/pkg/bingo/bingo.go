package bingo

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day06/pkg/utils"
)

const (
	inputSeparator = ","
	boardSeparator = " "
	boardSize      = 5
)

type board [boardSize][boardSize]int

type bingoBoard struct {
	GameBoard  *board
	ScoreBoard *board
}

type BingoSubsystem struct {
	rawData   []string
	input     []int
	boards    []*bingoBoard
	gameScore int
}

func NewBingoSubsystem(path string) *BingoSubsystem {
	bs := &BingoSubsystem{}
	bs.rawData = utils.LoadData(path)
	bs.input = convertInput(bs.rawData[0], inputSeparator)
	bs.gameScore = 0
	bs.convertBoards(bs.rawData[1:])
	return bs
}

func convertInput(input, separator string) []int {
	var converted []int
	separated := strings.Split(input, separator)
	for _, val := range separated {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
			os.Exit(666)
		}
		converted = append(converted, v)
	}
	return converted
}

func (bs *BingoSubsystem) convertBoards(input []string) {
	i := 0
	b := &board{}
	for _, line := range input {
		space := regexp.MustCompile(`\s+`)
		line = space.ReplaceAllString(line, boardSeparator)
		line = strings.TrimLeft(line, boardSeparator)
		if line != "" {
			l := convertInput(line, boardSeparator)
			for j := 0; j < len(b[i]); j++ {
				b[i][j] = l[j]
			}
			i++
			if i > 4 {
				bb := bingoBoard{
					GameBoard:  b,
					ScoreBoard: &board{},
				}
				bs.boards = append(bs.boards, &bb)
				b = &board{}
				i = 0
			}
		}
	}
}

func (bs *BingoSubsystem) PlayGame(findLoosing bool) {
	for _, val := range bs.input {
		result := bs.playRound(val, findLoosing)
		if result > -1 {
			bs.gameScore = result
			break
		}
	}
}

func (bs *BingoSubsystem) playRound(value int, remove bool) int {
	for index, board := range bs.boards {
		boardUnmarkedSum := 0
		for row := 0; row < boardSize; row++ {
			for column := 0; column < boardSize; column++ {
				if board.GameBoard[row][column] == value {
					board.ScoreBoard[row][column] = 1
				}
			}
		}
		if won, _ := board.ScoreBoard.isRowWinning(); won {
			boardUnmarkedSum = board.unmarkedSum()
		}
		if won, _ := board.ScoreBoard.isColumnWinning(); won {
			boardUnmarkedSum = board.unmarkedSum()
		}
		if boardUnmarkedSum > 0 {
			oldLen := len(bs.boards)
			if remove && oldLen > 1 {
				bs.removeAt(index)
			}
			if remove && oldLen == 1 || !remove {
				return boardUnmarkedSum * value
			}
		}
	}
	return -1
}

func (bs *BingoSubsystem) PrintScore() {
	fmt.Printf("Game score: %d\n", bs.gameScore)
}

func (bs *BingoSubsystem) WreckThisCasino() {
	for _, gb := range bs.boards {
		gb.ScoreBoard = &board{}
	}
	bs.gameScore = 0
}

func (bs *BingoSubsystem) removeAt(index int) {
	bb := []*bingoBoard{}
	for i := range bs.boards {
		if i != index {
			bb = append(bb, bs.boards[i])
		}
	}
	bs.boards = bb
}

func (b *board) isRowWinning() (bool, int) {
	sum := 0
	for row := 0; row < boardSize; row++ {
		for column := 0; column < boardSize; column++ {
			sum += b[row][column]
		}
		if sum == boardSize {
			return true, row
		}
		sum = 0
	}
	return false, -1
}

func (b *board) isColumnWinning() (bool, int) {
	sum := 0
	for column := 0; column < boardSize; column++ {
		for row := 0; row < boardSize; row++ {
			sum += b[row][column]
		}
		if sum == boardSize {
			return true, column
		}
		sum = 0
	}
	return false, -1
}

func (b *bingoBoard) unmarkedSum() int {
	sum := 0
	for row := 0; row < boardSize; row++ {
		for column := 0; column < boardSize; column++ {
			if b.ScoreBoard[row][column] == 0 {
				sum += b.GameBoard[row][column]
			}
		}
	}
	return sum
}
