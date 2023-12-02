package main

import (
  "fmt"
  "os"
  "strings"
  "strconv"
)

var MAX_VALUES = map[string]int{
  "red": 12,
  "green": 13,
  "blue": 14,
}

func main() {
  bytes, err := os.ReadFile("input.txt")
  if err != nil {
    panic(err)
  }

  lines := strings.Split(string(bytes), "\n")
  
  var sum int
  for n, line := range lines {
    id := n + 1
    
    fullgame := strings.Split(line, ": ")[1]
    games := strings.Split(fullgame, "; ")

    possible := true
    for _, game := range games {
      cubes := strings.Split(game, ", ")
      for _, cube := range cubes {
        parts := strings.Split(cube, " ")
        number, err := strconv.Atoi(parts[0])
        if err != nil {
          panic(err)
        }
        color := parts[1]

        if number > MAX_VALUES[color] {
          possible = false
          break
        }
      }
    }

    if possible {
      sum += id
    }
  }

  fmt.Println(sum)
}
