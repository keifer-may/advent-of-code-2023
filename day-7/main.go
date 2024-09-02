package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var (
	outfile, _ = os.Create("./log")
	l = log.New(outfile, "", 0)
)

func deepCopy(src, dst *Hand) {
	dst.bid = src.bid
	dst.handRank = src.handRank
	dst.handRankValue = src.handRankValue

	dst.cards = make([]uint16, len(src.cards))
	copy(dst.cards, src.cards)
}


type Hand struct {
	cards    []uint16
	bid      uint16
	handRank string
	handRankValue uint16
}
	
type HandGroup []Hand

func (h HandGroup) Len() int {
	return len(h)
}
func (h HandGroup) Less(i, j int) bool {
	num := slices.Compare(h[i].cards, h[j].cards)
	if num >= 0 {
		return false
	} else {
		return true
	}
}

func (h HandGroup) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

var rankMap = map[string]uint16{
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
	bid, _ := strconv.ParseUint(bidString, 10, 16)
	hand.bid = uint16(bid)

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
			hand.cards = append(hand.cards, uint16(intcard))
		}
	}
	return hand
}

func createListHands(lines []string) (hands []Hand) {
	for _, line := range lines {
		hand := createHand(line)
		hands = append(hands, hand)
	}
	return hands
}

func rankHand(hand Hand) (newHand Hand){
	newHand = hand
	cards := make([]uint16, len(hand.cards))
	_ = copy(cards, hand.cards)
	slices.Sort(cards)

	uniqueCards := []uint16{}
	for _, card := range cards {
		if slices.Contains(uniqueCards, card) {
			continue
		} else {
			uniqueCards = append(uniqueCards, card)
		}
	}

	cardCounts := make(map[uint16]uint16)
	for _, card := range uniqueCards {
		cardCounts[card] = uint16(0)
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

func groupHandRanks(handList []Hand) (listsHands []HandGroup) {
	uniqueHandRanks := []uint16{}
	for _, hand := range handList {
		rankInt := hand.handRankValue
		if slices.Contains(uniqueHandRanks, rankInt) {
			continue
		} else {
			uniqueHandRanks = append(uniqueHandRanks, rankInt)
		}
	}

	slices.Sort(uniqueHandRanks)

	for _, handRank := range uniqueHandRanks {
		listHands := []Hand{}
		for _, hand := range handList {
			handRankInt := hand.handRankValue
			if handRank == handRankInt {
				listHands = append(listHands, hand)
			} else {
				continue
			}
		}
		listsHands = append(listsHands, HandGroup(listHands))
	}
	return listsHands
}

func sortHandGroup(hands HandGroup) {
	sort.Sort(hands)
}

func processGroups(groupedHands []HandGroup) (listHands []Hand) {

	for i, group := range groupedHands {
		if i == len(groupedHands) - 1 {
			continue
		} else {
			nextGroup := groupedHands[i + 1]
			groupRank := group[0].handRankValue
			nextRank := nextGroup[0].handRankValue

			if nextRank < groupRank {
				tempGroup := group
				groupedHands[i] = nextGroup
				groupedHands[i + 1] = tempGroup
			}
		}
	}

	firstListHands := []Hand{}
	
	for _, group := range groupedHands {
		sortHandGroup(group)
		for _, hand := range group {
			firstListHands = append(firstListHands, hand)
		}
	}
	
	for i, hand := range firstListHands {
		hand.handRankValue = uint16(i + 1)
		listHands = append(listHands, hand)
	}
	return listHands
}

func getTotalSolutionOne(listHands []Hand) (total int64) {
	for _, hand := range listHands {
		product := int64(hand.bid) * int64(hand.handRankValue)
		total += product
	}
	return total
}

func replaceJacks(hand *Hand) {
	cards := hand.cards 
	for i, c := range cards {
		if c == uint16(11) {
			cards[i] = uint16(1)
		}
	}
}

func processJacks(hand *Hand) /*(newHand Hand)*/ {
	checkCards := hand.cards
	if slices.Contains(checkCards, uint16(11)) == false {
	//	newHand = hand
		return //newHand
	} else {
		replaceJacks(hand)
		var copiedHand Hand

		deepCopy(hand, &copiedHand)
		l.Println(copiedHand, hand)

		cards := copiedHand.cards

		uniqueCards := make(map[uint16]int)
		for _, card := range cards {
			if uniqueCards[card] == 0 {
				uniqueCards[card] = 1
			} else {
				count := uniqueCards[card]
				uniqueCards[card] = count + 1
			}
		}

		if len(uniqueCards) == 5 {
			maxCard := slices.Max(cards)
			locateJack := slices.Index(cards, uint16(1))
			cards = slices.Replace(cards, locateJack, locateJack + 1, maxCard)
			rankedHand := rankHand(copiedHand)
			rankName := rankedHand.handRank
			rankValue := rankedHand.handRankValue

			l.Println(rankName, rankValue)

			hand.handRank = rankName
			hand.handRankValue = rankValue
			l.Println(hand)
			return

		} else if len(uniqueCards) == 1 {
			for i, _ := range cards {cards[i] = uint16(1)}
			rankedHand := rankHand(copiedHand)
			rankName := rankedHand.handRank
			rankValue := rankedHand.handRankValue

			hand.handRank = rankName
			hand.handRankValue = rankValue
			return

		} else if len(uniqueCards) == 2 {
			cardNotJack := uint16(0)
			for _, card := range cards {
				if card == uint16(1) {
					continue
				} else {
					cardNotJack = card
				}
			}
			for i, card := range cards {
				if card == uint16(1) {
					cards[i] = cardNotJack
				} else {
					continue
				}
			}
			rankedHand := rankHand(copiedHand)
			rankName := rankedHand.handRank
			rankValue := rankedHand.handRankValue

			hand.handRank = rankName
			hand.handRankValue = rankValue
			return

		} else {

			//fmt.Println(uniqueCards)
			uniqueCounts := []int{}
			for _, v := range uniqueCards {uniqueCounts = append(uniqueCounts, v)}
			slices.Sort(uniqueCounts)
			slices.Reverse(uniqueCounts)
	//		fmt.Println(uniqueCounts)

			maxCard := slices.Max(cards)
			greatestInst := slices.Max(uniqueCounts)

			if uniqueCards[uint16(1)] == greatestInst {
				for card, count := range uniqueCards {
					if card == uint16(1) {
						continue
					} else {
						if count == greatestInst {
							maxCard = card
							for i, c := range cards {
								if c == uint16(1) {
									cards[i] = maxCard
								} else {
									continue
								}
							}
						}
					}
				}

				for i, c := range cards {
					if c == uint16(1) {
						cards[i] = maxCard
					}
				}
			}

			for card, count := range uniqueCards {
				if count == greatestInst {
					for i, c := range cards {
						if c == uint16(1) {
							cards[i] = card
						} else {
							continue
						}
					}
				}

			}
			rankedHand := rankHand(copiedHand)
			rankName := rankedHand.handRank
			rankValue := rankedHand.handRankValue

			hand.handRank = rankName
			hand.handRankValue = rankValue
			return
		}
		rankedHand := rankHand(copiedHand)
		rankName := rankedHand.handRank
		rankValue := rankedHand.handRankValue

		hand.handRank = rankName
		hand.handRankValue = rankValue
		return
	}
}

func solutionOne() {
	lines, _ := readFileToListStrings("input.txt")
//	fmt.Println(lines)

	// hand := createHand("AKKKK 333")
	// hand2 := createHand("AA254 234")
	hands := createListHands(lines)
	//fmt.Println(hands)

	for ind, hand := range hands {
		newHand := rankHand(hand)
		hands[ind] = newHand
		// fmt.Println(hand, newHand)
	}
	// fmt.Println(hands)
	groupedHands := groupHandRanks(hands)
	finalHandList := processGroups(groupedHands)
	//fmt.Println(finalHandList)
	solution := getTotalSolutionOne(finalHandList)
	fmt.Println("Final answer for problem one:", solution)
	// fmt.Println(groupedHands)
	// sortedHands := orderHandGroups(groupedHands)
	// fmt.Println(sortedHands)
}

func solutionTwo() {
	lines, _ := readFileToListStrings("input.txt")
	hands := createListHands(lines)

//	fmt.Println(hands)

	for ind, hand := range hands {
		l.Println(hand)
		newHand := rankHand(hand)
		processJacks(&newHand)
		hands[ind] = newHand
		l.Println(newHand)
	}

//	fmt.Println(hands)
	groupedHands := groupHandRanks(hands)
	finalHandList := processGroups(groupedHands)
	//fmt.Println(finalHandList)
	solution := getTotalSolutionOne(finalHandList)
	fmt.Println("Final answer for problem two:", solution)
	// fmt.Println(groupedHands)
	// sortedHands := orderHandGroups(groupedHands)
	// fmt.Println(sortedHands)
}

func main() {
	// fmt.Println("Hello, World!")
	solutionOne()
	solutionTwo()
}
