package main

import (
  "AoC/Common/Reader"
  "fmt"
  "regexp"
  "strings"
  "strconv"
)

type Bag struct {
  numberOfBags int
  bagType string
}


func main() {
	cleanInput := reader.ReadLineArray()

  parsedInput := parseInput(cleanInput)

  solvePart1(parsedInput)
  solvePart2(parsedInput)
}


func solvePart1(input map[string][]Bag) {
  counter := 0
  for key,_ := range input {
    if (recursiveFinder(key, input) && key != "shiny gold") {
      counter ++
    }
  }
  fmt.Println(counter)
}

func recursiveFinder(parent string, input map[string][]Bag) bool {
  if (strings.Contains(parent, "shiny gold")) {
    return true
  }
  for _,child := range input[parent] {
    if (recursiveFinder(child.bagType, input)) {
      return true;
    }
  }
  return false
}


func solvePart2(input map[string][]Bag) {
  fmt.Println(recursiveCounter("shiny gold", input) - 1) //we don't count the shyni gold one again
}

func recursiveCounter(parent string, input map[string][]Bag) (totalBags int) {
  totalBags = 1

  if (len(input[parent]) == 0) {
    return 1
  }

  for _,subBags := range input[parent] {
    totalBags += subBags.numberOfBags * recursiveCounter(subBags.bagType, input)
  }
  return
}


func parseInput(inputs []string) (mapping map[string][]Bag) {
  mapping = make(map[string][]Bag)
  primarySplit := regexp.MustCompile(`(?m) bags? contain `)
  secondarySplit := regexp.MustCompile(`(?m) bags?, `)
  replaceSecondary := regexp.MustCompile(`(?m) bags?[.]`)
  splitSecondarySplit := regexp.MustCompile(`(?m) `)

  for _,input := range inputs {
    primary := primarySplit.Split(input, -1)
    var secondary []Bag
    for _,content := range secondarySplit.Split(primary[1], -1) {
      secondaryString := splitSecondarySplit.Split(replaceSecondary.ReplaceAllString(content, ""), 2)
      if (strings.Contains(secondaryString[0], "no")) {
        continue;
      }
      numberOfBags,_ := strconv.Atoi(secondaryString[0])
      secondary = append(secondary, Bag{
        numberOfBags: numberOfBags,
        bagType: secondaryString[1],
      })
    }
    mapping[primary[0]] = secondary
  }

  return
}