package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type stack []rune

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *stack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	}
	l := len(*s)
	h := (*s)[l-1]
	*s = (*s)[:l-1]
	return h, true
}

func sameType(r1, r2 rune) bool {
	return int(r1) == int(r2)+32 || int(r1) == int(r2)-32
}

func part1(s string) int {
	polymer := s
	var stack stack
	for _, r := range polymer {
		if stack.IsEmpty() {
			stack.Push(r)
			continue
		}
		l, _ := stack.Pop()
		if sameType(l, r) {
			continue
		}
		stack.Push(l)
		stack.Push(r)
	}
	return len(stack)
}

func part2(s string) int {
	polymer := s
	min := len(polymer)
	for i := 65; i < 91; i++ {
		polymer := s
		var stack stack
		for _, r := range polymer {
			if r == rune(i) || r == rune(i+32) {
				continue
			}
			if stack.IsEmpty() {
				stack.Push(r)
				continue
			}
			l, _ := stack.Pop()
			if sameType(l, r) {
				continue
			}
			stack.Push(l)
			stack.Push(r)
		}
		l := len(stack)
		if l < min {
			min = l
		}
	}
	return min
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day05/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2018/day05/input.txt")

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
