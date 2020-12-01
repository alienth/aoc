package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	digits := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, "\t")

		digits = append(digits, strsToInts(strs))
	}

	sum := 0
	for _, row := range digits {

		divisor := 0
		dividend := 0
		for i, digit := range row {

			for x, digit2 := range row {
				if i == x {
					continue
				}

				var larger, smaller int
				if digit > digit2 {
					larger = digit
					smaller = digit2
				} else {
					larger = digit2
					smaller = digit
				}

				if larger%smaller == 0 {
					divisor = larger
					dividend = smaller
					break
				}

			}
			if dividend != 0 {
				break
			}

		}
		sum += divisor / dividend
	}

	fmt.Println(sum)

}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func strsToInts(s []string) []int {
	ints := make([]int, len(s))

	for i, str := range s {
		ints[i] = strToInt(str)
	}

	return ints
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}
