package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func count(w string) map[rune]int {
	c := make(map[rune]int)
	for _, r := range w {
		c[r]++
	}
	return c
}

func isRearranged(w1, w2 string) bool {
	c1 := count(w1)
	c2 := count(w2)
	for key, value := range c1 {
		value2, ok := c2[key]
		if !(ok && (value == value2)) {
			return false
		}
	}
	return len(c1) == len(c2)
}

func part1(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		l := len(line)
		c++
	loop:
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if line[i] == line[j] {
					c--
					break loop
				}
			}
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		l := len(line)
		c++
	loop:
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				if isRearranged(line[i], line[j]) {
					c--
					break loop
				}
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day04/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day04/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day04/input.txt")

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
