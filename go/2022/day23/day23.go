package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func format(s string) map[point]bool {
	res := make(map[point]bool)
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				res[point{i, j}] = true
			}
		}
	}
	return res
}

func propose(p point, dir string, mapping map[point]bool) (point, bool) {
	x, y := p.x, p.y
	switch dir {
	case "north":
		for _, pnt := range []point{{x - 1, y}, {x - 1, y - 1}, {x - 1, y + 1}} {
			_, ok := mapping[pnt]
			if ok {
				return point{}, false
			}
		}
		return point{x - 1, y}, true
	case "south":
		for _, pnt := range []point{{x + 1, y}, {x + 1, y - 1}, {x + 1, y + 1}} {
			_, ok := mapping[pnt]
			if ok {
				return point{}, false
			}
		}
		return point{x + 1, y}, true
	case "west":
		for _, pnt := range []point{{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1}} {
			_, ok := mapping[pnt]
			if ok {
				return point{}, false
			}
		}
		return point{x, y - 1}, true
	case "east":
		for _, pnt := range []point{{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}} {
			_, ok := mapping[pnt]
			if ok {
				return point{}, false
			}
		}
		return point{x, y + 1}, true
	}
	return point{}, false
}

func check(p point, mapping map[point]bool) bool {
	x, y := p.x, p.y
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			pnt := point{x + i, y + j}
			if p != pnt {
				_, ok := mapping[pnt]
				if ok {
					return true
				}
			}
		}
	}
	return false
}

func emptyTiles(mapping map[point]bool) int {
	res := 0
	xList := []int{}
	yList := []int{}
	for key := range mapping {
		xList = append(xList, key.x)
		yList = append(yList, key.y)
	}
	for i := slices.Min(xList); i <= slices.Max(xList); i++ {
		for j := slices.Min(yList); j <= slices.Max(yList); j++ {
			_, ok := mapping[point{i, j}]
			if !ok {
				res++
			}
		}
	}
	return res
}

func part1(s string) int {
	mapping := format(s)
	directions := []string{"north", "south", "west", "east"}
	rounds := 0
	for rounds < 10 {
		moves := make(map[point]point)
		positions := make(map[point]int)
	loop:
		for key := range mapping {
			if !check(key, mapping) {
				moves[key] = key
				positions[key]++
				continue
			}
			for i := 0; i < 4; i++ {
				dir := directions[(rounds+i)%4]
				newPos, ok := propose(key, dir, mapping)
				if ok {
					moves[key] = newPos
					positions[newPos]++
					continue loop
				}
			}
			moves[key] = key
			positions[key]++
		}
		newMapping := make(map[point]bool)
		for key, value := range moves {
			if positions[value] == 1 {
				newMapping[value] = true
			} else {
				newMapping[key] = true
			}
		}
		mapping = newMapping
		rounds++
	}
	return emptyTiles(mapping)
}

func part2(s string) int {
	mapping := format(s)
	directions := []string{"north", "south", "west", "east"}
	rounds := 0
	for {
		numberMoves := 0
		moves := make(map[point]point)
		positions := make(map[point]int)
	loop:
		for key := range mapping {
			if !check(key, mapping) {
				moves[key] = key
				positions[key]++
				continue
			}
			for i := 0; i < 4; i++ {
				dir := directions[(rounds+i)%4]
				newPos, ok := propose(key, dir, mapping)
				if ok {
					numberMoves++
					moves[key] = newPos
					positions[newPos]++
					continue loop
				}
			}
			moves[key] = key
			positions[key]++
		}
		if numberMoves == 0 {
			return rounds + 1
		}
		newMapping := make(map[point]bool)
		for key, value := range moves {
			if positions[value] == 1 {
				newMapping[value] = true
			} else {
				newMapping[key] = true
			}
		}
		mapping = newMapping
		rounds++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day23/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day23/input.txt")

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
