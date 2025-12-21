package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	intervals, ids := processInput()
	numFresh := 0

	for _, id := range ids {
		fresh := false

		for _, interval := range intervals {
			left, right := interval[0], interval[1]
			if left <= id && id <= right {
				fresh = true
				break
			}
		}

		if fresh {
			numFresh += 1
		}
	}

	fmt.Println(numFresh)
}

func processInput() ([][]int, []int) {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	intervals := [][]int{}
	ids := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		leftRaw, rightRaw := split[0], split[1]
		left, err := strconv.Atoi(leftRaw)
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(rightRaw)
		if err != nil {
			panic(err)
		}

		intervals = append(intervals, []int{left, right})
	}

	for scanner.Scan() {
		line := scanner.Text()
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	return intervals, ids
}
