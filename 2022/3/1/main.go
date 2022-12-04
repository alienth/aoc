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

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	sum := 0
	for _, line := range lines {
		var compartments []string
		compartments = append(compartments, line[0:len(line)/2])
		compartments = append(compartments, line[len(line)/2:len(line)])

		for _, c := range strings.Split(compartments[0], "") {
			if strings.Index(compartments[1], c) != -1 {
				sum += prio(c)
				break
			}
		}
	}

	fmt.Println(sum)

}

func prio(s string) int {
	prioString := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	return strings.Index(prioString, s) + 1
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
