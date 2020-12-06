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

	result := 0
	groups := make([]string, 0)
	for scanner.Scan() {
		newLine := scanner.Text()

		if newLine == "" {
			result += check(groups)
			groups = nil
		} else {
			groups = append(groups, newLine)
		}
	}
	result += check(groups)

	fmt.Println(result)
}

func check(groups []string) int {
	answer := 0
	seen := make(map[string]int)
	for _, g := range groups {
		for _, c := range strings.Split(g, "") {
			seen[c] += 1
		}
	}

	for _, s := range seen {
		if s == len(groups) {
			answer++
		}
	}

	return answer
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
