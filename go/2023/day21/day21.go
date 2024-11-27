package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func correct(p point, n, m int) bool {
	return p.x > -1 && p.x < n && p.y > -1 && p.y < m
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func modulo(i int, n int) int {
	if i > 0 {
		return i % n
	}
	return modulo(i+n, n)
}

func Lagrange(x int, points []point) int {
	res := 0
	for _, p := range points {
		temp := 1
		for _, q := range points {
			if p != q {
				temp *= (x - q.x) / (p.x - q.x)
			}
		}
		res += p.y * temp
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	n := len(tab)
	m := len(tab[0])
	possible := make(map[point]bool)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == "S" {
				possible[point{i, j}] = true
			}
		}
	}
	var d int
	if len(tab) == 131 {
		d = 64
	} else {
		d = 6
	}
	for k := 0; k < d; k++ {
		next := make(map[point]bool)
		for p := range possible {
			x := p.x
			y := p.y
			if correct(point{x + 1, y}, n, m) && (tab[x+1][y] == "." || tab[x+1][y] == "S") {
				next[point{x + 1, y}] = true
			}
			if correct(point{x - 1, y}, n, m) && (tab[x-1][y] == "." || tab[x-1][y] == "S") {
				next[point{x - 1, y}] = true
			}
			if correct(point{x, y + 1}, n, m) && (tab[x][y+1] == "." || tab[x][y+1] == "S") {
				next[point{x, y + 1}] = true
			}
			if correct(point{x, y - 1}, n, m) && (tab[x][y-1] == "." || tab[x][y-1] == "S") {
				next[point{x, y - 1}] = true
			}
		}
		possible = next
	}
	return len(possible)
}

func part2(s string) int {
	tab := format(s)
	var d int
	if len(tab) == 131 {
		d = 26501365
	} else {
		d = 5000
	}
	n := len(tab)
	m := len(tab[0])
	var c point
	possible := make(map[point]bool)
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == "S" {
				possible[point{i, j}] = true
				tab[i][j] = "."
				c.x = i
				c.y = j
			}
		}
	}
	function := make(map[int]int)
	for k := 0; k < d; k++ {
		next := make(map[point]bool)
		if k == (d%n) && n == 131 {
			function[k] = len(possible)
		}
		if k == n+(d%n) {
			function[k] = len(possible)
		}
		if k == 2*n+(d%n) {
			function[k] = len(possible)
		}
		if k == 3*n+(d%n) && n != 131 {
			function[k] = len(possible)
		}
		if len(function) == 3 {
			points := []point{}
			for key, value := range function {
				points = append(points, point{key, value})
			}
			return Lagrange(d, points)
		}
		for p := range possible {
			x := p.x
			y := p.y
			if tab[modulo(x+1, n)][modulo(y, m)] == "." {
				next[point{x + 1, y}] = true
			}
			if tab[modulo(x-1, n)][modulo(y, m)] == "." {
				next[point{x - 1, y}] = true
			}
			if tab[modulo(x, n)][modulo(y+1, m)] == "." {
				next[point{x, y + 1}] = true
			}
			if tab[modulo(x, n)][modulo(y-1, m)] == "." {
				next[point{x, y - 1}] = true
			}
		}
		possible = next
	}
	return len(possible)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day21/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day21/input.txt")

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
