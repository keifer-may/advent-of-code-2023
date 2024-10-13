package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"slices"
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
		if (value == '-') || (value == 'L') || (value == 'F') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.y != 0 {
		newX := start.x
		newY := start.y - 1
		value := grid[newY][newX]
		if (value == '|') || (value == '7') || (value == 'F') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.x != xRange {
		newX := start.x + 1
		newY := start.y
		value := grid[newY][newX]
		if (value == '-') || (value == '7') || (value == 'J') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}
	if start.y != yRange {
		newX := start.x
		newY := start.y + 1
		value := grid[newY][newX]
		if (value == '|') || (value == 'J') || (value == 'L') {
			possibleLocations = append(possibleLocations, location{x: newX, y: newY, val: value})
		}
	}

	//	slices.Reverse(possibleLocations)

	//	for _, loc := range possibleLocations {
	//		keep := validSecConnect(start, loc)
	//		if !keep {
	//			fmt.Println(loc)
	//			ind := slices.Index(possibleLocations, loc)
	//			possibleLocations = slices.Delete(possibleLocations, ind, ind+1)
	//		}
	//	}

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

func createPaths(start location, possibleLocs []location) (paths [][]location) {
	for _, loc := range possibleLocs {
		path := []location{start, loc}
		paths = append(paths, path)
	}
	return paths
}

//TODO: create functions to crawl path returning and int for the farthest point away

func crawlPath(path []location, grid [][]rune) int {
	currSect := path[len(path)-1]

	for !(currSect.val == 'S') {
		currSect = path[len(path)-1]
		prevSect := path[len(path)-2]

		prevDeltaX := currSect.x - prevSect.x
		prevDeltaY := currSect.y - prevSect.y

		currVal := currSect.val

		switch currVal {
		//-
		case '-':
			if prevDeltaX == 1 {
				path = append(path, location{x: currSect.x + 1, y: currSect.y, val: grid[currSect.y][currSect.x+1]})
			} else {
				path = append(path, location{x: currSect.x - 1, y: currSect.y, val: grid[currSect.y][currSect.x-1]})
			}
		//|
		case '|':
			if prevDeltaY == 1 {
				path = append(path, location{x: currSect.x, y: currSect.y + 1, val: grid[currSect.y+1][currSect.x]})
			} else {
				path = append(path, location{x: currSect.x, y: currSect.y - 1, val: grid[currSect.y-1][currSect.x]})
			}
		//L
		case 'L':
			if prevDeltaY == 1 {
				path = append(path, location{x: currSect.x + 1, y: currSect.y, val: grid[currSect.y][currSect.x+1]})
			} else {
				path = append(path, location{x: currSect.x, y: currSect.y - 1, val: grid[currSect.y-1][currSect.x]})
			}
		//J
		case 'J':
			if prevDeltaY == 1 {
				path = append(path, location{x: currSect.x - 1, y: currSect.y, val: grid[currSect.y][currSect.x-1]})
			} else {
				path = append(path, location{x: currSect.x, y: currSect.y - 1, val: grid[currSect.y-1][currSect.x]})
			}
		//F
		case 'F':
			if prevDeltaY == -1 {
				path = append(path, location{x: currSect.x + 1, y: currSect.y, val: grid[currSect.y][currSect.x+1]})
			} else {
				path = append(path, location{x: currSect.x, y: currSect.y + 1, val: grid[currSect.y+1][currSect.x]})
			}
		//7
		case '7':
			if prevDeltaY == -1 {
				path = append(path, location{x: currSect.x - 1, y: currSect.y, val: grid[currSect.y][currSect.x-1]})
			} else {
				path = append(path, location{x: currSect.x, y: currSect.y + 1, val: grid[currSect.y+1][currSect.x]})
			}
		}
	}

	if len(path)%2 == 0 {
		return len(path) / 2
	} else {
		return (len(path) - 1) / 2
	}
}

func solutionOne() {
	lines, _ := readFileToListStrings("./input.txt")
	//fmt.Println(lines)
	l.Println(lines)
	grid := linesToGrid(lines)
	//fmt.Println(grid)
	l.Println(grid)
	start := findStart(grid)
	//fmt.Println(start)
	l.Println(start)
	surroundingStart := surroundingStart(grid, start)
	l.Println(surroundingStart)
	//fmt.Println(surroundingStart)
	paths := createPaths(start, surroundingStart)

	for _, path := range paths {
		ans := crawlPath(path, grid)
		fmt.Println("First answer calculated:", ans)
		l.Println("First answer calculated:", ans)
	}
}

func main() {
	solutionOne()
}
