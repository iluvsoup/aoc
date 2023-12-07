package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardType int

const (
	_ CardType = iota
	HighCard
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (ct CardType) String() string {
	switch ct {
	case HighCard:
		return "HighCard"
	case OnePair:
		return "OnePair"
	case TwoPair:
		return "TwoPair"
	case ThreeOfAKind:
		return "ThreeOfAKind"
	case FullHouse:
		return "FullHouse"
	case FourOfAKind:
		return "FourOfAKind"
	case FiveOfAKind:
		return "FiveOfAKind"
	}

	return ""
}

type Card struct {
	hand     string
	cardType CardType
	bid      int
}

var letterStrengths = map[byte]int{
	'2': 1,
	'3': 2,
	'4': 3,
	'5': 4,
	'6': 5,
	'7': 6,
	'8': 7,
	'9': 8,
	'T': 9,
	'J': 10,
	'Q': 11,
	'K': 12,
	'A': 13,
}

var cards []Card

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bidString := parts[1]

		cardType := getType(hand)

		bid, err := strconv.Atoi(bidString)
		if err != nil {
			panic(err)
		}

		cards = append(cards, Card{
			hand:     hand,
			cardType: cardType,
			bid:      bid,
		})
	}

	sort.SliceStable(cards, cardSortFunction)

	var totalWinnings int
	for i, card := range cards {
		rank := i + 1
		totalWinnings += card.bid * rank
	}
	fmt.Println(totalWinnings)
}

func getType(hand string) CardType {
	var counts = map[byte]int{}
	for _, c := range []byte(hand) {
		counts[c] += 1
	}

	if len(counts) == 1 {
		return FiveOfAKind
	} else if len(counts) == 2 {
		for _, count := range counts {
			if count == 4 {
				return FourOfAKind
			} else if count == 3 {
				return FullHouse
			}
		}
	} else if len(counts) == 3 {
		for _, count := range counts {
			if count == 3 {
				return ThreeOfAKind
			} else if count == 2 {
				return TwoPair
			}
		}
	} else if len(counts) == 4 {
		return OnePair
	} else if len(counts) == 5 {
		return HighCard
	}

	return 0
}

func cardSortFunction(xIndex int, yIndex int) bool {
	x := cards[xIndex]
	y := cards[yIndex]

	if x.cardType < y.cardType {
		return true
	} else if x.cardType > y.cardType {
		return false
	} else {
		for i := 0; i < 5; i++ {
			xChar := x.hand[i]
			yChar := y.hand[i]
			if letterStrengths[xChar] < letterStrengths[yChar] {
				return true
			} else if letterStrengths[xChar] > letterStrengths[yChar] {
				return false
			}
		}
		return false
	}
}
