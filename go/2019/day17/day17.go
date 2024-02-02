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

type robot struct {
	x, y int
	d    string
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

func solve(s string, inputList []int) string {
	var output string
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
			input := inputList[0]
			inputList = inputList[1:]
			switch mode1 {
			case 0:
				p[p[i+1]] = input
			case 2:
				p[relativeBase+p[i+1]] = input
			}
			i += 2
			continue mainLoop
		case 4:
			//fmt.Println(output)
			output += string(rune(a))
			i += 2
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

func isValid(i, j int, n, m int) bool {
	return i >= 0 && j >= 0 && i < n && j < m
}

func next(r robot) (int, int) {
	switch r.d {
	case "^":
		return r.x - 1, r.y
	case "v":
		return r.x + 1, r.y
	case "<":
		return r.x, r.y - 1
	case ">":
		return r.x, r.y + 1
	}
	return -1, -1
}

func simpleDirection(d string) string {
	switch d {
	case "<", ">":
		return "<"
	case "^", "v":
		return "^"
	}
	return ""
}

func concatenatePath(s string) string {
	res := ""
	k := 0
	d := '^'
	i := 0
	for i < len(s) {
		r := rune(s[i])
		if r == '/' {
			res += strconv.Itoa(k) + ","
			k = 0
			newDirection := rune(s[i+1])
			switch d {
			case '^':
				switch newDirection {
				case '>':
					res += "R"
				case '<':
					res += "L"
				}
			case 'v':
				switch newDirection {
				case '>':
					res += "L"
				case '<':
					res += "R"
				}
			case '>':
				switch newDirection {
				case '^':
					res += "L"
				case 'v':
					res += "R"
				}
			case '<':
				switch newDirection {
				case '^':
					res += "R"
				case 'v':
					res += "L"
				}
			}
			res += ","
			d = newDirection
			i += 3
			continue
		} else {
			k++
			i++
		}
	}
	res += strconv.Itoa(k)
	return res[2:]
}

func calculateRoutine(s string) (string, string, string, string) {
	cpy := s
	res := []string{}
	for i := 0; i < 3; i++ {
		tab := strings.Split(s, ",")
		i := 6
		function := tab[0:6]
		k := strings.Count(s, strings.Join(function, ","))
		for {
			tempTab := tab[0 : i+2]
			temp := strings.Count(s, strings.Join(tempTab, ","))
			if temp < k {
				break
			} else {
				k += 2
				function = tempTab
			}

		}
		res = append(res, strings.Join(function, ","))
		s = strings.Replace(s, strings.Join(function, ",")+",", "", -1)
		s = strings.Replace(s, strings.Join(function, ","), "", -1)
	}
	s = cpy
	s = strings.Replace(s, res[0], "A", -1)
	s = strings.Replace(s, res[1], "B", -1)
	s = strings.Replace(s, res[2], "C", -1)
	return s + "\n", res[0] + "\n", res[1] + "\n", res[2] + "\n"
}

func part1(s string) int {
	c := 0
	str := solve(s, []int{})
	tab := strings.Split(str, "\n")
	tab = tab[:len(tab)-2]
	n, m := len(tab), len(tab[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if isValid(i+1, j, n, m) && isValid(i-1, j, n, m) && isValid(i, j+1, n, m) && isValid(i, j-1, n, m) {
				if tab[i+1][j] == '#' && tab[i-1][j] == '#' && tab[i][j+1] == '#' && tab[i][j-1] == '#' && tab[i][j] == '#' {
					c += i * j
				}
			}
		}
	}
	return c
}

func part2(s string) int {
	str := solve(s, []int{})
	tab := strings.Split(str, "\n")
	tab = tab[:len(tab)-2]
	n, m := len(tab), len(tab[0])
	visited := make(map[robot]bool)
	var curr robot
	for i, line := range tab {
		for j, r := range line {
			if r == '^' {
				curr = robot{i, j, "^"}
			}
		}
	}
	path := ""
	var dq deque.Deque[robot]
	dq.PushBack(curr)
	for dq.Len() != 0 {
		curr = dq.PopFront()
		visited[robot{curr.x, curr.y, simpleDirection(curr.d)}] = true
		i, j := next(curr)
		_, ok := visited[robot{i, j, simpleDirection(curr.d)}]
		if !ok && isValid(i, j, n, m) && tab[i][j] == '#' {

			dq.PushBack(robot{i, j, curr.d})
			path += curr.d
			continue
		}
		for _, d := range []string{"^", "<", ">", "v"} {
			i, j := next(robot{curr.x, curr.y, d})
			_, ok := visited[robot{i, j, simpleDirection(curr.d)}]
			if !ok && isValid(i, j, n, m) && tab[i][j] == '#' {
				path += "/" + d + "/"
				dq.PushBack(robot{curr.x, curr.y, d})
			}
		}
	}
	path = concatenatePath(path)
	main, A, B, C := calculateRoutine(path)
	s = "2" + s[1:]
	inputList := []int{}
	for _, r := range main + A + B + C {
		inputList = append(inputList, int(r))
	}
	inputList = append(inputList, int('n'), int('\n'))
	res := solve(s, inputList)
	var dust int
	for _, r := range res {
		dust = int(r)
	}
	return dust
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day17/input.data")

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
