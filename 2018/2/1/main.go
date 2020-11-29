package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	twice := 0
	thrice := 0
	for scanner.Scan() {
		line := scanner.Text()

		letterMap := make(map[string]int)

		letters := strings.Split(line, "")

		for _, l := range letters {
			letterMap[l] += 1
		}

		var foundTwice bool
		var foundThrice bool
		for _, count := range letterMap {
			if count == 2 {
				foundTwice = true
			}
			if count == 3 {
				foundThrice = true
			}
		}

		if foundTwice {
			twice += 1
		}
		if foundThrice {
			thrice += 1
		}
	}

	fmt.Println(twice * thrice)
}
