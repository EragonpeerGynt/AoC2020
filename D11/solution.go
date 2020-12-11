package main

import (
	"AoC/Common/Reader"
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	input := reader.ReadLineArray()

	mappedInput := parseToMap(input)
	solvePart1(mappedInput)

	mappedInput = parseToMap(input)
	solvePart2(mappedInput)
}


func solvePart1(input map[point]string) {
	for iterationsOfCycle(&input) != 0 {
		
	}
	counter := 0
	for _,val := range input {
		if val == "#" {
			counter++
		}
	}
	fmt.Println(counter)
}

func iterationsOfCycle(input *map[point]string) (changed int) {
	changed = 0
	mapping := make(map[point]string)
	for p,val := range *input {
		if val == "L" && countNeighbors(p, *input) == 0 {
			mapping[p] = "#"
			changed++
			continue
		}
		if val == "#" && countNeighbors(p, *input) >= 4 {
			mapping[p] = "L"
			changed++
			continue
		}
		mapping[p] = val
	}
	*input = mapping
	return
}

func countNeighbors(seat point, input map[point]string) (counter int) {
	counter = 0
	for _,neighbor := range neighbors(seat) {
		if input[neighbor] == "#" {
			counter ++
		}
	}
	return
}

func neighbors(seat point) []point {
	points := make([]point, 0)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(j == 0 && i == 0) {
				points = append(points, point{x:seat.x+i,y:seat.y+j,})
			}
		}
	}
	return points
}


func solvePart2(input map[point]string) {
	for iterationsOfCycleDirect(&input) != 0 {
		
	}
	counter := 0
	for _,val := range input {
		if val == "#" {
			counter++
		}
	}
	fmt.Println(counter)
}

func iterationsOfCycleDirect(input *map[point]string) (changed int) {
	changed = 0
	mapping := make(map[point]string)
	for p,val := range *input {
		if val == "L" && advancedNeighbors(p, *input) == 0 {
			mapping[p] = "#"
			changed++
			continue
		}
		if val == "#" && advancedNeighbors(p, *input) >= 5 {
			mapping[p] = "L"
			changed++
			continue
		}
		mapping[p] = val
	}
	*input = mapping
	return
}

func advancedNeighbors(seat point, input map[point]string) (counter int) {
	counter = 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if j == 0 && i == 0 {
				continue
			}
			for m := 1; m < 100; m++ {
				val,ok := input[point{x:seat.x+(i*m),y:seat.y+(j*m),}]
				if !ok {
					break
				}
				if val == "." {
					continue
				}
				if val == "L" {
					break
				}
				if val == "#" {
					counter ++
					break
				}
			}
		}
	}
	return
}


func parseToMap(input []string) map[point]string {
	mapping := make(map[point]string)
	for i,vali := range input {
		for j,val := range vali {
			mapping[point{x:i,y:j,}] = string(val)
		}
	}
	return mapping
}