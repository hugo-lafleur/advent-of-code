package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type Point struct {
	x, y int
}

func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var grid = make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func scoreBFS(start Point, grid [][]byte) int {
	var result int
	var m, n = len(grid), len(grid[0])
	var dq deque.Deque[Point]
	var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var visited = make([]bool, m*n)
	dq.PushBack(start)
	visited[start.x*n+start.y] = true
	for dq.Len() != 0 {
		curr := dq.PopBack()
		for _, dir := range dirs {
			next := Point{curr.x + dir[0], curr.y + dir[1]}
			if next.x < 0 || next.y < 0 || next.x >= m || next.y >= n { //check bounds
				continue
			}
			if visited[next.x*n+next.y] { //check visited
				continue
			}
			if grid[next.x][next.y] != 1+grid[curr.x][curr.y] { //check gradual uphill slope
				continue
			}
			if grid[next.x][next.y] == '9' {
				result++
			}
			dq.PushBack(next)
			visited[next.x*n+next.y] = true

		}
	}
	return result
}

func ratingBFS(start Point, grid [][]byte) int {
	var result int
	var m, n = len(grid), len(grid[0])
	var dq deque.Deque[Point]
	var dirs = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	dq.PushBack(start)
	for dq.Len() != 0 {
		curr := dq.PopBack()
		for _, dir := range dirs {
			next := Point{curr.x + dir[0], curr.y + dir[1]}
			if next.x < 0 || next.y < 0 || next.x >= m || next.y >= n { //check bounds
				continue
			}
			if grid[next.x][next.y] != 1+grid[curr.x][curr.y] { //check gradual uphill slope
				continue
			}
			if grid[next.x][next.y] == '9' {
				result++
			}
			dq.PushBack(next)

		}
	}
	return result
}

func part1(s string) int {
	var grid = parse(s)
	var starts []Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				starts = append(starts, Point{i, j})
			}
		}
	}
	var result int
	for i := range starts {
		result += scoreBFS(starts[i], grid)
	}
	return result
}

func part2(s string) int {
	var grid = parse(s)
	var starts []Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				starts = append(starts, Point{i, j})
			}
		}
	}
	var result int
	for i := range starts {
		result += ratingBFS(starts[i], grid)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day10/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day10/input.txt")

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
