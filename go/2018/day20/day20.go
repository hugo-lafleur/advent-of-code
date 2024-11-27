package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gammazero/deque"
)

func directionToComplex(r rune) complex128 {
	switch r {
	case 'N':
		return 1i
	case 'S':
		return -1i
	case 'E':
		return 1
	case 'W':
		return -1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxDistance(p map[complex128]int) int {
	m := 0
	for _, value := range p {
		if value > m {
			m = value
		}
	}
	return m
}

func atLeast1000Doors(p map[complex128]int) int {
	res := 0
	for _, value := range p {
		if value >= 1000 {
			res++
		}
	}
	return res
}

func solve(s string) map[complex128]int {
	positions := make(map[complex128]int)
	var dq deque.Deque[complex128]
	curr := complex(0, 0)
	positions[curr] = 0
	for _, r := range s {
		if r == '(' {
			dq.PushBack(curr)
		}
		if r == ')' {
			curr = dq.PopBack()
		}
		if r == '|' {
			curr = dq.Back()
		} else {
			new := curr + directionToComplex(r)
			if _, visited := positions[new]; visited {
				positions[new] = min(positions[new], positions[curr]+1)
			} else {
				positions[new] = positions[curr] + 1
			}
			curr = new
		}
	}
	return positions
}

func part1(s string) int {
	input := s[1 : len(s)-1]
	return maxDistance(solve(input))
}

func part2(s string) int {
	input := s[1 : len(s)-1]
	return atLeast1000Doors(solve(input))
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day20/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	content, err = os.ReadFile("../../../inputs/2018/day20/input.txt")

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
