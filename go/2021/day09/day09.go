package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
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
	return []point{{p.x + 1, p.y}, {p.x - 1, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}}
}

func explore(start point, mapping map[point]int, visit map[point]bool) map[point]bool {
	var q deque.Deque[point]
	q.PushBack(start)
	visit[start] = true
	for q.Len() != 0 {
		curr := q.PopFront()
		for _, neigh := range neighs(curr) {
			_, visited := visit[neigh]
			_, ok := mapping[neigh]
			if ok && !visited && mapping[neigh] < 9 {
				q.PushBack(neigh)
				visit[neigh] = true
			}
		}
	}
	return visit
}

func part1(s string) int {
	c := 0
	mapping := format(s)
points:
	for p := range mapping {
		for _, neigh := range neighs(p) {
			n, ok := mapping[neigh]
			if ok && n <= mapping[p] {
				continue points
			}
		}
		c += 1 + mapping[p]
	}
	return c
}

func part2(s string) int {
	mapping := format(s)
	lows := []point{}
points:
	for p := range mapping {
		for _, neigh := range neighs(p) {
			n, ok := mapping[neigh]
			if ok && n <= mapping[p] {
				continue points
			}
		}
		lows = append(lows, p)
	}
	bassins := []int{}
	for _, p := range lows {
		bassin := explore(p, mapping, make(map[point]bool))
		bassins = append(bassins, len(bassin))
	}
	sort.Ints(bassins)
	slices.Reverse(bassins)
	return bassins[0] * bassins[1] * bassins[2]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day09/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day09/input.data")

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
