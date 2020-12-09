package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	preamble := make([]int, 0)
	target := 1212510616
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		preamble = append(preamble, num)

	}

	total := 0
	set := make([]int, 0)
	start := 0
	for i := 0; i < len(preamble); {
		num := preamble[i]
		total += num
		set = append(set, num)
		if total < target {
			i++
		} else if total > target {
			start++
			i = start
			set = nil
			total = 0
		} else if total == target {
			sort.Ints(set)
			fmt.Println(set[0] + set[len(set)-1])
			return
		}
	}
}

func checkPreamble(preamble []int, check int) bool {
	fmt.Println(preamble)
	for _, x := range preamble {
		for _, y := range preamble {
			if x+y == check {
				return true
			}
		}
	}
	return false

}
