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

type cups []int

func format(s string) []int {
	split := strings.Split(s, "")
	res := []int{}
	for _, x := range split {
		n, _ := strconv.Atoi(x)
		res = append(res, n)
	}
	return res
}

func move(q deque.Deque[int]) deque.Deque[int] {
	front := q.PopFront()
	a, b, c := q.PopFront(), q.PopFront(), q.PopFront()
	q.PushFront(front)
	destination := front - 1
	for destination == a || destination == b || destination == c || destination < 1 {
		destination--
		if destination < 0 {
			destination = 9
		}
	}
	for q.Front() != destination {
		q.Rotate(1)
	}
	q.PopFront()
	q.PushFront(c)
	q.PushFront(b)
	q.PushFront(a)
	q.PushFront(destination)
	for q.Front() != front {
		q.Rotate(1)
	}
	q.Rotate(1)
	return q
}

func queueToString(q deque.Deque[int]) string {
	res := ""
	for i := 1; i < q.Len(); i++ {
		res += strconv.Itoa(q.At(i))
	}
	return res
}

func (l *cups) move(start int) int {
	a := (*l)[start]
	b := (*l)[a]
	c := (*l)[b]
	dest := start - 1
	for dest == a || dest == b || dest == c || dest < 1 {
		dest--
		if dest < 1 {
			dest = len((*l)) - 1
		}
	}
	(*l)[start] = (*l)[c]
	(*l)[c] = (*l)[dest]
	(*l)[dest] = a
	return (*l)[start]
}

func part1(s string) string {
	list := format(s)
	var q deque.Deque[int]
	for _, x := range list {
		q.PushBack(x)
	}
	for i := 0; i < 100; i++ {
		q = move(q)
	}
	for q.Front() != 1 {
		q.Rotate(1)
	}
	return queueToString(q)
}

func part2(s string) int {
	list := format(s)
	linkedList := cups(make([]int, 1000001))
	for i := range list {
		if i == len(list)-1 {
			break
		}
		linkedList[list[i]] = list[i+1]
	}
	linkedList[list[len(list)-1]] = 10
	for i := 10; i < 1000000; i++ {
		linkedList[i] = i + 1
	}
	linkedList[1000000] = list[0]
	start := list[0]
	for i := 0; i < 10000000; i++ {
		start = linkedList.move(start)
	}
	return linkedList[1] * linkedList[linkedList[1]]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day23/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day23/input.data")

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
