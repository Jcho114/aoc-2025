package main

import (
	"bufio"
	"fmt"
	"os"
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

	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	NUM_ROWS, NUM_COLUMNS := len(grid), len(grid[0])

	col := 0
	for col < NUM_COLUMNS {
		temp := col
		for temp < NUM_COLUMNS && grid[0][temp] == ' ' {
			temp += 1
		}
		for temp < NUM_COLUMNS && grid[0][temp] != ' ' {
			temp += 1
		}
		for temp < NUM_COLUMNS && !isSeparatorColumn(grid, temp) {
			temp += 1
		}
		length := temp - col

		problem := Problem{
			values:   make([]int, length),
			operator: "",
		}

		for count := 0; count < length; count++ {
			for row := 0; row < NUM_ROWS-1; row++ {
				if grid[row][col+count] == ' ' {
					continue
				}
				value := int(grid[row][col+count] - '0')
				problem.values[count] = problem.values[count]*10 + value
			}
		}

		problems = append(problems, problem)
		col = temp + 1
	}

	fields := strings.Fields(string(grid[NUM_ROWS-1]))
	for i, operator := range fields {
		problems[i].operator = operator
	}

	return problems
}

func isSeparatorColumn(grid [][]rune, col int) bool {
	for row := 0; row < len(grid)-1; row++ {
		if grid[row][col] != ' ' {
			return false
		}
	}
	return true
}
