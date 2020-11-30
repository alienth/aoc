package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type queue map[string]bool

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	nodes := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		u := ""
		v := ""
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &u, &v)

		var n []string
		var ok bool

		if n, ok = nodes[u]; !ok {
			n = make([]string, 0)
		}

		n = append(n, v)
		nodes[u] = n
	}

	inEdges := make(map[string]int)

	for u := range nodes {
		for _, v := range nodes[u] {
			inEdges[v]++
		}
	}

	queue := make([]string, 0)

	for u := range nodes {
		if _, ok := inEdges[u]; !ok {
			queue = append(queue, u)
		}
	}

	result := ""

	sort.Strings(queue)

	for len(queue) > 0 {
		sort.Strings(queue)
		u := queue[0]
		queue = queue[1:]
		result += u
		for _, v := range nodes[u] {
			inEdges[v]--

			if inEdges[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	fmt.Println(result)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func strsToInts(s []string) []int {
	ints := make([]int, len(s))

	for i, str := range s {
		ints[i] = strToInt(str)
	}

	return ints
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}
