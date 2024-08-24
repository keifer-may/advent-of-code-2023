package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

type Race64 struct {
	time     int64
	distance int64
}

type Option struct {
	holdTime   int
	travelTime int
	distance   int
}

type Option64 struct {
	holdTime   int64
	travelTime int64
	distance   int64
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

func getRaces(lines []string) (races []Race) {
	times := []int{}
	distances := []int{}
	for ind, line := range lines {
		if ind == 0 {
			timeString := strings.Split(line, ":")[1]
			timeString = strings.TrimSpace(timeString)
			timeStrings := strings.Split(timeString, " ")
			for _, timeString := range timeStrings {
				time, err := strconv.Atoi(timeString)
				if err != nil {
					continue
				}
				times = append(times, time)
			}
		} else if ind == 1 {
			distanceString := strings.Split(line, ":")[1]
			distanceString = strings.TrimSpace(distanceString)
			distanceStrings := strings.Split(distanceString, " ")
			for _, distanceString := range distanceStrings {
				distance, err := strconv.Atoi(distanceString)
				if err != nil {
					continue
				}
				distances = append(distances, distance)
			}
		}
	}

	for ind, time := range times {
		distance := distances[ind]
		race := Race{time, distance}
		races = append(races, race)
	}
	return races
}

func getSingleRace(lines []string) (race Race64) {
	time := int64(0)
	distance := int64(0)
	for ind, line := range lines {
		if ind == 0 {
			timeString := strings.Split(line, ":")[1]
			timeString = strings.TrimSpace(timeString)
			timeString = strings.Replace(timeString, " ", "", -1)
			time, _ = strconv.ParseInt(timeString, 10, 64)

		} else if ind == 1 {
			distanceString := strings.Split(line, ":")[1]
			distanceString = strings.TrimSpace(distanceString)
			distanceString = strings.Replace(distanceString, " ", "", -1)
			distance, _ = strconv.ParseInt(distanceString, 10, 64)
		}
	}
	race = Race64{time, distance}
	return race
}

func getRaceOptions(race Race) (options []Option) {
	for i := 1; i < race.time; i++ {
		option := Option{
			holdTime:   i,
			travelTime: race.time - i,
			distance:   i * (race.time - i),
		}
		options = append(options, option)
	}
	return options
}

func getRace64Options(race Race64) (options []Option64) {
	for i := int64(1); i < race.time; i++ {
		option := Option64{
			holdTime:   i,
			travelTime: race.time - i,
			distance:   i * (race.time - i),
		}
		options = append(options, option)
	}
	return options
}

func processRace64(race Race64) (count int64) {
	for i := int64(1); i < race.time; i++ {
		travelTime := race.time - i
		distance := i * travelTime
		if distance >= race.distance {
			count += 1
		} else {
			continue
		}
	}
	return count
}

func getWinningRaceOptions(race Race, options []Option) (count int) {
	for _, option := range options {
		if option.distance > race.distance {
			count++
		}
	}
	return count
}

func getWinningRace64Options(race Race64, options []Option64) (count int64) {
	for _, option := range options {
		if option.distance > race.distance {
			count++
		}
	}
	return count
}

func processRaces(races []Race) (counts []int) {
	for _, race := range races {
		options := getRaceOptions(race)
		count := getWinningRaceOptions(race, options)
		counts = append(counts, count)
	}
	return counts
}

func multiplyCounts(counts []int) (total int) {
	total = 1
	for _, count := range counts {
		total *= count
	}
	return total
}

func solutionOne() {
	lines, err := readFileToListStrings("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lines)

	races := getRaces(lines)
	// fmt.Println(races)

	// raceOneOptions := getRaceOptions(races[0])
	// fmt.Println(raceOneOptions)

	// winningOptionCount := getWinningRaceOptions(races[0], raceOneOptions)
	// fmt.Println(winningOptionCount)

	winningCounts := processRaces(races)
	// fmt.Println(winningCounts)

	totalProduct := multiplyCounts(winningCounts)

	fmt.Println("Solution 1:", totalProduct)
}

func solutionTwo() {
	lines, err := readFileToListStrings("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lines)

	race := getSingleRace(lines)
	fmt.Println(race)

	// options := getRace64Options(race)
	winningOptions := processRace64(race)

	fmt.Println("Solution 2:", winningOptions)
}

func main() {
	fmt.Println("Hello, World!")
	solutionOne()
	solutionTwo()
}
