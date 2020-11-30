package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type guardLog struct {
	timestamp time.Time
	guard     int
	event     string
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	guardSleep := make(map[int]int)

	// guards to minutes to frequency
	guardSleepMinutes := make(map[int]map[int]int)

	var ignore string
	guard := 0

	var asleep time.Time
	for scanner.Scan() {
		line := scanner.Text()

		elements := strings.Split(line, " ")
		timestampStr := strings.Join(elements[0:2], " ")
		event := strings.Join(elements[2:], " ")

		t, _ := time.Parse("[2006-01-02 15:04]", timestampStr)

		if strings.HasSuffix(event, "begins shift") {
			fmt.Sscanf(event, "%s #%d %s %s", &ignore, &guard, &ignore, &ignore)
		} else if strings.HasSuffix(event, "asleep") {
			asleep = t
		} else if strings.HasSuffix(event, "up") {
			guardSleep[guard] += int(t.Sub(asleep).Minutes())

			if _, ok := guardSleepMinutes[guard]; !ok {
				guardSleepMinutes[guard] = make(map[int]int)
			}

			for i := asleep.Minute(); i < t.Minute(); i++ {
				guardSleepMinutes[guard][i] += 1
			}
		}
	}

	mostFrequentGuard := 0
	mostFrequentMinute := 0
	highMark := 0
	for guard, sleepMinutes := range guardSleepMinutes {
		for minute, frequency := range sleepMinutes {
			if frequency > highMark {
				mostFrequentGuard = guard
				mostFrequentMinute = minute
				highMark = frequency
			}
		}
	}

	fmt.Println(mostFrequentGuard * mostFrequentMinute)
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
