package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	points := parseInput()
	distances := []DistEntry{}

	for i := range points {
		for j := i + 1; j < len(points); j++ {
			distances = append(distances, DistEntry{
				p1:   points[i],
				p2:   points[j],
				dist: points[i].dist(points[j]),
			})
		}
	}

	slices.SortFunc(distances, func(a, b DistEntry) int {
		switch {
		case a.dist < b.dist:
			return -1
		case a.dist > b.dist:
			return 1
		default:
			return 0
		}
	})

	indexMap := map[Point]int{}
	for i, point := range points {
		indexMap[point] = i
	}

	set := NewDisjoinSet(len(points))
	for _, entry := range distances {
		p1Index := indexMap[entry.p1]
		p2Index := indexMap[entry.p2]
		set.Join(p1Index, p2Index)

		if set.Converged() {
			res := entry.p1.X * entry.p2.X
			fmt.Println(int(res))
			return
		}
	}
}

type DisjointSet struct {
	parents []int
	ranks   []int
}

func NewDisjoinSet(size int) *DisjointSet {
	parents := make([]int, size)
	ranks := make([]int, size)

	for i := range parents {
		parents[i] = i
	}

	for i := range ranks {
		ranks[i] = 1
	}

	return &DisjointSet{
		parents: parents,
		ranks:   ranks,
	}
}

func (set *DisjointSet) Join(a int, b int) {
	ap, bp := set.Parent(a), set.Parent(b)
	if set.ranks[ap] > set.ranks[bp] {
		set.ranks[ap] += set.ranks[bp]
		set.parents[bp] = ap
	} else {
		set.ranks[bp] += set.ranks[ap]
		set.parents[ap] = bp
	}
}

func (set *DisjointSet) Parent(a int) int {
	if set.parents[a] == a {
		return a
	}

	parent := set.Parent(set.parents[a])
	set.parents[a] = parent
	return parent
}

func (set *DisjointSet) Converged() bool {
	clusterMap := map[int]int{}

	for i := range set.parents {
		parentIndex := set.Parent(i)
		if _, ok := clusterMap[parentIndex]; !ok {
			clusterMap[parentIndex] = 0
		}
		clusterMap[parentIndex] += 1
	}

	return len(clusterMap) == 1
}

type DistEntry struct {
	p1   Point
	p2   Point
	dist float64
}

type Point struct {
	X float64
	Y float64
	Z float64
}

func NewPoint(x float64, y float64, z float64) Point {
	return Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p1 Point) dist(p2 Point) float64 {
	xd := math.Pow(p1.X-p2.X, 2)
	yd := math.Pow(p1.Y-p2.Y, 2)
	zd := math.Pow(p1.Z-p2.Z, 2)
	return math.Sqrt(xd + yd + zd)
}

func parseInput() []Point {
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

		z, err := strconv.Atoi(split[2])
		if err != nil {
			panic(err)
		}

		points = append(points, NewPoint(float64(x), float64(y), float64(z)))
	}

	return points
}
