package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Start = "you"
	End   = "out"
)

func main() {
	graph := processInput()
	paths := graph.findPaths(Start)
	fmt.Println(paths)
}

type Graph map[string][]string

func (graph Graph) findPaths(curr string) int {
	if curr == End {
		return 1
	}

	neighbors := graph[curr]
	paths := 0
	for _, neighbor := range neighbors {
		paths += graph.findPaths(neighbor)
	}
	return paths
}

func processInput() Graph {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := Graph{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ": ")
		source := split[0]
		destinations := strings.Split(split[1], " ")
		if _, ok := graph[source]; !ok {
			graph[source] = []string{}
		}

		for _, destination := range destinations {
			if _, ok := graph[destination]; !ok {
				graph[destination] = []string{}
			}
			graph[source] = append(graph[source], destination)
		}
	}

	return graph
}
