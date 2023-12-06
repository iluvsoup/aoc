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
  timeString := strings.Replace(lines[0][5:], " ", "", -1)
  distanceString := strings.Replace(lines[1][9:], " ", "", -1)

  time, err := strconv.Atoi(timeString)
  if err != nil {
    panic(err)
  }

  bestDistance, err := strconv.Atoi(distanceString)
  if err != nil {
    panic(err)
  }

  fmt.Println(time, bestDistance)
  
  var numberOfWaysToWin int
  for t := 1; t < time; t++ {
    distanceToBeTraveled := (time - t) * t
    if distanceToBeTraveled > bestDistance {
      numberOfWaysToWin++
    }
  }
  fmt.Println(numberOfWaysToWin)
}