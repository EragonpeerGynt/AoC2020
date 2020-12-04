package reader

import (
	"fmt"
	"io/ioutil"
  "os"
  "strings"
)

func check(e error) {
	if e != nil {
        	panic(e)
    	}
}

func ReadAllLines() string {
	dat, err := ioutil.ReadFile("./" + os.Args[1] + "/input.txt")
  check(err)
  fmt.Print(string(dat) + "\n")
  return string(dat)
}

func ReadLineArray() []string {
  return strings.Split(ReadAllLines(), "\n")
}