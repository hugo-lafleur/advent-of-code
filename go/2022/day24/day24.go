package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

type state struct {
	pos   point
	steps int
	goals int
}

type couple struct {
	tab    [][]string
	minute int
}

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func next(tab [][]string) [][]string {
	res := make([][]string, len(tab))
	for i, line := range tab {
		res[i] = make([]string, len(line))
		//fmt.Println(res)
		for j, c := range line {
			if c != "#" {
				res[i][j] = "."
			} else {
				res[i][j] = "#"
			}
		}
	}
	for i, line := range tab {
		for j, c := range line {
			if strings.Contains(c, ">") {
				if j == len(tab[i])-2 {
					res[i][1] += ">"
					res[i][1] = strings.Trim(res[i][1], ".")
				} else {
					res[i][j+1] += ">"
					res[i][j+1] = strings.TrimPrefix(res[i][j+1], ".")
				}
			}
			if strings.Contains(c, "<") {
				if j == 1 {
					res[i][len(tab[i])-2] += "<"
					res[i][len(tab[i])-2] = strings.Trim(res[i][len(tab[i])-2], ".")
				} else {
					res[i][j-1] += "<"
					res[i][j-1] = strings.Trim(res[i][j-1], ".")
				}
			}
			if strings.Contains(c, "v") {
				if i == len(tab)-2 {
					res[1][j] += "v"
					res[1][j] = strings.Trim(res[1][j], ".")
				} else {
					res[i+1][j] += "v"
					res[i+1][j] = strings.Trim(res[i+1][j], ".")
				}
			}
			if strings.Contains(c, "^") {
				if i == 1 {
					res[len(tab)-2][j] += "^"
					res[len(tab)-2][j] = strings.Trim(res[len(tab)-2][j], ".")
				} else {
					res[i-1][j] += "^"
					res[i-1][j] = strings.Trim(res[i-1][j], ".")
				}
			}
		}
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	start := state{pos: point{0, 1}, steps: 0}
	var currentBlizzard couple
	currentBlizzard.tab = next(tab)
	currentBlizzard.minute = 0
	var q deque.Deque[state]
	visited := make(map[state]bool)
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopFront()
		if curr.pos.x == len(tab)-1 && curr.pos.y == len(tab[0])-2 {
			return curr.steps
		}
		var currTab [][]string
		if curr.steps == currentBlizzard.minute {
			currTab = currentBlizzard.tab
		} else {
			currentBlizzard.minute++
			currentBlizzard.tab = next(currentBlizzard.tab)
			currTab = currentBlizzard.tab
		}
		x, y := curr.pos.x, curr.pos.y
		for _, p := range []point{{x + 1, y}, {x - 1, y}, {x, y}, {x, y + 1}, {x, y - 1}} {
			if p.x >= 0 && currTab[p.x][p.y] == "." {
				newState := state{pos: p, steps: curr.steps + 1}
				_, seen := visited[newState]
				if !seen {
					q.PushBack(newState)
					visited[newState] = true
				}
			}
		}
	}
	return 0
}

func part2(s string) int {
	tab := format(s)
	start := state{pos: point{0, 1}, steps: 0}
	var currentBlizzard couple
	currentBlizzard.tab = next(tab)
	currentBlizzard.minute = 0
	var q deque.Deque[state]
	visited := make(map[state]bool)
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopFront()
		if curr.pos.x == len(tab)-1 && curr.pos.y == len(tab[0])-2 {
			if curr.goals == 0 {
				curr.goals++
			}
			if curr.goals == 2 {
				return curr.steps
			}
		}
		if curr.pos.x == 0 && curr.pos.y == 1 && curr.goals == 1 {
			curr.goals++
		}
		var currTab [][]string
		if curr.steps == currentBlizzard.minute {
			currTab = currentBlizzard.tab
		} else {
			currentBlizzard.minute++
			currentBlizzard.tab = next(currentBlizzard.tab)
			currTab = currentBlizzard.tab
		}
		x, y := curr.pos.x, curr.pos.y
		for _, p := range []point{{x + 1, y}, {x - 1, y}, {x, y}, {x, y + 1}, {x, y - 1}} {
			if p.x >= 0 && p.x < len(tab) && currTab[p.x][p.y] == "." {
				newState := state{pos: p, steps: curr.steps + 1, goals: curr.goals}
				_, seen := visited[newState]
				if !seen {
					q.PushBack(newState)
					visited[newState] = true
				}
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day24/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day24/input.txt")

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
