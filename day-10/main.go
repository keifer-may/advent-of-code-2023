package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
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
				start := location{x: x, y: y}
				return start
			}
		}
	}
	return location{x: 0, y: 0}
}

func surroundingStart(grid [][]rune, start location) ([]location) {
	yRange := len(grid) - 1
	xRange := len(grid[0]) - 1
	possibleLocations := []location{}
	if start.x != 0 {
		possibleLocations = append(possibleLocations, location{x: (start.x - 1), y: start.y})
	}
	if start.y != 0 {
		possibleLocations = append(possibleLocations, location{x: start.x, y: (start.y - 1)})
	}
	if start.x != xRange {
		possibleLocations = append(possibleLocations, location{x: (start.x + 1), y: start.y})
	}
	if start.y != yRange {
		possibleLocations = append(possibleLocations, location{x: start.x, y: (start.y + 1)})
	}

	delInd := []int{}
	for i, loc := range possibleLocations {
		val := grid[loc.y][loc.x]
		if val == '.' {
			delInd = append(delInd, i)
		}
	}

	for i, delInd := range delInd {
		
	}

	// otherwise we need to check what the value is in x +/- 1 or y +/- 1 and test if that is feasible at all......


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
