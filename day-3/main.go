package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Part struct {
	number        int
	startColIndex int
	endColIndex   int
	rowIndex      int
	include       bool
}

type Symbol struct {
	sym           string
	startColIndex int
	endColIndex   int
	rowIndex      int
}

type Gear struct {
	firstAdj  int
	secondAdj int
	colIndex  int
	rowIndex  int
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

func getSymbols(line string, rowIndex int) (symbols []Symbol) {
	replacer := strings.NewReplacer("0", ".", "1", ".", "2", ".", "3", ".", "4", ".", "5", ".", "6", ".", "7", ".", "8", ".", "9", ".")
	syms := []string{}
	indices := []int{}

	newString := replacer.Replace(line)

	splitLine := strings.Split(newString, ".")
	for ind, item := range splitLine {
		if item == "" {
			continue
		} else {
			syms = append(syms, item)
			indices = append(indices, ind)
		}
	}

	if len(syms) == 0 {
		return symbols
	}

	newLine := line
	for ind, item := range syms {
		index := strings.Index(newLine, fmt.Sprint(item))
		endNum := index + len(fmt.Sprint(item))
		newLine = newLine[endNum:]
		if ind > 0 {
			lastInd := symbols[ind-1].endColIndex + 1
			symbols = append(symbols,
				Symbol{
					sym:           item,
					startColIndex: index + lastInd,
					endColIndex:   index + lastInd + len(fmt.Sprint(item)) - 1,
					rowIndex:      rowIndex,
				})
		} else {
			lastInd := 0
			symbols = append(symbols,
				Symbol{
					sym:           item,
					startColIndex: index + lastInd,
					endColIndex:   index + lastInd + len(fmt.Sprint(item)) - 1,
					rowIndex:      rowIndex,
				})
		}
	}

	return symbols
}

func getPossibleGears(line string, rowIndex int) (gears []Gear) {
	replacer := strings.NewReplacer("0", ".", "1", ".", "2", ".", "3", ".", "4", ".", "5", ".", "6", ".", "7", ".", "8", ".", "9", ".", "!", ".", "@", ".", "#", ".", "$", ".", "%", ".", "^", ".", "&", ".", "(", ".", ")", ".", "_", ".", "+", ".", "-", ".", "=", ".", "/", ".")
	syms := []string{}
	indices := []int{}

	newString := replacer.Replace(line)

	splitLine := strings.Split(newString, ".")
	for ind, item := range splitLine {
		if item == "" {
			continue
		} else {
			syms = append(syms, item)
			indices = append(indices, ind)
		}
	}

	if len(syms) == 0 {
		return gears
	}

	newLine := line
	for ind, item := range syms {
		index := strings.Index(newLine, fmt.Sprint(item))
		endNum := index + len(fmt.Sprint(item))
		newLine = newLine[endNum:]
		if ind > 0 {
			lastInd := gears[ind-1].colIndex + 1
			gears = append(gears,
				Gear{
					firstAdj:      0,
					secondAdj:     0,
					colIndex: index + lastInd,
					rowIndex:      rowIndex,
				})
		} else {
			lastInd := 0
			gears = append(gears,
				Gear{
					firstAdj:      0,
					secondAdj:     0,
					colIndex: index + lastInd,
					rowIndex:      rowIndex,
				})
		}
	}

	return gears
}

func getNums(line string, rowIndex int) (parts []Part) {
	replacer := strings.NewReplacer("!", " ", "@", " ", "#", " ", "$", " ", "%", " ", "^", " ", "&", " ", "*", " ", "(", " ", ")", " ", "_", " ", "+", " ", "-", " ", "=", " ", "/", " ", ".", " ")
	nums := []int{}

	newString := replacer.Replace(line)
	splitLine := strings.Split(newString, " ")

	for _, item := range splitLine {
		if item == "" {
			continue
		} else {

			num, err := strconv.Atoi(item)
			if err != nil {
				continue
			} else {
				nums = append(nums, num)
			}
		}
	}

	if len(nums) == 0 {
		return parts
	}

	newLine := line
	for ind, item := range nums {
		index := strings.Index(newLine, fmt.Sprint(item))
		endNum := index + len(fmt.Sprint(item))
		newLine = newLine[endNum:]
		if ind > 0 {
			lastInd := parts[ind-1].endColIndex + 1
			parts = append(parts,
				Part{
					number:        item,
					startColIndex: index + lastInd,
					endColIndex:   index + lastInd + len(fmt.Sprint(item)) - 1,
					rowIndex:      rowIndex,
				})
		} else {
			lastInd := 0
			parts = append(parts,
				Part{
					number:        item,
					startColIndex: index + lastInd,
					endColIndex:   index + lastInd + len(fmt.Sprint(item)) - 1,
					rowIndex:      rowIndex,
				})
		}
	}

	// fmt.Println(parts, "\n")
	return parts
}

func getFinalListNums(finalList *[]Part, rowList []Part) {
	if len(rowList) == 0 {
		return
	} else {
		for _, part := range rowList {
			*finalList = append(*finalList, part)
		}
	}
}

func getFinalListSymbols(finalList *[]Symbol, rowList []Symbol) {
	if len(rowList) == 0 {
		return
	} else {
		for _, sym := range rowList {
			*finalList = append(*finalList, sym)
		}
	}
}

func getFinalListPossibleGears(finalList *[]Gear, rowList []Gear) {
	if len(rowList) == 0 {
		return
	} else {
		for _, gear := range rowList {
			*finalList = append(*finalList, gear)
		}
	}
}

func checkLeft(finalPartList []Part, finalSymbolList []Symbol) (checkedPartList []Part) {
	for _, part := range finalPartList {
		leftSymbolCheckColInd := part.startColIndex - 1
		leftSymbolCheckRowInd := part.rowIndex
		for _, sym := range finalSymbolList {
			if sym.rowIndex == leftSymbolCheckRowInd {
				if sym.endColIndex == leftSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				}
			}
		}
	}
	return checkedPartList
}

