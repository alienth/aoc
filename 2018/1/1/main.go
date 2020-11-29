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
	for scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			log.Fatal(err)
		}

		switch op := string(line[0]); op {
		case "-":
			frequency -= value
		case "+":
			frequency += value
		default:
			log.Fatal("unknown op: ", op)
		}
	}

	fmt.Println(frequency)
}
