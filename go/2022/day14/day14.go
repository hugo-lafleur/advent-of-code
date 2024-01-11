package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.FieldsFunc(x, Split))
	}
	return tab
}

func next(i int, j int, grid map[point]bool) (int, int) {
	if !grid[point{i + 1, j}] {
		return i + 1, j
	} else {
		if !grid[point{i + 1, j - 1}] {
			return i + 1, j - 1
		} else {
			if !grid[point{i + 1, j + 1}] {
				return i + 1, j + 1
			} else {
				return i, j
			}
		}
	}
}

func part1(s string) int {
	tab := format(s)
	grid := make(map[point]bool)
	j := 0
	max := 0
	for j < len(tab) {
		line := tab[j]
		i := 0
		for i < len(line)-4 {
			a, _ := strconv.Atoi(line[i])
			b, _ := strconv.Atoi(line[i+1])
			m, _ := strconv.Atoi(line[i+3])
			n, _ := strconv.Atoi(line[i+4])
			if b > max {
				max = b
			}
			if n > max {
				max = n
			}
			if a == m {
				if b < n {
					k := b
					for k <= n {
						grid[point{x: k, y: a}] = true
						k++
					}
				}
				if n < b {
					k := n
					for k <= b {
						grid[point{x: k, y: a}] = true
						k++
					}
				}
			}
			if b == n {
				if a < m {
					k := a
					for k <= m {
						grid[point{x: b, y: k}] = true
						k++
					}
				}
				if m < a {
					k := m
					for k <= a {
						grid[point{x: b, y: k}] = true
						k++
					}
				}
			}
			i = i + 3
		}
		j++
	}
	max = max + 2
	i := 0
	for {
		a, b := 0, 500
		n, m := next(a, b, grid)
		for a != n || b != m {
			if a == max {
				return i
			}
			a, b = n, m
			n, m = next(n, m, grid)
		}
		grid[point{n, m}] = true
		i++
	}
}

func part2(s string) int {
	tab := format(s)
	grid := make(map[point]bool)
	j := 0
	max := 0
	for j < len(tab) {
		line := tab[j]
		i := 0
		for i < len(line)-4 {
			a, _ := strconv.Atoi(line[i])
			b, _ := strconv.Atoi(line[i+1])
			m, _ := strconv.Atoi(line[i+3])
			n, _ := strconv.Atoi(line[i+4])
			if b > max {
				max = b
			}
			if n > max {
				max = n
			}
			if a == m {
				if b < n {
					k := b
					for k <= n {
						grid[point{x: k, y: a}] = true
						k++
					}
				}
				if n < b {
					k := n
					for k <= b {
						grid[point{x: k, y: a}] = true
						k++
					}
				}
			}
			if b == n {
				if a < m {
					k := a
					for k <= m {
						grid[point{x: b, y: k}] = true
						k++
					}
				}
				if m < a {
					k := m
					for k <= a {
						grid[point{x: b, y: k}] = true
						k++
					}
				}
			}
			i = i + 3
		}
		j++
	}
	max = max + 2
	i := 0
	for i < 1000 {
		grid[point{max, i}] = true
		i++
	}
	i = 0
	for {
		a, b := 0, 500
		n, m := next(a, b, grid)
		if n == 0 && m == 500 {
			return i + 1
		}
		for a != n || b != m {
			a, b = n, m
			n, m = next(n, m, grid)
		}
		grid[point{n, m}] = true
		i++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day14/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day14/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
