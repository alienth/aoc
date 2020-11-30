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

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	origLine := line
	size := len(line)

	for _, lowerRune := range "abcdefghijklmnopqrstuvwxyz" {
		line = strings.ReplaceAll(origLine, string(lowerRune), "")
		line = strings.ReplaceAll(line, strings.ToUpper(string(lowerRune)), "")

		for {
			oldLine := line
			for _, lowerRune := range "abcdefghijklmnopqrstuvwxyz" {
				lower := string(lowerRune)

				proteinA := fmt.Sprintf("%s%s", lower, strings.ToUpper(lower))
				proteinB := fmt.Sprintf("%s%s", strings.ToUpper(lower), lower)

				line = strings.ReplaceAll(line, proteinA, "")
				line = strings.ReplaceAll(line, proteinB, "")
			}

			if oldLine == line {
				if len(line) < size {
					size = len(line)
				}
				break
			}
		}
	}

	fmt.Println(size)
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
