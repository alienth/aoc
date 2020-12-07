package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bagCount struct {
	name  string
	count int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile(`(\d+) (\w+) (\w+)`)
	bags := make(map[string][]bagCount)
	for scanner.Scan() {
		newLine := scanner.Text()

		var adj, color string

		fmt.Sscanf(newLine, "%s %s bags contain", &adj, &color)
		bag := adj + color

		for _, match := range re.FindAllString(newLine, -1) {
			submatch := re.FindStringSubmatch(match)
			count, _ := strconv.Atoi(submatch[1])
			contained := bagCount{name: submatch[2] + submatch[3], count: count}
			bags[bag] = append(bags[bag], contained)
		}
	}

	fmt.Println(checkBags(bags, "shinygold"))
}

func checkBags(bags map[string][]bagCount, bag string) int {
	depth := 0
	for _, contained := range bags[bag] {
		depth += contained.count
		depth += checkBags(bags, contained.name) * contained.count
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
