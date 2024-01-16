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

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func nextLeft(p point, d string) (point, string) {
	x, y := p.x, p.y
	switch d {
	case "u":
		return point{x, y - 1}, "l"
	case "l":
		return point{x + 1, y}, "d"
	case "d":
		return point{x, y + 1}, "r"
	case "r":
		return point{x - 1, y}, "u"
	}
	return point{}, ""
}

func nextRight(p point, d string) (point, string) {
	x, y := p.x, p.y
	switch d {
	case "u":
		return point{x, y + 1}, "r"
	case "l":
		return point{x - 1, y}, "u"
	case "d":
		return point{x, y - 1}, "l"
	case "r":
		return point{x + 1, y}, "d"
	}
	return point{}, ""
}

func nextForward(p point, d string) (point, string) {
	x, y := p.x, p.y
	switch d {
	case "u":
		return point{x - 1, y}, d
	case "l":
		return point{x, y - 1}, d
	case "d":
		return point{x + 1, y}, d
	case "r":
		return point{x, y + 1}, d
	}
	return point{}, ""
}

func nextBackward(p point, d string) (point, string) {
	x, y := p.x, p.y
	switch d {
	case "u":
		return point{x + 1, y}, "d"
	case "l":
		return point{x, y + 1}, "r"
	case "d":
		return point{x - 1, y}, "u"
	case "r":
		return point{x, y - 1}, "l"
	}
	return point{}, ""
}

func part1(s string) int {
	c := 0
	input := format(s)
	grid := make(map[point]int)
	l := len(input)
	limit := 10000
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			x := i - (l / 2)
			y := j - (l / 2)
			p := point{x, y}
			switch input[i][j] {
			case ".":
				grid[p] = 0
			case "#":
				grid[p] = 1
			}
		}
	}
	curr := point{0, 0}
	direction := "u"
	burst := 0
	for burst < limit {
		mode, ok := grid[curr]
		if !ok || (ok && mode == 0) {
			grid[curr] = 1
			curr, direction = nextLeft(curr, direction)
			c++
		} else {
			grid[curr] = 0
			curr, direction = nextRight(curr, direction)
		}
		burst++
	}
	return c
}

func part2(s string) int {
	c := 0
	input := format(s)
	grid := make(map[point]int)
	l := len(input)
	limit := 10000000
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			x := i - (l / 2)
			y := j - (l / 2)
			p := point{x, y}
			switch input[i][j] {
			case ".":
				grid[p] = 0
			case "#":
				grid[p] = 2
			}
		}
	}
	curr := point{0, 0}
	direction := "u"
	burst := 0
	for burst < limit {
		mode, ok := grid[curr]
		if !ok || (ok && mode == 0) {
			grid[curr] = 1
			curr, direction = nextLeft(curr, direction)
		}
		if mode == 1 {
			grid[curr] = 2
			curr, direction = nextForward(curr, direction)
			c++
		}
		if mode == 2 {
			grid[curr] = 3
			curr, direction = nextRight(curr, direction)
		}
		if mode == 3 {
			grid[curr] = 0
			curr, direction = nextBackward(curr, direction)
		}
		burst++
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day22/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day22/input.data")

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
