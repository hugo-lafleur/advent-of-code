package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var result = make([][]byte, len(lines))
	for i := range lines {
		result[i] = []byte(lines[i])
	}
	return result
}

func XMAS(k int) byte {
	switch k {
	case 0:
		return 'X'
	case 1:
		return 'M'
	case 2:
		return 'A'
	case 3:
		return 'S'
	}
	return 0
}

func testXMAS(grid [][]byte, i, j int, dir [2]int) bool {
	var m, n = len(grid), len(grid[0])
	var last = []int{i + 3*dir[0], j + 3*dir[1]}
	if last[0] >= 0 && last[1] >= 0 && last[0] < m && last[1] < n {
		for k := range 4 {
			if grid[i+k*dir[0]][j+k*dir[1]] != XMAS(k) {
				return false
			}
		}
		return true
	}
	return false
}

func part1(s string) int {
	var grid = parse(s)
	var result int
	var dirs = [][2]int{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
	for i := range grid {
		for j := range grid[i] {
			for _, dir := range dirs {
				if testXMAS(grid, i, j, dir) {
					result++
				}
			}
		}
	}
	return result
}

func testShapeXMAS(grid [][]byte, i, j int) bool {
	var m, n = len(grid), len(grid[0])
	var neighbors = [][2]int{{1, 1}, {-1, -1}, {1, -1}, {-1, 1}}
	if i >= 1 && j >= 1 && i < m-1 && j < n-1 {
		var sum1, sum2 byte
		for _, n := range neighbors {
			char := grid[i+n[0]][j+n[1]]
			if char != 'M' && char != 'S' {
				return false
			}
			if n[0] == n[1] {
				sum1 += char
			} else {
				sum2 += char
			}
		}
		return sum1 == 160 && sum2 == 160

	}
	return false
}

func part2(s string) int {
	var grid = parse(s)
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'A' && testShapeXMAS(grid, i, j) {
				result++
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day04/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day04/input.txt")

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
