package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")
	instructionString := lines[0]
	nodeStrings := lines[2:]

	var positions []string
	var nodes = map[string][]string{}
	for _, nodeString := range nodeStrings {
		parts := strings.Split(nodeString, " = ")

		location := parts[0]
		destinations := strings.Split(strings.Trim(parts[1], "()"), ", ")

		if location[len(location)-1] == 'A' {
			positions = append(positions, location)
		}

		nodes[location] = destinations
	}

	var steps int
	var instructionIndex int

	var stepCounts []int

	for {
		if len(positions) == 0 {
			break
		}

		steps++

		instruction := instructionString[instructionIndex]

		var newPositions []string
		for i, position := range positions {
			destinations := nodes[position]
			if instruction == 'L' {
				positions[i] = destinations[0]
			} else {
				positions[i] = destinations[1]
			}

			if positions[i][len(positions[i])-1] == 'Z' {
				stepCounts = append(stepCounts, steps)
			} else {
				newPositions = append(newPositions, positions[i])
			}
		}
		positions = newPositions

		instructionIndex = (instructionIndex + 1) % len(instructionString)
	}

	// No idea why this works, I just wanna be done with it
	stepCount := LCM(stepCounts)
	fmt.Println(stepCount)
}

// goofy ahh
func LCM(numbers []int) int {
	if len(numbers) > 2 {
		return calculateLCM(numbers[0], LCM(numbers[1:]))
	} else {
		return calculateLCM(numbers[0], numbers[1])
	}
}

func calculateGCD(num1 int, num2 int) int {
	if num2 == 0 {
		return num1
	} else {
		return calculateGCD(num2, num1%num2)
	}
}

func calculateLCM(num1 int, num2 int) int {
	gcd := calculateGCD(num1, num2)
	lcm := (num1 * num2) / gcd
	return lcm
}
