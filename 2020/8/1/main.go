package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	inst := make([]string, 0)
	for scanner.Scan() {
		newLine := scanner.Text()

		inst = append(inst, newLine)
	}

	seen := make(map[int]bool)
	acc := 0
	i := 0
	for {
		in := inst[i]
		split := strings.Split(in, " ")
		num, _ := strconv.Atoi(split[1])
		if _, ok := seen[i]; ok {
			break
		}
		seen[i] = true

		if split[0] == "nop" {
			i += 1
		} else if split[0] == "acc" {
			acc += num
			i += 1
		} else if split[0] == "jmp" {
			i += num
		}
	}
	fmt.Println(acc)

}
