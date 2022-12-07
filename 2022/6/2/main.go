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

	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		done := false

		for i := 0; i <= len(chars)-14; i += 1 {
			charMap := make(map[string]bool)

			for _, c := range chars[i : i+14] {
				charMap[c] = true
				if len(charMap) == 14 {
					fmt.Println(i + 14)
					done = true
				}

			}
			if done {
				break
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
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}
