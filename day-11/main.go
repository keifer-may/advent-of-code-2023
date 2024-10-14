package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	//"strconv"
	//"strings"

	"github.com/keifer-may/advent-of-code-2023/utils"
)

var (
	outfile, _ = os.Create("./log")
	l          = log.New(outfile, "", 0)
)

func expandRow(grid [][]rune) [][]rune {
	indices := []int{}
	insRow := []rune{}
	for i, row := range grid {
		check := emptySet(row)
		if check == true {
			indices = append(indices, i)
			insRow = row
		}
	}

	for i, ind := range indices {
		fmt.Println(ind)
		grid = slices.Insert(grid, ind+i, insRow)
	}
	return grid
}

func expandCol(grid [][]rune) [][]rune {
	indices := []int{}
	for i, _ := range grid[0] {
		set := []rune{}
		for _, row := range grid {
			set = append(set, row[i])
		}
		check := emptySet(set)
		if check == true {
			indices = append(indices, i)
		}
	}

	for i, ind := range indices {
		fmt.Println(ind)
		for j, row := range grid {
			row = slices.Insert(row, ind+i, '.')
			grid[j] = row
		}
	}
	return grid
}

func emptySet(set []rune) bool {
	check := slices.Index(set, '#')
	if check == -1 {
		return true
	} else {
		return false
	}
}

func solutionOne() {
	grid := utils.FileToRuneGrid("./example1.txt")
	fmt.Println(grid, len(grid))
	l.Println(grid)
	grid = expandRow(grid)
	l.Println(grid, len(grid), len(grid[0]), len(grid[len(grid)-1]))
	grid = expandCol(grid)
	l.Println(grid, len(grid), len(grid[0]), len(grid[len(grid)-1]))
	locs := utils.LocItemsInGrid(grid, '#')
	l.Println(locs)
}

func main() {
	solutionOne()
}
