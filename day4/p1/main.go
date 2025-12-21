package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}

		for _, char := range line {
			row = append(row, char)
		}

		grid = append(grid, row)
	}

	accessible := 0
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
		{1, 1},
		{1, -1},
		{-1, -1},
		{-1, 1},
	}
	N, M := len(grid), len(grid[0])
	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
			if grid[row][col] != '@' {
				continue
			}

			numAdjacent := 0
			for _, direction := range directions {
				rowDiff, colDiff := direction[0], direction[1]
				rowNew, colNew := row+rowDiff, col+colDiff

				if rowNew >= 0 && rowNew < N && colNew >= 0 && colNew < M && grid[rowNew][colNew] == '@' {
					numAdjacent += 1
				}
			}

			if numAdjacent < 4 {
				accessible += 1
			}
		}
	}

	fmt.Println(accessible)
}
