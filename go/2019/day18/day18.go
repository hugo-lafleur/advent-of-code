package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

type state struct {
	p     point
	keys  string
	steps int
}

type multistate struct {
	p     []point
	keys  string
	steps int
}

type path struct {
	needed string
	steps  int
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func hash(s state) string {
	return strconv.Itoa(s.p.x) + "," + strconv.Itoa(s.p.y)
}

func hash2(s state) string {
	tab := strings.Split(s.keys, "")
	sort.Strings(tab)
	res := strconv.Itoa(s.p.x) + "," + strconv.Itoa(s.p.y) + "," + strings.Join(tab, "")
	return res
}

func hash3(s multistate) string {
	res := ""
	for _, p := range s.p {
		res += strconv.Itoa(p.x) + "," + strconv.Itoa(p.y) + ","
	}
	tab := strings.Split(s.keys, "")
	sort.Strings(tab)
	res += strings.Join(tab, "")
	return res
}

func doorToKey(s string) string {
	r := s[0]
	n := int(r) + 32
	return string(rune(n))
}

func isReachable(keys string, doors string) bool {
	for _, door := range doors {
		var key string
		if int(door) >= 97 && int(door) <= 122 {
			key = string(door)
		} else {
			key = doorToKey(string(door))
		}
		if !strings.Contains(keys, key) {
			return false
		}
	}
	return true
}

func addPoint(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func part1(s string) int {
	tab := format(s)
	keys := make(map[string]point)
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			if str == "@" {
				keys[str] = point{i, j}
			}
			if int(rune(str[0])) >= 97 && int(rune(str[0])) <= 122 {
				keys[str] = point{i, j}
			}
		}
	}
	paths := make(map[string]map[string]path)
	for key, value := range keys {
		paths[key] = make(map[string]path)
		var dq deque.Deque[state]
		done := make(map[string]bool)
		done[key] = true
		visited := make(map[string]bool)
		dq.PushBack(state{value, "", 0})
		for dq.Len() != 0 {
			curr := dq.PopFront()
			str := tab[curr.p.x][curr.p.y]
			if !strings.Contains(".@", str) {
				if str[0] >= 'a' {
					paths[key][str] = path{curr.keys, curr.steps}
					curr.keys += str
					done[str] = true
				}
				if str[0] <= 'Z' {
					curr.keys += str
				}
			}
			if len(done) == len(keys) {
				break
			}
			visited[hash(curr)] = true
			for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				p := addPoint(curr.p, offset)
				str := tab[p.x][p.y]
				var next state
				if str == "#" {
					continue
				} else {
					next = state{p, curr.keys, curr.steps + 1}
					_, ok := visited[hash(next)]
					if !ok {
						dq.PushBack(next)
					}
				}
			}
		}
	}
	res := 10000
	var dq deque.Deque[state]
	dq.PushBack(state{keys["@"], "@", 0})
	visited := make(map[string]int)
	for dq.Len() != 0 {
		curr := dq.PopFront()
		currKey := tab[curr.p.x][curr.p.y]
		if len(curr.keys) == len(paths) {
			res = min(res, curr.steps)
		}
		for key := range keys {
			if !strings.Contains(curr.keys, key) {
				if isReachable(curr.keys, paths[currKey][key].needed) {
					next := state{keys[key], curr.keys + key, curr.steps + paths[currKey][key].steps}
					steps, ok := visited[hash2(next)]
					if !ok || steps > next.steps {
						dq.PushBack(next)
						visited[hash2(next)] = next.steps
					}
				}
			}
		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	keys := make(map[string]point)
	k := 1
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			if str == "@" {
				tab[i][j] = strconv.Itoa(k)
				keys[strconv.Itoa(k)] = point{i, j}
				k++
			}
			if int(rune(str[0])) >= 97 && int(rune(str[0])) <= 122 {
				keys[str] = point{i, j}
			}
		}
	}
	paths := make(map[string]map[string]path)
	for key, value := range keys {
		paths[key] = make(map[string]path)
		var dq deque.Deque[state]
		done := make(map[string]bool)
		done[key] = true
		visited := make(map[string]bool)
		dq.PushBack(state{value, "", 0})
		for dq.Len() != 0 {
			curr := dq.PopFront()
			str := tab[curr.p.x][curr.p.y]
			if !strings.Contains(".1234", str) {
				if str[0] >= 'a' {
					paths[key][str] = path{curr.keys, curr.steps}
					curr.keys += str
					done[str] = true
				}
				if str[0] <= 'Z' {
					curr.keys += str
				}
			}
			if len(done) == len(keys) {
				break
			}
			visited[hash(curr)] = true
			for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				p := addPoint(curr.p, offset)
				str := tab[p.x][p.y]
				var next state
				if str == "#" {
					continue
				} else {
					next = state{p, curr.keys, curr.steps + 1}
					_, ok := visited[hash(next)]
					if !ok {
						dq.PushBack(next)
					}
				}
			}
		}
	}
	res := 10000
	var dq deque.Deque[multistate]
	dq.PushBack(multistate{[]point{keys["1"], keys["2"], keys["3"], keys["4"]}, "1234", 0})
	visited := make(map[string]int)
	for dq.Len() != 0 {
		curr := dq.PopFront()
		if len(curr.keys) == len(paths) {
			res = min(res, curr.steps)
		}
		for index := 1; index <= 4; index++ {
			str := strconv.Itoa(index)
			currKey := tab[curr.p[index-1].x][curr.p[index-1].y]
			for key := range paths[str] {
				if !strings.Contains(curr.keys, key) {
					if isReachable(curr.keys, paths[currKey][key].needed) {
						next := multistate{make([]point, 4), curr.keys + key, curr.steps + paths[currKey][key].steps}
						for i := range curr.p {
							if i == index-1 {
								next.p[i] = keys[key]
								continue
							}
							next.p[i] = curr.p[i]
						}
						steps, ok := visited[hash3(next)]
						if !ok || steps > next.steps {
							dq.PushBack(next)
							visited[hash3(next)] = next.steps
						}
					}
				}
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day18/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))

	content, err = os.ReadFile("../../../inputs/2019/day18/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day18/input1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))

	content, err = os.ReadFile("../../../inputs/2019/day18/input2.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
