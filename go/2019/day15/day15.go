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

type point struct {
	x, y int
}

func format(s string) map[int]int {
	res := make(map[int]int)
	lines := strings.Split(s, ",")
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		res[i] = n
	}
	return res
}

func opcodeSolve(n int) (int, int, int, int) {
	s := strconv.Itoa(n)
	for len(s) != 5 {
		s = "0" + s
	}
	var a, b, c, de int
	a, _ = strconv.Atoi(string(s[0]))
	b, _ = strconv.Atoi(string(s[1]))
	c, _ = strconv.Atoi(string(s[2]))
	de, _ = strconv.Atoi(s[3:])
	return de, c, b, a
}

func solve(s string, directions string) int {
	var output int
	p := format(s)
	i := 0
	relativeBase := 0
mainLoop:
	for {
		opcode, mode1, mode2, mode3 := opcodeSolve(p[i])
		var a, b, c, res int
		if opcode == 99 {
			break
		}
		switch mode1 {
		case 0:
			a = p[p[i+1]]
		case 1:
			a = p[i+1]
		case 2:
			a = p[relativeBase+p[i+1]]
		}
		switch mode2 {
		case 0:
			b = p[p[i+2]]
		case 1:
			b = p[i+2]
		case 2:
			b = p[relativeBase+p[i+2]]
		}
		switch mode3 {
		case 0:
			c = p[i+3]
		case 2:
			c = p[i+3] + relativeBase
		}
		switch opcode {
		case 1:
			res = a + b
			i += 3
		case 2:
			res = a * b
			i += 3
		case 3:
			input, _ := strconv.Atoi(directions[0:1])
			directions = directions[1:]
			switch mode1 {
			case 0:
				p[p[i+1]] = input
			case 2:
				p[relativeBase+p[i+1]] = input
			}
			i += 2
			continue mainLoop
		case 4:
			output = a
			i += 2
			if len(directions) == 0 {
				break mainLoop
			}
			continue mainLoop
		case 5:
			if a != 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 6:
			if a == 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 7:
			if a < b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 8:
			if a == b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 9:
			relativeBase += a
			i += 2
			continue mainLoop
		}
		p[c] = res
		i++
	}
	return output
}

func stringToPoint(s string) point {
	var res point
	for _, x := range s {
		switch x {
		case '1':
			res.x++
		case '2':
			res.x--
		case '3':
			res.y++
		case '4':
			res.y--
		}
	}
	return res
}

func part1(s string) int {
	var dq deque.Deque[string]
	visited := make(map[point]bool)
	dq.PushBack("")
	visited[point{0, 0}] = true
	for dq.Len() != 0 {
		curr := dq.PopFront()
		for i := 1; i <= 4; i++ {
			next := curr + strconv.Itoa(i)
			_, ok := visited[stringToPoint(next)]
			if !ok {
				visited[stringToPoint(next)] = true
				res := solve(s, next)
				if res == 2 {
					return len(next)
				}
				if res == 1 {
					dq.PushBack(next)
				}
			}
		}
	}
	return 0
}

func part2(s string) int {
	var dq deque.Deque[string]
	visited := make(map[point]bool)
	visited[point{0, 0}] = true
	oxygen := ""
	dq.PushBack("")
	visited[point{0, 0}] = true
loop:
	for dq.Len() != 0 {
		curr := dq.PopFront()
		for i := 1; i <= 4; i++ {
			next := curr + strconv.Itoa(i)
			_, ok := visited[stringToPoint(next)]
			if !ok {
				visited[stringToPoint(next)] = true
				res := solve(s, next)
				if res == 2 {
					oxygen = next
					break loop
				}
				if res == 1 {
					dq.PushBack(next)
				}
			}
		}
	}
	var dq2 deque.Deque[string]
	visited = make(map[point]bool)
	dq2.PushBack(oxygen)
	visited[stringToPoint(oxygen)] = true
	var minutes int
	for dq2.Len() != 0 {
		curr := dq2.PopFront()
		for i := 1; i <= 4; i++ {
			next := curr + strconv.Itoa(i)
			_, ok := visited[stringToPoint(next)]
			if !ok {
				visited[stringToPoint(next)] = true
				res := solve(s, next)
				if res == 1 {
					minutes = len(next)
					dq2.PushBack(next)
				}
			}
		}
	}
	return minutes - len(oxygen)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day15/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
