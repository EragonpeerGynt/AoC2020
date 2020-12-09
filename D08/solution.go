package main

import (
	"AoC/Common/Reader"
	"fmt"
	"strings"
	"strconv"
)

type Console struct {
	Index int
	Accumulator int
	Program []string
	History map[int]bool
}

func consoleFactory(program []string) (console Console) {
	console = Console {
		Index: 0,
		Accumulator: 0,
		Program: make([]string, len(program)),
		History: make(map[int]bool),
	}
	copy(console.Program, program)
	return
}


func main() {
	input := reader.ReadLineArray()

	solvePart1(input)
	solvePart2(input)
}


func solvePart1(input []string) {
	console := consoleFactory(input)
	for console.ExecuteNextCommand() == "continue" {

	}
	fmt.Println(console.Accumulator)
	return
}


func solvePart2(input []string) {
	fmt.Println(part2Looper(input))
}


func (console *Console)ExecuteNextCommand() string {
	if console.Index >= len(console.Program) {
		return "terminate"
	}

	if _,exists := console.History[console.Index]; exists {
		return "loop"
	}
	
	instruction := strings.Split(console.Program[console.Index], " ")
	command := instruction[0]
	modifier,_ := strconv.Atoi(instruction[1])
	console.History[console.Index] = true

	if strings.Contains(command, "nop") {
		console.Index ++
	} else if strings.Contains(command, "acc") {
		console.Accumulator += modifier
		console.Index ++
	} else if strings.Contains(command, "jmp") {
		console.Index += modifier
	}
	return "continue"
}

func part2Looper(input []string) int {
	for i := 0; i < len(input); i++ {
		console := consoleFactory(input)
		if strings.Contains(input[i], "nop") {
			console.Program[i] = strings.ReplaceAll(console.Program[i], "nop", "jmp")
		} else if strings.Contains(input[i], "jmp") {
			console.Program[i] = strings.ReplaceAll(console.Program[i], "jmp", "nop")
		} else {
			continue
		}
		for console.ExecuteNextCommand() == "continue" {

		}
		if console.ExecuteNextCommand() == "terminate" {
			return console.Accumulator
		}
	}
	return 0
}