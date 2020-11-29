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

		topLeft := coordinate{x: coords[0], y: coords[1]}

		bottomRight := coordinate{}

		bottomRight.x = topLeft.x + size[0]
		bottomRight.y = topLeft.y + size[1]

		s := square{topLeft: topLeft, bottomRight: bottomRight}

		squares = append(squares, s)
	}

	for i, a := range squares {
		overlap := false
		for i2, b := range squares {
			if i == i2 {
				continue
			}
			if a.topLeft.x < b.bottomRight.x && a.bottomRight.x > b.topLeft.x {
				if a.topLeft.y < b.bottomRight.y && a.bottomRight.y > b.topLeft.y {
					overlap = true
				}
			}
		}
		if overlap == false {
			fmt.Println(i + 1)
			fmt.Println(a)
			return
		}
	}

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
