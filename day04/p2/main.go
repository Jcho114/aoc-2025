package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var Directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, -1},
	{-1, 1},
}
var N, M int
var Removed = 'x'

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
	N, M = len(grid), len(grid[0])

	queue := NewPointQueue()
	initQueue(grid, queue)

	removed := 0
	for queue.Len() > 0 {
		tempRemoved := 0
		length := queue.Len()
		points := []Point{}
		for i := 0; i < length; i++ {
			point := queue.Dequeue()
			points = append(points, point)
		}

		for _, point := range points {
			if grid[point.Row][point.Col] == Removed {
				continue
			}
			grid[point.Row][point.Col] = Removed
			enqueueAdjacentRolls(queue, grid, point)
			removed += 1
			tempRemoved += 1
		}
	}
	fmt.Println(removed)
}

func debugPrintGrid(grid [][]rune) {
	var sb = strings.Builder{}
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

func enqueueAdjacentRolls(queue *PointQueue, grid [][]rune, point Point) {
	for _, direction := range Directions {
		rowDiff, colDiff := direction[0], direction[1]
		rowNew, colNew := point.Row+rowDiff, point.Col+colDiff

		if rowNew >= 0 && rowNew < N && colNew >= 0 && colNew < M && grid[rowNew][colNew] == '@' && canRemove(grid, rowNew, colNew) {
			queue.Enqueue(Point{
				Row: rowNew,
				Col: colNew,
			})
		}
	}
}

func canRemove(grid [][]rune, row int, col int) bool {
	numAdjacent := 0
	for _, direction := range Directions {
		rowDiff, colDiff := direction[0], direction[1]
		rowNew, colNew := row+rowDiff, col+colDiff

		if rowNew >= 0 && rowNew < N && colNew >= 0 && colNew < M && grid[rowNew][colNew] == '@' {
			numAdjacent += 1
		}
	}

	return numAdjacent < 4
}

func initQueue(grid [][]rune, queue *PointQueue) {
	for row := 0; row < N; row++ {
		for col := 0; col < M; col++ {
			if grid[row][col] == '@' && canRemove(grid, row, col) {
				queue.Enqueue(Point{
					Row: row,
					Col: col,
				})
			}
		}
	}
}
