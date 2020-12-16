package main

import (
	"AoC/Common/Reader"
	"AoC/Common/Arrays"
	"strings"
	"fmt"
	"strconv"
)

type Range struct {
	From int
	To int
}

type BoardingPassField struct {
	Ranges []Range
	PossibleIndexes map[int]int
}

type Field struct {
	Value int
	PossibleTypes []string
}

func main() {
	input := strings.Split(reader.ReadAllLines(), "\n\n")
	fields,possibleRanges := parseRanges(input[0])
	solvePart1(possibleRanges, strings.Split(input[2], "\n"))
	solvePart2(fields, possibleRanges, strings.Split(input[2], "\n"), strings.Split(input[1], "\n")[1])
}

func solvePart1(possible []Range, tickets []string) {
	totalMissing := 0
	for _,ticket := range tickets {
		intTicket := parseTicket(ticket)
		for _,num := range intTicket {
			if !isInRange(possible, num) {
				totalMissing += num
			}
		}
	}
	fmt.Println(totalMissing)
}

func isInRange(possible []Range, number int) bool {
	for _,ranger := range possible {
		if (number >= ranger.From && number <= ranger.To) {
			return true
		}
	}
	return false
}

func isWholeTicketInRange(possible []Range, fields []int) bool {
	for _,num := range fields {
		if !isInRange(possible, num) {
			return false
		}
	}
	return true
}

func solvePart2(fields map[string]*BoardingPassField, possible []Range, tickets []string, ourPass string) {
	var validPass [][]int
	ourPassInt := parseTicket(ourPass)
	for _,ticket := range tickets {
		intTicket := parseTicket(ticket)
		if isWholeTicketInRange(possible, intTicket) {
			validPass = append(validPass, intTicket)
		}
	}
	possibleTypeCombinations(validPass, fields)
	possibleIndexesRemaining := arraygen.GenerateProtoArray(1,len(ourPassInt))

	actualIndexes := make(map[string]int)

	for len(possibleIndexesRemaining) > 0 {
		name,index,_ := nextSingleValue(fields, possibleIndexesRemaining)
		actualIndexes[name] = index
		delete(possibleIndexesRemaining, index)
		for _,val := range fields {
			if _,ok := (*val).PossibleIndexes[index]; ok {
				delete((*val).PossibleIndexes, index)
			}
		}
	}
	multiply := int64(1)
	for key,val := range actualIndexes {
		if strings.Contains(key, "departure") {
			multiply *= int64(ourPassInt[val-1])
		}
	}
	fmt.Println(multiply)
}

func possibleTypeCombinations(tickets [][]int, fields map[string]*BoardingPassField) {
	for _,ticket := range tickets {
		for j,field := range ticket {
			possibleType(field, fields, j)
		}
	}
	return
} 

func possibleType(field int, fields map[string]*BoardingPassField, index int) {
	for key,val := range fields {
		if !isInRange(val.Ranges, field) {
			fields[key].PossibleIndexes[index+1] += 1
		}
	}
	return
}


func nextSingleValue(fullPossible map[string]*BoardingPassField, possibleIndexesRemaining map[int]bool) (name string, index int, exist bool) {
	exist = false
	for key,possibility := range fullPossible {
		if len((*possibility).PossibleIndexes)+1 == len(possibleIndexesRemaining) {
			name = key
			index = missingIndexMap((*possibility).PossibleIndexes, possibleIndexesRemaining)
			exist = true
			return
		}
	}
	return
}

func missingIndexMap(possibleIndexes map[int]int, possibleIndexesRemaining map[int]bool) int {
	for key,_ := range possibleIndexesRemaining {
		if _,ok := possibleIndexes[key]; !ok {
			return key
		}
	}
	return -1
}


func parseTicket(ticket string) (output []int) {
	splitTicket := strings.Split(ticket, ",")
	for _,split := range splitTicket {
		num,_ := strconv.Atoi(split)
		output = append(output, num)
	}
	return
}


func parseRanges(input string) (output map[string]*BoardingPassField, onlyRanges []Range) {
	output = make(map[string]*BoardingPassField)
	onlyRanges = make([]Range, 0)
	splitInput := strings.Split(input, "\n")
	for _,val := range splitInput {
		splitVal := strings.Split(val, ": ")
		name := splitVal[0]
		ranges := strings.Split(splitVal[1], " or ")
		var intRange []Range
		for _,rangeVal := range ranges {
			fromTo := strings.Split(rangeVal, "-")
			from,_ := strconv.Atoi(fromTo[0])
			to,_ := strconv.Atoi(fromTo[1])
			intRange = append(intRange, Range{From:from, To:to,})
			onlyRanges = append(onlyRanges, Range{From:from, To:to,})
		}
		output[name] = &BoardingPassField{Ranges:intRange,PossibleIndexes:make(map[int]int)}
	}
	return
}