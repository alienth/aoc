package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	wires := make([][]point, 2)

	s := bufio.NewScanner(f)
	for i := 0; i <= 1; i++ {
		s.Scan()
		t := s.Text()
		wires[i] = graph(strings.Split(t, ","))
	}

	intersections := make([]point, 0)

	wireAsteps := 0
	closestSteps := 0
	var closestStepsPoint point
	for i := 0; i < len(wires[0])-1; i++ {
		pointA := wires[0][i]
		wireAStart := pointA
		pointB := wires[0][i+1]
		segA := segment{pointA, pointB}

		wireBsteps := 0

		for x := 0; x < len(wires[1])-1; x++ {
			pointA := wires[1][x]
			pointB := wires[1][x+1]
			segB := segment{pointA, pointB}

			intersection := getIntersection(segA, segB)
			if intersection != nil {
				totalSteps := wireAsteps + wireBsteps
				finalStepsA := int(math.Abs(float64(wireAStart.x-intersection.x)) + math.Abs(float64(wireAStart.y-intersection.y)))
				finalStepsB := int(math.Abs(float64(pointA.x-intersection.x)) + math.Abs(float64(pointA.y-intersection.y)))
				totalSteps += finalStepsA + finalStepsB
				fmt.Printf("Final A steps: %d\nFinal B steps: %d\n", finalStepsA, finalStepsB)
				if closestSteps == 0 || closestSteps > totalSteps {
					closestSteps = totalSteps
					closestStepsPoint = *intersection
				}
				intersections = append(intersections, *intersection)
			}
			steps := int(math.Abs(float64(pointA.x-pointB.x)) + math.Abs(float64(pointA.y-pointB.y)))
			wireBsteps += steps
		}
		steps := int(math.Abs(float64(pointA.x-pointB.x)) + math.Abs(float64(pointA.y-pointB.y)))
		// fmt.Printf("Steps between points %s and %s: %d\n", pointA, pointB, steps)
		wireAsteps += steps
	}
	fmt.Printf("Closest steps %d at point %s.\n", closestSteps, closestStepsPoint)

	var closest point
	closestDistance := 0
	for _, i := range intersections {
		fmt.Println(i)
		distance := int(math.Abs(float64(i.x)) + math.Abs(float64(i.y)))
		if closestDistance == 0 || closestDistance > distance {
			closestDistance = distance
			closest = i
		}
	}
	fmt.Println(closest, closestDistance)
}

type point struct {
	x int
	y int
}

type segment struct {
	pointA point
	pointB point
}

func graph(wire []string) []point {
	var x, y int

	points := make([]point, 0)
	start := point{0, 0}
	points = append(points, start)
	for _, i := range wire {
		dir := string(i[0])
		distance, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case "R":
			x += distance
		case "L":
			x -= distance
		case "U":
			y += distance
		case "D":
			y -= distance
		default:
			log.Fatal("Unknown direction ", dir)
		}

		p := point{x, y}
		points = append(points, p)
	}
	// fmt.Println(wire)
	// fmt.Println(points)
	return points
}

func getIntersection(segA, segB segment) *point {
	segAPoints := segA.getAllPoints()
	segBPoints := segB.getAllPoints()

	for _, pA := range segAPoints {
		for _, pB := range segBPoints {
			if pA.x == pB.x && pA.y == pB.y {
				return &pA
			}
		}
	}

	return nil
}

func (seg *segment) getAllPoints() []point {
	points := make([]point, 0)
	if seg.pointA.x == seg.pointB.x {
		// vertcal line
		var lower, higher int
		if seg.pointA.y < seg.pointB.y {
			lower = seg.pointA.y
			higher = seg.pointB.y
		} else {
			lower = seg.pointB.y
			higher = seg.pointA.y
		}
		for y := lower; y <= higher; y += 1 {
			p := point{x: seg.pointA.x, y: y}
			points = append(points, p)
		}
	} else if seg.pointA.y == seg.pointB.y {
		// horizontal line
		var lower, higher int
		if seg.pointA.x < seg.pointB.x {
			lower = seg.pointA.x
			higher = seg.pointB.x
		} else {
			lower = seg.pointB.x
			higher = seg.pointA.x
		}

		for x := lower; x <= higher; x += 1 {
			p := point{x: x, y: seg.pointA.y}
			points = append(points, p)
		}
	}

	return points
}
