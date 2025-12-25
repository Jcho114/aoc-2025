package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	points, edges := processInput()

	max := 0
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]
			gridArea := p1.GridAreaWith(p2)
			if gridArea > max && !intersectsPolygon(p1, p2, edges) && withinPolygon(p1, p2, edges) {
				max = gridArea
			}
		}
	}

	fmt.Println(max)
}

func sortTwo(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func pointInOrOnPolygon(p Point, edges []Edge) bool {
	// Check if on vertex
	for _, e := range edges {
		if (p.X == e.Start.X && p.Y == e.Start.Y) || (p.X == e.End.X && p.Y == e.End.Y) {
			return true
		}
	}

	inside := false

	x, y := p.X, p.Y
	for _, e := range edges {
		x1, y1 := e.Start.X, e.Start.Y
		x2, y2 := e.End.X, e.End.Y

		if (y < y1) != (y < y2) && x < (x2-x1)*(y-y1)/(y2-y1)+x1 {
			inside = !inside
		}

		// Check if on an edge
		dx, dy := x2-x1, y2-y1
		if dx*(y-y1) == dy*(x-x1) {
			minX, maxX := sortTwo(x1, x2)
			minY, maxY := sortTwo(y1, y2)
			if x >= minX && x <= maxX && y >= minY && y <= maxY {
				return true
			}
		}
	}

	return inside
}

func withinPolygon(p1 Point, p2 Point, edges []Edge) bool {
	minX, maxX := sortTwo(p1.X, p2.X)
	minY, maxY := sortTwo(p1.Y, p2.Y)

	corners := []Point{
		{minX, minY},
		{minX, maxY},
		{maxX, minY},
		{maxX, maxY},
	}

	for _, c := range corners {
		if !pointInOrOnPolygon(c, edges) {
			return false
		}
	}

	return true
}

func pointOnSegment(p Point, e Edge) bool {
	if orientation(e.Start, e.End, p) != Collinear {
		return false
	}

	return p.X >= min(e.Start.X, e.End.X) &&
		p.X <= max(e.Start.X, e.End.X) &&
		p.Y >= min(e.Start.Y, e.End.Y) &&
		p.Y <= max(e.Start.Y, e.End.Y)
}

func edgesIntersect(e1 Edge, e2 Edge) bool {
	if e1.Start == e2.Start || e1.Start == e2.End || e1.End == e2.Start || e1.End == e2.End {
		return false
	}

	if pointOnSegment(e1.Start, e2) ||
		pointOnSegment(e1.End, e2) ||
		pointOnSegment(e2.Start, e1) ||
		pointOnSegment(e2.End, e1) {
		return false
	}

	o1 := orientation(e1.Start, e1.End, e2.Start)
	o2 := orientation(e1.Start, e1.End, e2.End)
	o3 := orientation(e2.Start, e2.End, e1.Start)
	o4 := orientation(e2.Start, e2.End, e1.End)

	return o1 != o2 && o3 != o4
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func onSegment(p Point, e Edge) bool {
	return (e.Start.X <= max(p.X, e.End.X) && e.Start.X >= min(p.X, e.End.X)) && (e.Start.Y <= max(p.Y, e.End.Y) && e.Start.Y >= min(p.Y, e.End.Y))
}

type Orientation int

const (
	Collinear Orientation = iota
	Clockwise
	CounterClockwise
)

func orientation(p, q, r Point) Orientation {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return Collinear
	}

	if val > 0 {
		return Clockwise
	}

	return CounterClockwise
}

func intersectsPolygon(p1 Point, p2 Point, edges []Edge) bool {
	minX, maxX := sortTwo(p1.X, p2.X)
	minY, maxY := sortTwo(p1.Y, p2.Y)

	rectEdges := []Edge{
		{Point{minX, minY}, Point{minX, maxY}},
		{Point{minX, maxY}, Point{maxX, maxY}},
		{Point{maxX, maxY}, Point{maxX, minY}},
		{Point{maxX, minY}, Point{minX, minY}},
	}

	for _, re := range rectEdges {
		for _, pe := range edges {
			if edgesIntersect(re, pe) {
				return true
			}
		}
	}

	return false
}

type Point struct {
	X int
	Y int
}

type Edge struct {
	Start Point
	End   Point
}

func (p1 Point) GridAreaWith(p2 Point) int {
	minX, maxX := sortTwo(p1.X, p2.X)
	length := maxX - minX + 1

	minY, maxY := sortTwo(p1.Y, p2.Y)
	width := maxY - minY + 1

	return length * width
}

func processInput() ([]Point, []Edge) {
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

	edges := []Edge{}
	for i, point := range points {
		if i == 0 {
			continue
		}

		edges = append(edges, Edge{
			Start: points[i-1],
			End:   point,
		})
	}

	edges = append(edges, Edge{
		Start: points[len(points)-1],
		End:   points[0],
	})

	return points, edges
}
