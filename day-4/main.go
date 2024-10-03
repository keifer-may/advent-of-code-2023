package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	lineNumber int
	winningCards int
	countCopy int
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

func removeCardNumber(line string) (cards string) {
	/*
		We are going to take each line, 
		remove the prefix, and 
		return the uncleaned/unsplit cards
		--we'll split that in a different function.
	*/
	cards = strings.Split(line, ": ")[1]
	return
}

func splitWinningAndCard(cards string) (winningNumbers []int, cardsInHandNumbers []int) {
	/*
		We are going to take the cards,
		split them into the winning number string,
		the card number string,
		and then we need to convert them both to ints
		and return them
	*/
	cardStrings := strings.Split(cards, " | ")
	winningStrings := strings.Split(cardStrings[0], " ")
	handStrings := strings.Split(cardStrings[1], " ")

	for _, winningString := range winningStrings {
		winningNumber, _ := strconv.Atoi(winningString)
		if winningNumber == 0 {
			continue
		} else {
			winningNumbers = append(winningNumbers, winningNumber)
		}
	}

	for _, handString := range handStrings {
		handNumber, _ := strconv.Atoi(handString)
		if handNumber == 0 {
			continue
		} else {
			cardsInHandNumbers = append(cardsInHandNumbers, handNumber)
		}
	}
	return
}

func getNumberOfWinningNumbers(winningNumbers []int, handNumbers []int) (numberOfWinningNumbers int) {
	/*
		We will take the winning numbers,
		the numbers in our hand,
		compare the two, and get return the number of matches.
	*/

	for _, winningNumber := range winningNumbers {
		if slices.Contains(handNumbers, winningNumber) {
			// fmt.Println(winningNumber)
			numberOfWinningNumbers++
		}
	}
	return
}

func createCard(winningNumbers []int, handNumbers []int, rowLine int) (card Card) {
	/*
		We are going to take the input from the split card numbers,
		get the number of winning numbers,
		and then create them as card structs.
	*/
	card.lineNumber = rowLine
	card.winningCards = getNumberOfWinningNumbers(winningNumbers, handNumbers)
	card.countCopy = 1
	return
}

func addCardsToSlice(cards []Card, card Card) (cardsSlice []Card) {
	/*
		We are going to take the cards slice, the card struct,
		and append the card to the slice.
	*/
	cardsSlice = append(cards, card)
	return
}

func getCardRec(cards *[]Card) {
	for ind, card := range *cards {
		for i := 0; i < card.countCopy; i++{
			if card.winningCards > 0 {
				for j := 1; j <= card.winningCards; j++ {
					copyCardIndex := ind + j
					(*cards)[copyCardIndex].countCopy++
					
				}
			}
		}
	}
	return
}

func getCountOfCards(cards []Card) (count int) {
	for _, cards := range cards {
		count += cards.countCopy
	}
	return
}

func partOneScoreFormula(numberOfWinningNumbers int) (score int) {
	/*
		We will take the number of winning numbers, and return the score
	*/
	if numberOfWinningNumbers == 0 {
		return
	} else if numberOfWinningNumbers == 1 {
		score = 1
		return
	} else if numberOfWinningNumbers > 1 {
		score = 1
		for i := 1; i < numberOfWinningNumbers; i++ {
			score = score * 2
		}
	}
	return
}

func partOneSolution() {
	lines, err := readFileToListStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	finalScore := 0
	for _, line := range lines {
		rawCards := removeCardNumber(line)
		winningNumbers, cardsInHandNumbers := splitWinningAndCard(rawCards)
		numberOfWinningCards := getNumberOfWinningNumbers(winningNumbers, cardsInHandNumbers)
		lineScore := partOneScoreFormula(numberOfWinningCards)
		finalScore += lineScore
	}
	 fmt.Println("Part one final score:", finalScore)
}

func partTwoSolution() {
	lines, err := readFileToListStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	listCards := []Card{}
	for ind, line := range lines {
		rawCards := removeCardNumber(line)
		winningNumbers, cardsInHandNumbers := splitWinningAndCard(rawCards)
		lineCard := createCard(winningNumbers, cardsInHandNumbers, ind)
		listCards = addCardsToSlice(listCards, lineCard)
	}

	getCardRec(&listCards)

	finalCount := getCountOfCards(listCards)
	fmt.Println("Part two final count of original cards:", finalCount)
}

func main() {
	partOneSolution()
	partTwoSolution()
}
