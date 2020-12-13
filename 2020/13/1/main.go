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

	target := 0
	nums := make([]int, 0)
	for scanner.Scan() {
		newLine := scanner.Text()
		nums = strsToInts(strings.Split(newLine, ","))

		if target == 0 {
			target = nums[0]
		}
	}

	quickest := 0
	quickestArrive := 0
	arrive := 0
	for _, bus := range nums {
		if bus == 0 {
			continue
		}

		arrive = ((target / bus) * bus) + bus
		if quickest == 0 || arrive-target < quickestArrive {
			quickest = bus
			quickestArrive = arrive - target
		}
	}
	fmt.Println(((quickestArrive) * quickest))
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
