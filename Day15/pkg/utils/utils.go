package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LoadData(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func GetDomLeastValues(data []string, index int) (string, string) {
	var values []int
	for _, d := range data {
		v, err := strconv.Atoi(string(d[index]))
		if err != nil {
			log.Fatal(err)
			os.Exit(44)
		}
		values = append(values, v)
	}

	ones := sumArray(values)
	zeros := len(values) - ones

	if ones >= zeros {
		return "1", "0"
	}
	return "0", "1"
}

func ConvertBinToDec(value string) int {
	val, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		log.Fatal(err)
		os.Exit(46)
	}
	return int(val)
}

func sumArray(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func TrimData(data []string, value string, index int) []string {
	var newData []string
	for _, d := range data {
		if string(d[index]) == value {
			newData = append(newData, d)
		}
	}
	return newData
}

func GetHigher(checked, checker *int) {
	if *checked > *checker {
		*checker = *checked
	}
}

func Sum2DIntSlice(toSum *[][]int) int {
	sum := 0
	for _, row := range *toSum {
		for _, col := range row {
			sum += col
		}
	}
	return sum
}

func PrintStringSlice(s []string) {
	for _, is := range s {
		fmt.Println(is)
	}
}

func Prepare2DInt(sizeX, sizeY int) *[][]int {
	b := make([][]int, sizeY)
	for i := range b {
		b[i] = make([]int, sizeX)
	}
	return &b
}

func Print2DIntSlice(s *[][]int) {
	for _, row := range *s {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}
