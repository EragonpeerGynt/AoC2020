package main

import (
	"AoC/Common/Reader"
  "strings"
  "fmt"
  "regexp"
)

func main() {
	dirty := reader.ReadAllLines()
  splitInput := strings.Split(dirty, "\n\n")

  //part 1
  var cleanInput []string
  for _,customs := range splitInput {
    cleanInput = append(cleanInput, string(unique(strings.Replace(customs, "\n", "", -1))))
  }

  all := 0
  for _,customs := range cleanInput {
    all += len(customs)
  }
  fmt.Println(all)

  //part 2
  onlyYes := 0
  for _,customs := range splitInput {
    onlyYes += findNonUnique(customs)
  }
  fmt.Println(onlyYes)
}

func unique(intSlice string) []rune {
    keys := make(map[rune]bool)
    list := []rune{} 
    for _, entry := range intSlice {
        if _, value := keys[entry]; !value {
            keys[entry] = true
            list = append(list, entry)
        }
    }    
    return list
}

func findNonUnique(customs string) (occurances int) {
  occurances = 0
  people := len(strings.Split(customs, "\n"))
  for _,character := range strings.Split(customs, "\n")[0] {
    if (string(character) == "\n") {
      continue
    }
    occurances += isThereEnoughCharacters(customs, string(character), people)
  }
  return
}

func isThereEnoughCharacters(customs string, character string, people int) int {
  reg := regexp.MustCompile(character)
  if (len(reg.FindAllString(customs, -1)) == people) {
    return 1
  }
  return 0
}