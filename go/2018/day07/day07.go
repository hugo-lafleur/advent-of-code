package main

import (
	"container/heap"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type stringHeap []string

func (h stringHeap) Len() int {
	return len(h)
}

func (h stringHeap) Less(i, j int) bool {
	return []rune(h[i])[0] < []rune(h[j])[0]
}

func (h stringHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *stringHeap) Push(x any) {
	*h = append(*h, x.(string))
}

func (h *stringHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func format(s string) map[string][]string {
	lines := strings.Split(s, "\n")
	res := make(map[string][]string)
	for _, line := range lines {
		strs := strings.Split(line, " ")
		res[strs[1]] = []string{}

	}
	for _, line := range lines {
		strs := strings.Split(line, " ")
		res[strs[7]] = append(res[strs[7]], strs[1])
	}
	return res
}

func remove(l []string, x string) []string {
	res := []string{}
	for _, s := range l {
		if x != s {
			res = append(res, s)
		}
	}
	return res
}

func part1(s string) string {
	requirements := format(s)
	res := ""
	h := &stringHeap{}
	heap.Init(h)
	for {
		if len(requirements) == 0 {
			return res
		}
		for key, value := range requirements {
			if len(value) == 0 {
				heap.Push(h, key)
				delete(requirements, key)
			}
		}
		curr := heap.Pop(h)
		res += curr.(string)
		for k, v := range requirements {
			requirements[k] = remove(v, curr.(string))
		}
	}
}

func part2(s string) int {
	requirements := format(s)
	h := &stringHeap{}
	var time, workersMax int
	if len(requirements) == 6 {
		time = 0
		workersMax = 2
	} else {
		time = 60
		workersMax = 5
	}
	workers := make(map[string]int)
	ticks := 0
	heap.Init(h)
	for {
		if len(requirements) == 0 && len(workers) == 0 && len(*h) == 0 {
			return ticks
		}
		for key, value := range requirements {
			if len(value) == 0 {
				heap.Push(h, key)
				delete(requirements, key)
			}
		}
		for len(workers) < workersMax && len(*h) != 0 {
			curr := heap.Pop(h)
			workers[curr.(string)] = int([]rune(curr.(string))[0]) - 64 + time
		}
		for key := range workers {
			workers[key]--
		}
		finished := []string{}
		for key, value := range workers {
			if value == 0 {
				finished = append(finished, key)
				delete(workers, key)
			}
		}
		for _, toRemove := range finished {
			for k, v := range requirements {
				requirements[k] = remove(v, toRemove)
			}
		}
		ticks++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day07/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2018/day07/input.txt")

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
