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

	// grid := make(map[int]map[int]*coord)

	for x := 0; x <= largestX; x++ {
		// grid[x] = make(map[int]*coord)
		for y := 0; y <= largestY; y++ {
			closest := closestCoord(x, y, coords)
			// grid[x][y] = closest
			if closest != nil {
				closest.area += 1
				if x == smallestX || x == largestX || y == smallestY || y == largestY {
					closest.infinite = true
				}
			}

		}
	}

	// for y := 0; y <= largestY; y++ {
	// 	for x := 0; x <= largestX; x++ {
	// 		coord := grid[x][y]
	// 		if coord == nil {
	// 			fmt.Print(".")
	// 		} else if x == coord.x && y == coord.y {
	// 			fmt.Print(coord.name)
	// 		} else {
	// 			fmt.Print(strings.ToLower(coord.name))
	// 		}
	// 	}
	// 	fmt.Print("\n")
	// }

	largestArea := 0
	var largestCoord coord
	for _, c := range coords {
		if c.infinite {
			continue
		}
		if c.area >= largestArea {
			largestArea = c.area
			largestCoord = *c
		}
	}
	fmt.Println(largestArea)
	prettyPrint(largestCoord)
}

// Returns nil if we match more than 1
func closestCoord(x, y int, coords []*coord) *coord {
	var shortestC *coord
	shortestDistance := 0

	distances := make(map[int]int)
	for i, c := range coords {
		d := distance(x, y, *c)
		distances[d] += 1

		if i == 0 || d < shortestDistance {
			shortestC = c
			shortestDistance = d
		}
	}
	if distances[shortestDistance] > 1 {
		// fmt.Printf("x: %d, y: %d is equidistant (%d) to two coordinates!\n", x, y, shortestDistance)
		return nil
	}

	return shortestC
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
