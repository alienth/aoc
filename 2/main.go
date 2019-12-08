package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const OpcodeAdd = 1
const OpcodeMultiply = 2
const OpcodeEnd = 99

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]int, 0)
	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	items := strings.Split(line, ",")
	for _, item := range items {
		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}

	for x := 0; x <= 99; x += 1 {
		for y := 0; y <= 99; y += 1 {
			m := make([]int, len(input))
			copy(m, input)
			m[1] = x
			m[2] = y
			result := execute(m)
			if result == 19690720 {
				fmt.Println(x, y)
				fmt.Println((100 * x) + y)
				return
			}
		}
	}
}

func execute(input []int) int {
	last := len(input) - 1

	for pos := 0; pos <= last; pos += 4 {
		param1 := input[input[pos+1]]
		param2 := input[input[pos+2]]
		store := input[pos+3]
		switch instruction := input[pos]; instruction {
		case OpcodeAdd:
			result := param1 + param2
			input[store] = result
		case OpcodeMultiply:
			result := param1 * param2
			input[store] = result
		case OpcodeEnd:
			return input[0]
		default:
			log.Fatal("Unknown opcode ", instruction)
		}
	}
	return 0
}
