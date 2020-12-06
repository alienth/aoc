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
	line := ""
	for scanner.Scan() {
		newLine := scanner.Text()
		line += newLine

		if newLine == "" {
			result += check(line)
			line = ""
		}

	}
	result += check(line)

	fmt.Println(result)
}

func check(group string) int {
	fmt.Println(group)

	seen := make(map[string]bool)
	for _, c := range strings.Split(group, "") {
		seen[c] = true
	}

	return len(seen)
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
