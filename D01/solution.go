package main

import (
	"fmt"
	"AoC/Common/Reader"
  "strings"
  "strconv"
)

func stringToIntArray(old []string) []int {
  var clean []int
  for _,number := range old {
    n, _ := strconv.Atoi(number)
    clean = append(clean, n)
  }
  return clean
}

func recursiveFinder(ranger []int) int {
  var current int = ranger[0]
  var forward []int = ranger[1:]

  for _,number := range forward {
    if (number + current == 2020) {
      return number * current
    }
  }

  return recursiveFinder(forward)
}

func trippleDiveInit(ranger []int) int {
  for i := 0; i < len(ranger); i++ {
    for j := i+1; j < len(ranger); j++ {
      for k := j+1; k < len(ranger); k++ {
        if (ranger[i] + ranger[j] + ranger[k] == 2020) {
          return ranger[i] * ranger[j] * ranger[k]
        }
      }
    }
  }
  return 0
}

func main() {
	input := reader.ReadAllLines()
  splitInput := strings.Split(input, "\n")
  cleanInput := stringToIntArray(splitInput)
  fmt.Println(recursiveFinder(cleanInput))
  fmt.Println(trippleDiveInit(cleanInput))
}