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

func solvePart2(busIds []Busses) {
	remainderTable := make([]chineseRemainder, 0)
	totalId := allReminder(busIds)
	for _,busId := range busIds {
		tmpReminder := chineseRemainder{
			bi:busId.modul,
			Ni:totalId/busId.id,
			xi:calculateXi(busId.id, totalId/busId.id),
		}
		remainderTable = append(remainderTable, tmpReminder)
	}
	totalChinese := int64(0)
	for _,remainder := range remainderTable {
		totalChinese += (remainder.bi*remainder.Ni*remainder.xi)
	}
	fmt.Println(remainderTable)
	fmt.Println(totalChinese%totalId)
}

func allReminder (busIds []Busses) int64 {
	all := int64(1)
	for _,bus := range busIds {
		all *= bus.id
	}
	return all
}

func calculateXi(base, bonus int64) int64{
	bonus = bonus % base
	i := int64(1)
	for {
		if (bonus * i) % base == 1 {
			break
		}
		i++
	}
	return (bonus * i)
}

func parseIds2(input string) (parsed []Busses) {
	parsed = make([]Busses, 0)
	for i,id := range strings.Split(input, ",") {
		if id == "x" {
			continue
		}
		idInt,_ := strconv.Atoi(id)
		parsed = append(parsed, Busses{id:int64(idInt),modul:int64(i),})
	}
	return
}


type chineseRemainder struct {
	bi int64
	Ni int64
	xi int64
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