package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

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

func solve(s string, inputList []int) int {
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
			output = a
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

func isLineBottomLine(s string, start, i int) (bool, int) {
	for j := start; ; j++ {
		if solve(s, []int{i, j}) == 1 {
			if solve(s, []int{i - 99, j + 99}) == 1 {
				return true, (i-99)*10000 + j
			}
			return false, j
		}
	}
}

func part1(s string) int {
	c := 0
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			c += solve(s, []int{i, j})
		}
	}
	return c
}

func part2(s string) int {
	var start int
	for i := 10; ; i++ {
		ok, res := isLineBottomLine(s, start, i)
		if ok {
			return res
		}
		start = res
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day19/input.data")

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
