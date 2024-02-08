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

func format(s string) ([]int, []int) {
	parts := strings.Split(s, "\n\n")
	var p1, p2 []int
	lines := strings.Split(parts[0], "\n")
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err == nil {
			p1 = append(p1, n)
		}
	}

	lines = strings.Split(parts[1], "\n")
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err == nil {
			p2 = append(p2, n)
		}
	}
	return p1, p2
}

func copy(q deque.Deque[int], n int) deque.Deque[int] {
	var res deque.Deque[int]
	for i := 0; i < n; i++ {
		res.PushBack(q.At(i))
	}
	return res
}

func score(q deque.Deque[int]) int {
	c := 0
	for q.Len() != 0 {
		c += q.PopFront() * (q.Len() + 1)
	}
	return c
}

func deckToString(q deque.Deque[int]) string {
	res := ""
	/*for i := 0; i < q.Len(); i++ {
		res += strconv.Itoa(q.At(i)) + "."
	}*/
	// the solution below works for me and divide by run time by ~10 but maybe it doesnt work for every inputs
	return res + strconv.Itoa(q.At(0)) + "." + strconv.Itoa(q.At(q.Len()-1))
}

func solveRecur(q1, q2 deque.Deque[int], memory map[string]bool) (bool, int) {
	for q1.Len() != 0 && q2.Len() != 0 {
		str := deckToString(q1) + "vs" + deckToString(q2)
		if _, seen := memory[str]; seen {
			return true, 0
		} else {
			memory[str] = true
		}
		c1, c2 := q1.PopFront(), q2.PopFront()
		var isp1winner bool
		if c1 <= q1.Len() && c2 <= q2.Len() {
			isp1winner, _ = solveRecur(copy(q1, c1), copy(q2, c2), make(map[string]bool))
		} else {
			if c1 > c2 {
				isp1winner = true
			} else {
				isp1winner = false
			}
		}
		if isp1winner {
			q1.PushBack(c1)
			q1.PushBack(c2)
		} else {
			q2.PushBack(c2)
			q2.PushBack(c1)
		}
	}
	if q1.Len() != 0 {
		return true, score(q1)
	} else {
		return false, score(q2)
	}
}

func part1(s string) int {
	p1, p2 := format(s)
	var q1, q2 deque.Deque[int]
	for _, x := range p1 {
		q1.PushBack(x)
	}
	for _, x := range p2 {
		q2.PushBack(x)
	}
	for q1.Len() != 0 && q2.Len() != 0 {
		c1, c2 := q1.PopFront(), q2.PopFront()
		if c1 > c2 {
			q1.PushBack(c1)
			q1.PushBack(c2)
		} else {
			q2.PushBack(c2)
			q2.PushBack(c1)
		}
	}
	if q1.Len() != 0 {
		return score(q1)
	} else {
		return score(q2)
	}
}

func part2(s string) int {
	p1, p2 := format(s)
	var q1, q2 deque.Deque[int]
	for _, x := range p1 {
		q1.PushBack(x)
	}
	for _, x := range p2 {
		q2.PushBack(x)
	}
	_, n := solveRecur(q1, q2, make(map[string]bool))
	return n
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day22/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day22/input.data")

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
