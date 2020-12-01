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

	var digits []int
	for scanner.Scan() {
		line := scanner.Text()

		strs := strings.Split(line, "")

		digits = strsToInts(strs)
	}

	sum := 0
	for i, d := range digits {
		next := 0
		if i+1 == len(digits) {
			next = digits[0]
		} else {
			next = digits[i+1]
		}

		if d == next {
			sum += next
		}

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
