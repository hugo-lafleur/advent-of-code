package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

type state struct {
	amphiods [][]point
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func (s state) finished() []bool {
	res := []bool{}
loop:
	for j := range s.amphiods {
		for i := range s.amphiods[j] {
			if s.amphiods[j][i].y != 3+j*2 && s.amphiods[j][i].y != 0 {
				res = append(res, false)
				continue loop
			}
		}
		res = append(res, true)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func hash(s state) int {
	var res strings.Builder
loop1:
	for y := 1; y <= 11; y++ {
		for i, amph := range s.amphiods {
			if slices.Contains(amph, point{1, y}) {
				res.WriteString(strconv.Itoa(i + 1))
				continue loop1
			}
		}
		res.WriteString("0")
	}
	for x := 2; x <= 5; x++ {
	loop2:
		for _, y := range []int{3, 5, 7, 9} {
			for i, amph := range s.amphiods {
				if slices.Contains(amph, point{x, y}) {
					res.WriteString(strconv.Itoa(i + 1))
					continue loop2
				}
			}
			res.WriteString("0")
		}
	}
	n, _ := strconv.ParseInt(res.String(), 5, 64)
	return int(n)
}

func copyState(s state) state {
	newState := state{}
	newA, newB, newC, newD := make([]point, 4), make([]point, 4), make([]point, 4), make([]point, 4)
	copy(newA, s.amphiods[0])
	copy(newB, s.amphiods[1])
	copy(newC, s.amphiods[2])
	copy(newD, s.amphiods[3])
	newState.amphiods = make([][]point, 4)
	newState.amphiods[0] = newA
	newState.amphiods[1] = newB
	newState.amphiods[2] = newC
	newState.amphiods[3] = newD
	return newState
}

func accessDenied(p1, p2 point, l []point) bool {
	var column int
	if p1.x == 1 {
		column = p2.y
	} else {
		column = p1.y
	}
	for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
		if slices.Contains(l, point{min(p1.x, p2.x), y}) {
			return true
		}
	}
	for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
		if slices.Contains(l, point{x, column}) {
			return true
		}
	}
	return false
}

func energy_per_step(i int) int {
	switch i {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	default:
		return 1000
	}
}

func except(l []point, index int) []point {
	res := []point{}
	for j, x := range l {
		if j != index {
			res = append(res, x)
		}
	}
	return res
}

func part1(s string) int {
	res := 100000000000
	tab := format(s)
	max_depth := len(tab) - 2
	start := state{}
	start.amphiods = make([][]point, 4)
	for i, line := range tab {
		for j, c := range line {
			switch c {
			case "A":
				start.amphiods[0] = append(start.amphiods[0], point{i, j})
				tab[i][j] = "."
			case "B":
				start.amphiods[1] = append(start.amphiods[1], point{i, j})
				tab[i][j] = "."
			case "C":
				start.amphiods[2] = append(start.amphiods[2], point{i, j})
				tab[i][j] = "."
			case "D":
				start.amphiods[3] = append(start.amphiods[3], point{i, j})
				tab[i][j] = "."
			}
		}
	}
	shortest := make(map[int]int)
	var q deque.Deque[state]
	q.PushBack(start)
	shortest[hash(start)] = 0
mainLoop:
	for q.Len() != 0 {
		curr := q.PopFront()
		curr_hash := hash(curr)
		if shortest[curr_hash] > res {
			continue
		}
		organized := curr.finished()
		if organized[0] && organized[1] && organized[2] && organized[3] {
			res = min(res, shortest[curr_hash])
			continue
		}
		for j, amphi := range curr.amphiods {
			if !organized[j] {
				for i, a := range amphi {
					if a.x == 1 {
						for x := max_depth; x >= 2; x-- {
							low := point{max_depth, 3 + j*2}
							p := point{x, 3 + j*2}
							if !accessDenied(a, low, curr.amphiods[(j+1)%4]) && !accessDenied(a, low, curr.amphiods[(j+2)%4]) && !accessDenied(a, low, curr.amphiods[(j+3)%4]) && !accessDenied(a, p, except(curr.amphiods[j], i)) {
								newState := copyState(curr)
								newState.amphiods[j][i] = p
								newValue := dist(a, p)*energy_per_step(j) + shortest[curr_hash]
								new_hash := hash(newState)
								_, seen := shortest[new_hash]
								if !seen {
									q.PushBack(newState)
									shortest[new_hash] = newValue
									continue mainLoop
								} else {
									shortest[new_hash] = min(newValue, shortest[new_hash])
								}
							}
						}
					}
					if a.x >= 2 {
						for _, y := range []int{1, 2, 4, 6, 8, 10, 11} {
							p := point{1, y}
							if !accessDenied(a, p, curr.amphiods[(j+1)%4]) && !accessDenied(a, p, curr.amphiods[(j+2)%4]) && !accessDenied(a, p, curr.amphiods[(j+3)%4]) && !accessDenied(a, p, except(curr.amphiods[j], i)) {
								newState := copyState(curr)
								newState.amphiods[j][i] = p
								newValue := dist(a, p)*energy_per_step(j) + shortest[curr_hash]
								new_hash := hash(newState)
								_, seen := shortest[new_hash]
								if !seen {
									q.PushBack(newState)
									shortest[new_hash] = newValue
								} else {
									shortest[new_hash] = min(newValue, shortest[new_hash])
								}
							}
						}
					}
				}
			}
		}
	}
	return res
}

func part2(s string) int {
	lines := strings.Split(s, "\n")
	newStr := strings.Join(lines[:3], "\n")
	newStr += "\n  #D#C#B#A#\n  #D#B#A#C#\n"
	newStr += strings.Join(lines[3:], "\n")
	return part1(newStr)
}

func main() {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	content, err := os.ReadFile("../../../inputs/2021/day23/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day23/input.data")

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
