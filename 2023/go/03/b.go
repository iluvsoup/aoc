package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	x int
	y int
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	var gears = map[Position][]int{}
	for y, line := range lines {
		var parseNumber string
		var adjacentGears []Position
		for x, char := range []byte(line) {
			if isNumber(char) {
				parseNumber += string(char)

				// check surrounding characters
				for checkX := max(0, x-1); checkX <= min(len(line)-1, x+1); checkX++ {
					for checkY := max(0, y-1); checkY <= min(len(lines)-1, y+1); checkY++ {
						gear := Position{checkX, checkY}
						if lines[checkY][checkX] == '*' && !isGearInList(&adjacentGears, &gear) {
							adjacentGears = append(adjacentGears, gear)
						}
					}
				}

				if x+1 == len(line) || !isNumber(line[x+1]) {
					number, err := strconv.Atoi(parseNumber)
					if err != nil {
						panic(err)
					}
					parseNumber = ""

					for _, gear := range adjacentGears {
						gears[gear] = append(gears[gear], number)
					}
					adjacentGears = nil
				}
			}
		}
	}

	var sum int
	for _, numbers := range gears {
		if len(numbers) > 1 {
			gearRatio := 1
			for _, number := range numbers {
				gearRatio *= number
			}
			sum += gearRatio
		}
	}
	fmt.Println(sum)
}

func isGearInList(gears *[]Position, gear *Position) bool {
	for _, v := range *gears {
		if v == *gear {
			return true
		}
	}
	return false
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
