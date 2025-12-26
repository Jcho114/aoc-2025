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
		total += findBestJoltage(bank)
	}

	fmt.Println(total)
}

var NumDigits = 12

func findBestJoltage(bank string) int {
	bankList := convertToIntList(bank)
	left, res, remaining := 0, 0, NumDigits

	for left < len(bankList) && remaining > 0 {
		maxValue, maxIndex := -1, -1

		for i := left; i < len(bankList)-remaining+1; i++ {
			if bankList[i] > maxValue {
				maxValue = bankList[i]
				maxIndex = i
			}
		}

		left = maxIndex + 1
		res = res*10 + maxValue
		remaining -= 1
	}

	return res
}

func convertToIntList(bank string) []int {
	res := []int{}
	for _, char := range bank {
		val := int(char - '0')
		res = append(res, val)
	}
	return res
}
