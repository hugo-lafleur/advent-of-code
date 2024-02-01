package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

/*func screen(m map[point]int) string {
	maxX := 41
	maxY := 22
	s := "\n"
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			t := m[point{j, i}]
			switch t {
			case 0:
				s += " "
			case 1:
				s += "|"
			case 2:
				s += "*"
			case 3:
				s += "_"
			case 4:
				s += "o"
			}
			s += ""
		}
		s += "\n"
	}
	return s
}*/

func posBallposPaddle(m map[point]int) (int, int) {
	var ball, paddle int
	for key, value := range m {
		if value == 3 {
			paddle = key.x
		}
		if value == 4 {
			ball = key.x
		}
	}
	return ball, paddle
}

func solve(s string) (map[point]int, int) {
	var temp []int
	var input int
	var score int
	output := make(map[point]int)
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
			xBall, xPaddle := posBallposPaddle(output)
			if xBall < xPaddle {
				input = -1
			}
			if xBall > xPaddle {
				input = 1
			}
			if xBall == xPaddle {
				input = 0
			}
			switch mode1 {
			case 0:
				p[p[i+1]] = input
			case 2:
				p[relativeBase+p[i+1]] = input
			}
			i += 2
			continue mainLoop
		case 4:
			temp = append(temp, a)
			if len(temp) == 3 {
				if temp[0] == -1 && temp[1] == 0 {
					score = temp[2]
				} else {
					output[point{temp[0], temp[1]}] = temp[2]
				}
				temp = []int{}
			}
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
	return output, score
}

func part1(s string) int {
	c := 0
	output, _ := solve(s)
	for _, value := range output {
		if value == 2 {
			c++
		}
	}
	return c
}

func part2(s string) int {
	newS := "2" + s[1:]
	_, score := solve(newS)
	return score
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day13/input.data")

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
