package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []int {
	res := []int{}
	lines := strings.Split(s, ",")
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func opcodeSolve(n int) (int, int, int, int) {
	s := strconv.Itoa(n)
	for len(s) != 5 {
		s = "0" + s
	}
	var a, b, c, de int
	if string(s[0]) == "0" {
		a = 0
	} else {
		a = 1
	}
	if string(s[1]) == "0" {
		b = 0
	} else {
		b = 1
	}
	if string(s[2]) == "0" {
		c = 0
	} else {
		c = 1
	}
	de, _ = strconv.Atoi(s[3:])
	return de, c, b, a
}

func solve(s string, input int) int {
	var output int
	p := format(s)
	i := 0
mainLoop:
	for {
		opcode, mode1, mode2, _ := opcodeSolve(p[i])
		var a, b, res int
		if opcode == 99 {
			break
		}
		if opcode <= 2 || opcode >= 5 {
			switch mode1 {
			case 0:
				a = p[p[i+1]]
			case 1:
				a = p[i+1]
			}
			switch mode2 {
			case 0:
				b = p[p[i+2]]
			case 1:
				b = p[i+2]
			}
		}
		switch opcode {
		case 1:
			res = a + b
			i += 3
		case 2:
			res = a * b
			i += 3
		case 3:
			res = input
			i += 1
		case 4:
			res = p[p[i+1]]
			output = res
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
		}
		p[p[i]] = res
		i++
	}
	return output
}

func part1(s string) int {
	return solve(s, 1)
}

func part2(s string) int {
	return solve(s, 5)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day05/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day05/input.data")

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
