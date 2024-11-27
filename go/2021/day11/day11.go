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

func format(s string) map[point]int {
	lines := strings.Split(s, "\n")
	res := make(map[point]int)
	for i, line := range lines {
		lineSplit := strings.Split(line, "")
		for j, x := range lineSplit {
			p := point{i, j}
			n, _ := strconv.Atoi(x)
			res[p] = n
		}
	}
	return res
}

func neighs(p point) []point {
	return []point{{p.x + 1, p.y + 1}, {p.x - 1, p.y - 1}, {p.x + 1, p.y - 1}, {p.x - 1, p.y + 1}, {p.x + 1, p.y}, {p.x - 1, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}}
}

func hasTen(m map[point]int) bool {
	for _, value := range m {
		if value >= 10 {
			return true
		}
	}
	return false
}

func part1(s string) int {
	c := 0
	mapping := format(s)
	steps := 0
	for steps < 100 {
		flash := make(map[point]bool)
		for key := range mapping {
			mapping[key]++
		}
		for hasTen(mapping) {
			for key := range mapping {
				if mapping[key] > 9 {
					flash[key] = true
					c++
					mapping[key] = 0
					for _, neigh := range neighs(key) {
						_, flashed := flash[neigh]
						_, ok := mapping[neigh]
						if ok && !flashed {
							mapping[neigh]++
						}
					}
				}
			}
		}
		steps++
	}
	return c
}

func part2(s string) int {
	mapping := format(s)
	steps := 0
	for {
		flash := make(map[point]bool)
		for key := range mapping {
			mapping[key]++
		}
		for hasTen(mapping) {
			for key := range mapping {
				if mapping[key] > 9 {
					flash[key] = true
					mapping[key] = 0
					for _, neigh := range neighs(key) {
						_, flashed := flash[neigh]
						_, ok := mapping[neigh]
						if ok && !flashed {
							mapping[neigh]++
						}
					}
				}
			}
		}
		steps++
		if len(flash) == len(mapping) {
			return steps
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day11/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day11/input.txt")

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
