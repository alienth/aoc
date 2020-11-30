package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type coord struct {
	x int
	y int
	// name string
	area     int
	infinite bool
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	coords := make([]*coord, 0)

	var line string
	for scanner.Scan() {
		line = scanner.Text()

		c := coord{}

		fmt.Sscanf(line, "%d, %d", &c.x, &c.y)

		coords = append(coords, &c)
	}

	largestX := 0
	largestY := 0

	for _, c := range coords {
		if c.x > largestX {
			largestX = c.x
		}

		if c.y > largestY {
			largestY = c.y
		}
	}

	smallestX := largestX
	smallestY := largestY

	for _, c := range coords {
		if c.x < smallestX {
			smallestX = c.x
		}

		if c.y < smallestY {
			smallestY = c.y
		}
	}

	fmt.Println(smallestX, smallestY, largestX, largestY)

	safeRegion := 0
	for x := 0; x <= largestX; x++ {
		for y := 0; y <= largestY; y++ {
			d := distanceToAllCoords(x, y, coords)

			if d < 10000 {
				safeRegion += 1
			}
		}
	}

	fmt.Println(safeRegion)
}

func distanceToAllCoords(x, y int, coords []*coord) int {
	total := 0
	for _, c := range coords {
		total += distance(x, y, *c)
	}

	return total
}

func distance(x, y int, c coord) int {
	return abs(x-c.x) + abs(y-c.y)
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
