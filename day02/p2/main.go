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

			for length := 1; length < len(str)/2+1; length += 1 {
				base := str[:length]
				if len(str)%length != 0 {
					continue
				}

				invalid := true
				for start := length; start < len(str)-length+1; start += length {
					if base != str[start:start+length] {
						invalid = false
						break
					}
				}

				if invalid {
					invalids += curr
					break
				}
			}
		}
	}

	fmt.Println(invalids)
}
