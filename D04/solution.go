package main

import (
	"AoC/Common/Reader"
  "strings"
  "fmt"
  "regexp"
)

func CheckPassword(password string, required []string) bool {
  for _,parameter := range required {
    if (!strings.Contains(password, parameter)) {
      return false
    }
  }
  return true
}

func CheckPasswords(passwords []string, required []string) int {
  valid := 0;
  for _,password := range passwords {
    if (CheckPassword(password, required)) {
      valid ++
    }
  }
  return valid
}

func ValidatePassword(password string, required []*regexp.Regexp) bool {
  for _,reg := range required {
    if (len(reg.FindAllString(password, -1)) == 0) {
      return false
    }
  }
  return true
}

func ValidatePasswords(passwords []string, required []*regexp.Regexp) int {
  valid := 0;
  for _,password := range passwords {
    if (ValidatePassword(password, required)) {
      valid ++
    }
  }
  return valid
}

func main() {
	dirtyInput := reader.ReadAllLines()
  splitInput := strings.Split(dirtyInput, "\n\n")
  required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
  fmt.Println(CheckPasswords(splitInput, required))
  var regexps []*regexp.Regexp
  reg := regexp.MustCompile(`(?m)byr:(19[2-9]\d|200[0-2])`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)iyr:(20(1[0-9]|20))`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)eyr:(20(2[0-9]|30))`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)hgt:((1([5-8][0-9]|9[0-4]))cm|(59|6[0-9]|7[0-6])in)`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)hcl:#[0-9a-f]{6}`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
  regexps = append(regexps, reg)
  reg = regexp.MustCompile(`(?m)pid:\d{9}([^\d]|$)`)
  regexps = append(regexps, reg)
  fmt.Println(ValidatePasswords(splitInput, regexps))
}