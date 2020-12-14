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

	mem := make(map[int64]int64)
	var mask int64
	var wildcardMask []int
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask =") {
			mask = 0
			wildcardMask = make([]int, 0)
			var maskStr string
			fmt.Sscanf(line, "mask = %s", &maskStr)

			mask, _ = strconv.ParseInt(strings.ReplaceAll(maskStr, "X", "0"), 2, 64)
			split := strings.Split(maskStr, "")
			for i := len(split) - 1; i >= 0; i-- {
				x := split[i]
				if x == "X" {
					wildcardMask = append(wildcardMask, len(maskStr)-i-1)
				}
			}
		} else {
			var addr int64
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)

			addr |= mask
			origAddr := addr
			for i := 0; i <= 1<<len(wildcardMask)-1; i++ {
				addr = origAddr

				for sourcePosition, destPosition := range wildcardMask {
					bit := (i & (1 << sourcePosition))
					if bit != 0 {
						addr |= 1 << destPosition
					} else {
						addr &= ^(1 << destPosition)
					}
				}
				mem[addr] = value
			}

		}
	}
	var result int64
	for _, v := range mem {
		result += v
	}
	fmt.Println(result)
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
