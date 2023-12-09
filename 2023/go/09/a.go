package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	var sum int

	for _, line := range lines {
		if line != "" {
			numberStrings := strings.Split(line, " ")
			var numbers []int
			for _, numberString := range numberStrings {
				number, err := strconv.Atoi(numberString)
				if err != nil {
					panic(err)
				}
				numbers = append(numbers, number)
			}

			d := numbers
			var differenceSequences [][]int
			for {
				allZeroes := true
				for _, v := range d {
					if v != 0 {
						allZeroes = false
						break
					}
				}
				if allZeroes {
					break
				} else {
					differenceSequences = append(differenceSequences, d)
				}

				d = getDifferences(d)
			}

			previousValue := 0
			for i := len(differenceSequences) - 1; i >= 0; i-- {
				differences := differenceSequences[i]
				newValue := differences[len(differences)-1] + previousValue
				previousValue = newValue
			}
			sum += previousValue
		}
	}

	fmt.Println(sum)
}

func getDifferences(numbers []int) []int {
	var differences []int
	for i := 0; i < len(numbers)-1; i++ {
		number := numbers[i]
		nextNumber := numbers[i+1]

		diff := nextNumber - number
		differences = append(differences, diff)
	}
	return differences
}
