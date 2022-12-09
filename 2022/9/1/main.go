package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	visited := make(map[int]map[int]bool)

	var head, tail pos
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		dir := words[0]
		length := strToInt(words[1])

		for i := 0; i < length; i += 1 {
			oldHead := head
			if dir == "L" {
				head.x -= 1
			} else if dir == "R" {
				head.x += 1
			} else if dir == "U" {
				head.y += 1
			} else if dir == "D" {
				head.y -= 1
			}

			if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
				// close enough already
			} else if dir == "U" || dir == "D" {
				// head moved and tail was right behind it
				if tail.x == head.x {
					tail.y = oldHead.y
				} else {
					// diagonal

					if dir == "U" {
						tail.x = head.x
						tail.y = head.y - 1

					} else if dir == "D" {
						tail.x = head.x
						tail.y = head.y + 1
					}
				}

			} else if dir == "R" || dir == "L" {
				if tail.y == head.y {
					tail.x = oldHead.x
				} else {
					// diagonal
					if dir == "R" {
						tail.y = head.y
						tail.x = head.x - 1

					} else if dir == "L" {
						tail.y = head.y
						tail.x = head.x + 1
					}
				}

			}

			if _, ok := visited[tail.x]; !ok {
				visited[tail.x] = make(map[int]bool)
			}
			visited[tail.x][tail.y] = true

		}

	}

	total := 0
	for _, m := range visited {
		total += len(m)
	}

	fmt.Println(total)
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
