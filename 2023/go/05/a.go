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

	parts := strings.Split(string(bytes), "\n\n")

	var values []int
	seedString := strings.Split(parts[0], ": ")[1]
	for _, seed := range strings.Split(seedString, " ") {
		number, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		values = append(values, number)
	}

	for _, part := range parts[1:] {
		ranges := strings.Split(part, "\n")[1:]

		var newValues []int
		for _, value := range values {
			newValues = append(newValues, value)
		}

		for _, mapRange := range ranges {
			numbers := strings.Split(mapRange, " ")

			destinationRangeStart, err1 := strconv.Atoi(numbers[0])
			if err1 != nil {
				panic(err)
			}

			sourceRangeStart, err2 := strconv.Atoi(numbers[1])
			if err2 != nil {
				panic(err)
			}

			rangeLength, err3 := strconv.Atoi(numbers[2])
			if err3 != nil {
				panic(err)
			}

			for i, value := range values {
				if value >= sourceRangeStart && value < sourceRangeStart+rangeLength {
					diff := value - sourceRangeStart
					newValues[i] = destinationRangeStart + diff
				}
			}
		}

		values = newValues
	}

	lowestLocation := values[0]
	for _, location := range values {
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	fmt.Println(lowestLocation)
}
