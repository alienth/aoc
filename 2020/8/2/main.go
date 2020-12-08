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

	for x, ins := range inst {
		split := strings.Split(ins, " ")
		it := split[0]
		num, _ := strconv.Atoi(split[1])

		newInst := make([]string, len(inst))

		if it == "nop" {
			if num == x {
				continue
			}

			copy(newInst, inst)
			newInst[x] = "jmp " + split[1]
		} else if it == "jmp" {
			copy(newInst, inst)
			newInst[x] = "nop " + split[1]
		} else {
			continue
		}

		seen := make(map[int]bool)
		i := 0
		acc := 0
		for {
			if i >= len(inst) {
				fmt.Println(acc)
				return
			}
			in := newInst[i]
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
	}

}
