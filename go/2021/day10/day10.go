package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

func format(s string) []string {
	return strings.Split(s, "\n")
}

func errorScore(s string) int {
	var q deque.Deque[rune]
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' || r == '<' {
			q.PushBack(r)
		} else {
			pop := q.PopBack()
			if r == ')' && pop != '(' {
				return 3
			}
			if r == ']' && pop != '[' {
				return 57
			}
			if r == '}' && pop != '{' {
				return 1197
			}
			if r == '>' && pop != '<' {
				return 25137
			}
		}
	}
	return 0
}

func incompleteScore(s string) int {
	var q deque.Deque[rune]
	for _, r := range s {
		if r == '(' || r == '[' || r == '{' || r == '<' {
			q.PushBack(r)
		} else {
			pop := q.PopBack()
			if r == ')' && pop != '(' {
				return 0
			}
			if r == ']' && pop != '[' {
				return 0
			}
			if r == '}' && pop != '{' {
				return 0
			}
			if r == '>' && pop != '<' {
				return 0
			}
		}
	}
	score := 0
	for q.Len() != 0 {
		score *= 5
		switch q.PopBack() {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, str := range list {
		c += errorScore(str)
	}
	return c
}

func part2(s string) int {
	scores := []int{}
	list := format(s)
	for _, str := range list {
		n := incompleteScore(str)
		if n > 0 {
			scores = append(scores, n)
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day10/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day10/input.data")

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
