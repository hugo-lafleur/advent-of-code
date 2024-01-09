package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "")
}

func newTrap(a, b, c string) bool {
	if a == "^" && b == "^" && c == "." {
		return true
	}
	if a == "." && b == "^" && c == "^" {
		return true
	}
	if a == "^" && b == "." && c == "." {
		return true
	}
	if a == "." && b == "." && c == "^" {
		return true
	}
	return false
}

func nextLine(line []string) []string {
	res := []string{}
	if newTrap(".", line[0], line[1]) {
		res = append(res, "^")
	} else {
		res = append(res, ".")
	}
	for i := 1; i < len(line)-1; i++ {
		if newTrap(line[i-1], line[i], line[i+1]) {
			res = append(res, "^")
		} else {
			res = append(res, ".")
		}
	}
	if newTrap(line[len(line)-2], line[len(line)-1], ".") {
		res = append(res, "^")
	} else {
		res = append(res, ".")
	}
	return res
}

func part1(s string) int {
	c := 0
	line := format(s)
	var l int
	if len(line) == 10 {
		l = 10
	} else {
		l = 40
	}
	for i := 0; i < l; i++ {
		for _, tile := range line {
			if tile == "." {
				c++
			}
		}
		line = nextLine(line)
	}
	return c
}

func part2(s string) int {
	c := 0
	line := format(s)
	l := 400000
	for i := 0; i < l; i++ {
		for _, tile := range line {
			if tile == "." {
				c++
			}
		}
		line = nextLine(line)
	}
	return c
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
