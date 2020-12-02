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

	// lines := make([]string, 0)
	type password struct {
		startRange int
		endRange   int
		letter     string
		password   string
	}

	passwords := make([]password, 0)
	for scanner.Scan() {
		line := scanner.Text()

		p := password{}
		fmt.Sscanf(line, "%d-%d %s: %s", &p.startRange, &p.endRange, &p.letter, &p.password)

		p.password = strings.Split(line, " ")[2]
		p.letter = strings.TrimSuffix(p.letter, ":")

		passwords = append(passwords, p)
	}

	count := 0
	for _, p := range passwords {
		split := strings.Split(p.password, "")
		one := split[p.startRange-1] == p.letter
		other := split[p.endRange-1] == p.letter

		if one && !other {
			count++
		}
		if other && !one {
			count++
		}
	}

	fmt.Println(count)

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
