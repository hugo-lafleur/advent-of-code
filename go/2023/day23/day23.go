package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

type state struct {
	pos     point
	history uint64
	steps   int
}

type pair struct {
	p point
	d int
}

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func pointInPairs(p point, l []pair) bool {
	for _, x := range l {
		if p == x.p {
			return true
		}
	}
	return false
}

func compactMapping(mapping map[point][]pair) map[point][]pair {
	change := true
	for change {
		change = false
	loop:
		for key1, value1 := range mapping {
			for key2, value2 := range mapping {
				if key1 != key2 {
					for i, v1 := range value1 {
						for j, v2 := range value2 {
							if v1.p == v2.p && len(mapping[v1.p]) == 2 && pointInPairs(key1, mapping[v1.p]) && pointInPairs(key2, mapping[v1.p]) {
								delete(mapping, v1.p)
								mapping[key1][i] = pair{key2, v1.d + v2.d}
								mapping[key2][j] = pair{key1, v1.d + v2.d}
								change = true
								break loop
							}
							if len(value2) == 2 && pointInPairs(key1, value2) && v1.p == key2 && !pointInPairs(key2, mapping[v2.p]) {
								delete(mapping, key2)
								mapping[key1][i] = pair{v2.p, v1.d + v2.d}
								change = true
								break loop
							}
							if len(value1) == 1 && len(value2) == 1 && len(mapping[v2.p]) == 1 && v1.p == key2 && key2 != v2.p {
								delete(mapping, key2)
								for key3 := range mapping {
									for i, p := range mapping[key3] {
										if p.p == key2 {
											mapping[key3][i] = pair{v2.p, v2.d + p.d}
										}
									}
								}
								change = true
								break loop
							}
						}
					}
				}
			}
		}
	}
	change = true
	for change {
		change = false
	loop2:
		for key1, value1 := range mapping {
			for key2, value2 := range mapping {
				if key1 != key2 {
					for i, v1 := range value1 {
						for _, v2 := range value2 {
							if len(value2) == 1 && v1.p == key2 && v1.d == 2 && len(value1) >= 2 && v2.p != key1 {
								delete(mapping, key2)
								mapping[key1][i] = pair{v2.p, v1.d + v2.d}
								change = true
								break loop2
							}
						}
					}
				}
			}
		}
	}
	return mapping
}

func part1(s string) int {
	tab := format(s)
	mapping := make(map[point][]pair)
	for i, line := range tab {
		for j, c := range line {
			pnt := point{i, j}
			if c == "." {
				for _, p := range []point{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
					if p.x >= 0 && p.x < len(tab) && tab[p.x][p.y] == "." {
						mapping[pnt] = append(mapping[pnt], pair{p, 1})
					}
				}
			}
			if c == "^" {
				mapping[point{i + 1, j}] = append(mapping[point{i + 1, j}], pair{point{i - 1, j}, 2})
			}
			if c == "v" {
				mapping[point{i - 1, j}] = append(mapping[point{i - 1, j}], pair{point{i + 1, j}, 2})
			}
			if c == ">" {
				mapping[point{i, j - 1}] = append(mapping[point{i, j - 1}], pair{point{i, j + 1}, 2})
			}
			if c == "<" {
				mapping[point{i, j + 1}] = append(mapping[point{i, j + 1}], pair{point{i, j - 1}, 2})
			}
		}
	}
	mapping[point{len(tab) - 1, len(tab[0]) - 2}] = []pair{}
	mapping = compactMapping(mapping)
	index := make(map[point]uint64)
	i := 0
	for key := range mapping {
		index[key] = uint64(i)
		i++
	}
	var res int
	start := state{pos: point{0, 1}, history: 0, steps: 0}
	var q deque.Deque[state]
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopBack()
		curr.history |= (1 << index[curr.pos])
		if curr.pos.x == len(tab)-1 && curr.pos.y == len(tab[0])-2 {
			res = max(res, curr.steps)
			continue
		}
		for _, v := range mapping[curr.pos] {
			if curr.history&(1<<index[v.p]) == 0 {
				newState := state{pos: v.p, history: curr.history, steps: curr.steps + v.d}
				q.PushBack(newState)
			}
		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	mapping := make(map[point][]pair)
	for i, line := range tab {
		for j, c := range line {
			pnt := point{i, j}
			if c != "#" {
				for _, p := range []point{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
					if p.x >= 0 && p.x < len(tab) && tab[p.x][p.y] != "#" {
						mapping[pnt] = append(mapping[pnt], pair{p, 1})
					}
				}
			}
		}
	}
	mapping[point{len(tab) - 1, len(tab[0]) - 2}] = []pair{}
	mapping = compactMapping(mapping)
	index := make(map[point]uint64)
	i := 0
	for key := range mapping {
		index[key] = uint64(i)
		i++
	}
	var res int
	start := state{pos: point{0, 1}, history: 0, steps: 0}
	var q deque.Deque[state]
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopBack()
		curr.history |= (1 << index[curr.pos])
		if curr.pos.x == len(tab)-1 && curr.pos.y == len(tab[0])-2 {
			res = max(res, curr.steps)
		}
		for _, v := range mapping[curr.pos] {
			if curr.history&(1<<index[v.p]) == 0 {
				newState := state{pos: v.p, history: curr.history, steps: curr.steps + v.d}
				q.PushBack(newState)
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day23/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day23/input.txt")

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
