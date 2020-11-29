package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	sort.Strings(lines)

	for i, line := range lines {
		nextLine := lines[i+1]

		firstLetters := strings.Split(line, "")
		nextLetters := strings.Split(nextLine, "")

		diff := 0
		diffLetter := 0

		for i, l := range firstLetters {
			if l != nextLetters[i] {
				diff += 1
				diffLetter = i
			}
		}

		if diff == 1 {
			fmt.Println(line)
			fmt.Println(nextLine)
			fmt.Printf("%s%s\n", line[0:diffLetter], line[diffLetter+1:])
			return
		}
	}
}
