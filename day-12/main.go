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

//func checkTryAgainstRequirements(try string, require []int) bool {
//	split := strings.Split(try, ".")
//	final := []int{}
//
//	for _, str := range split {
//		if strings.Contains(str, "#") {
//			final = append(final, len(str))
//		}
//	}
//
//	return slices.Equal(final, require.Counts)
//}

//func createTries(items string, req Requirement) {
//	//count := strings.Count(items, "?")
//	//inds := []int{}
//	//fmt.Println(items, req)
//	//removed := strings.ReplaceAll(items, ".", "")
//	//fmt.Println(removed)
//	validCount := 0
//	fmt.Println(items, req)

//	chars := []rune(items)

//	for i, _ := range chars {
//		if i+req.Counts[0] > len(chars)-1 {
//			//return
//			fmt.Println("")
//		} else {
//			firstWind := chars[i : i+req.Counts[0]]
//			nextWindStart := i + req.Counts[0] + 1
//			fmt.Println("First window", string(firstWind), i, nextWindStart)
//			if !(slices.Contains(firstWind, '.')) && !(len(req.Counts) == 1) {
//				for j, winLen := range req.Counts[1:] {
//					if (nextWindStart < len(chars)-1) && (nextWindStart+winLen < len(chars)-1) {
//						subChars := chars[nextWindStart:]
//						for k, _ := range subChars {
//							nextWind := subChars[k : k+winLen]
//							fmt.Println("Next window", string(nextWind), k)
//							if !(j+1 == len(req.Counts)) && !(slices.Contains(nextWind, '.')) {
//								nextWindStart = k + winLen + 1
//								if (nextWindStart < len(subChars)-1) && (nextWindStart+req.Counts[j+1] < len(subChars)-1) {
//									subChars = subChars[nextWindStart:]
//								}
//							} else if j+1 == len(req.Counts) {
//								validCount++
//							}
//						}
//					}
//				}
//			}
//		}
//	}

//	fmt.Println(validCount)
//	for i := 0; i <= count; i++ {

//		broke := strings.Replace(items, "?", "#", i)
//		broke = strings.ReplaceAll(broke, "?", ".")
//		fmt.Println(checkTryAgainstRequirements(broke, req), broke)
//		working := strings.Replace(items, "?", ".", i)
//		working = strings.ReplaceAll(working, "?", "#")
//		fmt.Println(checkTryAgainstRequirements(working, req), working)
//	}
//}
//

func containsWorking(chars []rune) bool {
	if slices.Contains(chars, '.') {
		return false
	} else {
		return true
	}
}

func createWindows(chars []rune, req []int, valid *int) int {
	if len(chars) >= req[0] {
		window := chars[:req[0]]
		if containsWorking(window) {
			if len(req) == 1 {
				*valid = *valid + 1
				fmt.Println(string(window))
				//return 0
			} else if len(chars) > req[0]+1 {
				nextChars := chars[req[0]+1:]
				createWindows(nextChars, req[1:], valid)
				if len(nextChars) >= req[1] {
					nextChars = nextChars[1:]
					nextReq := req[1:]
					fmt.Println(nextReq, string(nextChars))
					createWindows(nextChars, nextReq, valid)
				}
			}
		} else if len(chars) > req[0]+1 {
			nextChars := chars[1:]
			if len(nextChars) >= req[0] {
				createWindows(nextChars, req, valid)
			}
		}
	}
	return 0
}

func processWholeString(chars []rune, req []int, valid *int) {
	for i, _ := range chars {
		createWindows(chars[i:], req, valid)

	}
}

func solutionOne() {
	lines, _ := utils.FileToStringArray("./example1.txt")
	items, req := lineToItemsAndRequirements(lines[3])
	count := 0
	createWindows(items, req, &count)
	fmt.Println(count)

}

func main() {
	fmt.Println("Hello")
	solutionOne()
}
