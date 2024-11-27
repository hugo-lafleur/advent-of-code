package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"
)

type state struct {
	x, y     int
	sequence string
}

type queue []state

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Push(s state) {
	*q = append(*q, s)
}

func (q *queue) Pop() (state, bool) {
	if q.IsEmpty() {
		return state{0, 0, ""}, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func open(r rune) bool {
	return r == 'b' || r == 'c' || r == 'd' || r == 'e' || r == 'f'
}

func next(st state) []state {
	res := []state{}
	if st.x > 0 && open(rune(GetMD5Hash(st.sequence)[2])) {
		res = append(res, state{st.x - 1, st.y, st.sequence + "L"})
	}
	if st.x < 3 && open(rune(GetMD5Hash(st.sequence)[3])) {
		res = append(res, state{st.x + 1, st.y, st.sequence + "R"})
	}
	if st.y > 0 && open(rune(GetMD5Hash(st.sequence)[0])) {
		res = append(res, state{st.x, st.y - 1, st.sequence + "U"})
	}
	if st.y < 3 && open(rune(GetMD5Hash(st.sequence)[1])) {
		res = append(res, state{st.x, st.y + 1, st.sequence + "D"})
	}
	return res
}

func end(st state) bool {
	return st.x == 3 && st.y == 3
}

func part1(s string) string {
	start := state{0, 0, s}
	var queue queue
	queue.Push(start)
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		if end(curr) {
			return curr.sequence[len(s):]
		} else {
			toAdd := next(curr)
			for _, x := range toAdd {
				queue.Push(x)
			}
		}
	}
	return ""
}

func part2(s string) int {
	start := state{0, 0, s}
	var queue queue
	queue.Push(start)
	max := 0
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		if end(curr) {
			if len(curr.sequence)-len(s) > max {
				max = len(curr.sequence) - len(s)
			}
			continue
		}
		toAdd := next(curr)
		for _, x := range toAdd {
			queue.Push(x)
		}
	}
	return max
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day17/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2016/day17/input.txt")

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
