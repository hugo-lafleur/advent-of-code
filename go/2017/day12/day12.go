package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type list []int

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	res := [][]int{}
	for _, line := range lines {
		intLine := []int{}
		strs := strings.FieldsFunc(line, Split)
		for _, x := range strs {
			n, err := strconv.Atoi(x)
			if err == nil {
				intLine = append(intLine, n)
			}
		}
		res = append(res, intLine)
	}
	return res
}

func isIn(n int, l []int) bool {
	for _, x := range l {
		if x == n {
			return true
		}
	}
	return false
}

func (l *list) connected(n int, programs map[int][]int) {
	*l = append(*l, n)
	for _, p := range programs[n] {
		if !isIn(p, *l) {
			(*l).connected(p, programs)
		}
	}
}

func part1(s string) int {
	listPrograms := format(s)
	programs := make(map[int][]int)
	for _, line := range listPrograms {
		programs[line[0]] = line[1:]
	}
	var res list
	res.connected(0, programs)
	return len(res)
}

func part2(s string) int {
	c := 0
	programList := make(map[int]bool)
	listPrograms := format(s)
	programs := make(map[int][]int)
	for _, line := range listPrograms {
		programs[line[0]] = line[1:]
		programList[line[0]] = true
	}
	for len(programList) != 0 {
		var curr int
		for key := range programList {
			curr = key
		}
		var toDelete list
		toDelete.connected(curr, programs)
		for _, x := range toDelete {
			delete(programList, x)
		}
		c++
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day12/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day12/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
