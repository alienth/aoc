package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// result := 0
	maxId := 0
	seenSeats := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")

		frontBack := strings.Join(split[0:7], "")
		prettyPrint(frontBack)

		min := 1
		max := 128
		mid := max / 2
		row := 0
		seat := 0
		for _, f := range split[0:7] {
			if f == "F" {
				max = mid
			} else {
				min = mid + 1
			}
			mid = min + ((max - min) / 2)

			if max == min {
				row = min - 1
				break
			}
		}

		min = 1
		max = 8
		mid = max / 2
		for _, f := range split[7:] {
			if f == "L" {
				max = mid
			} else {
				min = mid + 1
			}
			mid = min + ((max - min) / 2)

			if max == min {
				seat = min - 1
				break
			}
		}

		id := (row * 8) + seat
		seenSeats = append(seenSeats, id)
		if id > maxId {
			maxId = id
		}

	}

	lastK := 0
	sort.Ints(seenSeats)
	for _, k := range seenSeats {
		fmt.Println(k)
		if k > lastK+1 && lastK != 0 {
			fmt.Println(k)
			return
		}
		lastK = k
	}
	fmt.Println(maxId)

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
