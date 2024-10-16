package main

import (
	"fmt"
	"github.com/keifer-may/advent-of-code-2023/utils"
	"log"
	"math"
	"os"
	"slices"
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

func calculateDistance(galOne utils.Location, galTwo utils.Location) int {
	xDistance := float64(galTwo.X - galOne.X)
	yDistance := float64(galTwo.Y - galOne.Y)
	totalDist := math.Abs(xDistance) + math.Abs(yDistance)
	return int(totalDist)
}

func processAllPairs(galaxies []utils.Location) (allDist int) {
	for i, galaxy := range galaxies {
		if i == len(galaxies)-1 {
			break
		} else {
			for _, nextGal := range galaxies[i+1:] {
				distance := calculateDistance(galaxy, nextGal)
				allDist += distance
			}
		}
	}
	return allDist
}

func calcDistArray(galaxies []utils.Location) (distances []int) {
	for i, galaxy := range galaxies {
		if i == len(galaxies)-1 {
			break
		} else {
			for _, nextGal := range galaxies[i+1:] {
				distance := calculateDistance(galaxy, nextGal)
				distances = append(distances, distance)
			}
		}
	}
	return distances
}

func deltaDistArrays(distArrOne []int, distArrTwo []int) (deltas []int) {
	for i, dist := range distArrOne {
		secDist := distArrTwo[i]
		diff := math.Abs(float64(secDist - dist))
		deltas = append(deltas, int(diff))
	}
	return deltas
}

func calcBasedOnDeltas(distArr []int, deltas []int, multiple int) int {
	total := 0
	for i, dist := range distArr {
		sub := dist + (deltas[i] * multiple)
		total += sub
	}
	return total
}

func solutionOne() {
	grid := utils.FileToRuneGrid("./input.txt")
	l.Println(grid, len(grid))
	l.Println(grid)
	grid = expandRow(grid)
	l.Println(grid, len(grid), len(grid[0]), len(grid[len(grid)-1]))
	grid = expandCol(grid)
	l.Println(grid, len(grid), len(grid[0]), len(grid[len(grid)-1]))
	locs := utils.LocItemsInGrid(grid, '#')
	l.Println(locs)
	finalAns := processAllPairs(locs)
	l.Println(finalAns)
	fmt.Println("Solution one:", finalAns)
}

func solutionTwo() {
	grid := utils.FileToRuneGrid("./input.txt")
	locs := utils.LocItemsInGrid(grid, '#')
	initDists := calcDistArray(locs)
	grid = expandRow(grid)
	grid = expandCol(grid)
	locsTwo := utils.LocItemsInGrid(grid, '#')
	secDists := calcDistArray(locsTwo)
	deltas := deltaDistArrays(initDists, secDists)
	answer := calcBasedOnDeltas(initDists, deltas, 999999)
	fmt.Println("Solution two:", answer)
}

func main() {
	solutionOne()
	solutionTwo()
}
