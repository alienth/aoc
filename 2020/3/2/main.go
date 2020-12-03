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

	type move struct {
		right int
		down  int
	}
	moves := make([]move, 0)
	moves = append(moves, move{1, 1})
	moves = append(moves, move{3, 1})
	moves = append(moves, move{5, 1})
	moves = append(moves, move{7, 1})
	moves = append(moves, move{1, 2})

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	endResult := 0
	for _, m := range moves {
		right := m.right
		down := m.down
		result := 0
		last := 0
		for i := 0; i < len(lines); i += down {
			line := lines[i]
			lineSize := len(line)

			split := strings.Split(line, "")
			if split[last] == "#" {
				result++
			}

			last += right
			if last >= lineSize {
				last = last - lineSize
			}
		}
		if endResult == 0 {
			endResult = result
		} else {
			endResult = endResult * result
		}

	}

	fmt.Println(endResult)

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
