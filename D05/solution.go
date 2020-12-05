package main

import (
	"AoC/Common/Reader"
  "fmt"
  "sort"
)

func main() {
	input := reader.ReadLineArray()
  fmt.Println(findHighestId(input))
  fmt.Println(findMySeat(input))
}

func findHighestId(seattings []string) (max int) {
  max = 0
  var current int
  for _,seatting := range seattings {
    current = binaryConverter(seatting[:7]) * 8 + binaryConverter(seatting[7:])
    if current > max {
      max = current
    }
  }
  return
}

func findMySeat(seattings []string) (seat int) {
  var allSeats []int
  for _,seatting := range seattings {
    allSeats = append(allSeats, binaryConverter(seatting[:7]) * 8 + binaryConverter(seatting[7:]))
  }

  sort.Ints(allSeats)

  for i,id := range allSeats {
    if (id + 2 == allSeats[i+1]) {
      seat = id + 1
      return
    }
  }
  return 0
}

func binaryConverter(binaryString string) (converted int) {
  step := 1
  converted = 0
  for i := len(binaryString) - 1; i >= 0; i-- {
    if (binaryString[i] == "B"[0] || binaryString[i] == "R"[0]) {
      converted += step
    }
    step = step * 2
  }
  return
}