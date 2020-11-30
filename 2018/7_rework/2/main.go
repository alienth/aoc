package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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

	timeElapsed := 0

	type worker struct {
		started int
		work    string
	}

	workers := [5]*worker{}

	workersFree := func() []int {
		free := make([]int, 0)
		for i, w := range workers {
			if w.work == "" {
				free = append(free, i)
			}
		}
		return free
	}

	doWork := func() {
		done := make([]string, 0)
		for _, w := range workers {
			if w.work != "" && timeRequired(w.work) == (timeElapsed-w.started) {
				done = append(done, w.work)
				for _, v := range nodes[w.work] {
					inEdges[v]--

					if inEdges[v] == 0 {
						queue = append(queue, v)
					}
				}

				sort.Strings(queue)
				w.work = ""
			}
		}

		sort.Strings(done)
		result += strings.Join(done, "")
	}

	for i := 0; i < 5; i++ {
		workers[i] = &worker{}
	}

	for len(queue) > 0 || len(workersFree()) < 5 {
		for len(queue) > 0 && len(workersFree()) > 0 {
			u := queue[0]
			queue = queue[1:]
			w := workers[workersFree()[0]]
			w.work = u
			w.started = timeElapsed
		}
		timeElapsed += 1
		doWork()
	}

	fmt.Println(timeElapsed)
}

// returns 61 for step A
func timeRequired(s string) int {
	return int([]byte(s)[0]) - int([]byte("A")[0]) + 1 + 60
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
