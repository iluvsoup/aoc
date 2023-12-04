package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var matchingNumbersCache = map[int]int{}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numberScratchCards := 0
	games := strings.Split(string(bytes), "\n")

	var copiesToBeProcessed = map[int]int{}

	for i, game := range games {
		cardNr := i + 1
		numberScratchCards++

		data := strings.Split(game, ": ")[1]
		lists := strings.Split(data, " | ")

		winningNumbers := parseNumberList(&lists[0])
		ourNumbers := parseNumberList(&lists[1])

		sort.Ints(winningNumbers)
		sort.Ints(ourNumbers)

		matchingNumbers := 0
		minIndex := 0
		for _, winningNumber := range winningNumbers {
			for i := minIndex; i < len(ourNumbers); i++ {
				if winningNumber == ourNumbers[i] {
					matchingNumbers++
					minIndex = i + 1
				}
			}
		}
		matchingNumbersCache[i] = matchingNumbers

		for copyNr := cardNr + 1; copyNr <= cardNr+matchingNumbers; copyNr++ {
			copiesToBeProcessed[copyNr] += 1
		}
	}

	for copyNr, count := range copiesToBeProcessed {
		for i := 0; i < count; i++ {
			numberScratchCards += processCopy(copyNr)
		}
	}

	fmt.Println(numberScratchCards)
}

func processCopy(cardNr int) int {
	numberScratchCards := 1
	matchingNumbers := matchingNumbersCache[cardNr-1]
	for copyNr := cardNr + 1; copyNr <= cardNr+matchingNumbers; copyNr++ {
		numberScratchCards += processCopy(copyNr)
	}
	return numberScratchCards
}

func parseNumberList(str *string) []int {
	var list []int

	parsingNumber := ""

	for _, char := range []byte(*str) {
		if char == ' ' {
			if parsingNumber != "" {
				number, err := strconv.Atoi(parsingNumber)
				if err != nil {
					panic(err)
				}
				list = append(list, number)
				parsingNumber = ""
			}
		} else {
			parsingNumber += string(char)
		}
	}

	if parsingNumber != "" {
		number, err := strconv.Atoi(parsingNumber)
		if err != nil {
			panic(err)
		}
		list = append(list, number)
		parsingNumber = ""
	}

	return list
}
