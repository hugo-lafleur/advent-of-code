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

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == 'x' || r == ':' || r == '#'
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func multipleClaims(f map[point]int) int {
	res := 0
	for _, value := range f {
		if value > 1 {
			res++
		}
	}
	return res
}

func part1(s string) int {
	claims := format(s)
	fabric := make(map[point]int)
	for _, claim := range claims {
		a, _ := strconv.Atoi(claim[2])
		b, _ := strconv.Atoi(claim[3])
		n, _ := strconv.Atoi(claim[4])
		m, _ := strconv.Atoi(claim[5])
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				p := point{a + i, b + j}
				fabric[p]++
			}
		}
	}
	return multipleClaims(fabric)
}

func part2(s string) string {
	claims := format(s)
	fabric := make(map[point]int)
	for _, claim := range claims {
		a, _ := strconv.Atoi(claim[2])
		b, _ := strconv.Atoi(claim[3])
		n, _ := strconv.Atoi(claim[4])
		m, _ := strconv.Atoi(claim[5])
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				p := point{a + i, b + j}
				fabric[p]++
			}
		}
	}
loop:
	for _, claim := range claims {
		a, _ := strconv.Atoi(claim[2])
		b, _ := strconv.Atoi(claim[3])
		n, _ := strconv.Atoi(claim[4])
		m, _ := strconv.Atoi(claim[5])
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				p := point{a + i, b + j}
				if fabric[p] > 1 {
					continue loop
				}
			}
		}
		return claim[0]
	}
	return ""
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day03/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day03/input.data")

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
