package main

import (
	"AoC/Common/Reader"
	"sort"
	"fmt"
	"strconv"
)

func main() {
	parsed := parseIntInput(reader.ReadLineArray())

	solvePart1(parsed)
	solvePart2(parsed)
}


func solvePart1(input []int) {
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)
	diff := [4]int{0,0,0,0}
	previous := 0
	for _,jolt := range input {
		diff[jolt-previous] ++
		previous = jolt
	}
	fmt.Println(diff[1] * diff[3])
}

func solvePart2(input []int) {
	//setting up input and initial state
	input = append(input, 0)
	sort.Ints(input)
	input = reverseInts(input)
	routes := make(map[int]int)
	//actual logic
	///I used dictionary as stack simulator for this program. It works probably somewhat as fast as array but with a lot less overhead
	routes[input[0]+3] = 1
	for _,jolt := range input {
		routes[jolt] = 0
		for i := 1; i <= 3; i++ {
			if val,ok := routes[jolt+i]; ok {
				fmt.Println(jolt+i, val)
				routes[jolt] += val
			}
		}
	}
	fmt.Println(routes[0])
}


func reverseInts(input []int) []int {
    if len(input) == 0 {
        return input
    }
	return append(reverseInts(input[1:]), input[0]) 
}


func parseIntInput(unparsed []string) (parsed []int) {
	parsed = make([]int, 0)
	for _,valstr := range unparsed {
		val,_ := strconv.Atoi(valstr)
		parsed = append(parsed, val)
	}
	return
}