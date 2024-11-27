package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type state struct {
	cave              string
	smallCavesVisited map[string]int
	history           string
	visitedTwice      bool
}

func format(s string) map[string][]string {
	lines := strings.Split(s, "\n")
	res := make(map[string][]string)
	for _, line := range lines {
		lineSplit := strings.Split(line, "-")
		A, B := lineSplit[0], lineSplit[1]
		res[A] = append(res[A], B)
		res[B] = append(res[B], A)
	}
	return res
}

func part1(s string) int {
	c := 0
	mapping := format(s)
	var q deque.Deque[state]
	start := state{"start", make(map[string]int), "start", false}
	start.smallCavesVisited["start"] = 1
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopFront()
		if curr.cave == "end" {
			c++
			continue
		}
		for _, close := range mapping[curr.cave] {
			_, visited := curr.smallCavesVisited[close]
			if close[0] < 'a' {
				q.PushBack(state{close, curr.smallCavesVisited, curr.history + ";" + close, false})
			}
			if close[0] >= 'a' && !visited {
				newVistied := make(map[string]int)
				for key, value := range curr.smallCavesVisited {
					newVistied[key] = value
				}
				newVistied[close] = 1
				q.PushBack(state{close, newVistied, curr.history + ";" + close, false})
			}
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	mapping := format(s)
	var q deque.Deque[state]
	start := state{"start", make(map[string]int), "start", false}
	start.smallCavesVisited["start"] = 2
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopFront()
		if curr.cave == "end" {
			c++
			continue
		}
		for _, close := range mapping[curr.cave] {
			visits, visited := curr.smallCavesVisited[close]
			if close[0] < 'a' {
				q.PushBack(state{close, curr.smallCavesVisited, curr.history + ";" + close, curr.visitedTwice})
			}
			if close[0] >= 'a' && !visited {
				newVistied := make(map[string]int)
				for key, value := range curr.smallCavesVisited {
					newVistied[key] = value
				}
				newVistied[close]++
				q.PushBack(state{close, newVistied, curr.history + ";" + close, curr.visitedTwice})
			}
			if close[0] >= 'a' && visited && !curr.visitedTwice && visits < 2 {
				newVistied := make(map[string]int)
				for key, value := range curr.smallCavesVisited {
					newVistied[key] = value
				}
				newVistied[close]++
				q.PushBack(state{close, newVistied, curr.history + ";" + close, true})
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day12/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day12/input.txt")

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
