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

func format(s string) (int, int) {
	line := strings.Split(s, " ")
	n, _ := strconv.Atoi(line[0])
	m, _ := strconv.Atoi(line[6])
	return n, m
}

func max(l []int) int {
	m := l[0]
	for _, x := range l {
		if x > m {
			m = x
		}
	}
	return m
}

func part1(s string) int {
	numberOfPlayers, lastMarble := format(s)
	score := make([]int, numberOfPlayers)
	var queue deque.Deque[int]
	queue.PushBack(0)
	for i := 1; i < lastMarble+1; i++ {
		j := i % numberOfPlayers
		if i%23 == 0 {
			queue.Rotate(7)
			score[j] += i + queue.PopFront()
			queue.Rotate(-1)
			continue
		}
		queue.Rotate(-1)
		queue.PushFront(i)

	}
	return max(score)
}

func part2(s string) int {
	numberOfPlayers, lastMarble := format(s)
	lastMarble *= 100
	score := make([]int, numberOfPlayers)
	var queue deque.Deque[int]
	queue.PushBack(0)
	for i := 1; i < lastMarble+1; i++ {
		j := i % numberOfPlayers
		if i%23 == 0 {
			queue.Rotate(7)
			score[j] += i + queue.PopFront()
			queue.Rotate(-1)
			continue
		}
		queue.Rotate(-1)
		queue.PushFront(i)

	}
	return max(score)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day09/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day09/input.txt")

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
