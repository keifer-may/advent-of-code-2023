package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	//"slices"
	//"strconv"
	//"strings"
)

var (
	outfile, _ = os.Create("./log")
	l          = log.New(outfile, "", 0)
)

type location struct {
	x   int
	y   int
	val rune
}

type connection struct {
	current  location
	last     *location
	next     *location
	value    []rune
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

func linesToGrid(lines []string) (grid [][]rune) {
	for _, line := range lines {
		list := []rune{}
		for _, char := range line {
			list = append(list, char)
		}
		grid = append(grid, list)
	}
	return grid
}

func findStart(grid [][]rune) location {
	for y, row := range grid {
		for x, char := range row {
			if char == 'S' {
				start := location{x: x, y: y, val: 'S'}
				return start
			}
		}
	}
	return location{x: 0, y: 0, val: ' '}
}

func surroundingStart(grid [][]rune, start location) []location {
	yRange := len(grid) - 1
	xRange := len(grid[0]) - 1
	possibleLocations := []location{}

	if start.x != 0 {
		newX := start.x - 1
		newY := start.y
		value := grid[newY][newX]
		if !(value == '.') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.y != 0 {
		newX := start.x
		newY := start.y - 1
		value := grid[newY][newX]
		if !(value == '.') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.x != xRange {
		newX := start.x + 1
		newY := start.y
		value := grid[newY][newX]
		if !(value == '.') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.y != yRange {
		newX := start.x
		newY := start.y + 1
		value := grid[newY][newX]
		if !(value == '.') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}

	for i, loc := range possibleLocations {
		keep := validSecConnect(start, loc)
		if !keep {
			fmt.Println(loc)
			possibleLocations = slices.Delete(possibleLocations, i, i+1)
		}
	}

	//| is a vertical pipe connecting north and south.
	//- is a horizontal pipe connecting east and west.
	//L is a 90-degree bend connecting north and east.
	//J is a 90-degree bend connecting north and west.
	//7 is a 90-degree bend connecting south and west.
	//F is a 90-degree bend connecting south and east.
	//. is ground; there is no pipe in this tile.
	//S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

	return possibleLocations
}

func validSecConnect(first location, second location) bool {
	deltaX := second.x - first.x
	deltaY := second.y - first.y
	currentVal := second.val

	if (currentVal == '|') && (deltaY == 0) {
		return false
	} else if (currentVal == '-') && (deltaX == 0) {
		return false
	} else if (currentVal == 'L') && ((deltaX != -1) && (deltaY != 1)) {
		return false
	} else if (currentVal == 'J') && ((deltaX != 1) && (deltaY != 1)) {
		return false
	} else if (currentVal == '7') && ((deltaX != 1) && (deltaY != -1)) {
		return false
	} else if (currentVal == 'F') && ((deltaX != -1) && (deltaY != -1)) {
		return false
	} else if (deltaX == 0) && (deltaY == 0) {
		return false
	} else {
		return true
	}
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
	surroundingStart := surroundingStart(grid, start)
	l.Println(surroundingStart)
	fmt.Println(surroundingStart)
}

func main() {
	solutionOne()
}
