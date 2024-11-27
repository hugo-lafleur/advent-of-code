package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var path map[string]string

type state struct {
	s []int
	c int
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
		return state{[]int{}, 0}, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func format(s string) state {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	g := 1
	m := 2
	dict := make(map[int]int)
	for i, line := range res {
		for _, word := range line {
			if word == "generator" || word == "generator." || word == "generator," {
				dict[g] = i + 1
				g += 2
			}
			if word == "microchip" || word == "microchip." || word == "microchip," {
				dict[m] = i + 1
				m += 2
			}
		}
	}
	st := make([]int, len(dict)+1)
	st[0] = 1
	for key, value := range dict {
		st[key] = value
	}
	return state{st, 0}
}

func stateToString(state state) string {
	s := state.s
	res := ""
	for _, k := range s {
		res += strconv.Itoa(k)
	}
	return res
}

func unsafeTest(state state) bool {
	s := state.s
	if s[2] == s[3] && s[2] != s[1] {
		return true
	}
	if s[4] == s[1] && s[4] != s[3] {
		return true
	}
	return false
}

func unsafe(state state) bool {
	s := state.s
	if s[2] != s[1] && (s[2] == s[3] || s[2] == s[5] || s[2] == s[7] || s[2] == s[9]) {
		return true
	}
	if s[4] != s[3] && (s[4] == s[1] || s[4] == s[5] || s[4] == s[7] || s[4] == s[9]) {
		return true
	}
	if s[6] != s[5] && (s[6] == s[3] || s[6] == s[1] || s[6] == s[7] || s[6] == s[9]) {
		return true
	}
	if s[8] != s[7] && (s[8] == s[1] || s[8] == s[5] || s[8] == s[3] || s[8] == s[9]) {
		return true
	}
	if s[10] != s[9] && (s[10] == s[1] || s[10] == s[5] || s[10] == s[3] || s[10] == s[7]) {
		return true
	}
	return false
}

func unsafe2(state state) bool {
	s := state.s
	if s[2] != s[1] && (s[2] == s[3] || s[2] == s[5] || s[2] == s[7] || s[2] == s[9] || s[2] == s[11] || s[2] == s[13]) {
		return true
	}
	if s[4] != s[3] && (s[4] == s[1] || s[4] == s[5] || s[4] == s[7] || s[4] == s[9] || s[4] == s[11] || s[4] == s[13]) {
		return true
	}
	if s[6] != s[5] && (s[6] == s[3] || s[6] == s[1] || s[6] == s[7] || s[6] == s[9] || s[6] == s[11] || s[6] == s[13]) {
		return true
	}
	if s[8] != s[7] && (s[8] == s[1] || s[8] == s[5] || s[8] == s[3] || s[8] == s[9] || s[8] == s[11] || s[8] == s[13]) {
		return true
	}
	if s[10] != s[9] && (s[10] == s[1] || s[10] == s[5] || s[10] == s[3] || s[10] == s[7] || s[10] == s[11] || s[10] == s[13]) {
		return true
	}
	if s[12] != s[11] && (s[12] == s[1] || s[12] == s[5] || s[12] == s[3] || s[12] == s[9] || s[12] == s[7] || s[12] == s[13]) {
		return true
	}
	if s[14] != s[13] && (s[14] == s[1] || s[14] == s[5] || s[14] == s[3] || s[14] == s[7] || s[14] == s[11] || s[14] == s[9]) {
		return true
	}
	return false
}

func nextTest(currentState state) []state {
	st := currentState.s
	res := []state{}
	elevator := st[0]
	var add state
	if elevator < 4 {
		curr := []int{elevator + 1, st[1], st[2], st[3], st[4]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator + 1, st[1], st[2], st[3], st[4]}
					new[i]++
					new[j]++
					add = state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafeTest(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)
					}
				}
			}
		}
		curr = []int{elevator + 1, st[1], st[2], st[3], st[4]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator + 1, st[1], st[2], st[3], st[4]}
				new[i]++
				add = state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafeTest(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
	}
	if elevator > 1 {
		curr := []int{elevator - 1, st[1], st[2], st[3], st[4]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator - 1, st[1], st[2], st[3], st[4]}
				new[i]--
				add = state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafeTest(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
		curr = []int{elevator - 1, st[1], st[2], st[3], st[4]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator - 1, st[1], st[2], st[3], st[4]}
					new[i]--
					new[j]--
					add = state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafeTest(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)

					}
				}
			}
		}
	}
	return res
}

func next(currentState state) []state {
	st := currentState.s
	res := []state{}
	elevator := st[0]
	if elevator < 4 {
		curr := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
					new[i]++
					new[j]++
					add := state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafe(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)
					}
				}
			}
		}
		curr = []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
				new[i]++
				add := state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafe(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
	}
	if elevator > 1 {
		curr := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
				new[i]--
				add := state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafe(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
		curr = []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10]}
					new[i]--
					new[j]--
					add := state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafe(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)
					}
				}
			}
		}
	}
	return res
}

func next2(currentState state) []state {
	st := currentState.s
	res := []state{}
	elevator := st[0]
	if elevator < 4 {
		curr := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
					new[i] += 1
					new[j] += 1
					add := state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafe2(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)
					}
				}
			}
		}
		curr = []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator + 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
				new[i] += 1
				add := state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafe2(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
	}
	if elevator > 1 {
		curr := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
		for i, value := range curr {
			if value == elevator {
				new := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
				new[i] -= 1
				add := state{new, currentState.c + 1}
				_, ok := path[stateToString(add)]
				if !unsafe2(add) && !ok {
					path[stateToString(add)] = stateToString(currentState)
					res = append(res, add)
				}
			}
		}
		curr = []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
		for i, value := range curr {
			for j, value2 := range curr {
				if value == elevator && value2 == elevator && i != j {
					new := []int{elevator - 1, st[1], st[2], st[3], st[4], st[5], st[6], st[7], st[8], st[9], st[10], st[11], st[12], st[13], st[14]}
					new[i] -= 1
					new[j] -= 1
					add := state{new, currentState.c + 1}
					_, ok := path[stateToString(add)]
					if !unsafe2(add) && !ok {
						path[stateToString(add)] = stateToString(currentState)
						res = append(res, add)
					}
				}
			}
		}
	}
	return res
}

func done(state state) bool {
	st := state.s
	for i := 1; i < len(st); i++ {
		if st[i] != 4 {
			return false
		}
	}
	return true
}

func part1(s string) int {
	st := format(s)
	var queue queue
	path = make(map[string]string)
	if len(st.s) == 5 {
		queue.Push(st)
		path[stateToString(st)] = "b"
		for !queue.IsEmpty() {
			curr, _ := queue.Pop()
			if done(curr) {
				return curr.c
			} else {
				adds := nextTest(curr)
				for _, st := range adds {
					queue.Push(st)
				}
			}
		}
	}
	if len(st.s) == 11 {
		queue.Push(st)
		path[stateToString(st)] = "b"
		for !queue.IsEmpty() {
			curr, _ := queue.Pop()
			if done(curr) {
				return curr.c
			} else {
				adds := next(curr)
				for _, st := range adds {
					queue.Push(st)
				}
			}
		}
	}
	return 0
}

func part2(s string) int {
	st := format(s)
	for i := 0; i < 4; i++ {
		st.s = append(st.s, 1)
	}
	var queue queue
	path = make(map[string]string)
	queue.Push(st)
	path[stateToString(st)] = "b"
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		if done(curr) {
			return curr.c
		} else {
			adds := next2(curr)
			for _, s := range adds {
				queue.Push(s)
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day11/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day11/input.txt")

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
