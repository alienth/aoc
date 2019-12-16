package main

import "bufio"
import "fmt"
import "log"
import "math"
import "os"
import "strconv"

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	input := make([]int, 0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, i)
	}

	totalFuel := 0
	for _, m := range input {
		// fmt.Printf("Module is %d mass.\n", m)
		f := getFuel(m)
		// fmt.Printf("Requires %d fuel.\n", f)
		totalFuel += f
	}

	fmt.Println(totalFuel)
}

func getFuel(mass int) int {
	fuel := float64(mass / 3)
	fuel = math.Floor(fuel) - 2
	// fmt.Printf("Requiring %.0f fuel.\n", fuel)
	if fuel >= 0 {
		fuel += float64(getFuel(int(fuel)))
	} else if fuel < 0 {
		fuel = 0
	}

	return int(fuel)
}
