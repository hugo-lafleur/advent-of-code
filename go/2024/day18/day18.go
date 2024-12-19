package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra/v2"
)

type point struct {
	x, y int
}

func parse(s string, bytes int) [][]byte {
	var m, n = 71, 71
	var lines = strings.Split(s, "\n")
	if len(lines) == 25 {
		m, n = 7, 7
	}
	var grid = make([][]byte, m)
	for i := range grid {
		grid[i] = make([]byte, n)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for i := range lines[:bytes] {
		var x, y int
		fmt.Sscanf(lines[i], "%d,%d", &x, &y)
		grid[x][y] = '#'
	}
	return grid
}

func solve(s string, bytes int) int {
	var grid = parse(s, bytes)
	var m, n = len(grid), len(grid[0])
	var graph = dijkstra.NewMappedGraph[point]()
	for i := range m {
		for j := range n {
			graph.AddEmptyVertex(point{i, j})
		}
	}
	var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := range m {
		for j := range n {
			if grid[i][j] == '.' {
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					if x >= 0 && y >= 0 && x < m && y < n && grid[x][y] == '.' {
						graph.AddArc(point{i, j}, point{x, y}, 1)
						graph.AddArc(point{x, y}, point{i, j}, 1)
					}
				}
			}
		}
	}
	best, _ := graph.Shortest(point{0, 0}, point{m - 1, n - 1})
	return int(best.Distance)
}

func part1(s string) int {
	if len(s) == 99 {
		return solve(s, 12)
	}
	return solve(s, 1024)
}

func part2(s string) string {
	var left, right = 0, len(strings.Split(s, "\n")) - 1
	for left < right {
		mid := (left + right) / 2
		if solve(s, mid) == 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	fmt.Println(solve(s, left-1))
	return strings.Split(s, "\n")[left-1]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day18/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day18/input.txt")

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
