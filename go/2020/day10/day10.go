package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []int {
	lines := strings.Split(s, "\n")
	res := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func allPaths(n int, p map[int][]int, cache map[int]int) int {
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	if len(p[n]) == 0 {
		return 1
	}
	res := 0
	for _, x := range p[n] {

		res += allPaths(x, p, cache)
	}
	cache[n] = res
	return res
}

func part1(s string) int {
	list := format(s)
	differences := make(map[int]int)
	jolt := 0
	for {
		curr := 10000
		for i := 0; i < len(list); i++ {
			n := list[i]
			if n <= jolt+3 && n > jolt && n < curr {
				curr = n
			}
		}
		if curr == 10000 {
			break
		} else {
			differences[curr-jolt]++
			jolt = curr
		}
	}
	return differences[1] * (differences[3] + 1)
}

func part2(s string) int {
	list := format(s)
	paths := make(map[int][]int)
	jolt := 0
	for {
		curr := 10000
		paths[jolt] = []int{}
		for i := 0; i < len(list); i++ {
			n := list[i]
			if n <= jolt+3 && n > jolt {
				paths[jolt] = append(paths[jolt], n)
				if n < curr {
					curr = n
				}
			}
		}
		if curr == 10000 {
			break
		} else {
			jolt = curr
		}
	}
	return allPaths(0, paths, make(map[int]int))
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day10/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day10/input.txt")

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
