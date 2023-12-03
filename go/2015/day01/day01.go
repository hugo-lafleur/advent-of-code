package main

import (
	"fmt"
	"log"
	"os"
	"time"
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
