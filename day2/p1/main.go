package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(bytes)

	intervals := strings.Split(input, ",")
	invalids := 0
	for _, interval := range intervals {
		split := strings.Split(interval, "-")

		left, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		right, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for curr := left; curr <= right; curr++ {
			str := strconv.Itoa(curr)
			if len(str)%2 == 1 {
				continue
			}

			half := len(str) / 2
			if str[:half] == str[half:] {
				invalids += curr
			}
		}
	}

	fmt.Println(invalids)
}
