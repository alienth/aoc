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

	maxId := 0
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, "")

		codePos := 0
		lowerPosition := ""
		binaryPositionCheck := func(i int) bool {
			result := false
			if split[codePos] == lowerPosition {
				result = true
			}
			codePos += 1
			return result
		}

		lowerPosition = "F"
		row := sort.Search(127, binaryPositionCheck)
		lowerPosition = "L"
		column := sort.Search(7, binaryPositionCheck)
		id := (row * 8) + column
		if id > maxId {
			maxId = id
		}
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
