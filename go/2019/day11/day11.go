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

func solve(s string, m map[point]string) (map[point]bool, map[point]string) {
	var output int
	paintedOnce := make(map[point]bool)
	currentPoint := point{0, 0}
	direction := "^"
	nextIsPaint := true
	p := format(s)
	i := 0
	relativeBase := 0
	k := 0
mainLoop:
	for k < 100000 {
		opcode, mode1, mode2, mode3 := opcodeSolve(p[i])
		//fmt.Println(i, p[i], p[8], p[10])
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
			var input int
			str, ok := m[currentPoint]
			if str == "." || !ok {
				input = 0
			} else {
				input = 1
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
			output = a
			if nextIsPaint {
				if output == 0 {
					m[currentPoint] = "."
				} else {
					m[currentPoint] = "#"
					paintedOnce[currentPoint] = true
				}
			} else {
				k++
				if output == 0 {
					switch direction {
					case "^":
						currentPoint.y--
						direction = "<"
					case "<":
						currentPoint.x++
						direction = "v"
					case "v":
						currentPoint.y++
						direction = ">"
					case ">":
						currentPoint.x--
						direction = "^"
					}
				} else {
					switch direction {
					case "^":
						currentPoint.y++
						direction = ">"
					case "<":
						currentPoint.x--
						direction = "^"
					case "v":
						currentPoint.y--
						direction = "<"
					case ">":
						currentPoint.x++
						direction = "v"
					}
				}

			}
			nextIsPaint = !nextIsPaint
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
	return paintedOnce, m
}

func part1(s string) int {
	c := 0
	m := make(map[point]string)
	p, _ := solve(s, m)
	for _, value := range p {
		if value {
			c++
		}
	}
	return c
}

func part2(s string) string {
	m := make(map[point]string)
	m[point{0, 0}] = "#"
	_, m = solve(s, m)
	var minX, maxX, minY, maxY int
	for key := range m {
		minX = key.x
		maxX = key.x
		minY = key.y
		maxY = key.y
	}
	for key := range m {
		if key.x < minX {
			minX = key.x
		}
		if key.x > maxX {
			maxX = key.x
		}
		if key.y < minY {
			minY = key.y
		}
		if key.y > maxY {
			maxY = key.y
		}
	}
	res := "\n"
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			str, ok := m[point{x, y}]
			if !ok || str == "." {
				res += "."
			} else {
				res += "#"
			}
			res += " "
		}
		res += "\n"
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day11/input.data")

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
