package main

import (
	"AoC/Common/Reader"
	"fmt"
	"strconv"
	"strings"
)

type Direction struct {
	x int
	y int
}

type Ship struct {
	ns int
	ew int
	direction int
	directions []Direction
}

var directionOptions = []Direction{
	Direction{x:0,y:1,},
	Direction{x:1,y:0,},
	Direction{x:0,y:-1,},
	Direction{x:-1,y:0,},
}

var cardinalDirections = map[string]Direction {
	"E":Direction{x:0,y:1,},
	"N":Direction{x:1,y:0,},
	"W":Direction{x:0,y:-1,},
	"S":Direction{x:-1,y:0,},
}

func (ship *Ship) rotate (direction string, degrees int) {
	degrees = degrees / 90
	if (direction == "R") {
		degrees *= 3
	}

	ship.direction = (ship.direction + degrees) % 4
}

func main() {
	input := reader.ReadLineArray()
	solvePart1(input)
	solvePart2(input)
}


func solvePart1(input []string) {
	ship := Ship{
		ns:0,
		ew:0,
		direction:0,
		directions:directionOptions,
	}
	for _,command := range input {
		movement := string(command[0])
		repeats,_ := strconv.Atoi(string(command[1:]))
		if strings.Contains("LR", movement) {
			ship.rotate(movement, repeats)
		} else if movement == "F" {
			ship.ns = ship.ns + (ship.directions[ship.direction].x * repeats)
			ship.ew = ship.ew + (ship.directions[ship.direction].y * repeats)
		} else {
			ship.ns = ship.ns + (cardinalDirections[movement].x * repeats)
			ship.ew = ship.ew + (cardinalDirections[movement].y * repeats)
		}
	}
	fmt.Println(Abs(ship.ew) + Abs(ship.ns))
}


func solvePart2(input []string) {
	ship := Ship{
		ns:0,
		ew:0,
		direction:0,
		directions:directionOptions,
	}
	waypoint := Ship{
		ns:1,
		ew:10,
		direction:0,
		directions:directionOptions,
	}
	for _,command := range input {
		movement := string(command[0])
		repeats,_ := strconv.Atoi(string(command[1:]))
		if strings.Contains("LR", movement) {
			waypoint.rotate2(movement, repeats)
		} else if movement == "F" {
			ship.ns = ship.ns + (waypoint.ns * repeats)
			ship.ew = ship.ew + (waypoint.ew * repeats)
		} else {
			waypoint.ns = waypoint.ns + (cardinalDirections[movement].x * repeats)
			waypoint.ew = waypoint.ew + (cardinalDirections[movement].y * repeats)
		}
	}
	fmt.Println(Abs(ship.ew) + Abs(ship.ns))
}

func (waypoint *Ship) rotate2 (direction string, degrees int) {
	if direction == "L" {
		degrees *= 3
	}
	for i := 0; i < degrees/90; i++ {
		waypoint.ns, waypoint.ew = -waypoint.ew, waypoint.ns
	}
}


func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}