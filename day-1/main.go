package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

var digitStrings = map[string]string{"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func find_nums_agnostic(line string, regex *regexp.Regexp) int64 {
	var matches []string
	string_copy := line
	int_subtotal := int64(0)

	for index, _ := range string_copy {
		sub_length := 6
		final_char := len(string_copy) - 1
		var substring_bytes string
		if final_char - index == 5 {
			substring_bytes = string_copy[index:]
		} else if index == final_char{
			substring_bytes = string(string_copy[index])
		} else {
			sub_length = final_char+1 - index
			substring_bytes = string_copy[index:index+sub_length]
		}
		substring := string(substring_bytes)
		found_match := regex.FindAllString(substring,-1)
		if len(found_match) > 0{
			for _, match := range found_match{
				// fmt.Println(match)
				matches = append(matches, match)
			}
		}
	}
	if len(matches) > 0 {
		first_num := matches[0]
		last_num := matches[len(matches)-1]

		for k, v := range digitStrings {
			if k == first_num {
				first_num = v
			} 
			if k == last_num {
				last_num = v
			}
		}
		concat := first_num + last_num
		int_subtotal, err := strconv.ParseInt(concat, 10, 64)
		if err != nil {
			fmt.Println(err, line)
		} else {
			// fmt.Println(line,int_subtotal)
			return int_subtotal
		}
	} else {
		fmt.Println(line)
	}
	return int_subtotal
}

func main() {
	start := time.Now()
	readFile, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	
	sum_one := int64(0)
	sum_two := int64(0)
	
	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(line) > 0 {
			reg := regexp.MustCompile(`[0-9]`)
			get_total := find_nums_agnostic(line, reg)
			sum_one = sum_one + get_total

			reg2 := regexp.MustCompile(`[0-9]|one|two|three|four|five|six|seven|eight|nine`)
			get_total2 := find_nums_agnostic(line, reg2)
			sum_two = sum_two + get_total2
		}
	}
	fmt.Printf("Problem 1 solution: %d\n",sum_one)
	fmt.Printf("Problem 2 solution: %d\n",sum_two)
	readFile.Close()
	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed for program: %s\n", timeElapsed)
}
