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

func solve(s string, noun, verb int) int {
	p := format(s)
	if len(p) > 12 {
		p[1] = noun
		p[2] = verb
	}
	i := 0
	for {
		opcode := p[i]
		if opcode == 99 {
			break
		}
		a, b, c := p[i+1], p[i+2], p[i+3]
		switch opcode {
		case 1:
			p[c] = p[a] + p[b]
		case 2:
			p[c] = p[a] * p[b]
		}
		i += 4
	}
	return p[0]
}

func part1(s string) int {
	return solve(s, 12, 2)
}

func part2(s string) int {
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			if solve(s, n, v) == 19690720 {
				return 100*n + v
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day02/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day02/input.data")

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
