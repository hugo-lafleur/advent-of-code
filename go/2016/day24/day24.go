package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type state struct {
	x, y             int
	pointsOfInterest string
	steps            int
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
		return state{}, false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func isValid(i, j int, tab [][]string) bool {
	return i >= 0 && i < len(tab) && j >= 0 && j < len(tab[0]) && tab[i][j] != "#"
}

func isNumber(i, j int, tab [][]string) bool {
	p := tab[i][j]
	return p == "1" || p == "2" || p == "3" || p == "4" || p == "5" || p == "6" || p == "7" || p == "8" || p == "9"
}

func isIn(n string, points string) bool {
	for _, x := range points {
		if string(x) == n {
			return true
		}
	}
	return false
}

func stateToString(s state) string {
	return strconv.Itoa(s.x) + "," + strconv.Itoa(s.y) + "," + s.pointsOfInterest
}

func next(s state, tab [][]string, visited map[string]bool) []state {
	res := []state{}
	i := s.x + 1
	j := s.y
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x - 1
	j = s.y
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x
	j = s.y + 1
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x
	j = s.y - 1
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	return res
}

func next2(s state, tab [][]string, visited map[string]bool, goal int) []state {
	res := []state{}
	i := s.x + 1
	j := s.y
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if n == "0" && done(s, goal) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		}
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x - 1
	j = s.y
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if n == "0" && done(s, goal) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		}
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x
	j = s.y + 1
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if n == "0" && done(s, goal) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		}
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	i = s.x
	j = s.y - 1
	if isValid(i, j, tab) {
		n := tab[i][j]
		_, ok := visited[stateToString(state{i, j, s.pointsOfInterest, s.steps})]
		_, okAdd := visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps})]
		if n == "0" && done(s, goal) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		}
		if !isIn(n, s.pointsOfInterest) && !okAdd && isNumber(i, j, tab) {
			res = append(res, state{i, j, s.pointsOfInterest + n, s.steps + 1})
			visited[stateToString(state{i, j, s.pointsOfInterest + n, s.steps + 1})] = true
		} else {
			if !ok {
				res = append(res, state{i, j, s.pointsOfInterest, s.steps + 1})
				visited[stateToString(state{i, j, s.pointsOfInterest, s.steps + 1})] = true
			}
		}
	}
	return res
}

func done(s state, goal int) bool {
	return len(s.pointsOfInterest) == goal
}

func done2(s state, goal int) bool {
	return len(s.pointsOfInterest) == goal+1 && string(s.pointsOfInterest[len(s.pointsOfInterest)-1]) == "0"
}

func part1(s string) int {
	visited := make(map[string]bool)
	tab := format(s)
	var curr state
	var goal int
	for i, line := range tab {
		for j, p := range line {
			if p == "0" {
				curr = state{i, j, "", 0}
			}
			if p == "1" || p == "2" || p == "3" || p == "4" || p == "5" || p == "6" || p == "7" || p == "8" || p == "9" {
				goal++
			}
		}
	}
	var queue queue
	queue.Push(curr)
	visited[stateToString(curr)] = true
	i := 0
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		if done(curr, goal) {
			return curr.steps
		} else {
			toAdd := next(curr, tab, visited)
			for _, s := range toAdd {
				queue.Push(s)
			}
		}
		i++
	}
	return 0
}

func part2(s string) int {
	visited := make(map[string]bool)
	tab := format(s)
	var curr state
	var goal int
	for i, line := range tab {
		for j, p := range line {
			if p == "0" {
				curr = state{i, j, "", 0}
			}
			if p == "1" || p == "2" || p == "3" || p == "4" || p == "5" || p == "6" || p == "7" || p == "8" || p == "9" {
				goal++
			}
		}
	}
	var queue queue
	queue.Push(curr)
	visited[stateToString(curr)] = true
	i := 0
	for !queue.IsEmpty() {
		curr, _ := queue.Pop()
		if done2(curr, goal) {
			return curr.steps
		} else {
			toAdd := next2(curr, tab, visited, goal)
			for _, s := range toAdd {
				queue.Push(s)
			}
		}
		i++
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day24/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day24/input.data")

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
