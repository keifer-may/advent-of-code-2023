package utils

import (
	"bufio"
	"os"
)

type Location struct {
	X   int
	Y   int
	Val rune
}

func FileToStringArray(path string) (lines []string, err error) {

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

func FileToRuneGrid(path string) (grid [][]rune) {
	lines, err := FileToStringArray(path)
	if err != nil {
		return
	}

	for _, line := range lines {
		list := []rune{}
		for _, char := range line {
			list = append(list, char)
		}
		grid = append(grid, list)
	}
	return grid
}

func LocItemsInGrid(grid [][]rune, value rune) (locations []Location) {
	for y, row := range grid {
		for x, char := range row {
			if char == value {
				loc := Location{X: x, Y: y, Val: value}
				locations = append(locations, loc)
			}
		}
	}
	return locations
}
