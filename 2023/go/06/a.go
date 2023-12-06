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
  timesString := strings.TrimSpace(lines[0][5:])
  distancesString := strings.TrimSpace(lines[1][9:])
  
  times := parseIntegerList(&timesString)
  distances := parseIntegerList(&distancesString)

  product := 1
  
  for raceNr := 0; raceNr < len(times); raceNr++ {
    time := times[raceNr]
    bestDistance := distances[raceNr]

    var numberOfWaysToWin int
    for t := 1; t < time; t++ {
      distanceToBeTraveled := (time - t) * t
      if distanceToBeTraveled > bestDistance {
        numberOfWaysToWin++
      }
    }

    product *= numberOfWaysToWin
  }

  fmt.Println(product)
}

func isDigit(c byte) bool {
  return c >= '0' && c <= '9'
}

func parseIntegerList(str *string) []int {
  var list []int
  
  parseNumber := ""
  for _, c := range []byte(*str) {
    if isDigit(c) {
      parseNumber += string(c)
    } else {
      if parseNumber != "" {
        number, err := strconv.Atoi(parseNumber)
        if err != nil {
          panic(err)
        }

        list = append(list, number)
        parseNumber = ""
      }
    }
  }

  if parseNumber != "" {
    number, err := strconv.Atoi(parseNumber)
    if err != nil {
      panic(err)
    }

    list = append(list, number)
    parseNumber = ""
  }

  return list
}