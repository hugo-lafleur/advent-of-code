package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	index, child, metadata int
}

type stack []node

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(n node) {
	*s = append(*s, n)
}

func (s *stack) Pop() (node, bool) {
	if s.IsEmpty() {
		return node{}, false
	}
	l := len(*s)
	h := (*s)[l-1]
	*s = (*s)[:l-1]
	return h, true
}

func format(s string) []int {
	strs := strings.Fields(s)
	res := []int{}
	for _, x := range strs {
		n, _ := strconv.Atoi(x)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	c := 0
	list := format(s)
	var stack stack
	stack.Push(node{0, list[0], list[1]})
	i := 0
	for !stack.IsEmpty() {
		curr, _ := stack.Pop()
		if curr.child != 0 {
			curr.child--
			stack.Push(curr)
			i += 2
			stack.Push(node{i, list[i], list[i+1]})
			continue
		}
		if curr.child == 0 {
			for k := 0; k < curr.metadata; k++ {
				c += list[i+2+k]
			}
			i += 2 + curr.metadata - 2
		}
	}
	return c
}

func part2(s string) int {
	list := format(s)
	var stack stack
	stack.Push(node{0, list[0], list[1]})
	childs := make(map[int][]int)
	sums := make(map[int]int)
	i := 0
	for !stack.IsEmpty() {
		curr, _ := stack.Pop()
		if curr.child != 0 {
			curr.child--
			stack.Push(curr)
			i += 2
			stack.Push(node{i, list[i], list[i+1]})
			childs[curr.index] = append(childs[curr.index], i)
			continue
		}
		if curr.child == 0 {
			_, hasChild := childs[curr.index]
			if !hasChild {
				for k := 0; k < curr.metadata; k++ {
					sums[curr.index] += list[i+2+k]
				}
			} else {
				for k := 0; k < curr.metadata; k++ {
					index := list[i+2+k]
					if index >= 1 && index <= len(childs[curr.index]) {
						sums[curr.index] += sums[childs[curr.index][index-1]]
					}
				}
			}

			i += 2 + curr.metadata - 2
		}
	}
	return sums[0]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day08/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day08/input.data")

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
