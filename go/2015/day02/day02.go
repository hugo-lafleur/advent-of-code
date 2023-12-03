package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]int {
	data := [][]int{}
	tab := [][]string{}
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		tab = append(tab, strings.Split(line, "x"))
	}

	for i, line := range tab {
		data = append(data, []int{})
		for _, x := range line {
			n, _ := strconv.Atoi(x)
			data[i] = append(data[i], n)
		}
	}
	return data
}
func min(a int, b int, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func min2(a int, b int, c int) (int, int) {
	i := 0
	m := a
	if b > m {
		m = b
		i = 1
	}
	if c > m {
		m = c
		i = 2
	}
	switch i {
	case 0:
		return b, c
	case 1:
		return a, c
	case 2:
		return a, b
	}
	return 0, 0

}

func part1(s string) int {
	tab := format(s)
	r := 0
	for _, line := range tab {
		a1 := line[0] * line[1] * 2
		a2 := line[0] * line[2] * 2
		a3 := line[2] * line[1] * 2
		r += a1 + a2 + a3
		r += min(a1, a2, a3) / 2
	}
	return r
}

func part2(s string) int {
	tab := format(s)
	r := 0
	for _, line := range tab {
		a := line[0]
		b := line[1]
		c := line[2]
		x, y := min2(a, b, c)
		r += 2*x + 2*y
		r += a * b * c

	}
	return r
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