func checkRight(finalPartList []Part, finalSymbolList []Symbol) (checkedPartList []Part) {
	for _, part := range finalPartList {
		rightSymbolCheckColInd := part.endColIndex + 1
		rightSymbolCheckRowInd := part.rowIndex
		for _, sym := range finalSymbolList {
			if sym.rowIndex == rightSymbolCheckRowInd {
				if sym.startColIndex == rightSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				}
			}
		}
	}
	return checkedPartList
}

func checkUp(finalPartList []Part, finalSymbolList []Symbol) (checkedPartList []Part) {
	for _, part := range finalPartList {
		leftSymbolCheckColInd := part.startColIndex - 1
		rightSymbolCheckColInd := part.endColIndex + 1
		checkRowInd := part.rowIndex - 1
		for _, sym := range finalSymbolList {
			if sym.rowIndex == checkRowInd {
				if sym.endColIndex >= leftSymbolCheckColInd && sym.endColIndex <= rightSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				} else if sym.startColIndex >= leftSymbolCheckColInd && sym.startColIndex <= rightSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				}
			}
		}
	}
	return checkedPartList
}

func checkDown(finalPartList []Part, finalSymbolList []Symbol) (checkedPartList []Part) {
	for _, part := range finalPartList {
		leftSymbolCheckColInd := part.startColIndex - 1
		rightSymbolCheckColInd := part.endColIndex + 1
		checkRowInd := part.rowIndex + 1
		for _, sym := range finalSymbolList {
			if sym.rowIndex == checkRowInd {
				if sym.endColIndex >= leftSymbolCheckColInd && sym.endColIndex <= rightSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				} else if sym.startColIndex >= leftSymbolCheckColInd && sym.startColIndex <= rightSymbolCheckColInd {
					newPart := part
					newPart.include = true
					checkedPartList = append(checkedPartList, newPart)
					continue
				}
			}
		}
	}
	return checkedPartList
}

func getSum(finalPartList []Part) (sum int) {
	for _, part := range finalPartList {
		if part.include == true {
			sum += part.number
		}
	}
	return sum
}

func cleanList(partList []Part) (cleanList []Part) {

	for _, part := range partList {
		if part.include == true {
			cleanList = append(cleanList, part)
		}
	}
	// fmt.Println(cleanList)
	return cleanList
}

func removeDuplicates(finalList []Part) (cleanedList []Part) {
	trueValues := make(map[Part]bool)
	for _, part := range finalList {
		trueValues[part] = true
	}

	for part, _ := range trueValues {
		cleanedList = append(cleanedList, part)
	}

	return cleanedList
}

func countDuplicatePartNums(finalList []Part) (duplicatePartNums []int) {
	listNums := []int{}
	for _, part := range finalList {
		listNums = append(listNums, part.number)
	}

	size := len(listNums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if listNums[i] == listNums[j] {
				duplicatePartNums = append(duplicatePartNums, listNums[i])
			}
		}
	}

	fmt.Println(duplicatePartNums)
	return duplicatePartNums
}

