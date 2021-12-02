package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	minDataLenWoWindow   = 2
	minDataLenWithWindow = 5
)

type sonarData []int

func main() {
	data := loadSonarData(os.Args[1])
	c := data.countInc()
	fmt.Printf("Result: %d\n", c)

	c = data.countWindow()
	fmt.Printf("Result: %d\n", c)
}

func loadSonarData(path string) sonarData {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data sonarData

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		data = append(data, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func (sd *sonarData) countInc() int {
	counter := 0
	if len(*sd) < 2 {
		fmt.Println("Invalid data: too few values")
		os.Exit(42)
	}
	for i := 1; i < len(*sd); i++ {
		if (*sd)[i] > (*sd)[i-1] {
			counter++
		}
	}
	return counter
}

func (sd *sonarData) countWindow() int {
	counter := 0
	if len(*sd) < minDataLenWithWindow {
		fmt.Println("Invalid data: too few values")
		os.Exit(43)
	}
	for i := 1; i < len(*sd)-2; i++ {
		w1 := (*sd)[i-1] + (*sd)[i] + (*sd)[i+1]
		w2 := (*sd)[i] + (*sd)[i+1] + (*sd)[i+2]
		if w2 > w1 {
			counter++
		}
	}
	return counter
}
