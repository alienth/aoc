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

	frequency := 0

	hitFrequencies := make(map[int]bool)
	hitFrequencies[frequency] = true

	input := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err)
		}

		switch op := string(line[0]); op {
		case "-":
			input = append(input, -value)
		case "+":
			input = append(input, value)
		default:
			log.Fatal("unknown op: ", op)
		}
	}

	for {
		for _, v := range input {
			frequency += v
			if _, ok := hitFrequencies[frequency]; ok {
				fmt.Printf("Hit %d twice.", frequency)
				return
			}
			hitFrequencies[frequency] = true
		}
	}
}
