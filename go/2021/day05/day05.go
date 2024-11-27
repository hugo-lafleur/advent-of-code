package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func Split(r rune) bool {
	return r == ',' || r == ' '
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func format(s string) (map[point]int, map[point]int) {
	lines := strings.Split(s, "\n")
	nodiags := make(map[point]int)
	diags := make(map[point]int)
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		x1, _ := strconv.Atoi(lineSplit[0])
		y1, _ := strconv.Atoi(lineSplit[1])
		x2, _ := strconv.Atoi(lineSplit[3])
		y2, _ := strconv.Atoi(lineSplit[4])
		if x1 == x2 || y1 == y2 {
			if x1 == x2 {
				for j := min(y1, y2); j <= max(y1, y2); j++ {
					p := point{x1, j}
					diags[p]++
					nodiags[p]++
				}
			}
			if y1 == y2 {
				for i := min(x1, x2); i <= max(x1, x2); i++ {
					p := point{i, y1}
					diags[p]++
					nodiags[p]++
				}
			}
		} else {
			if y1 > y2 && x1 > x2 || (y1 < y2 && x1 < x2) {
				for i := 0; i <= abs(x1-x2); i++ {
					p := point{min(x1, x2) + i, min(y1, y2) + i}
					diags[p]++
				}
			} else {
				for i := 0; i <= abs(x1-x2); i++ {
					p := point{min(x1, x2) + i, max(y1, y2) - i}
					diags[p]++
				}
			}
		}
	}
	return nodiags, diags
}

func part1(s string) int {
	c := 0
	m, _ := format(s)
	for _, value := range m {
		if value > 1 {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	_, m := format(s)
	for _, value := range m {
		if value > 1 {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day05/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day05/input.txt")

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
