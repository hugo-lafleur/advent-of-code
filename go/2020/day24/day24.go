package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type hex struct {
	q, r int
}

func move(h hex, dir string) hex {
	new := h
	switch dir {
	case "e":
		new.q++
	case "se":
		new.r++
	case "sw":
		new.q--
		new.r++
	case "w":
		new.q--
	case "nw":
		new.r--
	case "ne":
		new.q++
		new.r--
	}
	return new
}

func neighbors(h hex) []hex {
	res := []hex{}
	for _, s := range []string{"e", "w", "sw", "nw", "se", "ne"} {
		res = append(res, move(h, s))
	}
	return res
}

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func countBlack(grid map[hex]int) int {
	c := 0
	for _, value := range grid {
		if value%2 == 1 {
			c++
		}
	}
	return c
}

func part1(s string) int {
	grid := make(map[hex]int)
	list := format(s)
	for _, path := range list {
		curr := hex{0, 0}
		for i := 0; i < len(path); i++ {
			if path[i] == "e" || path[i] == "w" {
				curr = move(curr, path[i])
			} else {
				curr = move(curr, path[i]+path[i+1])
				i++
			}
		}
		grid[curr]++
	}
	return countBlack(grid)
}

func part2(s string) int {
	grid := make(map[hex]int)
	list := format(s)
	for _, path := range list {
		curr := hex{0, 0}
		for i := 0; i < len(path); i++ {
			if path[i] == "e" || path[i] == "w" {
				curr = move(curr, path[i])
			} else {
				curr = move(curr, path[i]+path[i+1])
				i++
			}
		}
		grid[curr]++
	}
	day := 0
	for day < 100 {
		newGrid := make(map[hex]int)
		for h := range grid {
			for _, neigh := range neighbors(h) {
				_, done := newGrid[neigh]
				if !done {
					neighborsNeigh := neighbors(neigh)
					numberOfBlacks := 0
					for _, neighOfneigh := range neighborsNeigh {
						n, ok := grid[neighOfneigh]
						if ok && (n%2) == 1 {
							numberOfBlacks++
						}
					}
					n, ok := grid[neigh]
					if !ok || (n%2) == 0 {
						if numberOfBlacks == 2 {
							newGrid[neigh] = 1
						} else {
							newGrid[neigh] = 0
						}
					}
					if ok && (n%2) == 1 {
						if numberOfBlacks == 0 || numberOfBlacks > 2 {
							newGrid[neigh] = 0
						} else {
							newGrid[neigh] = 1
						}
					}
				}
			}
		}
		grid = newGrid
		day++
	}
	return countBlack(grid)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day24/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day24/input.data")

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
