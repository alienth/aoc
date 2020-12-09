package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	i := 0
	preamble := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		preamble = append(preamble, num)

		if len(preamble) > 25 {
			check := checkPreamble(preamble[i-25:i], num)
			if check == false {
				fmt.Println(num)
				return
			}
		}
		i++
	}
}

func checkPreamble(preamble []int, check int) bool {
	for _, x := range preamble {
		for _, y := range preamble {
			if x+y == check {
				return true
			}
		}
	}
	return false

}
