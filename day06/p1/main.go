package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	problems := processInput()
	total := 0
	for _, problem := range problems {
		total += problem.evaluate()
	}
	fmt.Println(total)
}

type Problem struct {
	values   []int
	operator string
}

func (problem *Problem) evaluate() int {
	if problem.operator == "+" {
		sum := 0
		for _, value := range problem.values {
			sum += value
		}
		return sum
	} else if problem.operator == "*" {
		product := 1
		for _, value := range problem.values {
			product *= value
		}
		return product
	}
	return -1
}

func processInput() []Problem {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	problems := []Problem{}

	scanner.Scan()
	line := scanner.Text()
	fields := strings.Fields(line)
	for _, field := range fields {
		value, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		problems = append(problems, Problem{
			values:   []int{value},
			operator: "",
		})
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] == "*" || fields[0] == "+" {
			for i, field := range fields {
				problems[i].operator = field
			}
		} else {
			for i, field := range fields {
				value, err := strconv.Atoi(field)
				if err != nil {
					panic(err)
				}
				problems[i].values = append(problems[i].values, value)
			}
		}
	}

	return problems
}
