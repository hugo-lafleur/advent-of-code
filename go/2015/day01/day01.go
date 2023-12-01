package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func part1(s string) int {
	c := 0
	for _, char := range s {
		if char == '(' {
			c++
		}
		if char == ')' {
			c--
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	r := 0
	for i, char := range s {
		if char == '(' {
			c++
		}
		if char == ')' {
			c--
		}
		if c == -1 {
			r = i
			break
		}
	}
	return r + 1
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
