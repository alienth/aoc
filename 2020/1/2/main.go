package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	ints := strsToInts(lines)

	for _, x := range ints {
		for _, y := range ints {
			for _, z := range ints {
				if x+y+z == 2020 {
					fmt.Println(x * y * z)
					return
				}
			}
		}
	}

	prettyPrint(ints)

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
