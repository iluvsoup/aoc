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
      firstindex := strings.IndexAny(line, "0123456789")
      lastindex := strings.LastIndexAny(line, "0123456789")
      
      if firstindex != -1 && lastindex != -1 {
        firstnumber := string(line[firstindex])
        lastnumber := string(line[lastindex])
        
        number, err := strconv.Atoi(firstnumber + lastnumber)
        if err != nil {
          panic(err)
        }
        
        sum += number
      }
    }
  }

  fmt.Println(sum)
}

