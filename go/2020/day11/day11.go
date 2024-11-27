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
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func addPoint(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func mulPoint(a int, p point) point {
	return point{a * p.x, a * p.y}
}

func tabToString(tab [][]string) string {
	res := ""
	for _, line := range tab {
		for _, s := range line {
			res += s
		}
	}
	return res
}

func isValid(p point, tab [][]string) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(tab) && p.y < len(tab[p.x])
}

func part1(s string) int {
	tab := format(s)
	mem := make(map[string]int)
	round := 0
	for {
		newTab := make([][]string, len(tab))
		for i := 0; i < len(tab); i++ {
			newTab[i] = make([]string, len(tab[i]))
			for j := 0; j < len(tab[i]); j++ {
				p := point{i, j}
				occupied := 0
				for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
					otherP := addPoint(p, offset)
					if isValid(otherP, tab) && tab[otherP.x][otherP.y] == "#" {
						occupied++
					}
				}
				if tab[p.x][p.y] == "L" && occupied == 0 {
					newTab[p.x][p.y] = "#"
				} else if tab[p.x][p.y] == "#" && occupied >= 4 {
					newTab[p.x][p.y] = "L"
				} else {
					newTab[p.x][p.y] = tab[p.x][p.y]
				}
			}
		}
		tab = newTab
		round++
		_, ok := mem[tabToString(tab)]
		if ok {
			c := 0
			for _, line := range tab {
				for _, s := range line {
					if s == "#" {
						c++
					}
				}
			}
			return c
		}
		mem[tabToString(tab)] = round
	}
}

func part2(s string) int {
	tab := format(s)
	mem := make(map[string]int)
	round := 0
	for {
		newTab := make([][]string, len(tab))
		for i := 0; i < len(tab); i++ {
			newTab[i] = make([]string, len(tab[i]))
			for j := 0; j < len(tab[i]); j++ {
				p := point{i, j}
				occupied := 0
				for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
					for n := 1; ; n++ {
						otherP := addPoint(p, mulPoint(n, offset))
						if !isValid(otherP, tab) {
							break
						}
						if tab[otherP.x][otherP.y] == "#" {
							occupied++
							break
						}
						if tab[otherP.x][otherP.y] == "L" {
							break
						}
					}
				}
				if tab[p.x][p.y] == "L" && occupied == 0 {
					newTab[p.x][p.y] = "#"
				} else if tab[p.x][p.y] == "#" && occupied >= 5 {
					newTab[p.x][p.y] = "L"
				} else {
					newTab[p.x][p.y] = tab[p.x][p.y]
				}
			}
		}
		tab = newTab
		round++
		_, ok := mem[tabToString(tab)]
		if ok {
			c := 0
			for _, line := range tab {
				for _, s := range line {
					if s == "#" {
						c++
					}
				}
			}
			return c
		}
		mem[tabToString(tab)] = round
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day11/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day11/input.txt")

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
