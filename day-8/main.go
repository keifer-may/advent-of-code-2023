package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	outfile, _ = os.Create("./log")
	l = log.New(outfile, "", 0)
)

// this solution inspired by https://github.com/bsadia/aoc_goLang/blob/master/day08/main.go

//CREATE FUNCTION THAT CREATES THE PATH FROM LINES AND THE INSTRUCTIONS FROM THE LINES 

func processLines(lines []string) (path map[string][2]string, instructions string) {
	instructions = lines[0]
	path = make(map[string][2]string)

	for _, line := range lines[2:] {
		line = strings.Replace(line, "(", "", -1)
		line = strings.Replace(line, ")", "", -1)
		line = strings.Replace(line, "=", "", -1)
		line = strings.Replace(line, ",", " ", -1)
		splitLine := strings.Split(line, "  ")
		path[splitLine[0]] = [2]string{splitLine[1], splitLine[2]}
	}
	return path, instructions
}

//CREATE FUNCTION THAT WILL TAKE PATH, INSTRUCTIONS, START, AND END AND WILL RETURN THE NUMBER OF STEPS TO GET TO THE END

func walkOne(path map[string][2]string, instructions string, start string, end string) (steps int) {
	for !(start == end) {
		for _, direct := range instructions {
			if direct == 'L' {
				start = path[start][0]
			} else {
				start = path[start][1]
			}
			steps++
			if start == end {
				break
			}
		}
	}
	return steps
}

func walkTwo(path map[string][2]string, instructions string, startNodes []string, endNodes []string) (stepsList []int) {
	for _, node := range startNodes {
		steps := 0
		for !(slices.Contains(endNodes, node)) {
			for _, direct := range instructions {
				if direct == 'L' {
					node = path[node][0]
				} else {
					node = path[node][1]
				}
				steps++
				if slices.Contains(endNodes, node) {
					stepsList = append(stepsList, steps)
					break
				}
			}
		}
	}
	return stepsList
}

func getListBySuffix(path map[string][2]string, suffix string) (listNodes []string) {
	for node, _ := range path {
		lastChar := node[2]
		if string(lastChar) == suffix {
			listNodes = append(listNodes, node)
		}
	}
	return listNodes
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

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}

func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / greatestCommonDivisor(a, b)
	l.Println(result)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

func solutionOne() {
	fmt.Println("Running problem 1")
	lines, _ := readFileToListStrings("input.txt")
	l.Println(lines)
	path, insts := processLines(lines)
	l.Println(path, insts)
	answer := walkOne(path, insts, "AAA", "ZZZ")
	fmt.Println("Solution 1:", answer)
}

func solutionTwo() {
	fmt.Println("Running problem 2")
	lines, _ := readFileToListStrings("input.txt")
	l.Println(lines)
	path, insts := processLines(lines)
	l.Println(path, insts)
	allAs := getListBySuffix(path, "A")
	l.Println(allAs)
	allZs := getListBySuffix(path, "Z")
	l.Println(allZs)
	stepList := walkTwo(path, insts, allAs, allZs)
	//fmt.Println(stepList)
	//params := make([]interface{}, 0)
	//for _, int := range stepList{
	//	params = append(params, int)
	//}
	answer := leastCommonMultiple(stepList[0], stepList[1], stepList[2:]...)
	fmt.Println("Solution 2:", answer)
}

func main() {
	solutionOne()
	solutionTwo()
}
