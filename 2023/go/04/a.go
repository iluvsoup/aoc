package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	points := 0

	games := strings.Split(string(bytes), "\n")
	for _, game := range games {
		data := strings.Split(game, ": ")[1]
		lists := strings.Split(data, " | ")

		winningNumbers := parseNumberList(&lists[0])
		ourNumbers := parseNumberList(&lists[1])

		sort.Ints(winningNumbers)
		sort.Ints(ourNumbers)

		gamePoints := 0

		minIndex := 0
		for _, winningNumber := range winningNumbers {
			for i := minIndex; i < len(ourNumbers); i++ {
				if winningNumber == ourNumbers[i] {
					if gamePoints == 0 {
						gamePoints = 1
					} else {
						gamePoints *= 2
					}

					minIndex = i + 1
				}
			}
		}

		points += gamePoints
	}

	fmt.Println(points)
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
