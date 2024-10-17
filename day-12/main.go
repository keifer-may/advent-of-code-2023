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

type Requirement struct {
	Counts []int
}

var (
	outfile, _ = os.Create("./log")
	l          = log.New(outfile, "", 0)
)

func lineToItemsAndRequirements(line string) (items string, require Requirement) {
	split := strings.Split(line, " ")
	items = split[0]
	requireStr := split[1]
	requireList := strings.Split(requireStr, ",")
	reqs := []int{}
	for _, req := range requireList {
		num, _ := strconv.Atoi(req)
		reqs = append(reqs, int(num))
	}
	require = Requirement{Counts: reqs}
	return items, require
}

func checkTryAgainstRequirements(try string, require Requirement) bool {
	split := strings.Split(try, ".")
	final := []int{}

	for _, str := range split {
		if strings.Contains(str, "#") {
			final = append(final, len(str))
		}
	}

	return slices.Equal(final, require.Counts)
}

func createTries(items string, req Requirement) {
	count := strings.Count(items, "?")
	//inds := []int{}

	for i := 0; i <= count; i++ {

		broke := strings.Replace(items, "?", "#", i)
		broke = strings.ReplaceAll(broke, "?", ".")
		fmt.Println(checkTryAgainstRequirements(broke, req), broke)
		working := strings.Replace(items, "?", ".", i)
		working = strings.ReplaceAll(working, "?", "#")
		fmt.Println(checkTryAgainstRequirements(working, req), working)
	}
}

func solutionOne() {
	lines, _ := utils.FileToStringArray("./example1.txt")
	l.Println(lines)
	try := "...###...##.#"
	items, req := lineToItemsAndRequirements(lines[0])
	createTries(items, req)
	//_:= Requirement{Counts: []int{5, 2, 1}}
	check := checkTryAgainstRequirements(try, req)
	fmt.Println(check)
}

func main() {
	fmt.Println("Hello")
	solutionOne()
}
