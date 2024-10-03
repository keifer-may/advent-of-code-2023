package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func get_draws(s string) map[string]int {
	draws := make(map[string]int)
	for _, draw := range strings.Split(s, ", ") {
		split_draw := strings.Split(draw, " ")
		count, err := strconv.Atoi(split_draw[0])
		if err != nil {
			fmt.Println(err)
			break
		}
		draws[split_draw[1]] = count
	}
	return draws
}

func get_games(line string) []string {
	games := strings.Split(line, ": ")[1]
	list_games := strings.Split(games, "; ")
	return list_games
}

func get_game_id(line string) int {
	prefix := strings.Split(line, ": ")[0]
	index_string := strings.Split(prefix, " ")[1]
	index, err := strconv.Atoi(index_string)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return index
}

func check_one(color string, count int, require map[string]int) bool {
	game_possible := true
	for group, max_val := range require {
		if color == group && count > max_val {
			game_possible = false
			break
		}
	}
	return game_possible
}

func check_two(color string, count int, require map[string]int) {
	for group, max_val := range require {
		if color == group && count > max_val {
			require[group] = count
		}
	}
	return
}

func get_product(maxes map[string]int) int {
	product := 1
	for _, max_val := range maxes {
		product *= max_val
	}
	return product
}

func main() {
	start := time.Now()

	readFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	count_one := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		game_id := get_game_id(line)
		games := get_games(line)
		game_possible_one := true
		for _, game := range games {
			draws := get_draws(game)
			for color, count := range draws {
				first_criteria := map[string]int{"red": 12, "green": 13, "blue": 14}

				game_possible_one = check_one(color, count, first_criteria)
				if !game_possible_one {
					break
				}
			}
			if !game_possible_one {
				break
			}
		}
		if game_possible_one {
			count_one += game_id
		}
	}

	readFile2, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile2.Close()

	fileScanner2 := bufio.NewScanner(readFile2)
	fileScanner2.Split(bufio.ScanLines)

	count_two := 0
	for fileScanner2.Scan() {
		line := fileScanner2.Text()
		games := get_games(line)
		second_maxes := map[string]int{"red": 1, "green": 1, "blue": 1}
		for _, game := range games {
			draws := get_draws(game)
			for color, count := range draws {
				check_two(color, count, second_maxes)
			}
		}
		product := get_product(second_maxes)
		count_two += product
	}

	fmt.Printf("The sum of IDs for problem one: %d \nThe sum of products for problem two: %d \n", count_one, count_two)
	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed for program: %s\n", timeElapsed)
}
