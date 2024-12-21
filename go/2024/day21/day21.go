package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

var numKeypad = [][]byte{{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{0, '0', 'A'},
}
var dirKeypad = [][]byte{{0, '^', 'A'},
	{'<', 'v', '>'},
}

var cache = map[struct {
	str           string
	depth, robots int
}]int{}

type State struct {
	x, y int
	path string
}

type Direction struct {
	x, y int
	sym  string
}

func parse(s string) []string {
	return strings.Split(s, "\n")
}

func dfs(keypad [][]byte, start, end byte) []string {
	var result []string
	var dirs = []Direction{{0, -1, "<"}, {1, 0, "v"}, {0, 1, ">"}, {-1, 0, "^"}}
	var dq deque.Deque[State]
	for i := range keypad {
		for j := range keypad[i] {
			if keypad[i][j] == start {
				dq.PushBack(State{i, j, ""})
			}
		}
	}
	for dq.Len() != 0 {
		curr := dq.PopFront()
		if keypad[curr.x][curr.y] == end {
			result = append(result, curr.path+"A")
		}
		if len(result) > 0 {
			continue
		}
		for _, dir := range dirs {
			next := State{curr.x + dir.x, curr.y + dir.y, curr.path + dir.sym}
			if next.x >= 0 && next.y >= 0 && next.x < len(keypad) && next.y < len(keypad[0]) && keypad[next.x][next.y] != 0 {
				dq.PushBack(next)
			}
		}
	}
	slices.SortFunc(result, func(a, b string) int { return len(a) - len(b) })
	return result
}

func dp(grid [][]byte, str string, depth, robots int) int {
	if val, ok := cache[struct {
		str    string
		depth  int
		robots int
	}{str, depth, robots}]; ok {
		return val
	}
	var curr = byte('A')
	var result int
	for i := range str {
		var moves []string
		moves = dfs(grid, curr, str[i])
		if depth == robots {
			result += len(moves[0])
		} else {
			var minLength = math.MaxInt
			for _, move := range moves {
				minLength = min(minLength, dp(dirKeypad, move, depth+1, robots))
			}
			result += minLength
		}
		curr = str[i]
	}
	cache[struct {
		str    string
		depth  int
		robots int
	}{str, depth, robots}] = result
	return result
}

func part1(s string) int {
	var result int
	for _, code := range parse(s) {
		n, _ := strconv.Atoi(code[:3])
		result += n * dp(numKeypad, code, 0, 2)
	}
	return result
}

func part2(s string) int {
	var result int
	for _, code := range parse(s) {
		n, _ := strconv.Atoi(code[:3])
		result += n * dp(numKeypad, code, 0, 25)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day21/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day21/input.txt")

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
