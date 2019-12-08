package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	count := 0
	for i := 193651; i <= 649729; i++ {
		digits := getDigits(i)
		if repeatingDigits(digits) {
			if ascendingDigits(digits) {
				count += 1
			}
		}
	}
	fmt.Println(count)
}

func repeatingDigits(digits []int) bool {
	var lastDigit int
	repeatCount := 1
	for i := 0; i < len(digits); i++ {
		d := digits[i]
		if i == 0 {
			lastDigit = d
			continue
		}

		if d == lastDigit {
			repeatCount += 1
		} else if repeatCount == 2 {
			return true
		} else {
			repeatCount = 1
		}

		lastDigit = d
	}
	// Get tailing case.
	if repeatCount == 2 {
		return true
	}

	return false
}

func ascendingDigits(digits []int) bool {
	for i, d := range digits {
		if i+1 < len(digits) && d > digits[i+1] {
			return false
		}
	}

	return true
}

func getDigits(i int) []int {
	digitChars := strings.Split(fmt.Sprintf("%d", i), "")
	digits := make([]int, 6)

	for i, d := range digitChars {
		digits[i], _ = strconv.Atoi(d)
	}
	return digits
}
