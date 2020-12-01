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

		largest := 0
		smallest := -1
		for _, digit := range row {

			if digit > largest {
				largest = digit
			}

			if digit < smallest || smallest == -1 {
				smallest = digit
			}
		}
		sum += largest - smallest
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
