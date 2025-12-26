package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	points := processInput()

	max := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			if p1.IsBottomLeftOf(p2) || p2.IsBottomLeftOf(p1) || p1.IsBottomRightOf(p2) || p1.IsBottomRightOf(p2) {
				gridArea := p1.GridAreaWith(p2)
				if gridArea > max {
					max = gridArea
				}
			}
		}
	}

	fmt.Println(max)
}

type Point struct {
	X int
	Y int
}

func (p1 Point) IsBottomLeftOf(p2 Point) bool {
	return p1.X <= p2.X && p1.Y >= p2.Y
}

func (p1 Point) IsBottomRightOf(p2 Point) bool {
	return p1.X >= p2.X && p1.Y >= p2.Y
}

func (p1 Point) GridAreaWith(p2 Point) int {
	var diffX = 0
	if p1.X < p2.X {
		diffX = p2.X - p1.X
	} else {
		diffX = p1.X - p2.X
	}

	var diffY = 0
	if p1.Y < p2.Y {
		diffY = p2.Y - p1.Y
	} else {
		diffY = p1.Y - p2.Y
	}

	length, width := diffX+1, diffY+1
	return length * width
}

func processInput() []Point {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	points := []Point{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		points = append(points, Point{
			X: x,
			Y: y,
		})
	}

	return points
}
