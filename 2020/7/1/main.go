package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var bags map[string][]string

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	result := 0
	bags = make(map[string][]string)
	for scanner.Scan() {
		newLine := scanner.Text()

		var adj, color string
		var containadj, containcolor string
		i := 0

		fmt.Sscanf(newLine, "%s %s bags contain %d %s %s", &adj, &color, &i, &containadj, &containcolor)
		bags[adj+color] = append(bags[adj+color], containadj+containcolor)

		for x, foo := range strings.Split(newLine, ",") {
			if x == 0 {
				continue
			}
			fmt.Sscanf(foo, "%d %s %s", &i, &containadj, &containcolor)
			bags[adj+color] = append(bags[adj+color], containadj+containcolor)
		}
	}

	canCarry := make(map[string]bool)
	for bag := range bags {
		depth := checkBags(bags, bag)
		if depth > 0 {
			canCarry[bag] = true
		}
		fmt.Println(bag, depth)
	}

	for range canCarry {
		result++
	}

	fmt.Println(result)
}

func checkBags(bags map[string][]string, bag string) int {
	depth := 0
	for _, contained := range bags[bag] {
		if contained == "shinygold" {
			return 1
		} else {
			depth += checkBags(bags, contained)
		}
	}
	return depth
}

func check(group string) int {
	fmt.Println(group)

	seen := make(map[string]bool)
	for _, c := range strings.Split(group, "") {
		seen[c] = true
	}

	return len(seen)
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
