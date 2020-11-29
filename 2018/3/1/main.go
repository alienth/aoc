package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type square struct {
	topLeft     coordinate
	bottomRight coordinate
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	squares := make([]square, 0)

	for scanner.Scan() {
		line := scanner.Text()

		elements := strings.Split(line, " ")

		coords := strsToInts(strings.Split(strings.TrimSuffix(elements[2], ":"), ","))
		size := strsToInts(strings.Split(elements[3], "x"))

		topLeft := coordinate{x: coords[0] + 1, y: coords[1] + 1}

		bottomRight := coordinate{}

		bottomRight.x = topLeft.x + size[0] - 1
		bottomRight.y = topLeft.y + size[1] - 1

		s := square{topLeft: topLeft, bottomRight: bottomRight}

		squares = append(squares, s)
	}

	twoOrMore := 0
	for x := 0; x <= 1000; x++ {
		for y := 0; y <= 1000; y++ {
			found := 0
			for _, s := range squares {
				if x >= s.topLeft.x && y >= s.topLeft.y {
					if x <= s.bottomRight.x && y <= s.bottomRight.y {
						found += 1
					}
				}
			}
			if found >= 2 {
				twoOrMore += 1
			}
		}
	}

	fmt.Println(twoOrMore)

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
