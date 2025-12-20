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
			first := current
			if current == 0 {
				first = 100
			}

			if amount >= first {
				password += 1 + (amount-first)/100
			}

			current = (current - (amount % 100) + 100) % 100
		} else {
			first := 100 - current

			if amount >= first {
				password += 1 + (amount-first)/100
			}

			current = (current + (amount % 100)) % 100
		}
	}

	fmt.Println(password)
}
