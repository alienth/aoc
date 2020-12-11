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

		distance := 0
		dirSeen := make(map[int]map[int]bool)
		dirSeen[-1] = make(map[int]bool)
		dirSeen[0] = make(map[int]bool)
		dirSeen[1] = make(map[int]bool)
		for {
			distance++

			valid := 0
			for dirX := -1; dirX <= 1; dirX++ {
				for dirY := -1; dirY <= 1; dirY++ {
					if dirY == 0 && dirX == 0 {
						continue
					}
					if dirSeen[dirX][dirY] {
						continue
					}
					newX := x + (dirX * distance)
					newY := y + (dirY * distance)

					if newX < 0 || newY < 0 || newX > len(current[y])-1 || newY > len(current)-1 {
						continue
					}
					valid++

					if current[newY][newX] == "#" {
						used++
						dirSeen[dirX][dirY] = true
					} else if current[newY][newX] == "L" {
						dirSeen[dirX][dirY] = true
					}

				}
			}
			if valid == 0 {
				break
			}
		}
		return used
	}

	for {
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
				} else if seat == "#" && used(x, y) >= 5 {
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
