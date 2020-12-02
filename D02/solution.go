package main

import (
	"AoC/Common/Reader"
  "fmt"
  "strings"
  "regexp"
  "strconv"
)

type password struct {
  numberRepeat string
  numberIndex1 int
  numberIndex2 int
  repeatType string
  password string
}

func ParsePassword(pass string) password {
  cleanPass := strings.Split(pass, " ")
  indexes := strings.Split(cleanPass[0], "-")
  var parsedPass password
  parsedPass.numberRepeat = "{" + strings.Replace(cleanPass[0], "-", ",", -1) + "}"
  parsedPass.numberIndex1,_ = strconv.Atoi(indexes[0])
  parsedPass.numberIndex2,_ = strconv.Atoi(indexes[1])
  parsedPass.repeatType = strings.Replace(cleanPass[1], ":", "", -1)
  parsedPass.password = cleanPass[2]
  return parsedPass
}

func (pass *password) DoesPassMatch() bool {
  reg := regexp.MustCompile(`(?m)^([^` + pass.repeatType + `]*` + pass.repeatType + `[^` + pass.repeatType + `]*)` + pass.numberRepeat + `$`)
  if(len(reg.FindAllString(pass.password, -1)) != 0) {
    return true
  }
  return false
}

func (pass *password) ExactlyOneIndex() bool {
  if ((pass.password[pass.numberIndex1-1] == pass.repeatType[0] && pass.password[pass.numberIndex2-1] != pass.repeatType[0]) || (pass.password[pass.numberIndex1-1] != pass.repeatType[0] && pass.password[pass.numberIndex2-1] == pass.repeatType[0])) {
    return true
  }
  return false
}

func main() {
	dirtyInput := reader.ReadAllLines()
  splitInput := strings.Split(dirtyInput, "\n")
  var parseInput []password
  for _,pass := range splitInput {
    parseInput = append(parseInput, ParsePassword(pass))
  }
  correctPass := 0
  for _,pass := range parseInput {
    if(pass.DoesPassMatch()) {
      correctPass++
    }
  }
  fmt.Println(correctPass)

  correctPass = 0
  for _,pass := range parseInput {
    if(pass.ExactlyOneIndex()) {
      correctPass++
    }
  }
  fmt.Println(correctPass)
  
}