func appendPartLists(list1 []Part, list2 []Part, list3 []Part, list4 []Part) (fullTruePartList []Part) {
	for _, part := range list1 {
		fullTruePartList = append(fullTruePartList, part)
	}
	for _, part := range list2 {
		if contained := slices.Contains(fullTruePartList, part); contained == false {
			fullTruePartList = append(fullTruePartList, part)
		}
	}
	for _, part := range list3 {
		if contained := slices.Contains(fullTruePartList, part); contained == false {
			fullTruePartList = append(fullTruePartList, part)
		}
	}
	for _, part := range list4 {
		if contained := slices.Contains(fullTruePartList, part); contained == false {
			fullTruePartList = append(fullTruePartList, part)
		}
	}
	return
}

func checkGears(partList []Part, gearList []Gear) (finalGearList []Gear) {
	for _, gear := range gearList {
		listAdjacent := []int{}
		rowAbove := gear.rowIndex - 1
		rowBelow := gear.rowIndex + 1
		colLeft := gear.colIndex - 1
		colRight := gear.colIndex + 1
		for _, part := range partList {
			if part.rowIndex == rowAbove || part.rowIndex == rowBelow {
				if part.endColIndex >= colLeft && part.endColIndex <= colRight {
					listAdjacent = append(listAdjacent, part.number)
				} else if part.startColIndex >= colLeft && part.startColIndex <= colRight {
					listAdjacent = append(listAdjacent, part.number)
				}
			} else if part.rowIndex == gear.rowIndex {
				if part.endColIndex == colLeft || part.startColIndex == colRight{
					listAdjacent = append(listAdjacent, part.number)
				}
			}
		}
		if len(listAdjacent) == 2 {
			gear.firstAdj = listAdjacent[0]
			gear.secondAdj = listAdjacent[1]
			finalGearList = append(finalGearList, gear)
		}
	}
	// fmt.Println(finalGearList)
	return
}

func gearSum(finalGearList []Gear) (sum int) {
	for _, gear := range finalGearList {
		product := gear.firstAdj * gear.secondAdj
		sum += product
	}
	return
}

func solutionOne() {
	lines, err := readFileToListStrings("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lines)

	finalPartList := []Part{}
	finalSymbolList := []Symbol{}

	for rowIndex, line := range lines {
		numsList := getNums(line, rowIndex)
		getFinalListNums(&finalPartList, numsList)

		symbolsList := getSymbols(line, rowIndex)
		getFinalListSymbols(&finalSymbolList, symbolsList)
	}

	// fmt.Println(finalPartList)
	// fmt.Println(finalSymbolList)

	// fmt.Println("Final part list:", finalPartList)
	// fmt.Println("Final symbol list:", finalSymbolList)
	totalPartList := []Part{}

	leftCheckedPartList := checkLeft(finalPartList, finalSymbolList)

	rightCheckedPartList := checkRight(finalPartList, finalSymbolList)

	topCheckedPartList := checkUp(finalPartList, finalSymbolList)

	downCheckedPartList := checkDown(finalPartList, finalSymbolList)

	totalPartList = appendPartLists(leftCheckedPartList, rightCheckedPartList, downCheckedPartList, topCheckedPartList)

	// fmt.Println("Total list:", totalPartList)
	// fmt.Println("Left list:", leftCheckedPartList)
	// fmt.Println("Right list:", rightCheckedPartList)
	// fmt.Println("Down list:", downCheckedPartList)
	// fmt.Println("Top list:", topCheckedPartList)

	finalSum := getSum(totalPartList)

	fmt.Println("Final sum one:", finalSum)
}

func solutionTwo() {
	lines, err := readFileToListStrings("./input.txt")

	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lines)

	finalPartList := []Part{}
	finalPossibleGearList := []Gear{}

	for rowIndex, line := range lines {
		numsList := getNums(line, rowIndex)
		getFinalListNums(&finalPartList, numsList)

		possibleGearsList := getPossibleGears(line, rowIndex)
		getFinalListPossibleGears(&finalPossibleGearList, possibleGearsList)
	}

	// fmt.Println(finalPartList)
	// fmt.Println(finalPossibleGearList)

	finalGears := checkGears(finalPartList, finalPossibleGearList)

	// fmt.Println(finalGears)

	getAnswer := gearSum(finalGears)

	fmt.Println("Final sum two:", getAnswer)
}

func main() {
	solutionOne()
	solutionTwo()
}
