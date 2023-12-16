package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type point struct {
	i, j      int
	direction string
}

type coord struct {
	i, j int
}

type grid [][]string

var energized map[coord]int
var done map[point]bool

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func isDone(p point) bool {
	_, ok := done[p]
	return ok
}

func (tab grid) nextPoint(p point) {
	x := p.i
	y := p.j
	d := p.direction
	n := len(tab)
	m := len(tab[0])
	if x == -1 || x == n || y == -1 || y == m || isDone(p) {
		return
	}
	s := tab[x][y]
	done[p] = true
	energized[coord{x, y}]++
	switch s {
	case ".":
		switch d {
		case ">":
			tab.nextPoint(point{x, y + 1, d})
		case "<":
			tab.nextPoint(point{x, y - 1, d})
		case "^":
			tab.nextPoint(point{x - 1, y, d})
		case "v":
			tab.nextPoint(point{x + 1, y, d})
		}
	case "/":
		switch d {
		case ">":
			tab.nextPoint(point{x - 1, y, "^"})
		case "<":
			tab.nextPoint(point{x + 1, y, "v"})
		case "v":
			tab.nextPoint(point{x, y - 1, "<"})
		case "^":
			tab.nextPoint(point{x, y + 1, ">"})
		}
	case "\\":
		switch d {
		case "<":
			tab.nextPoint(point{x - 1, y, "^"})
		case ">":
			tab.nextPoint(point{x + 1, y, "v"})
		case "^":
			tab.nextPoint(point{x, y - 1, "<"})
		case "v":
			tab.nextPoint(point{x, y + 1, ">"})
		}
	case "|":
		switch d {
		case ">", "<":
			tab.nextPoint(point{x + 1, y, "v"})
			tab.nextPoint(point{x - 1, y, "^"})
		case "v":
			tab.nextPoint(point{x + 1, y, "v"})
		case "^":
			tab.nextPoint(point{x - 1, y, "^"})
		}
	case "-":
		switch d {
		case "v", "^":
			tab.nextPoint(point{x, y + 1, ">"})
			tab.nextPoint(point{x, y - 1, "<"})
		case "<":
			tab.nextPoint(point{x, y - 1, "<"})
		case ">":
			tab.nextPoint(point{x, y + 1, ">"})
		}
	}
}

func maxTilesConfiguration(dict map[point]int) int {
	max := 0
	var maxPoint point
	for key, value := range dict {
		if value > max {
			max = value
			maxPoint = key
		}
	}
	fmt.Println(maxPoint)
	return max
}

func calculateEnergized(p point, tab grid) int {
	energized = make(map[coord]int)
	done = make(map[point]bool)
	tab.nextPoint(p)
	return len(energized)
}

func part1(s string) int {
	tab := grid(format(s))
	return calculateEnergized(point{0, 0, ">"}, tab)
}

func part2(s string) int {
	tab := grid(format(s))
	tilesConfiguration := make(map[point]int)
	n := len(tab)
	m := len(tab[0])
	for i := 0; i < n; i++ {
		a := point{0, i, "v"}
		b := point{n - 1, i, "^"}
		tilesConfiguration[a] = calculateEnergized(a, tab)
		tilesConfiguration[b] = calculateEnergized(b, tab)

	}
	for i := 0; i < m; i++ {
		a := point{i, 0, ">"}
		b := point{i, m - 1, "<"}
		tilesConfiguration[a] = calculateEnergized(a, tab)
		tilesConfiguration[b] = calculateEnergized(b, tab)
	}
	return maxTilesConfiguration(tilesConfiguration)
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
