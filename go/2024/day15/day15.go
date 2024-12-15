package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

func parse(s string) ([][]byte, string) {
	var lines = strings.Split(s, "\n")
	var grid [][]byte
	var i int
	for i < len(lines) {
		if lines[i] == "" {
			i++
			break
		}
		grid = append(grid, []byte(lines[i]))
		i++
	}
	var moves string
	for i < len(lines) {
		moves += lines[i]
		i++
	}
	return grid, moves
}

func nextPoint(p Point, m rune) Point {
	switch m {
	case '^':
		return Point{p.x - 1, p.y}
	case 'v':
		return Point{p.x + 1, p.y}
	case '<':
		return Point{p.x, p.y - 1}
	case '>':
		return Point{p.x, p.y + 1}
	}
	return p
}

func backward(m rune) rune {
	switch m {
	case '^':
		return 'v'
	case 'v':
		return '^'
	case '<':
		return '>'
	case '>':
		return '<'
	}
	return rune(0)
}

func sumGPSCoordinates(grid [][]byte) int {
	var result int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' || grid[i][j] == '[' {
				result += 100*i + j
			}
		}
	}
	return result
}

func widerGrid(grid [][]byte) [][]byte {
	var result = make([][]byte, len(grid))
	for i := range grid {
		result[i] = make([]byte, 2*len(grid[i]))
		for j := range grid[i] {
			switch grid[i][j] {
			case '#':
				result[i][2*j] = '#'
				result[i][2*j+1] = '#'
			case 'O':
				result[i][2*j] = '['
				result[i][2*j+1] = ']'
			case '.':
				result[i][2*j] = '.'
				result[i][2*j+1] = '.'
			case '@':
				result[i][2*j] = '@'
				result[i][2*j+1] = '.'
			}
		}
	}
	return result
}

func boxesList(grid [][]byte, p Point, m rune) (bool, []Point) {
	switch grid[p.x][p.y] {
	case '@':
		b, l := boxesList(grid, nextPoint(p, m), m)
		return b, append([]Point{p}, l...)
	case '#':
		return false, nil
	case '.':
		return true, []Point{p}
	case 'O':
		b, l := boxesList(grid, nextPoint(p, m), m)
		return b, append([]Point{p}, l...)
	case '[':
		if m == '^' || m == 'v' {
			b1, l1 := boxesList(grid, nextPoint(p, m), m)
			b2, l2 := boxesList(grid, nextPoint(Point{p.x, p.y + 1}, m), m)
			list := append([]Point{p, {p.x, p.y + 1}}, l1...)
			list = append(list, l2...)
			return b1 && b2, list
		}
		b, l := boxesList(grid, nextPoint(p, m), m)
		return b, append([]Point{p}, l...)
	case ']':
		if m == '^' || m == 'v' {
			b1, l1 := boxesList(grid, nextPoint(p, m), m)
			b2, l2 := boxesList(grid, nextPoint(Point{p.x, p.y - 1}, m), m)
			list := append([]Point{p, {p.x, p.y - 1}}, l1...)
			list = append(list, l2...)
			return b1 && b2, list
		}
		b, l := boxesList(grid, nextPoint(p, m), m)
		return b, append([]Point{p}, l...)
	}
	return true, nil
}

func solve(grid [][]byte, moves string) int {
	var curr Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				curr = Point{i, j}
			}
		}
	}
	for _, m := range moves {
		b, l := boxesList(grid, curr, m)
		if b {
			var mapping = make(map[Point]byte)
			for _, p := range l {
				prev := nextPoint(p, backward(m))
				if !slices.Contains(l, prev) {
					mapping[p] = '.'
				} else {
					mapping[p] = grid[prev.x][prev.y]
				}
			}
			for key, val := range mapping {
				grid[key.x][key.y] = val
			}
			curr = nextPoint(curr, m)
		}
	}
	return sumGPSCoordinates(grid)
}

func part1(s string) int {
	var grid, moves = parse(s)
	return solve(grid, moves)
}

func part2(s string) int {
	var grid, moves = parse(s)
	return solve(widerGrid(grid), moves)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day15/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day15/input.txt")

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
