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

	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()

		instructions = strings.Split(line, ", ")
	}

	startX := 0
	startY := 0

	visited := make(map[string]bool)

	facing := 0 // North == 0, left
	for _, ins := range instructions {
		parts := strings.Split(ins, "")

		direction := parts[0]
		distance, _ := strconv.Atoi(strings.Join(parts[1:], ""))

		if direction == "L" {
			facing -= 1
			if facing < 0 {
				facing = 4 + facing
			}
		} else {
			facing += 1
			if facing > 3 {
				facing = facing - 4
			}
		}

		origX := startX
		origY := startY

		visit := func(x, y int) bool {
			locale := fmt.Sprintf("%d,%d", x, y)
			if visited[locale] {
				fmt.Printf("visited %s twice.\n", locale)
				fmt.Println(abs(x) + abs(y))
				return true
			} else {
				visited[locale] = true
			}

			return false
		}

		switch facing {
		case 0:
			startY += distance
			for y := origY + 1; y <= startY; y++ {
				if visit(startX, y) {
					return
				}
			}
		case 1:
			startX += distance
			for x := origX + 1; x <= startX; x++ {
				if visit(x, startY) {
					return
				}
			}
		case 2:
			startY -= distance
			for y := origY - 1; y >= startY; y-- {
				if visit(startX, y) {
					return
				}
			}
		case 3:
			startX -= distance
			for x := origX - 1; x >= startX; x-- {
				if visit(x, startY) {
					return
				}
			}
		}

	}

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
