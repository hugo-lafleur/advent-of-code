package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func format(s string) []string {
	return strings.Split(s, "")
}

func newRock(shape int, i int) []point {
	switch shape {
	case 0:
		return []point{{i, 2}, {i, 3}, {i, 4}, {i, 5}}
	case 1:
		return []point{{i, 3}, {i + 1, 2}, {i + 1, 3}, {i + 1, 4}, {i + 2, 3}}
	case 2:
		return []point{{i, 2}, {i, 3}, {i, 4}, {i + 1, 4}, {i + 2, 4}}
	case 3:
		return []point{{i, 2}, {i + 1, 2}, {i + 2, 2}, {i + 3, 2}}
	default:
		return []point{{i, 2}, {i, 3}, {i + 1, 2}, {i + 1, 3}}
	}
}

func shift(c string, rock []point) []point {
	new := []point{}
	for _, p := range rock {
		switch c {
		case ">":
			new = append(new, point{p.x, p.y + 1})
		case "<":
			new = append(new, point{p.x, p.y - 1})
		}
	}
	return new
}

func fall(rock []point) []point {
	new := []point{}
	for _, p := range rock {
		new = append(new, point{p.x - 1, p.y})
	}
	return new
}

func highest(m map[point]string) int {
	var high int
	for p := range m {
		high = p.x
	}
	for p := range m {
		if p.x > high {
			high = p.x
		}
	}
	return high + 4
}

func leftMost(l []point) int {
	res := l[0].y
	for _, p := range l {
		if p.y < res {
			res = p.y
		}
	}
	return res
}

func notValid(rock []point, m map[point]string) bool {
	for _, p := range rock {
		_, ok := m[p]
		if ok || p.x < 0 || p.y < 0 || p.y >= 7 {
			return true
		}
	}
	return false
}

func solve(s string, max int) int {
	pattern := format(s)
	pattern_index := 0
	mapping := make(map[point]string)
	state := make([]int, 6)
	memory := make(map[string][]int)
	var height_increase int
	var cycle_found bool
	shape := 0
	top := 3
	for shape < max {
		rock := newRock(shape%5, top)
		for {
			new_rock := shift(pattern[pattern_index%len(pattern)], rock)
			pattern_index++
			if notValid(new_rock, mapping) {
				new_rock = rock
			}
			rock = new_rock
			new_rock = fall(rock)
			if notValid(new_rock, mapping) {
				break
			}
			rock = new_rock
		}
		for _, p := range rock {
			mapping[p] = "#"
		}
		top = highest(mapping)
		state[shape%5] = leftMost(rock)
		state[5] = pattern_index % len(pattern)
		shape++
		mem, ok := memory[fmt.Sprint(state)]
		if ok && !cycle_found {
			cycle_found = true
			rocks_per_cycle := shape - mem[0]
			height_per_cycle := top - mem[1]
			remaining_rocks := max - shape + 1
			cycles_remaining := remaining_rocks / rocks_per_cycle
			rocks_remainder := remaining_rocks % rocks_per_cycle
			height_increase = height_per_cycle * cycles_remaining
			shape = max - rocks_remainder + 1
		}
		memory[fmt.Sprint(state)] = []int{shape, top}

	}
	return top - 3 + height_increase
}

func part1(s string) int {
	return solve(s, 2022)
}

func part2(s string) int {
	return solve(s, 1000000000000)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day17/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day17/input.txt")

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
