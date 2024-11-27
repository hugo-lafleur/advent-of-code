package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type queue []int

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Push(n int) {
	*q = append(*q, n)
}

func (q *queue) Pop() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func part1(s string) int {
	elfs := []int{}
	n, _ := strconv.Atoi(s)
	for i := 1; i < n+1; i++ {
		elfs = append(elfs, i)
	}
	odd := true
	for len(elfs) > 1 {
		switch odd {
		case true:
			if len(elfs)%2 == 1 {
				odd = !odd
			}
			new := []int{}
			for i, elf := range elfs {
				if i%2 == 0 {
					new = append(new, elf)
				}
			}
			elfs = new
		case false:
			if len(elfs)%2 == 1 {
				odd = !odd
			}
			new := []int{}
			for i, elf := range elfs {
				if i%2 == 1 {
					new = append(new, elf)
				}
			}
			elfs = new
		}
	}
	return elfs[0]
}

func part2(s string) int {
	n, _ := strconv.Atoi(s)
	var right queue
	var left queue
	for i := 1; i < n+1; i++ {
		if i < (n/2)+1 {
			left.Push(i)
		} else {
			right.Push(i)
		}
	}
	for i := 1; i < n; i++ {
		right.Pop()
		if len(right) == len(left) {
			a, _ := left.Pop()
			b, _ := right.Pop()
			right.Push(a)
			left.Push(b)
		} else {
			a, _ := left.Pop()
			right.Push(a)
		}
	}
	res, _ := right.Pop()
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day19/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2016/day19/input.txt")

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
