package main

import (
	"AoC/Common/Reader"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	input := parseInput(strings.Split(reader.ReadAllLines(), ","))

	solvePart1(input)
	solvePart2(input)
}

func solvePart1(input []int) {
	window := make([]int, len(input))
	copy(window, input)
	inputLength := len(input)
	lastNumber := input[len(input)-1]
	for i := inputLength ; i < 2020; i++ {
		diff := areThereTwoOccurances(window, lastNumber)
		lastNumber = diff
		window = append(window, diff)
	}
	fmt.Println(lastNumber)
}

func solvePart2(input []int) {
	history := map[int]int{}
	last := 0
	for i,val := range input {
		last = val
		history[last] = i + 1
	}
	for i := len(input); i < 30000000; i++ {
		if val,ok := history[last]; ok {
			history[last] = i
			last = i-val
		} else {
			history[last] = i
			last = 0
		}
	}
	fmt.Println(last)
}

func areThereTwoOccurances(input []int, lastNumber int) (diff int) {
	diff = 0
	for i := len(input)-1; i >= 0; i-- {
		if input[i] == lastNumber {
			if diff == 0 {
				diff = i
			} else {
				diff = diff - i
				return
			}
		}
	}
	diff = 0
	return
}


func parseInput(input []string) (output []int) {
	output = make([]int, 0)
	for _,val := range input {
		parsed,_ := strconv.Atoi(val)
		output = append(output, parsed)
	}
	return
}