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

type grid struct {
	empty point
	goal  point
	steps int
}

type queue []grid

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Push(g grid) {
	*q = append(*q, g)
}

func (q *queue) Pop() (grid, bool) {
	if q.IsEmpty() {
		return grid{}, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

type node struct {
	x, y, size, used, avail, use int
}

func Split(r rune) bool {
	return r == ' ' || r == '-' || r == 'x' || r == 'y' || r == 'T' || r == '%'
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func done(g grid) bool {
	return g.goal.x == 0 && g.goal.y == 0
}

func isValid(p point, maxX int, maxY int, full []point) bool {
	return p.x >= 0 && p.x <= maxX && p.y >= 0 && p.y <= maxY && !isIn(p, full)
}

func isIn(p point, l []point) bool {
	for _, q := range l {
		if p.x == q.x && p.y == q.y {
			return true
		}
	}
	return false
}

func next(g grid, maxX int, maxY int, visited map[string]bool, full []point) []grid {
	res := []grid{}
	nextPoint := point{g.empty.x + 1, g.empty.y}
	if isValid(nextPoint, maxX, maxY, full) {
		nextGrid := grid{nextPoint, g.goal, g.steps + 1}
		ifSwap := grid{nextPoint, g.empty, g.steps + 1}
		_, ok := visited[gridToString(nextGrid)]
		_, okSwap := visited[gridToString(ifSwap)]
		if nextPoint.x == g.goal.x && nextPoint.y == g.goal.y && !okSwap {
			res = append(res, ifSwap)
			visited[gridToString(ifSwap)] = true
		}
		if !ok && (nextPoint.x != g.goal.x || nextPoint.y != g.goal.y) {
			res = append(res, nextGrid)
			visited[gridToString(nextGrid)] = true

		}
	}
	nextPoint = point{g.empty.x - 1, g.empty.y}
	if isValid(nextPoint, maxX, maxY, full) {
		nextGrid := grid{nextPoint, g.goal, g.steps + 1}
		ifSwap := grid{nextPoint, g.empty, g.steps + 1}
		_, ok := visited[gridToString(nextGrid)]
		_, okSwap := visited[gridToString(ifSwap)]
		if nextPoint.x == g.goal.x && nextPoint.y == g.goal.y && !okSwap {
			res = append(res, ifSwap)
			visited[gridToString(ifSwap)] = true
		}
		if !ok && (nextPoint.x != g.goal.x || nextPoint.y != g.goal.y) {
			res = append(res, nextGrid)
			visited[gridToString(nextGrid)] = true
		}
	}
	nextPoint = point{g.empty.x, g.empty.y + 1}
	if isValid(nextPoint, maxX, maxY, full) {
		nextGrid := grid{nextPoint, g.goal, g.steps + 1}
		ifSwap := grid{nextPoint, g.empty, g.steps + 1}
		_, ok := visited[gridToString(nextGrid)]
		_, okSwap := visited[gridToString(ifSwap)]
		if nextPoint.x == g.goal.x && nextPoint.y == g.goal.y && !okSwap {
			res = append(res, ifSwap)
			visited[gridToString(ifSwap)] = true
		}
		if !ok && (nextPoint.x != g.goal.x || nextPoint.y != g.goal.y) {
			res = append(res, nextGrid)
			visited[gridToString(nextGrid)] = true
		}
	}
	nextPoint = point{g.empty.x, g.empty.y - 1}
	if isValid(nextPoint, maxX, maxY, full) {
		nextGrid := grid{nextPoint, g.goal, g.steps + 1}
		ifSwap := grid{nextPoint, g.empty, g.steps + 1}
		_, ok := visited[gridToString(nextGrid)]
		_, okSwap := visited[gridToString(ifSwap)]
		if nextPoint.x == g.goal.x && nextPoint.y == g.goal.y && !okSwap {
			res = append(res, ifSwap)
			visited[gridToString(ifSwap)] = true
		}
		if !ok && (nextPoint.x != g.goal.x || nextPoint.y != g.goal.y) {
			res = append(res, nextGrid)
			visited[gridToString(nextGrid)] = true
		}

	}
	return res
}

func gridToString(g grid) string {
	return strconv.Itoa(g.empty.x) + ";" + strconv.Itoa(g.empty.y) + ";" + strconv.Itoa(g.goal.x) + ";" + strconv.Itoa(g.goal.x)
}

func part1(s string) int {
	c := 0
	nodeList := []node{}
	tab := format(s)
	for _, line := range tab {
		if line[0] == "/dev/grid/node" {
			x, _ := strconv.Atoi(line[1])
			y, _ := strconv.Atoi(line[2])
			size, _ := strconv.Atoi(line[3])
			used, _ := strconv.Atoi(line[4])
			avail, _ := strconv.Atoi(line[5])
			use, _ := strconv.Atoi(line[6])
			n := node{x, y, size, used, avail, use}
			nodeList = append(nodeList, n)
		}
	}
	for i := 0; i < len(nodeList); i++ {
		for j := 0; j < len(nodeList); j++ {
			node1 := nodeList[i]
			node2 := nodeList[j]
			if node1.used != 0 && (node1.x != node2.x || node1.y != node2.y) && node1.used <= node2.avail {
				c++
			}
		}
	}
	return c
}

func part2(s string) int {
	tab := format(s)
	maxX := 0
	maxY := 0
	startX := 0
	startY := 0
	full := []point{}
	for _, line := range tab {
		if line[0] == "/dev/grid/node" {
			x, _ := strconv.Atoi(line[1])
			y, _ := strconv.Atoi(line[2])
			size, _ := strconv.Atoi(line[3])
			used, _ := strconv.Atoi(line[4])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			if used == 0 {
				startX = x
				startY = y
			}
			if size > 200 {
				full = append(full, point{x, y})
			}
		}
	}
	visited := make(map[string]bool)
	curr := grid{point{startX, startY}, point{maxX, 0}, 0}
	var queue queue
	queue.Push(curr)
	for !queue.IsEmpty() {
		curr, _ = queue.Pop()
		//fmt.Println(curr)
		visited[gridToString(curr)] = true
		if done(curr) {
			return curr.steps
		} else {
			toAdd := next(curr, maxX, maxY, visited, full)
			for _, g := range toAdd {
				queue.Push(g)
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
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
