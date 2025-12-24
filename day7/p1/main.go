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
	splits := 0

	pointQueue := NewPointQueue()
	pointQueue.Enqueue(Point{
		Row: startRow,
		Col: startCol,
	})

	for pointQueue.Len() > 0 {
		currPoint := pointQueue.Dequeue()
		currRow, currCol := currPoint.Row, currPoint.Col
		for currRow < N-1 && grid[currRow+1][currCol] == Empty {
			currRow += 1
			grid[currRow][currCol] = Beam
		}
		if currRow < N-1 && grid[currRow+1][currCol] == Splitter {
			if currCol-1 >= 0 {
				pointQueue.Enqueue(Point{
					Row: currRow + 1,
					Col: currCol - 1,
				})
			}
			if currCol+1 < M {
				pointQueue.Enqueue(Point{
					Row: currRow + 1,
					Col: currCol + 1,
				})
			}
			splits += 1
		}
		// prettyPrintGrid(grid)
	}

	fmt.Println(splits)
}

func prettyPrintGrid(grid [][]rune) {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

type Point struct {
	Row int
	Col int
}

type PointQueue struct {
	points []Point
}

func (queue *PointQueue) Enqueue(point Point) {
	queue.points = append(queue.points, point)
}

func (queue *PointQueue) Dequeue() Point {
	point := queue.points[0]
	queue.points = queue.points[1:]
	return point
}

func (queue *PointQueue) Len() int {
	return len(queue.points)
}

func NewPointQueue() *PointQueue {
	return &PointQueue{
		points: []Point{},
	}
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
