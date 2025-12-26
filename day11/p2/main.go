package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	Start = "svr"
	End   = "out"
)

func main() {
	graph := processInput()
	paths := graph.findPaths(State{curr: Start, dacVisited: false, fftVisited: false}, map[State]int{})
	fmt.Println(paths)
}

type Graph map[string][]string

type State struct {
	curr       string
	dacVisited bool
	fftVisited bool
}

func (graph Graph) findPaths(state State, memo map[State]int) int {
	if _, ok := memo[state]; ok {
		return memo[state]
	}

	if state.curr == End {
		if !state.dacVisited {
			return 0
		}
		if !state.fftVisited {
			return 0
		}
		return 1
	}

	neighbors := graph[state.curr]
	paths := 0
	for _, neighbor := range neighbors {
		newState := state
		if neighbor == "dac" {
			newState.dacVisited = true
		} else if neighbor == "fft" {
			newState.fftVisited = true
		}
		newState.curr = neighbor

		paths += graph.findPaths(newState, memo)
	}

	memo[state] = paths
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
