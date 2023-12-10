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

	var nodes = map[string][]string{}
	for _, nodeString := range nodeStrings {
		parts := strings.Split(nodeString, " = ")

		location := parts[0]
		destinations := strings.Split(strings.Trim(parts[1], "()"), ", ")

		nodes[location] = destinations
	}

	var steps int

	var instructionIndex int
	var currentNode = "AAA"
	for {
		if currentNode == "ZZZ" {
			break
		}

		instruction := instructionString[instructionIndex]
		destinations := nodes[currentNode]
		if instruction == 'L' {
			currentNode = destinations[0]
		} else {
			currentNode = destinations[1]
		}
		instructionIndex = (instructionIndex + 1) % len(instructionString)
		steps++
	}

	fmt.Println(steps)
}
