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

	// ROCK     A
	// PAPER    B
	// SCISCORS C

	total := 0

	throws := []int{1, 2, 3}

	for _, line := range lines {
		game := []rune(line)
		opponent := int(game[0]) - 65
		outcome := string(game[2])
		score := 0

		lose := "X"
		draw := "Y"
		win := "Z"

		throw := 0
		if outcome == draw {
			throw = opponent
			score += 3
		} else if outcome == win {
			throw = opponent + 1
			score += 6
		} else if outcome == lose {
			throw = opponent - 1
		}

		if throw < 0 {
			throw = len(throws) + throw
		}
		throw = throw % len(throws)
		score += throws[throw]

		total += score
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
