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

  input := string(bytes)
  lines := strings.Split(input, "\n")

  var sum int
  for _, line := range lines {
    if line != "" {
      firstNumber, lastNumber := parseLine(line)
      // putting the two numbers together
      number := firstNumber * 10 + lastNumber
      sum += number
    }
  }

  fmt.Println(sum)
}

func parseLine(line string) (int, int) {
  var numberString string
  var firstNumber, lastNumber int
  
  for i := 0; i < len(line); i++ {
    s := line[i]
    if isNumber(s) {
      parsedNumber, err := strconv.Atoi(string(s))
      if err != nil {
        panic(err)
      }
      firstNumber = parsedNumber
      break
    } else {
      // Assume it's a letter
      numberString += string(s)
      parsedNumber := checkNumber(numberString)
      if parsedNumber != -1 {
        firstNumber = parsedNumber
        numberString = ""
        break
      }
    }
  }

  for i := len(line) - 1; i >= 0; i-- {
    s := line[i]
    if isNumber(s) {
      parsedNumber, err := strconv.Atoi(string(s))
      if err != nil {
        panic(err)
      }
      lastNumber = parsedNumber
      break
    } else {
      // Assume it's a letter
      numberString = string(s) + numberString
      parsedNumber := checkNumber(numberString)
      if parsedNumber != -1 {
        lastNumber = parsedNumber
        numberString = ""
        break
      }
    }
  }

  return firstNumber, lastNumber
}

func isNumber(b byte) bool {
  return b >= '0' && b <= '9'
}

func checkNumber(s string) int {
  if strings.Contains(s, "one") { return 1 }
  if strings.Contains(s, "two") { return 2 }
  if strings.Contains(s, "three") { return 3 }
  if strings.Contains(s, "four") { return 4 }
  if strings.Contains(s, "five") { return 5 }
  if strings.Contains(s, "six") { return 6 }
  if strings.Contains(s, "seven") { return 7 }
  if strings.Contains(s, "eight") { return 8 }
  if strings.Contains(s, "nine") { return 9 }
  return -1
}