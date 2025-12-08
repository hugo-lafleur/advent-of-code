package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y, z int
}

type state struct {
	d      int
	p1, p2 point
}

func format(s string) []point {
	result := []point{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		result = append(result, point{x, y, z})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(p1, p2 point) int {
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y) + (p1.z-p2.z)*(p1.z-p2.z)
}

func part1(s string) int {
	points := format(s)
	connections := 1000
	if len(points) == 20 {
		connections = 10
	}
	stack := []state{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			stack = append(stack, state{dist(points[i], points[j]), points[i], points[j]})
		}
	}
	slices.SortFunc(stack, func(a, b state) int { return a.d - b.d })
	parent := make(map[point]point)
	for _, p := range points {
		parent[p] = p
	}
	var find func(p point) point
	find = func(p point) point {
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}
	union := func(x, y point) {
		xR, yR := find(x), find(y)
		if xR != yR {
			if yR.x < xR.x {
				parent[xR] = yR
			} else {
				parent[yR] = xR
			}
		}
	}
	for i := range connections {
		curr := stack[i]
		_, p1, p2 := curr.d, curr.p1, curr.p2
		union(p1, p2)
	}
	for _, p := range points {
		find(p)
	}
	circuit := make(map[point]int)
	for p := range parent {
		circuit[parent[p]]++
	}
	sizes := []int{}
	for _, s := range circuit {
		sizes = append(sizes, s)
	}
	slices.SortFunc(sizes, func(a, b int) int { return b - a })
	return sizes[0] * sizes[1] * sizes[2]
}

func part2(s string) int {
	points := format(s)
	stack := []state{}
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			stack = append(stack, state{dist(points[i], points[j]), points[i], points[j]})
		}
	}
	slices.SortFunc(stack, func(a, b state) int { return a.d - b.d })
	parent := make(map[point]point)
	for _, p := range points {
		parent[p] = p
	}
	var find func(p point) point
	find = func(p point) point {
		if parent[p] != p {
			parent[p] = find(parent[p])
		}
		return parent[p]
	}
	union := func(x, y point) {
		xR, yR := find(x), find(y)
		if xR != yR {
			if yR.x < xR.x {
				parent[xR] = yR
			} else {
				parent[yR] = xR
			}
		}
	}
	checkCircuit := func() bool {
		for i := range points {
			for j := i + 1; j < len(points); j++ {
				if find(points[i]) != find(points[j]) {
					return false
				}
			}
		}
		return true
	}
	for _, curr := range stack {
		_, p1, p2 := curr.d, curr.p1, curr.p2
		union(p1, p2)
		if checkCircuit() {
			return curr.p1.x * curr.p2.x
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day08/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day08/input.txt")

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
