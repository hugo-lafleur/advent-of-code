package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type Point struct {
	x, y int
}

type PointFloat struct {
	x, y float64
}

func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var grid = make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func dfs(grid [][]byte) [][]Point {
	var m, n = len(grid), len(grid[0])
	var result [][]Point
	var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dq deque.Deque[Point]
	var visited = make(map[int]bool)
	for i := range grid {
		for j := range grid[i] {
			start := Point{i, j}
			var group []Point
			if visited[i*n+j] {
				continue
			}
			dq.PushBack(start)
			visited[i*n+j] = true
			for dq.Len() != 0 {
				curr := dq.PopBack()
				group = append(group, curr)
				for _, dir := range dirs {
					next := Point{curr.x + dir[0], curr.y + dir[1]}
					if next.x < 0 || next.y < 0 || next.x >= m || next.y >= n {
						continue
					}
					if visited[next.x*n+next.y] || grid[next.x][next.y] != grid[curr.x][curr.y] {
						continue
					}
					dq.PushBack(next)
					visited[next.x*n+next.y] = true
				}
			}
			result = append(result, group)
		}
	}
	return result
}

func bordersList(group []Point) []PointFloat {
	var result []PointFloat
	var dirs = [][2]float64{{0.25, 0}, {-0.25, 0}, {0, 0.25}, {0, -0.25}}
	for i := range group {
		x, y := float64(group[i].x), float64(group[i].y)
		for _, dir := range dirs {
			if !slices.Contains(group, Point{group[i].x + int(4*dir[0]), group[i].y + int(4*dir[1])}) {
				result = append(result, PointFloat{x + dir[0], y + dir[1]})
			}
		}
	}
	return result
}

func perimeter(group []Point) int {
	return len(bordersList(group))
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func numberOfSides(group []Point) int {
	var borders = bordersList(group)
	var parent = make(map[PointFloat]PointFloat)
	for _, b := range borders {
		parent[b] = b
	}
	var Find func(p PointFloat) PointFloat
	Find = func(p PointFloat) PointFloat {
		if parent[p] != p {
			parent[p] = Find(parent[p])
		}
		return parent[p]
	}
	var Union = func(p1, p2 PointFloat) {
		p1Parent := Find(p1)
		p2Parent := Find(p2)
		if p1Parent != p2Parent {
			parent[p1Parent] = p2Parent
		}
	}
	for i := range borders {
		for j := i + 1; j < len(borders); j++ {
			b1, b2 := borders[i], borders[j]
			if (float64(int(b1.x)) != b1.x && b1.x == b2.x && abs(b1.y-b2.y) == 1) || (float64(int(b1.y)) != b1.y && b1.y == b2.y && abs(b1.x-b2.x) == 1) {
				Union(b1, b2)
			}
		}
	}
	var result int
	for key, val := range parent {
		if key == val {
			result++
		}
	}
	return result
}

func part1(s string) int {
	var grid = parse(s)
	var groups = dfs(grid)
	var result int
	for _, g := range groups {
		result += len(g) * perimeter(g)
	}
	return result
}

func part2(s string) int {
	var grid = parse(s)
	var groups = dfs(grid)
	var result int
	for _, g := range groups {
		result += len(g) * numberOfSides(g)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day12/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day12/input.txt")

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
