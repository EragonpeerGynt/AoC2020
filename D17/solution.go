package main

import (
	"AoC/Common/Reader"
	"strings"
	"fmt"
)

type Point struct {
	X int
	Y int
	Z int
}

func addPoints(point1, point2 Point) Point {
	return Point{
		X:point1.X+point2.X,
		Y:point1.Y+point2.Y,
		Z:point1.Z+point2.Z,
	}
}

type Point2 struct {
	X int
	Y int
	Z int
	W int
}

func addPoints2(point1, point2 Point2) Point2 {
	return Point2{
		X:point1.X+point2.X,
		Y:point1.Y+point2.Y,
		Z:point1.Z+point2.Z,
		W:point1.W+point2.W,
	}
}



var NeighborConfig = make([]Point, 0)
var NeighborConfig2 = make([]Point2, 0)


func main() {
	configureGlobalNeighborRange()
	configureGlobalNeighborRange2()
	fmt.Println("-")
	solvePart1(parseInput())
	solvePart2(parseInput2())
}

//part1
func solvePart1(input map[Point]string) {
	for i := 0; i < 6; i++ {
		generateNextGeneration(&input)
	}
	fmt.Println(len(input))
}

func generateNextGeneration(input *map[Point]string) {
	checks := generateChangablePoints(*input)
	output := make(map[Point]string)
	for check,_ := range checks {
		neighbors := countNeighbors(check, (*input))
		if _,ok := (*input)[check]; ok {
			if (neighbors == 2 || neighbors == 3) {
				output[check] = "#"
			}
		} else {
			if neighbors == 3 {
				output[check] = "#"
			}
		}
	}
	*input = output
	return
}

func countNeighbors(point Point, points map[Point]string) (counter int) {
	counter = 0
	for _,neighbor := range actualNeighborsFromRelative(point) {
		if _,ok := points[neighbor]; ok {
			counter++
		}
	}
	return 
}

func actualNeighborsFromRelative(point Point) (output []Point) {
	output = make([]Point, 0)
	for _,relative := range NeighborConfig {
		output = append(output, addPoints(point, relative))
	}
	return
}


func generateChangablePoints(points map[Point]string) (output map[Point]bool) {
	output = make(map[Point]bool)
	for key,_ := range points {
		for _,point := range generateChangableRangeOfPoint(key) {
			output[point] = true
		}
	}
	return
}

func generateChangableRangeOfPoint(point Point) (points []Point){
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				points = append(points, Point{X:point.X+x,Y:point.Y+y,Z:point.Z+z,})
			}
		}
	}
	return
}



func parseInput() (output map[Point]string) {
	output = make(map[Point]string)
	for i,row := range reader.ReadLineArray() {
		for j, val := range strings.Split(row, "") {
			if (string(val) == "#") {
				output[Point{X:j,Y:i,Z:0,}] = string(val)
			}
		}
	}
	return
}

func configureGlobalNeighborRange() {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if !(x == 0 && y == 0 && z == 0) {
					NeighborConfig = append(NeighborConfig, Point{X:x,Y:y,Z:z,})
				}
			}
		}
	}
}


//part2
func solvePart2(input map[Point2]string) {
	for i := 0; i < 6; i++ {
		generateNextGeneration2(&input)
	}
	fmt.Println(len(input))
}

func generateNextGeneration2(input *map[Point2]string) {
	checks := generateChangablePoints2(*input)
	output := make(map[Point2]string)
	for check,_ := range checks {
		neighbors := countNeighbors2(check, (*input))
		if _,ok := (*input)[check]; ok {
			if (neighbors == 2 || neighbors == 3) {
				output[check] = "#"
			}
		} else {
			if neighbors == 3 {
				output[check] = "#"
			}
		}
	}
	*input = output
	return
}

func countNeighbors2(point Point2, points map[Point2]string) (counter int) {
	counter = 0
	for _,neighbor := range actualNeighborsFromRelative2(point) {
		if _,ok := points[neighbor]; ok {
			counter++
		}
	}
	return 
}

func actualNeighborsFromRelative2(point Point2) (output []Point2) {
	output = make([]Point2, 0)
	for _,relative := range NeighborConfig2 {
		output = append(output, addPoints2(point, relative))
	}
	return
}


func generateChangablePoints2(points map[Point2]string) (output map[Point2]bool) {
	output = make(map[Point2]bool)
	for key,_ := range points {
		for _,point := range generateChangableRangeOfPoint2(key) {
			output[point] = true
		}
	}
	return
}

func generateChangableRangeOfPoint2(point Point2) (points []Point2){
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					points = append(points, Point2{X:point.X+x,Y:point.Y+y,Z:point.Z+z,W:point.W+w})
				}
			}
		}
	}
	return
}



func parseInput2() (output map[Point2]string) {
	output = make(map[Point2]string)
	for i,row := range reader.ReadLineArray() {
		for j, val := range strings.Split(row, "") {
			if (string(val) == "#") {
				output[Point2{X:j,Y:i,Z:0,W:0}] = string(val)
			}
		}
	}
	return
}

func configureGlobalNeighborRange2() {
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if !(x == 0 && y == 0 && z == 0 && w == 0) {
						NeighborConfig2 = append(NeighborConfig2, Point2{X:x,Y:y,Z:z,W:w,})
					}
				}
			}
		}
	}
}