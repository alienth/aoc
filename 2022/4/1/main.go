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

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	total := 0
	for _, line := range lines {
		split := strings.Split(line, ",")
		numbersA := strsToInts(strings.Split(split[0], "-"))
		numbersB := strsToInts(strings.Split(split[1], "-"))

		if numbersA[0] <= numbersB[0] && numbersA[1] >= numbersB[1] {
			total += 1
		} else if numbersB[0] <= numbersA[0] && numbersB[1] >= numbersA[1] {
			total += 1
		}
	}

	fmt.Println(total)

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
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}
