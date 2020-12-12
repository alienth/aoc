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

	type move struct {
		direction string
		num       int
	}

	var moves []move
	for scanner.Scan() {
		newLine := scanner.Text()

		split := strings.Split(newLine, "")

		num, _ := strconv.Atoi(strings.Join(split[1:], ""))
		moves = append(moves, move{direction: split[0], num: num})
	}

	x := 0
	y := 0
	facing := 0 // East
	for _, m := range moves {

		switch m.direction {
		case "F":
			switch facing {
			case 0:
				x += m.num
			case 1:
				y -= m.num
			case 2:
				x -= m.num
			case 3:
				y += m.num
			}
		case "N":
			y += m.num
		case "S":
			y -= m.num
		case "E":
			x += m.num
		case "W":
			x -= m.num
		case "R":
			turn := m.num / 90
			facing += turn
			facing = wrapInt(facing, 0, 3)
		case "L":
			turn := m.num / 90
			facing -= turn
			facing = wrapInt(facing, 0, 3)
		}
	}
	fmt.Println(abs(x) + abs(y))
}

func wrapInt(i, lower, upper int) int {
	rangeSize := upper - lower + 1
	if i < lower {
		i += rangeSize * ((lower-i)/rangeSize + 1)
	}
	return lower + (i-lower)%rangeSize
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
