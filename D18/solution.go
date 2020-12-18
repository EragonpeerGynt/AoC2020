package main

import (
	"AoC/Common/Reader"
	"strings"
	"strconv"
	"fmt"
	"regexp"
)

func main() {
	input := reader.ReadLineArray()

	solvePart1(input)
	solvePart2(input)
}

func solvePart1(input []string) {
	total := 0
	for _,expression := range input {
		calculatedValue,_ := strconv.Atoi(fullExpressionParse(expression))
		total += calculatedValue
	}
	fmt.Println(total)
}

func fullExpressionParse(expression string) string {
	var re = regexp.MustCompile(`(?m)\([^(]*?\)`)
	for true {
		matches := re.FindAllString(expression, -1)
		if(len(matches) == 0) {
			break
		}
		for _,match := range matches {
			expression = strings.Replace(expression, match, evaluateExpression(match), -1)
		}
	}
	return evaluateExpression(expression)
}

func solvePart2(input []string) {
	total := 0
	for _,expression := range input {
		calculatedValue,_ := strconv.Atoi(fullExpressionParsePrioritizePlus(expression))
		total += calculatedValue
	}
	fmt.Println(total)
}

func fullExpressionParsePrioritizePlus(expression string) string {
	var re = regexp.MustCompile(`(?m)\([^(]*?\)`)
	for true {
		matches := re.FindAllString(expression, -1)
		if(len(matches) == 0) {
			break
		}
		for _,match := range matches {
			expression = strings.Replace(expression, match, evaluateExpressionSplit(match), -1)
		}
	}
	return evaluateExpressionSplit(expression)
}

func evaluateExpressionSplit(expression string) string {
	var re = regexp.MustCompile(`(?m)(\(|\))`)
	expression = re.ReplaceAllString(expression, "")
	elements := strings.Split(expression, " * ")
	addition := make([]string, 0)
	for _,element := range elements {
		addition = append(addition, evaluateExpression(element))
	}

	return evaluateExpression(strings.Join(addition, " * "))
}

func evaluateExpression(expression string) string {
	var re = regexp.MustCompile(`(?m)(\(|\))`)
	expression = re.ReplaceAllString(expression, "")
	elements := strings.Split(expression, " ")
	base := 0
	operator := "+"
	for _,element := range elements {
		if !(element == "+" || element == "*") {
			current,_ := strconv.Atoi(element)
			if operator == "+" {
				base += current
			} else {
				base *= current
			}
		} else {
			operator = element
		}
		
	}
	return strconv.Itoa(base)
}

