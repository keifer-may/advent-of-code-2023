package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards    []uint8
	bid      uint8
	handRank string
	handRankValue uint8
}

var rankMap = map[string]uint8{
	"High card":1,
	"One pair":2,
	"Two pair":3,
	"Trips":4,
	"Full house":5,
	"Quads":6,
	"Five":7,
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

func createHand(line string) (hand Hand) {
	splitString := strings.Split(line, " ")
	cards := splitString[0]
	bidString := splitString[1]
	bid, _ := strconv.ParseUint(bidString, 10, 8)
	hand.bid = uint8(bid)

	for _, card := range cards {
		switch card {
		case 'A':
			hand.cards = append(hand.cards, 14)
		case 'K':
			hand.cards = append(hand.cards, 13)
		case 'Q':
			hand.cards = append(hand.cards, 12)
		case 'J':
			hand.cards = append(hand.cards, 11)
		case 'T':
			hand.cards = append(hand.cards, 10)
		default:
			intcard, _ := strconv.ParseUint(string(card), 10, 8)
			hand.cards = append(hand.cards, uint8(intcard))
		}
	}
	return hand
}

func rankHand(hand Hand) (newHand Hand){
	newHand = hand
	cards := make([]uint8, len(hand.cards))
	_ = copy(cards, hand.cards)
	slices.Sort(cards)

	uniqueCards := []uint8{}
	for _, card := range cards {
		if slices.Contains(uniqueCards, card) {
			continue
		} else {
			uniqueCards = append(uniqueCards, card)
		}
	}

	cardCounts := make(map[uint8]uint8)
	for _, card := range uniqueCards {
		cardCounts[card] = uint8(0)
		for _, card2 := range cards {
			if card2 == card {
				cardCounts[card]++
			}
		}
	}

	pairsOrMore := []string{}
	for _, count := range cardCounts {
		switch count {
		case 1:
			continue
		case 2:
			pairsOrMore = append(pairsOrMore, "Pair")
		case 3:
			pairsOrMore = append(pairsOrMore, "Trips")
		case 4:
			pairsOrMore = append(pairsOrMore, "Quads")
		case 5:
			pairsOrMore = append(pairsOrMore, "Five")
		}
	}
	// fmt.Println(pairsOrMore)
	
	if len(pairsOrMore) == 0 {
		newHand.handRank = "High card"
		newHand.handRankValue = rankMap[newHand.handRank]
		return newHand
	} else if len(pairsOrMore) == 1 && pairsOrMore[0] == "Pair"{
		newHand.handRank = "One pair"
		newHand.handRankValue = rankMap[newHand.handRank]
		return newHand
	} else if len(pairsOrMore) == 1 {
		newHand.handRank = pairsOrMore[0]
		newHand.handRankValue = rankMap[newHand.handRank]
		return newHand
	} else if len(pairsOrMore) > 1 && slices.Contains(pairsOrMore, "Trips") {
		newHand.handRank = "Full house"
		newHand.handRankValue = rankMap[newHand.handRank]
		return newHand
	} else {
		newHand.handRank = "Two pair"
		newHand.handRankValue = rankMap[newHand.handRank]
		return newHand
	}

}

func revalueHandRankValue(hands []Hand) (updatedHands []Hand){
	/* Okay, let's think this shit through:
	So, we can easily get our hands and get what sort of hand we have.
	It is fairly trivial to be able to create an array of hands. 
	So we need to get the list off all unique hand types first.*/
	updatedHands = hands

	uniqueHandRanks := []uint8{}
	for _, hand := range hands {
		if slices.Contains(uniqueHandRanks, hand.handRankValue) {
			continue
		} else {
			uniqueHandRanks = append(uniqueHandRanks, hand.handRankValue)
		}
	}
	fmt.Println(uniqueHandRanks)
	
	slices.Sort(uniqueHandRanks)

	fmt.Println(updatedHands)
	if len(uniqueHandRanks) == 7 {
		return updatedHands
	} else {
		for i := 1; i <= len(uniqueHandRanks); i++ {
			ind := i - 1
			for handInd, hand := range updatedHands {
				if hand.handRankValue == uniqueHandRanks[ind] {
					fmt.Println(hand)
					newHand := hand
					newHand.handRankValue = uint8(i)
					updatedHands[handInd] = newHand
				}
				
			}
		}
	}
	fmt.Println(updatedHands)
	return updatedHands
}

func solutionOne() {
	lines, _ := readFileToListStrings("example.txt")
	fmt.Println(lines)

	hand := createHand("AKKKK 333")
	hand2 := createHand("AA254 234")
	hands := []Hand{hand, hand2}

	for ind, hand := range hands {
		newHand := rankHand(hand)
		hands[ind] = newHand
		// fmt.Println(hand, newHand)
	}
	fmt.Println(hands)
	newHands := revalueHandRankValue(hands)
	// fmt.Println(hand)
	// rankHand(&hand)
	fmt.Println(newHands)
}

func main() {
	fmt.Println("Hello, World!")
	solutionOne()
}
