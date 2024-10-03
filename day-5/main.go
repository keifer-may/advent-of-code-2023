package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Seed struct {
	number int64
}

type SeedRanges struct {
	start int64
	end   int64
}

type Directions struct {
	destinationStart int64
	sourceStart      int64
	sourceEnd        int64
	difference       int64
	rangeLength      int64
	order            int
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

func getSeeds(lines []string) (seeds []Seed) {
	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		if len(splitLine) > 1 {
			if splitLine[0] == "seeds" {
				seedsString := splitLine[1]
				seedsString = strings.Trim(seedsString, " ")
				splitSeedString := strings.Split(seedsString, " ")
				for _, seedString := range splitSeedString {
					seedInt, _ := strconv.ParseInt(seedString, 10, 64)
					seed := Seed{
						number: seedInt,
					}
					seeds = append(seeds, seed)
				}
			}
		}
	}
	return seeds
}

func getSeedRanges(lines []string) (seeds []Seed) {
	beginningOfRanges := []int64{}
	lengthOfRange := []int64{}
	for _, line := range lines {
		splitLine := strings.Split(line, ":")
		if len(splitLine) > 1 {
			if splitLine[0] == "seeds" {
				seedsString := splitLine[1]
				seedsString = strings.Trim(seedsString, " ")
				splitSeedString := strings.Split(seedsString, " ")
				for ind, seedString := range splitSeedString {
					if ind%2 == 0 {
						firstSeed, _ := strconv.ParseInt(seedString, 10, 64)
						beginningOfRanges = append(beginningOfRanges, firstSeed)
					} else {
						rangeLength, _ := strconv.ParseInt(seedString, 10, 64)
						lengthOfRange = append(lengthOfRange, rangeLength)
					}
				}
			}
		}
	}

	for i, seed := range beginningOfRanges {
		rangeLength := lengthOfRange[i]
		// seeds = append(seeds, Seed{seed})
		for j := int64(0); j < rangeLength; j++ {
			newSeed := seed + j
			seeds = append(seeds, Seed{newSeed})
		}
	}
	return seeds
}

func getDirections(lines []string) (locations [][]Directions) {
	listMapBegin := []int{}
	listMapEnd := []int{}

	for ind, line := range lines {
		if ind == 0 || ind == 1 {
			continue
		} else {
			if strings.Contains(line, "map") {
				// fmt.Println(line)
				listMapBegin = append(listMapBegin, ind+1)
			} else if len(line) == 0 {
				listMapEnd = append(listMapEnd, ind)
			} else if ind == len(lines)-1 {
				listMapEnd = append(listMapEnd, ind+1)
			}
		}
	}

	for ord, begin := range listMapBegin {
		directions := []Directions{}
		end := listMapEnd[ord]
		for i := begin; i < end; i++ {
			line := strings.Split(lines[i], " ")
			if len(line) == 3 {
				destinationStart, _ := strconv.ParseInt(line[0], 10, 64)
				sourceStart, _ := strconv.ParseInt(line[1], 10, 64)
				rangeLength, _ := strconv.ParseInt(line[2], 10, 64)
				difference := sourceStart - destinationStart
				sourceEnd := sourceStart + rangeLength
				order := ord
				directions = append(directions, Directions{destinationStart, sourceStart, sourceEnd, difference, rangeLength, order})
			}
			if i == end-1 {
				locations = append(locations, directions)
			}
		}
	}
	return locations
}

func getFinalLocations(locationDirections [][]Directions, seeds []Seed) (finalLocations []int64) {
	for _, seed := range seeds {
		location := seed.number
		for _, directionGroup := range locationDirections {
			for _, direction := range directionGroup {
				if location >= direction.sourceStart && location <= direction.sourceEnd {
					location = location - direction.difference
					break
				}
			}
		}
		finalLocations = append(finalLocations, location)
	}
	return finalLocations
}

func getLowestLocation(finalLocations []int64) (lowestLocation int64) {
	lowestLocation = finalLocations[0]
	for _, location := range finalLocations {
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}

func solutionOne() {
	lines, err := readFileToListStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	seeds := getSeeds(lines)

	listAllDirections := getDirections(lines)

	listFinalLocations := getFinalLocations(listAllDirections, seeds)

	lowestLocation := getLowestLocation(listFinalLocations)
	fmt.Println("Answer to part one is:", lowestLocation)
}

func solutionTwo() {
	lines, err := readFileToListStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	seeds := getSeedRanges(lines)
	fmt.Println(seeds)

	listAllDirections := getDirections(lines)

	listFinalLocations := getFinalLocations(listAllDirections, seeds)

	lowestLocation := getLowestLocation(listFinalLocations)
	fmt.Println("Answer to part two is:", lowestLocation)
}

func main() {
	solutionOne()
	solutionTwo()
}
