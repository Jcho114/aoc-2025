package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Start = 'S'
var Splitter = '^'
var Beam = '|'
var Empty = '.'

func main() {
	grid := processInput()
	N, M := len(grid), len(grid[0])
	startRow, startCol := findStart(grid)

	cache := make([][]int, N)
	for i := range cache {
		cache[i] = make([]int, M)
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}

	paths := traverseGrid(grid, cache, startRow, startCol)

	fmt.Println(paths)
}

func isValidCoordinates(row int, col int, N int, M int) bool {
	return row >= 0 && row < N && col >= 0 && col < M
}

func traverseGrid(grid [][]rune, cache [][]int, currRow int, currCol int) int {
	N, M := len(grid), len(grid[0])

	if currRow == N-1 {
		// prettyPrintGrid(grid)
		return 1
	}

	if cache[currRow][currCol] != -1 {
		return cache[currRow][currCol]
	}

	downRow, downCol := currRow+1, currCol
	if !isValidCoordinates(downRow, downCol, N, M) {
		return 0
	}

	paths := 0
	switch grid[downRow][downCol] {
	case Splitter:
		leftRow, leftCol := downRow, downCol-1
		if isValidCoordinates(leftRow, leftCol, N, M) && grid[leftRow][leftCol] != Splitter {
			grid[leftRow][leftCol] = Beam
			paths += traverseGrid(grid, cache, leftRow, leftCol)
		}

		rightRow, rightCol := downRow, downCol+1
		if isValidCoordinates(rightRow, rightCol, N, M) && grid[rightRow][rightCol] != Splitter {
			grid[rightRow][rightCol] = Beam
			paths += traverseGrid(grid, cache, rightRow, rightCol)
		}
	default:
		grid[downRow][downCol] = Beam
		paths += traverseGrid(grid, cache, downRow, downCol)
	}

	cache[currRow][currCol] = paths
	return paths
}

func prettyPrintGrid(grid [][]rune) {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func findStart(grid [][]rune) (int, int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == Start {
				return row, col
			}
		}
	}
	return -1, -1
}

func processInput() [][]rune {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}
