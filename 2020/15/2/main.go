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
	var nums []int
	for scanner.Scan() {
		nums = strsToInts(strings.Split(scanner.Text(), ","))
	}

	spokenCount := make(map[int]int)
	spoken := make(map[int][]int)
	lastSpoken := 0

	for i, num := range nums {
		spoken[num] = append(spoken[num], i+1)
		spokenCount[num]++
		lastSpoken = num
	}

	for i := len(nums) + 1; i <= 30000000; i++ {
		speak := 0
		if spokenCount[lastSpoken] == 1 {
			speak = 0
		} else {
			speak = spoken[lastSpoken][len(spoken[lastSpoken])-1] - spoken[lastSpoken][len(spoken[lastSpoken])-2]
		}

		lastSpoken = speak
		spoken[speak] = append(spoken[speak], i)

		spokenCount[speak]++
		result = speak
	}

	fmt.Println(result)
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
