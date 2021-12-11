package lcddisplay

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/CurlyQuokka/AdventOfCode2021/Day11/pkg/utils"
)

const (
	dataSeparator   = "|"
	numberSeparator = " "
)

func getDispNums() []string {
	return []string{"1110111", "0010010", "1011101", "1011011", "0111010", "1101011", "1101111", "1010010", "1111111", "1111011"}
}

type LCDDisplay struct {
	data []string
}

type digit struct {
	T, TL, TR, M, BL, BR, B string
	dispString              string
}

func removeWord(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}

func NewDigit(data string) digit {
	d := digit{}
	words := strings.Split(data, numberSeparator)
	var numbers [10]string

	for i := 0; i < len(words); i++ {
		switch len(words[i]) {
		case 2:
			numbers[1] = words[i]
			removeWord(words, i)
			i--
		case 3:
			numbers[7] = words[i]
			removeWord(words, i)
			i--
		case 4:
			numbers[4] = words[i]
			removeWord(words, i)
			i--
		case 7:
			numbers[8] = words[i]
			removeWord(words, i)
			i--
		}
	}

	for i := 0; i < len(words); i++ {
		switch len(words[i]) {
		case 5:
			// 3
			_, n := getCommonChars(words[i], numbers[1])
			if n == 2 {
				numbers[3] = words[i]
				removeWord(words, i)
				i--
			}
			// 2, 5

		case 6:
			// 6
			_, n := getCommonChars(words[i], numbers[1])
			if n == 1 {
				numbers[6] = words[i]
				removeWord(words, i)
				i--
			}
		}
	}

	for i := 0; i < len(words); i++ {
		switch len(words[i]) {
		case 5:
			// 5
			_, n := getCommonChars(words[i], numbers[6])
			if n == 5 {
				numbers[5] = words[i]
				removeWord(words, i)
				i--
			} else if n == 4 {
				numbers[2] = words[i]
				removeWord(words, i)
				i--
			}
		case 6:
			// 0 , 9
			_, n := getCommonChars(words[i], numbers[4])
			if n == 4 {
				numbers[9] = words[i]
				removeWord(words, i)
				i--
			} else {
				numbers[0] = words[i]
				removeWord(words, i)
				i--
			}
		}
	}

	d.T = getUniqueChar(numbers[7], numbers[1])
	d.M = getUniqueChar(numbers[8], numbers[0])

	TLcommon, _ := getCommonChars(numbers[4], numbers[1])
	d.TL = numbers[4]
	for _, l := range TLcommon {
		d.TL = strings.ReplaceAll(d.TL, string(l), "")
	}

	d.TL = strings.ReplaceAll(d.TL, d.M, "")
	d.TR = getUniqueChar(numbers[8], numbers[6])

	d.BL = getUniqueChar(numbers[8], numbers[9])
	d.BR, _ = getCommonChars(numbers[6], numbers[1])

	all := d.T + d.TL + d.TR + d.M + d.BL + d.BR
	d.B = getUniqueChar(numbers[8], all)

	d.dispString = all + d.B
	return d
}

func NewLCDDisplay(path string) *LCDDisplay {
	ha := &LCDDisplay{
		data: utils.LoadData(path),
	}
	return ha
}

func (lcd *LCDDisplay) Count1478() {
	count := 0
	for _, line := range lcd.data {
		splitted := strings.Split(line, dataSeparator)
		displayed := strings.TrimLeft(splitted[1], numberSeparator)
		digits := strings.Split(displayed, numberSeparator)
		for _, digit := range digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Printf("Count 1478: %d\n", count)
}

func getUniqueChar(longer, shorter string) string {
	for _, c := range longer {
		if !strings.Contains(shorter, string(c)) {
			return string(c)
		}
	}
	return ""
}

func getCommonChars(a, b string) (string, int) {
	common := ""
	for _, c := range a {
		if strings.Contains(b, string(c)) {
			common += string(c)
		}
	}
	return common, len(common)
}

func (lcd *LCDDisplay) Decode() {
	sum := 0
	for _, line := range lcd.data {
		splitted := strings.Split(line, dataSeparator)
		d := NewDigit(splitted[0])
		displayed := strings.TrimLeft(splitted[1], numberSeparator)
		values := strings.Split(displayed, numberSeparator)
		strVal := ""

		for _, v := range values {
			strVal += decodeDigit(d, v)
		}

		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(5454)
		}
		sum += intVal
	}
	fmt.Printf("Sum: %d\n", sum)
}

func decodeDigit(d digit, value string) string {
	s := ""
	for _, c := range d.dispString {
		if strings.Contains(value, string(c)) {
			s += "1"
		} else {
			s += "0"
		}
	}

	dispNums := getDispNums()
	for i := 0; i < len(dispNums); i++ {
		if dispNums[i] == s {
			return fmt.Sprint(i)
		}
	}
	return ""
}
