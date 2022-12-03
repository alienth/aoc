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

	// ROCK     A X  1
	// PAPER    B Y  2
	// SCISCORS C Z  3

	score := 0
	for _, line := range lines {
		game := []rune(line)
		opponent := int(game[0]) - 64
		you := int(game[2]) - 87

		score += you

		if you == opponent {
			score += 3
		} else if you-opponent == 1 || you-opponent == -2 {
			score += 6
		}
	}

	fmt.Println(score)
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
