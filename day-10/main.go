package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"slices"
	//"strconv"
	//"strings"
)

var (
	outfile, _ = os.Create("./log")
	l = log.New(outfile, "", 0)
)

type location struct {
	x int
	y int
}

type connection struct {
	current location
	last *location
	next *location
	value []rune
	possible [2]location
}

func readFileToListStrings(path string) (lines []string, err error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	return lines, nil
}

func linesToGrid (lines []string) (grid [][]rune) {
	for _, line := range lines {
		list := []rune{}
		for _, char := range line {
			list = append(list, char)
		}
		grid = append(grid, list)
	}
	return grid
}

func findStart(grid [][]rune) (location) {
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				start := location{x,y}
				return start
			}
		}
	}
	return location{0, 0}
}



func solutionOne() {
	lines, _ := readFileToListStrings("./example1.txt")
	fmt.Println(lines)
	l.Println(lines)
	grid := linesToGrid(lines)
	fmt.Println(grid)
	l.Println(grid)
	start := findStart(grid)
	fmt.Println(start)
	l.Println(start)
}

func main() {
	solutionOne()
}
