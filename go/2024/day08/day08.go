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

func add(a, b Point) Point {
	return Point{a.x + b.x, a.y + b.y}
}

func mul(m int, a Point) Point {
	return Point{m * a.x, m * a.y}
}

func diff(a, b Point) Point {
	return Point{a.x - b.x, a.y - b.y}
}
func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var grid = make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func part1(s string) int {
	var grid = parse(s)
	var m, n = len(grid), len(grid[0])
	var antennas = make(map[byte][]Point)
	var antinodes = make(map[Point]bool)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Point{i, j})
			}
		}
	}
	for freq := range antennas {
		var list = antennas[freq]
		for i := range list {
			for j := range list {
				if i != j {
					anti := add(list[i], mul(2, diff(list[j], list[i])))
					if anti.x >= 0 && anti.y >= 0 && anti.x < m && anti.y < n {
						antinodes[anti] = true
					}
				}
			}
		}
	}
	return len(antinodes)
}

func part2(s string) int {
	var grid = parse(s)
	var m, n = len(grid), len(grid[0])
	var antennas = make(map[byte][]Point)
	var antinodes = make(map[Point]bool)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				antennas[grid[i][j]] = append(antennas[grid[i][j]], Point{i, j})
			}
		}
	}
	for freq := range antennas {
		var list = antennas[freq]
		for i := range list {
			for j := range list {
				if i != j {
					for k := 1; ; k++ {
						anti := add(list[i], mul(k, diff(list[j], list[i])))
						if anti.x >= 0 && anti.y >= 0 && anti.x < m && anti.y < n {
							antinodes[anti] = true
						} else {
							break
						}
					}
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day08/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day08/input.txt")

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
