package main

import (
	"fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  bytes, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  lines := strings.Split(string(bytes), "\n")
  
  var sum int
  for _, line := range lines {
    fullgame := strings.Split(line, ": ")[1]
    games := strings.Split(fullgame, "; ")

    var min = map[string]int{
      "red": 0,
      "green": 0,
      "blue": 0,
    }
    
    for _, game := range games {
      cubes := strings.Split(game, ", ")
      for _, cube := range cubes {
        parts := strings.Split(cube, " ")
        number, err := strconv.Atoi(parts[0])
        if err != nil {
          panic(err)
        }
        color := parts[1]

        if number > min[color] {
          min[color] = number
        }
      }
    }

    power := min["red"] * min["green"] * min["blue"]
    sum += power
  }

  fmt.Println(sum)
}
