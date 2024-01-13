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

func part1(s string) int {
	c := 0
	input := format(s)
	groupLevel := 0
	isGarbage := false
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case "{":
			if !isGarbage {
				groupLevel++
			}
		case "}":
			if !isGarbage {
				c += groupLevel
				groupLevel--
			}
		case "<":
			isGarbage = true
		case ">":
			isGarbage = false
		case "!":
			i++
		}

	}
	return c
}

func part2(s string) int {
	c := 0
	input := format(s)
	isGarbage := false
	for i := 0; i < len(input); i++ {
		char := input[i]
		switch char {
		case "<":
			if isGarbage {
				c++
			}
			isGarbage = true
		case ">":
			isGarbage = false
		case "!":
			i++
		default:
			if isGarbage {
				c++
			}
		}

	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day09/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day09/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day09/input.data")

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
