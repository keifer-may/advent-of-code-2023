package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"slices"
	"strconv"
	"strings"
)

var (
	outfile, _ = os.Create("./log")
	l = log.New(outfile, "", 0)
)

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

func lineToSlice(line string) (nums []int) {
	splitLine := strings.Split(line, " ")
	for _, str := range splitLine {
		num, err := strconv.Atoi(str)
		if err == nil {
			nums = append(nums, num)
		}
	}
	return nums
}

func expandSlice(nums []int) (expanded [][]int) {
	expanded = append(expanded, nums)

	for !(checkForZero(expanded[len(expanded)-1])) {
		new_list := []int{}
		for i, num := range expanded[len(expanded)-1]{
			if i == len(expanded[len(expanded)-1]) - 1 {
				expanded = append(expanded, new_list)
				break
			} else {
				next_num := expanded[len(expanded)-1][i+1]
				new_list = append(new_list, (next_num - num))
			}
		}
	}
	return expanded
}

func contractSliceForward(expanded [][]int) (nums []int) {
	for i := len(expanded) - 2; i >= 0; i-- {
		if i == 0 {
			break
		} else {
			//next_row := expanded[i - 1]
			row := expanded[i]
			expanded[i - 1] = append(expanded[i-1], expanded[i-1][len(expanded[i-1])-1] + row[len(row)-1])
		}
	}
	nums = expanded[0]
	return nums
}

func contractSliceBackward(expanded [][]int) (nums []int) {
	for i := len(expanded) - 2; i >= 0; i-- {
		if i == 0 {
			break
		} else {
			next_row := expanded[i - 1]
			row := expanded[i]
			new_slice := []int{expanded[i-1][0]-row[0]}
			expanded[i - 1] = append(new_slice, next_row...)
		}
	}
	nums = expanded[0]
	return nums
}

/*
Alright.

So our task is to be able to take hold of the fucking input and extrapolate the next number in the array that we have on each line.

I think that this will end up mean having 2 main functions that we are going to use:

1 will need to telescope down, adding new slices into a slice of slices, each new list calculating the difference between the values of the last array to make up the values of the new list

2 will need to be able to add the last value in the first array plus the last value in the second array -- this will be our projected new value.

3 we need to take this new projected value and be able to add all of those new projected values up

The total of the example should be 114 when done correctly
*/

func checkForZero(row []int) bool {
	for _, num := range row {
		if num != 0 {
			return false
		}
	}
	return true
}

func problemOne() {
	lines, _ := readFileToListStrings("./input.txt")
	sum := 0
	for _, line := range lines {
		ints := lineToSlice(line)
		expanded := expandSlice(ints)
		contracted := contractSliceForward(expanded)
		sum += contracted[len(contracted) - 1]
		l.Println(ints, expanded, contracted, sum)
	}
	fmt.Println("Solution one:", sum)
}

func problemTwo() {
	lines, _ := readFileToListStrings("./input.txt")
	sum := 0
	for _, line := range lines {
		ints := lineToSlice(line)
		expanded := expandSlice(ints)
		contracted := contractSliceBackward(expanded)
		sum += contracted[0]
		l.Println(ints,expanded,contracted,sum)
	}
	fmt.Println("Solution two:", sum)
}

func main() {
	l.Println("Hello world")
	l.Println("The solution to the example should be 114")
	problemOne()
	problemTwo()
}

