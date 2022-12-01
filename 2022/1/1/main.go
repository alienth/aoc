package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	elves := make([]int, 0)
	holding := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			elves = append(elves, holding)
			holding = 0
			continue
		}

		i := strToInt(line)

		holding += i
	}

	elves = append(elves, holding)
	holding = 0

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("Holding: %d\n", elves[0])

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
