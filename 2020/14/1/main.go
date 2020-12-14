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

	mem := make(map[int]int64)
	var clearMask, orMask int64
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "mask =") {
			orMask = 0
			clearMask = 0
			var maskStr string
			fmt.Sscanf(line, "mask = %s", &maskStr)
			for i, x := range strings.Split(maskStr, "") {
				if x == "1" {
					orMask |= 1 << (len(maskStr) - 1 - i)
				} else if x == "0" {
					clearMask |= 1 << (len(maskStr) - 1 - i)
				}
			}
		} else {
			var addr int
			var value int64
			fmt.Sscanf(line, "mem[%d] = %d", &addr, &value)

			value |= orMask
			value &= ^clearMask
			mem[addr] = value
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
