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

	var trees [][]int

	for scanner.Scan() {
		line := scanner.Text()

		digits := strsToInts(strings.Split(line, ""))

		trees = append(trees, digits)
	}

	highest := 0
	for x, line := range trees {

		for y, tree := range line {

			xEdge := len(trees) - 1
			yEdge := len(line) - 1

			if x == 0 || x == xEdge || y == 0 || y == yEdge {
				continue
			}

			visibleLeft, visibleRight, visibleTop, visibleBottom := 0, 0, 0, 0

			for checkX := x - 1; checkX >= 0; checkX -= 1 {
				visibleLeft += 1
				if trees[checkX][y] >= tree {
					break
				}
			}

			for checkX := x + 1; checkX <= xEdge; checkX += 1 {
				visibleRight += 1
				if trees[checkX][y] >= tree {
					break
				}
			}

			for checkY := y - 1; checkY >= 0; checkY -= 1 {
				visibleTop += 1
				if trees[x][checkY] >= tree {
					break
				}
			}

			for checkY := y + 1; checkY <= yEdge; checkY += 1 {
				visibleBottom += 1
				if trees[x][checkY] >= tree {
					break
				}
			}

			score := visibleLeft * visibleRight * visibleTop * visibleBottom

			if score > highest {
				highest = score
			}
		}

	}

	fmt.Println(highest)
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
