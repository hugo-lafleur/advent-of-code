package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var result = make([][]byte, len(lines))
	for i := range lines {
		result[i] = []byte(lines[i])
	}
	return result
}

func simulation(grid [][]byte) int {
	var m, n = len(grid), len(grid[0])
	var curr Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' {
				curr = Point{i, j}
			}
		}
	}
	var dirs = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	var dirIndex = 0
	var visited = make([]int, m*n)
	var steps int
	for {
		visited[curr.x*n+curr.y]++
		if visited[curr.x*n+curr.y] > 4 {
			return -1
		}
		next := Point{curr.x + dirs[dirIndex][0], curr.y + dirs[dirIndex][1]}
		if next.x >= 0 && next.y >= 0 && next.x < n && next.y < m {
			for grid[next.x][next.y] == '#' {
				dirIndex = (dirIndex + 1) % 4
				next = Point{curr.x + dirs[dirIndex][0], curr.y + dirs[dirIndex][1]}
			}
			curr = next
			steps++
		} else {
			var result int
			for i := range visited {
				if visited[i] != 0 {
					result++
				}
			}
			return result
		}
	}
}

func part1(s string) int {
	var grid = parse(s)
	return simulation(grid)
}

func part2(s string) int {
	var grid = parse(s)
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				grid[i][j] = '#'
				if simulation(grid) == -1 {
					result++
				}
				grid[i][j] = '.'
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day06/input.txt")

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
