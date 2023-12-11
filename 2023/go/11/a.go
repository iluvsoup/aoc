package main

import (
	"fmt"
	"math"
	"os"
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

	var galaxies []Position

	rows := strings.Split(string(bytes), "\n")
	width := len(rows[0])
	height := len(rows)

	expandedX := 0
	for x := 0; x < width; x++ {
		isColumnEmpty := true
		expandedY := 0

		for y := 0; y < height; y++ {
			isRowEmpty := true
			for _x := 0; _x < width; _x++ {
				char := rows[y][_x]
				if char == '#' {
					isRowEmpty = false
					break
				}
			}

			char := rows[y][x]
			if char == '#' {
				galaxies = append(galaxies, Position{
					expandedX,
					expandedY,
				})
				isColumnEmpty = false
			}
			expandedY++
			if isRowEmpty {
				expandedY++
			}
		}
		expandedX++
		if isColumnEmpty {
			expandedX++
		}
	}

	var sum int

	// DERANGED for loop, PLEASE rewrite
	var distances = map[int]map[int]int{}
	for galaxyIndex, _ := range galaxies {
		for otherGalaxyIndex, _ := range galaxies {
			var aIndex, bIndex int
			if otherGalaxyIndex > galaxyIndex {
				aIndex = galaxyIndex
				bIndex = otherGalaxyIndex
			} else {
				aIndex = otherGalaxyIndex
				bIndex = galaxyIndex
			}

			galaxy := galaxies[aIndex]
			otherGalaxy := galaxies[bIndex]

			distMap, ok := distances[aIndex]
			if !ok {
				distances[aIndex] = make(map[int]int)
				distances[aIndex][bIndex] = manhattanDistance(galaxy, otherGalaxy)
			} else {
				_, ok := distMap[bIndex]
				if !ok {
					distMap[bIndex] = manhattanDistance(galaxy, otherGalaxy)
				}
			}
		}
	}

	for _, a := range distances {
		for _, b := range a {
			sum += b
		}
	}

	fmt.Println(sum)
}

func manhattanDistance(a Position, b Position) int {
	return int(math.Abs(float64(a.x-b.x)) + (math.Abs(float64(a.y - b.y))))
}
