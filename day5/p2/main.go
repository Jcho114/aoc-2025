package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	intervals := processInput()

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	merged := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		left, right := interval[0], interval[1]
		if left <= merged[len(merged)-1][1] {
			if merged[len(merged)-1][0] > left {
				merged[len(merged)-1][0] = left
			}
			if merged[len(merged)-1][1] < right {
				merged[len(merged)-1][1] = right
			}
		} else {
			merged = append(merged, interval)
		}
	}

	numFresh := 0
	for _, interval := range merged {
		left, right := interval[0], interval[1]
		numFresh += right - left + 1
	}

	fmt.Println(numFresh)
}

func processInput() [][]int {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	intervals := [][]int{}

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

	return intervals
}
