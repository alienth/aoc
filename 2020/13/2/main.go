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

	nums := make([]int, 0)
	for scanner.Scan() {
		newLine := scanner.Text()
		nums = strsToInts(strings.Split(newLine, ","))
	}

	largest := 0
	for _, b := range nums {
		if b > largest {
			largest = b
		}
	}

	for t := 0; ; t += jump {
		if check(t, len(nums)-1, nums) {
			fmt.Println(t)
			return
		}
	}
}

var jump = 1

func check(t int, end int, nums []int) bool {
	if end == -1 {
		return true
	}
	for e := end; e >= 0; e-- {
		if nums[e] != 0 {
			end = e
			break
		}
	}
	if (t+end)%nums[end] == 0 {
		jump = lcm(jump, nums[end])
		return check(t, end-1, nums)
	}
	return false
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

func lcm(a, b int) int {
	return abs(a*b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
