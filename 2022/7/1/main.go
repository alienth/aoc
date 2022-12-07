package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Directory struct {
	Path      string
	Parent    *Directory
	Subdirs   []*Directory
	Size      int
	TotalSize int
}

func main() {
	f, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	numbers := regexp.MustCompile("^[0-9]")

	dirName := ""
	var currDir, parent *Directory

	dirs := make(map[string]*Directory)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")

		if words[1] == "cd" {
			if words[2] == ".." {
				dirName = currDir.Parent.Path
			} else if strings.HasPrefix(words[2], "/") {
				dirName = words[2]
			} else {
				parent = currDir
				dirName += words[2] + "/"
			}
		} else if numbers.MatchString(words[0]) {
			currDir.Size += strToInt(words[0])
		}

		if _, ok := dirs[dirName]; !ok {
			d := Directory{Path: dirName}

			d.Parent = parent
			if parent != nil {
				parent.Subdirs = append(parent.Subdirs, &d)
			}

			dirs[dirName] = &d
		}

		currDir = dirs[dirName]
	}

	root := dirs["/"]

	root.recordTotalSize()

	answer := 0
	for _, dir := range dirs {
		if dir.TotalSize <= 100000 {
			answer += dir.TotalSize
		}
	}

	fmt.Println(answer)

}

func (d *Directory) recordTotalSize() int {

	d.TotalSize = d.Size
	for _, s := range d.Subdirs {
		d.TotalSize += s.recordTotalSize()
	}

	return d.TotalSize
}

func parentDirName(s string) string {
	split := strings.Split(s, "/")

	return strings.Join(split[0:len(split)-1], "/")
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
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func prettyPrint(x ...interface{}) {
	fmt.Printf("%+v\n", x)
}
