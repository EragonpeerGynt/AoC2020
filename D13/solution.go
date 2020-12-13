package main

import (
	"AoC/Common/Reader"
	"strings"
	"strconv"
	"fmt"
)

func main() {
	input := reader.ReadLineArray()
	ourTime,_ := strconv.Atoi(input[0])
	busIds1 := parseIds1(input[1])
	busIds2 := parseIds2(input[1])

	solvePart1(ourTime, busIds1)

	solvePart2(busIds2)
}

func solvePart1(ourTime int, busIds []int) {
	busId := 0
	waitTime := ourTime * 2
	for _,id := range busIds {
		var manyTimes int = (ourTime / id + 1) * id
		if manyTimes < waitTime {
			busId = id
			waitTime = manyTimes
		}
	}
	fmt.Println(busId * (waitTime - ourTime))
}

func solvePart2(busIds map[int64]int64) {
	minValue := int64(0)
	runningProduct := int64(1)
	for k, v := range busIds {
		for (minValue+v)%k != 0 {
			minValue += runningProduct
		}
		runningProduct *= k
	}
	fmt.Println(minValue)
}


func parseIds2(input string) (parsed map[int64]int64) {
	parsed = make(map[int64]int64)
	for i,id := range strings.Split(input, ",") {
		if id == "x" {
			continue
		}
		idInt,_ := strconv.Atoi(id)
		parsed[int64(idInt)] = int64(i)
	}
	return
}

type Busses struct {
	id int64
	modul int64
}

func parseIds1(input string) (parsed []int) {
	cleanInput := strings.Replace(input, "x,", "", -1)
	parsed = make([]int, 0)
	for _,id := range strings.Split(cleanInput, ",") {
		idInt,_ := strconv.Atoi(id)
		parsed = append(parsed, idInt)
	}
	return
}