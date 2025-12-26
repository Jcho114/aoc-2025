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

	"github.com/draffensperger/golp"
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
		wg.Add(1)

		go (func() {
			path := machine.findShortestPath()
			sum.Add(int64(path))
			wg.Done()
			counter.Add(1)
			fmt.Printf("Done With %d/%d Machines\n", counter.Load(), len(machines))
		})()
	}

	wg.Wait()

	fmt.Println(sum.Load())
}

func (machine Machine) findShortestPath() int {
	lp := golp.NewLP(0, machine.numButtons)
	grid := make([][]float64, machine.numJoltages)
	for i := range grid {
		grid[i] = make([]float64, machine.numButtons)
		for j := range grid[i] {
			grid[i][j] = 0.0
		}
	}

	for i, button := range machine.buttons {
		for _, target := range button {
			grid[target][i] = 1
		}
	}

	for i, row := range grid {
		lp.AddConstraint(row, golp.EQ, float64(machine.targetJoltages[i]))
	}

	for i := 0; i < machine.numButtons; i++ {
		lp.SetInt(i, true)
	}

	objFn := make([]float64, machine.numButtons)
	for i := range objFn {
		objFn[i] = 1.0
	}

	lp.SetObjFn(objFn)
	lp.Solve()
	vars := lp.Variables()

	sum := 0
	for _, val := range vars {
		sum += int(math.Round(val))
	}
	return sum
}

type Machine struct {
	lights         []rune
	targetJoltages []int
	numJoltages    int
	buttons        [][]int
	numButtons     int
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
		lights := []rune(first[1 : len(first)-1])

		last := split[len(split)-1]
		joltagesRaw := last[1 : len(last)-1]
		joltagesSplit := strings.Split(joltagesRaw, ",")
		targetJoltages := []int{}
		for _, joltageAscii := range joltagesSplit {
			joltage, err := strconv.Atoi(joltageAscii)
			if err != nil {
				panic(err)
			}
			targetJoltages = append(targetJoltages, joltage)
		}

		numJoltages := len(targetJoltages)

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
			lights,
			targetJoltages,
			numJoltages,
			buttons,
			numButtons,
		}

		machines = append(machines, machine)
	}

	return machines
}
