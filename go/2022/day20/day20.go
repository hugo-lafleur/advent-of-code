package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type tuple struct {
	p, x int
}

func format(s string) []int {
	lines := strings.Split(s, "\n")
	res := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	list := format(s)
	var q deque.Deque[tuple]
	for i, x := range list {
		q.PushBack(tuple{i, x})
	}
	for p, x := range list {
		t := tuple{p, x}
		for q.Front() != t {
			q.Rotate(1)
		}
		q.PopFront()
		q.Rotate(x % q.Len())
		q.PushFront(t)
	}
	for q.Front().x != 0 {
		q.Rotate(1)
	}
	return q.At(1000%q.Len()).x + q.At(2000%q.Len()).x + q.At(3000%q.Len()).x
}

func part2(s string) int {
	list := format(s)
	for i := range list {
		list[i] = 811589153 * list[i]
	}
	var q deque.Deque[tuple]
	for p, x := range list {
		q.PushBack(tuple{p, x})
	}
	for c := 0; c < 10; c++ {
		for p, x := range list {
			t := tuple{p, x}
			for q.Front() != t {
				q.Rotate(1)
			}
			q.PopFront()
			q.Rotate(x % q.Len())
			q.PushFront(t)
		}
	}
	for q.Front().x != 0 {
		q.Rotate(1)
	}
	return q.At(1000%q.Len()).x + q.At(2000%q.Len()).x + q.At(3000%q.Len()).x
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day20/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day20/input.txt")

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
