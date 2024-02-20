package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type cube struct {
	x, y, z int
}

func format(s string) map[cube]bool {
	res := make(map[cube]bool)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		lineSplit := strings.Split(line, ",")
		a, _ := strconv.Atoi(lineSplit[0])
		b, _ := strconv.Atoi(lineSplit[1])
		c, _ := strconv.Atoi(lineSplit[2])
		res[cube{a, b, c}] = true
	}
	return res
}

func neighs(c cube) []cube {
	res := []cube{}
	for i := -1; i <= 1; i += 2 {
		res = append(res, cube{c.x + i, c.y, c.z})
		res = append(res, cube{c.x, c.y + i, c.z})
		res = append(res, cube{c.x, c.y, c.z + i})
	}
	return res
}

func countNeighs(c cube, mapping map[cube]bool) int {
	res := 0
	for _, neigh := range neighs(c) {
		_, ok := mapping[neigh]
		if ok {
			res++
		}
	}
	return res
}

func maxCoords(mapping map[cube]bool) int {
	var res int
	for c := range mapping {
		res = max(res, c.x, c.y, c.z)
	}
	return res
}

func part1(s string) int {
	c := 0
	cubes := format(s)
	for cube := range cubes {
		c += 6 - countNeighs(cube, cubes)
	}
	return c
}

func part2(s string) int {
	c := 0
	cubes := format(s)
	var q deque.Deque[cube]
	visited := make(map[cube]bool)
	q.PushBack(cube{0, 0, 0})
	visited[cube{0, 0, 0}] = true
	max := maxCoords(cubes)
	for q.Len() != 0 {
		curr := q.PopFront()
		if curr.x >= max && curr.y >= max && curr.z >= max {
			break
		}
		for _, neigh := range neighs(curr) {
			_, seen := visited[neigh]
			_, isLava := cubes[neigh]
			if isLava {
				c++
			}
			if !seen && !isLava {
				q.PushBack(neigh)
				visited[neigh] = true
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day18/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	//fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day18/input.data")

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
