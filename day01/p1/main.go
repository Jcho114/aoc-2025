package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic("failed to open input file")
	}
	defer file.Close()

	current := 50
	password := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		amount, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic("failed to process line in file")
		}

		if direction == 'L' {
			current = (current - (amount % 100) + 100) % 100
		} else {
			current = (current + (amount % 100)) % 100
		}

		if current == 0 {
			password += 1
		}
	}

	fmt.Println(password)
}
