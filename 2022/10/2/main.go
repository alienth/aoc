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
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var queue []int
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")

		op := words[0]
		value := 0
		if len(words) == 2 {
			value = strToInt(words[1])
		}

		if op == "noop" {
			queue = append(queue, 0)

		} else {
			queue = append(queue, 0)
			queue = append(queue, value)
		}
	}

	reg := 1
	for cycle := 1; len(queue) > 0; cycle += 1 {
		draw := "."
		if abs(reg-((cycle-1)%40)) <= 1 {
			draw = "#"
		}
		fmt.Print(draw)

		if cycle%40 == 0 {
			fmt.Print("\n")
		}

		var apply int
		apply, queue = queue[0], queue[1:]
		if apply != 0 {
			reg += apply
		}
	}

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
