package main

import (
	"AoC/Common/Reader"
  "strings"
  "fmt"
)

type road struct {
  right int
  down int
}

func DiveDeeper(coordinates []string, move road, width int, treesEncounter int) int {
  if(coordinates[0][width] == "#"[0]) {
    treesEncounter ++
  }
  if (len(coordinates) > 1) {
    return DiveDeeper(coordinates[move.down:], move, (width + move.right) % len(coordinates[0]), treesEncounter)
  }
  return treesEncounter
}

func main() {
	dirtyInput := reader.ReadAllLines()
  splitInput := strings.Split(dirtyInput, "\n")
  move := road{right: 3, down: 1}
  fmt.Println(DiveDeeper(splitInput, move, 0, 0))
  var base int64 = int64(1)
  move = road{right: 1, down: 1}
  base = base * int64(DiveDeeper(splitInput, move, 0, 0))
  move = road{right: 3, down: 1}
  base = base * int64(DiveDeeper(splitInput, move, 0, 0))
  move = road{right: 5, down: 1}
  base = base * int64(DiveDeeper(splitInput, move, 0, 0))
  move = road{right: 7, down: 1}
  base = base * int64(DiveDeeper(splitInput, move, 0, 0))
  move = road{right: 1, down: 2}
  base = base * int64(DiveDeeper(splitInput, move, 0, 0))
  fmt.Println(base)
}