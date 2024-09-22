package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func walkOne(path map[string][2]string, instructions string, start string, end string) (steps uint16) {
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

func solutionOne() {
	fmt.Println("Running problem 1")
	lines, _ := readFileToListStrings("input.txt")
	l.Println(lines)
	path, insts := processLines(lines)
	l.Println(path, insts)
	answer := walkOne(path, insts, "AAA", "ZZZ")
	fmt.Println("Solution 1:", answer)
}

func main() {
	solutionOne()
}
