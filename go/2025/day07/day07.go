package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	result := [][]string{}
	for _, line := range lines {
		result = append(result, strings.Split(line, ""))
	}
	return result
}

type point struct {
	x, y int
}

func part1(s string) int {
	tab := format(s)
	var dq deque.Deque[point]
	for y, char := range tab[0] {
		if char == "S" {
			dq.PushBack(point{0, y})
			break
		}
	}
	result := 0
	visited := make(map[point]bool)
	for dq.Len() != 0 {
		curr := dq.PopFront()
		next := point{curr.x + 1, curr.y}
		if visited[next] || next.x >= len(tab) {
			continue
		}
		visited[next] = true
		switch tab[next.x][next.y] {
		case ".":
			dq.PushBack(next)
		case "^":
			result++
			dq.PushBack(point{next.x, next.y - 1})
			dq.PushBack(point{next.x, next.y + 1})
		}
	}
	return result
}

func part2(s string) int {
	tab := format(s)
	var dq deque.Deque[point]
	visited := make(map[point]int)
	for y, char := range tab[0] {
		if char == "S" {
			dq.PushBack(point{0, y})
			visited[point{0, y}] = 1
			break
		}
	}
	result := 0
	for dq.Len() != 0 {
		curr := dq.PopFront()
		next := point{curr.x + 1, curr.y}
		if next.x == len(tab) {
			result += visited[curr]
			continue
		}
		switch tab[next.x][next.y] {
		case ".":
			if visited[next] == 0 {
				dq.PushBack(next)
			}
			visited[next] += visited[curr]

		case "^":
			p1, p2 := point{next.x, next.y - 1}, point{next.x, next.y + 1}
			if visited[p1] == 0 {
				dq.PushBack(p1)
			}
			if visited[p2] == 0 {
				dq.PushBack(p2)
			}
			visited[p1] += visited[curr]
			visited[p2] += visited[curr]
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day07/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day07/input.txt")

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
