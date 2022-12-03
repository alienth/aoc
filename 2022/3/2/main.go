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

	sum := 0

	for i := 0; i+3 <= len(lines); i += 3 {
		sacks := lines[i : i+3]

		for _, c := range strings.Split(sacks[0], "") {
			if strings.Index(sacks[1], c) != -1 && strings.Index(sacks[2], c) != -1 {
				sum += prio(c)
				break
			}
		}
	}

	fmt.Println(sum)
}

func prio(s string) int {
	prioString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return strings.Index(prioString, s) + 1
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
