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

	for k, v := range orbits {
		fmt.Printf("%s %v\n", k, v)
		fmt.Println(orbits.countTotalChildren(k))
	}
	direct := orbits.countDirectOrbits()
	indirect := orbits.countIndirectOrbits()
	fmt.Println(direct + indirect)
}

type orbitMap map[string][]string

func (o orbitMap) addOrbit(from, to string) {
	if _, ok := o[from]; !ok {
		o[from] = make([]string, 0)
	}
	o[from] = append(o[from], to)
}

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
	for parent, children := range o {
		indirectChildren := 0
		for _, child := range children {
			indirectChildren += o.countTotalChildren(child)
		}
		total += indirectChildren
		fmt.Printf("%s has %d indirect children.\n", parent, indirectChildren)
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
