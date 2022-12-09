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

	// x y
	var trees [][]int

	for scanner.Scan() {
		line := scanner.Text()

		digits := strsToInts(strings.Split(line, ""))

		trees = append(trees, digits)
	}

	visible := 0
	for x, line := range trees {

		for y, tree := range line {

			xEdge := len(trees) - 1
			yEdge := len(line) - 1

			if x == 0 || x == xEdge || y == 0 || y == yEdge {
				visible += 1
				continue
			}

			visibleLeft, visibleRight, visibleTop, visibleBottom := true, true, true, true
			for checkX := x - 1; checkX >= 0; checkX -= 1 {
				if trees[checkX][y] >= tree {
					visibleLeft = false
				}
			}

			for checkX := x + 1; checkX <= xEdge; checkX += 1 {
				if trees[checkX][y] >= tree {
					visibleRight = false
				}
			}

			for checkY := y - 1; checkY >= 0; checkY -= 1 {
				if trees[x][checkY] >= tree {
					visibleTop = false
				}
			}

			for checkY := y + 1; checkY <= yEdge; checkY += 1 {
				if trees[x][checkY] >= tree {
					visibleBottom = false
				}
			}

			if visibleTop || visibleBottom || visibleLeft || visibleRight {
				visible += 1
			}
		}

	}

	fmt.Println(visible)
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
