package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) (int, int) {
	strs := strings.Split(s, "-")
	a, _ := strconv.Atoi(strs[0])
	b, _ := strconv.Atoi(strs[1])
	return a, b
}

func sixDigit(s string) bool {
	return len(s) == 6
}

func adjacent(s string) bool {
	if len(s) == 1 {
		return false
	}
	return s[0] == s[1] || adjacent(s[1:])
}

func adjacent2(s string) bool {
	if len(s) == 1 {
		return false
	}
	if len(s) == 2 {
		return s[0] == s[1]
	}
	if len(s) == 3 {
		return (s[0] == s[1] && s[0] != s[2]) || (s[1] == s[2] && s[0] != s[1])
	}
	if s[0] == s[1] && s[1] != s[2] {
		return true
	}
	if s[0] == s[1] && s[1] == s[2] {
		return adjacent2(s[3:]) && adjacent2(s[2:])
	}
	return adjacent2(s[1:])
}

func neverDecrease(s string) bool {
	if len(s) == 1 {
		return true
	}
	return int(s[0]) <= int(s[1]) && neverDecrease(s[1:])
}

func isValid(x int) bool {
	s := strconv.Itoa(x)
	return sixDigit(s) && adjacent(s) && neverDecrease(s)
}

func isValid2(x int) bool {
	s := strconv.Itoa(x)
	return sixDigit(s) && adjacent2(s) && neverDecrease(s)
}

func part1(s string) int {
	c := 0
	a, b := format(s)
	for i := a; i <= b; i++ {
		if isValid(i) {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	a, b := format(s)
	for i := a; i <= b; i++ {
		if isValid2(i) {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day04/input.data")

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
