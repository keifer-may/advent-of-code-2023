package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/keifer-may/advent-of-code-2023/utils"
)

var (
	outfile, _ = os.Create("./log")
	l          = log.New(outfile, "", 0)
)

func lineToItemsAndRequirements(line string) (items []rune, require []int) {
	split := strings.Split(line, " ")
	str := split[0]
	items = []rune(str)
	requireStr := split[1]
	requireList := strings.Split(requireStr, ",")
	for _, req := range requireList {
		num, _ := strconv.Atoi(req)
		require = append(require, int(num))
	}
	return items, require
}

func notWorking(chars []rune) bool {
	if slices.Contains(chars, '.') {
		return false
	} else {
		return true
	}
}

func checkWindows(chars []rune, req []int, valid *int) {
	enoughChars := len(chars) >= req[0]
	if enoughChars {

		window := chars[:req[0]]
		firstChar := chars[0]
		oneReq := len(req) == 1
		moreChars := len(chars) > req[0]

		if notWorking(window) && oneReq {
			*valid = *valid + 1
			if moreChars {
				checkWindows(chars[1:], req, valid)
			}
		} else if (firstChar == '#') && moreChars && (chars[req[0]] == '#') {
			checkWindows(chars[2:], req, valid)
		} else if notWorking(window) && !(oneReq) {
			if moreChars {
				if window[len(window)-1] == '#' {
					checkWindows(chars[2:], req, valid)
				} else {
					checkWindows(chars[1:], req, valid)
				}
				if !(chars[req[0]] == '#') {
					checkWindows(chars[req[0]+1:], req[1:], valid)
				}
			}
		} else if !(notWorking(window)) && moreChars {
			checkWindows(chars[1:], req, valid)
		}
	}
}

func solutionOne() {
	lines, _ := utils.FileToStringArray("./input.txt")
	//items, req := lineToItemsAndRequirements(lines[0])
	//count := 0
	//checkWindows(items, req, &count)
	//fmt.Println(count, lines[0])
	sum := 0
	for _, line := range lines {
		items, req := lineToItemsAndRequirements(line)
		count := 0
		checkWindows(items, req, &count)
		l.Println(count, line)
		sum += count
	}
	fmt.Println("Final answer one:", sum)
}

func main() {
	solutionOne()
}
