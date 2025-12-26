package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	On  = '#'
	Off = '.'
)

func main() {
	machines := processInput()

	wg := sync.WaitGroup{}
	var sum atomic.Int64
	var counter atomic.Int64

	for _, machine := range machines {
		go (func() {
			path := machine.findShortestPath(0, 0)
			sum.Add(int64(path))
			wg.Done()
			counter.Add(1)
			fmt.Printf("Done With %d/%d Machines\n", counter.Load(), len(machines))
		})()

		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(sum.Load())
}

func (machine Machine) swapButtons(a, b int) {
	machine.buttons[a], machine.buttons[b] = machine.buttons[b], machine.buttons[a]
}

func (machine Machine) pressButton(index int) {
	swaps := machine.buttons[index]
	for _, light := range swaps {
		if machine.currentLights[light] == On {
			machine.currentLights[light] = Off
		} else {
			machine.currentLights[light] = On
		}
	}
}

func (machine Machine) findShortestPath(index int, cost int) int {
	if index >= machine.numLights {
		if machine.checkLights() {
			return cost
		}
		return math.MaxInt
	}

	if machine.checkLights() {
		return cost
	}

	min := math.MaxInt
	for targetIndex := index; targetIndex < machine.numButtons; targetIndex++ {
		machine.swapButtons(index, targetIndex)
		machine.pressButton(index)

		path := machine.findShortestPath(index+1, cost+1)
		if path < min {
			min = path
		}

		machine.pressButton(index)
		machine.swapButtons(index, targetIndex)
	}

	return min
}

func (machine Machine) checkLights() bool {
	for i := range machine.currentLights {
		if machine.currentLights[i] != machine.targetLights[i] {
			return false
		}
	}
	return true
}

type Machine struct {
	currentLights []rune
	targetLights  []rune
	numLights     int
	joltages      []int
	buttons       [][]int
	numButtons    int
}

func processInput() []Machine {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	machines := []Machine{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		first := split[0]
		targetLights := []rune(first[1 : len(first)-1])

		numLights := len(targetLights)

		currentLights := make([]rune, numLights)
		for i := range currentLights {
			currentLights[i] = Off
		}

		last := split[len(split)-1]
		joltagesRaw := last[1 : len(last)-1]
		joltagesSplit := strings.Split(joltagesRaw, ",")
		joltages := []int{}
		for _, joltageAscii := range joltagesSplit {
			joltage, err := strconv.Atoi(joltageAscii)
			if err != nil {
				panic(err)
			}
			joltages = append(joltages, joltage)
		}

		buttons := [][]int{}
		for _, buttonRaw := range split[1 : len(split)-1] {
			buttonRaw = buttonRaw[1 : len(buttonRaw)-1]
			targetsSplit := strings.Split(buttonRaw, ",")
			targets := []int{}

			for _, targetAscii := range targetsSplit {
				target, err := strconv.Atoi(targetAscii)
				if err != nil {
					panic(err)
				}
				targets = append(targets, target)
			}

			buttons = append(buttons, targets)
		}

		numButtons := len(buttons)

		machine := Machine{
			currentLights,
			targetLights,
			numLights,
			joltages,
			buttons,
			numButtons,
		}

		machines = append(machines, machine)
	}

	return machines
}
