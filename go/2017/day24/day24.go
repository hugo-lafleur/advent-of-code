package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type component struct {
	a, b int
}

type bridge struct {
	last int
	used map[component]bool
}

type queue []bridge

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Push(s bridge) {
	*q = append(*q, s)
}

func (q *queue) Pop() (bridge, bool) {
	if q.IsEmpty() {
		return bridge{}, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func format(s string) []component {
	lines := strings.Split(s, "\n")
	res := []component{}
	for _, line := range lines {
		strs := strings.Split(line, "/")
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])
		res = append(res, component{a, b})
	}
	return res
}

func next(b bridge) []bridge {
	res := []bridge{}
	for key, value := range b.used {
		if !value && (key.a == b.last || key.b == b.last) {
			copy := make(map[component]bool)
			for k, v := range b.used {
				copy[k] = v
			}
			copy[key] = true
			var n int
			if key.a == b.last {
				n = key.b
			} else {
				n = key.a
			}
			res = append(res, bridge{n, copy})
		}
	}
	return res
}

func strenght(b bridge) int {
	res := 0
	for key, value := range b.used {
		if value {
			res += key.a + key.b
		}
	}
	return res
}

func length(b bridge) int {
	res := 0
	for _, value := range b.used {
		if value {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(s string) int {
	c := 0
	components := format(s)
	used := make(map[component]bool)
	for _, c := range components {
		used[c] = false
	}
	var queue queue
	queue.Push(bridge{0, used})
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		toAdd := next(curr)
		for _, b := range toAdd {
			queue.Push(b)
		}
		c = max(c, strenght(curr))
	}
	return c
}

func part2(s string) int {
	components := format(s)
	used := make(map[component]bool)
	for _, c := range components {
		used[c] = false
	}
	var queue queue
	queue.Push(bridge{0, used})
	best := bridge{0, used}
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		toAdd := next(curr)
		for _, b := range toAdd {
			queue.Push(b)
		}
		if length(curr) > length(best) {
			best = curr
		}
		if length(curr) == length(best) {
			if strenght(curr) > strenght(best) {
				best = curr
			}
		}
	}
	return strenght(best)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day24/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day24/input.data")

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
