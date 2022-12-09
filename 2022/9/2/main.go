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

	var head pos
	var tails []*pos

	for i := 1; i <= 9; i += 1 {
		tails = append(tails, &pos{})
	}

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		dir := words[0]
		length := strToInt(words[1])

		for i := 0; i < length; i += 1 {
			if dir == "L" {
				head.x -= 1
			} else if dir == "R" {
				head.x += 1
			} else if dir == "U" {
				head.y += 1
			} else if dir == "D" {
				head.y -= 1
			}

			leader := &head
			for _, tail := range tails {
				if abs(leader.x-tail.x) <= 1 && abs(leader.y-tail.y) <= 1 {
					// close enough already
				} else {
					if leader.x > tail.x {
						tail.x += 1
					} else if leader.x < tail.x {
						tail.x -= 1
					}

					if leader.y > tail.y {
						tail.y += 1
					} else if leader.y < tail.y {
						tail.y -= 1
					}
				}
				leader = tail
			}
			tail := tails[8]
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
