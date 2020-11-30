package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type step struct {
	name         string
	depends      []*step
	dependedOnBy map[string]*step
	done         bool
}

type queue map[string]bool

var readyQueue queue

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	steps := make(map[string]*step)
	for scanner.Scan() {
		line := scanner.Text()

		var s *step
		var ok bool
		stepName := ""
		dependsName := ""

		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &dependsName, &stepName)

		if s, ok = steps[stepName]; !ok {
			s = &step{name: stepName}
			steps[s.name] = s
		}

		var dependedStep *step

		if dependedStep, ok = steps[dependsName]; !ok {
			dependedStep = &step{name: dependsName}
			steps[dependedStep.name] = dependedStep
		}

		if dependedStep.dependedOnBy == nil {
			dependedStep.dependedOnBy = make(map[string]*step)
		}

		dependedStep.dependedOnBy[s.name] = s

		s.depends = append(s.depends, dependedStep)
	}

	stepNames := sortedKeys(steps)

	stepOrder := ""

	readyQueue = make(queue)

	for _, stepName := range stepNames {
		s := steps[stepName]
		if s == nil {
			log.Fatal(stepName)
		}
		if len(s.depends) == 0 {
			readyQueue.enqueue(s.name)
		}
	}

	var w workers
	w.launch(5)

	seconds := 0

	for len(stepOrder) < len(steps) {
		fmt.Println(readyQueue.size(), w.avail(), len(w))
		for readyQueue.size() > 0 && w.avail() > 0 {
			w.enqueue(readyQueue.pop())
		}
		w.workAll()
		for _, done := range w.fetchDone() {
			stepOrder += done
			steps[done].do()
		}
		seconds += 1
		fmt.Println(seconds)
	}

	fmt.Println(stepOrder)
	fmt.Println(seconds)
}

type jobWorker struct {
	work         string // current step
	time         int    // time spent on step
	timeRequired int
	done         string
}

type workers []*jobWorker

func (w *workers) launch(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("launched")
		worker := jobWorker{}
		*w = append(*w, &worker)
	}
	fmt.Println(len(*w))
}

func (w *workers) workAll() {
	for _, worker := range *w {
		worker.doWork()
	}
}

func (w *workers) avail() int {
	free := 0
	for _, worker := range *w {
		if worker.isFree() {
			free += 1
		}
	}

	return free
}

func (w *workers) enqueue(step string) {
	for _, worker := range *w {
		if worker.isFree() {
			worker.work = step
			worker.timeRequired = timeRequired(step)
			return
		}
	}
}

// Returns all completed work from workers
func (w *workers) fetchDone() []string {
	allDone := make([]string, 0)
	for _, worker := range *w {
		done := worker.fetchDone()
		if done != "" {
			allDone = append(allDone, done)
		}

	}

	sort.Strings(allDone)

	return allDone
}

// Fetches completed work from worker
func (w *jobWorker) fetchDone() string {
	if w.done != "" {
		done := w.done
		w.done = ""
		return done
	}

	return ""
}

func (w *jobWorker) isFree() bool {
	if w.work == "" {
		return true
	}

	return false
}

func (w *jobWorker) doWork() {
	if !w.isFree() {
		w.time += 1

		if w.time == w.timeRequired {
			w.done = w.work
			fmt.Printf("Worker finished %s!\n", w.work)
			w.work = ""
			w.time = 0
		}
	}
}

func doStep(s *step) string {
	completed := ""
	log.Printf("Doing step %s.\n", s.name)
	// prettyPrint(*s)

	if !s.done {
		s.do()

		completed = s.name
	} else {
		log.Printf("Step %s was already done.\n", s.name)
		return ""
	}

	for _, n := range sortedKeys(s.dependedOnBy) {
		log.Printf("Step %s is depended on by %s.\n", s.name, n)
		completed += doStep(s.dependedOnBy[n])
	}
	return completed
}

func (s *step) do() string {
	fmt.Printf("Marking %s done.\n", s.name)
	s.done = true
	s.enqueueDependencies()

	return s.name
}

func (s *step) enqueueDependencies() {
	for k, v := range s.dependedOnBy {
		ready := true
		for _, dep := range v.depends {
			if !dep.done {
				fmt.Printf("Can't do step %s yet because it depends on %s, which isn't done.\n", v.name, dep.name)
				ready = false
				break
			}
		}

		if ready {
			readyQueue.enqueue(k)
		}
	}
}

func (q queue) enqueue(stepName string) {
	q[stepName] = true
}

func (q queue) pop() string {
	next := ""
	for k := range q {
		if next == "" || k < next {
			next = k
		}
	}

	delete(q, next)
	return next
}

func (q queue) size() int {
	return len(q)
}

// returns 61 for step A
func timeRequired(s string) int {
	return int([]byte(s)[0]) - int([]byte("A")[0]) + 1 + 60
}

func sortedKeys(m map[string]*step) []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
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
