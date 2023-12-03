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

	for y, line := range lines {
		var parseNumber string
		var isPartNumber bool
		for x, char := range []byte(line) {
			if isNumber(char) {
				parseNumber += string(char)

				// check surrounding characters
				if isPartNumber == false {
					for checkX := max(0, x-1); checkX <= min(len(line)-1, x+1); checkX++ {
						for checkY := max(0, y-1); checkY <= min(len(lines)-1, y+1); checkY++ {
							if isSymbol(lines[checkY][checkX]) {
								isPartNumber = true
								break
							}
						}
					}
				}

				if x+1 == len(line) || !isNumber(line[x+1]) {
					if isPartNumber {
						number, err := strconv.Atoi(parseNumber)
						if err != nil {
							panic(err)
						}
						sum += number
					}
					parseNumber = ""
					isPartNumber = false
				}
			}
		}
	}

	fmt.Println(sum)
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isNumber(c)
}
