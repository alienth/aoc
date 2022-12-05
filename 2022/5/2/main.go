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

	var stacks, stackLines []string
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "[") {
			break
		}

		stackLines = append(stackLines, line)
	}

	stackHeader := strings.TrimSpace(scanner.Text())
	numOfStacks := strToInt(string(stackHeader[len(stackHeader)-1]))
	stacks = make([]string, numOfStacks)

	for _, line := range stackLines {
		for stack, pos := 0, 1; pos <= len(line); stack, pos = stack+1, pos+4 {
			crate := string(line[pos])
			if crate == " " {
				continue
			}

			stacks[stack] += crate
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "move") {
			continue
		}

		instructions := strings.Split(line, " ")
		numToMove := strToInt(instructions[1])
		from := strToInt(instructions[3])
		to := strToInt(instructions[5])

		take := string(stacks[from-1][0:numToMove])
		stacks[from-1] = strings.TrimPrefix(stacks[from-1], take)
		stacks[to-1] = fmt.Sprintf("%s%s", take, stacks[to-1])

	}

	for _, stack := range stacks {
		fmt.Printf("%s", string(stack[0]))
	}
	fmt.Println("")

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
