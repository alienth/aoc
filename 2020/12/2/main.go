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

	shipX := 0
	shipY := 0
	x := 10
	y := 1
	for _, m := range moves {
		switch m.direction {
		case "F":
			shipX += x * m.num
			shipY += y * m.num
		case "N":
			y += m.num
		case "S":
			y -= m.num
		case "E":
			x += m.num
		case "W":
			x -= m.num
		case "L", "R":
			startX := x
			startY := y
			if m.direction == "L" {
				m.num *= -1
			}
			if m.num == 90 || m.num == -270 {
				x = startY
				y = startX * -1
			} else if m.num == 180 || m.num == -180 {
				x = startX * -1
				y = startY * -1
			} else if m.num == 270 || m.num == -90 {
				x = startY * -1
				y = startX
				break
			}

		}
	}
	fmt.Println(abs(shipX) + abs(shipY))
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
