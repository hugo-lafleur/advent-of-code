package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

func format(s string) []point {
	lines := strings.Split(s, "\n")
	result := []point{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		result = append(result, point{a, b})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(s string) int {
	list := format(s)
	result := 0
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			result = max(result, abs(list[i].x-list[j].x+1)*abs(list[i].y-list[j].y+1))
		}
	}
	return result
}

func check(p1, p2 point, exterior map[point]bool) bool {
	xmin, xmax := min(p1.x, p2.x), max(p1.x, p2.x)
	ymin, ymax := min(p1.y, p2.y), max(p1.y, p2.y)
	for x := xmin; x <= xmax; x++ {
		if exterior[point{x, ymin}] || exterior[point{x, ymax}] {
			return false
		}
	}
	for y := ymin; y <= ymax; y++ {
		if exterior[point{xmin, y}] || exterior[point{xmax, y}] {
			return false
		}
	}
	return true
}

func part2(s string) int {
	list := format(s)
	list = append(list, list[0])
	xs := []int{}
	ys := []int{}
	for i := range list {
		xs = append(xs, list[i].x)
		ys = append(ys, list[i].y)
	}
	slices.Sort(xs)
	slices.Sort(ys)
	xs = slices.Compact(xs)
	ys = slices.Compact(ys)
	xmap := map[int]int{}
	xback := map[int]int{}
	ymap := map[int]int{}
	yback := map[int]int{}
	for i, val := range xs {
		xmap[val] = i
		xback[i] = val
	}
	for i, val := range ys {
		ymap[val] = i
		yback[i] = val
	}
	for i := range list {
		list[i].x = xmap[list[i].x]
		list[i].y = ymap[list[i].y]
	}
	green := map[point]bool{}
	result := 0
	xmin, xmax := -1, len(xs)
	ymin, ymax := -1, len(ys)
	for i := range len(list) - 1 {
		p1, p2 := list[i], list[i+1]
		if p1.x == p2.x {
			for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
				green[point{p1.x, y}] = true
			}
		} else {
			for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
				green[point{x, p1.y}] = true
			}
		}
	}
	exterior := map[point]bool{}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var dq deque.Deque[point]
	dq.PushBack(point{xmin, ymin})
	for dq.Len() != 0 {
		curr := dq.PopBack()
		if curr.x < xmin || curr.y < ymin || curr.x > xmax || curr.y > ymax {
			continue
		}
		exterior[curr] = true
		for _, dir := range dirs {
			nx, ny := curr.x+dir[0], curr.y+dir[1]
			next := point{nx, ny}
			if !green[next] && !exterior[next] {
				dq.PushBack(next)
			}
		}
	}
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			area := (abs(xback[list[i].x]-xback[list[j].x]) + 1) * (abs(yback[list[i].y]-yback[list[j].y]) + 1)
			if area > result && check(list[i], list[j], exterior) {
				result = area
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day09/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day09/input.txt")

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
