package main

import (
	"AoC/Common/Reader"
	"strconv"
	"fmt"
	"sort"
)

func main() {
	parsed := parseIntInput(reader.ReadLineArray())
	
	target := solvePart1(parsed)
	solvePart2(parsed, target)
}


func solvePart1(input []int) int {
	solution := findUncomputableNumber(input)
	fmt.Println(solution)
	return solution
}


func solvePart2(input []int, target int) {
	fmt.Println(rappidAddition(input, target))
}


func findUncomputableNumber(input []int) int {
	for i := 25; i < len(input); i ++ {
		if !(doesItAddUp(input[i-25:i], input[i])) {
			return input[i]
		}
	}
	return 0
}

func doesItAddUp(options []int, target int) bool {
	for i := 0; i < len(options); i++ {
		for j := i+1; j < len(options); j++ {
			if (options[i] + options[j]) == target {
				return true
			}
		}
	}
	return false
}


func rappidAddition(input []int, target int) int {
	for i := 2; i <= len(input); i++ {
		for j := 0; j < len(input)-i; j++ {
			if (target == sum(input[j:j+i])) {
				return min(input[j:j+i]) + max(input[j:j+i])
			}
		}
	}
	return 0
}

func sum(input []int) (i int){
	i = 0
	for _,val := range input {
		i += val
	}
	return
}

func min(input []int) int {
	sort.Ints(input)
	return input[0]
}

func max(input []int) int {
	sort.Ints(input)
	return input[len(input)-1]
}

func parseIntInput(unparsed []string) (parsed []int) {
	parsed = make([]int, 0)
	for _,valstr := range unparsed {
		val,_ := strconv.Atoi(valstr)
		parsed = append(parsed, val)
	}
	return
}