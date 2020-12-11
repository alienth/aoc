package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	seatmap := make([][]string, 0)

	result := 0
	row := 0
	for scanner.Scan() {
		newLine := scanner.Text()

		split := strings.Split(newLine, "")
		seatmap = append(seatmap, split)
		row++
	}

	current := make([][]string, 0)
	used := func(x, y int) int {
		used := 0
		for neighborX := max(0, x-1); neighborX <= min(len(current[x])-1, x+1); neighborX++ {
			for neighborY := max(0, y-1); neighborY <= min(len(current)-1, y+1); neighborY++ {
				if neighborX == x && neighborY == y {
					continue
				}
				if current[neighborY][neighborX] == "#" {
					used++
				}
			}
		}

		return used
	}

	for i := 0; ; i++ {
		current = nil
		for f := range seatmap {
			copyRow := make([]string, len(seatmap[f]))
			copy(copyRow, seatmap[f])
			current = append(current, copyRow)
		}

		changed := 0
		for y := range seatmap {
			for x, seat := range seatmap[y] {
				if seat == "L" && used(x, y) == 0 {
					seatmap[y][x] = "#"
					changed++
				} else if seat == "#" && used(x, y) >= 4 {
					seatmap[y][x] = "L"
					changed++
				}

			}
		}

		if changed == 0 {
			break
		}
	}

	for y := range seatmap {
		for _, seat := range seatmap[y] {
			if seat == "#" {
				result++
			}
		}
	}

	fmt.Println(result)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
