package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		largest, next := -1, -1
		for i, joltageRaw := range bank {
			joltage := int(joltageRaw - '0')
			if joltage > largest && i < len(bank)-1 {
				largest = joltage
				next = -1
			} else if joltage > next {
				next = joltage
			}
		}

		total += 10*largest + next
	}

	fmt.Println(total)
}
