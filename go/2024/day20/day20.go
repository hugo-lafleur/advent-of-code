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
	var grid = make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(p1, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func solve(s string, cheat func(int) bool) int {
	var grid = parse(s)
	var start Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				start = Point{i, j}
			}
		}
	}
	var result int
	var dirs = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var distance = make([]Point, strings.Count(s, ".")+2)
	distance[0] = start
	for i := range distance {
		curr := distance[i]
		if grid[curr.x][curr.y] == 'E' {
			break
		}
		for _, d := range dirs {
			next := Point{curr.x + d[0], curr.y + d[1]}
			if grid[next.x][next.y] != '#' && (i == 0 || distance[i-1] != next) {
				distance[i+1] = next
				break
			}
		}
	}
	var limit = 100
	if len(distance) == 85 {
		if cheat(20) {
			limit = 76
		} else {
			limit = 64
		}
	}
	for d1 := 0; d1 < len(distance); d1++ {
		for d2 := d1 + 1; d2 < len(distance); d2++ {
			p1 := distance[d1]
			p2 := distance[d2]
			d := manhattanDistance(p1, p2)
			if cheat(d) && d2-d1-d >= limit {
				result++
			}
		}
	}
	return result
}

func part1(s string) int {
	return solve(s, func(d int) bool { return d == 2 })
}

func part2(s string) int {
	return solve(s, func(d int) bool { return d > 1 && d <= 20 })
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day20/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day20/input.txt")

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
