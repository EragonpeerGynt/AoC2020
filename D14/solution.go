package main

import (
	"AoC/Common/Reader"
	"strings"
	"regexp"
	"strconv"
	"fmt"
)

type Transition struct {
	Command string
	Value string
}

func main() {
	input := reader.ReadLineArray()
	parsed := parseInput(input)
	
	solvePart1(parsed)
	solvePart2(parsed)
}

func solvePart1(input []Transition) {
	mask := make([]string, 0)
	memory := make(map[int]int64)
	for _,row := range input {
		if strings.Contains(row.Command, "mask") {
			mask = strings.Split(row.Value, "")
		} else {
			memorySlot := row.getMemoryAddress()
			memory[memorySlot] = row.maskNumber(mask)
		}
	}

	total := int64(0)
	for _,val := range memory {
		total += val
	}
	fmt.Println(total)
}


func solvePart2(input []Transition) {
	mask := make([]string, 0)
	memory := make(map[int]int64)
	for _,row := range input {
		if strings.Contains(row.Command, "mask") {
			mask = strings.Split(row.Value, "")
		} else {
			memorySlots := row.getPossibleMemoryAddresses(mask)
			value,_ := strconv.ParseInt(row.Value, 10, 64)
			for _,memorySlot := range memorySlots {
				memory[int(memorySlot)] = value
			}
		}
	}

	total := int64(0)
	for _,val := range memory {
		total += val
	}
	fmt.Println(total)
}


func (command *Transition) maskNumber(mask []string) (maskedValue int64) {
	unmasked := strings.Split(fmt.Sprintf("%036v", command.getBinaryRepresentation()), "")
	for i,val := range mask {
		if strings.ToUpper(val) != "X" {
			unmasked[i] = val
		}
	}
	maskedValue,_ = strconv.ParseInt(strings.Join(unmasked, ""), 2, 64)
	return
}

func (command *Transition) getMemoryAddress() (address int) {
	re := regexp.MustCompile(`(?m).*\[|\]`)
	address,_ = strconv.Atoi(re.ReplaceAllString(command.Command, ""))
	return
}

func (command *Transition) getBinaryRepresentation() (number string) {
	intNumber,_ := strconv.ParseInt(command.Value, 10, 64)
	number = strconv.FormatInt(intNumber, 2)
	return
}

func (command *Transition) getPossibleMemoryAddresses(mask []string) []int64 {
	re := regexp.MustCompile(`(?m).*\[|\]`)
	intAddress,_ := strconv.ParseInt(re.ReplaceAllString(command.Command, ""), 10, 64)
	address := strconv.FormatInt(intAddress, 2)
	splitAddress := strings.Split(fmt.Sprintf("%036v", address), "")
	for i,val := range mask {
		if val != "0" {
			splitAddress[i] = val
		}
	}
	return listFloatingAddresses("", splitAddress)
}

func listFloatingAddresses(complete string, address []string) []int64 {
	if len(address) == 0 {
		address,_ := strconv.ParseInt(complete, 2, 64)
		return []int64{address,}
	}
	if address[0] != "X" {
		return listFloatingAddresses(complete + address[0], address[1:])
	} else {
		return append(listFloatingAddresses(complete + "0", address[1:]), listFloatingAddresses(complete + "1", address[1:])...)
	}
	return make([]int64, 0)
}


func parseInput(input []string) (output []Transition) {
	output = make([]Transition, 0)
	for _,row := range input {
		rowContent := strings.Split(row, " = ")
		output = append(output, Transition{Command:rowContent[0],Value:rowContent[1],})
	}
	return
}