package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	orbits := make(orbitMap)

	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		pair := strings.Split(line, ")")
		orbits.addOrbit(pair[0], pair[1])
	}

	// For all bodies, find distance to target.
	// For all bodies having a connection to both targets, find body with the lowest sum to both targets.
	// Traverse upwards to source, then back down.

	lowest := -1
	lowestParent := ""
	for k, _ := range orbits {
		toYou := orbits.distanceToTarget(k, "YOU") - 1
		toSAN := orbits.distanceToTarget(k, "SAN") - 1
		fmt.Printf("%s is %d to YOU and %d to SAN.\n", k, toYou, toSAN)
		if toYou > 0 && toSAN > 0 {
			total := toYou + toSAN
			if lowest == -1 || total < lowest {
				lowest = total
				lowestParent = k
			}
		}
	}

	fmt.Printf("Lowest parent is %s at a distance of %d.\n", lowestParent, lowest)
}

type orbitMap map[string][]string

func (o orbitMap) addOrbit(from, to string) {
	if _, ok := o[from]; !ok {
		o[from] = make([]string, 0)
	}
	o[from] = append(o[from], to)
}

// Return child orbits
func (o orbitMap) getOrbits(from string) []string {
	if orbits, ok := o[from]; ok {
		return orbits
	}
	return nil
}

// Gets count of all direct orbits.
func (o orbitMap) countDirectOrbits() int {
	total := 0
	for _, children := range o {
		total += len(children)
	}
	return total
}

// Gets total count of indirect orbits.
func (o orbitMap) countIndirectOrbits() int {
	total := 0
	for _, children := range o {
		indirectChildren := 0
		for _, child := range children {
			indirectChildren += o.countTotalChildren(child)
		}
		total += indirectChildren
	}
	return total
}

// Gets recursive count of orbits for a given parent.
func (o orbitMap) countTotalChildren(parent string) int {
	total := 0
	if orbits := o.getOrbits(parent); orbits != nil {
		total += len(orbits)
		for _, child := range orbits {
			total += o.countTotalChildren(child)
		}
	}
	return total
}

// func (o orbitMap) findParent(target string) string {
// 	for parent, child := range o {
// 		if child == target {
// 			return parent
// 		}
// 	}
// 	return ""
// }

func (o orbitMap) distanceToTarget(source, target string) int {
	distance := 0
	// fmt.Println(source, target)
	if orbits := o.getOrbits(source); orbits != nil {
		for _, child := range orbits {
			if child == target {
				// fmt.Printf("%s is a child of %s.\n", target, source)
				return 1
			} else {
				// fmt.Printf("Checking children of %s.\n", child)
				distance += o.distanceToTarget(child, target)
				if distance > 0 {
					// fmt.Printf("%s is %d from %s.\n", target, distance, source)
					return distance + 1
				}

			}
		}
	}
	return distance
}
