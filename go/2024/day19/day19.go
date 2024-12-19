package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	cache map[string]int
)

func parse(s string) ([]string, []string) {
	var lines = strings.Split(s, "\n")
	return strings.Split(lines[0], ", "), lines[2:]
}

func countWays(design string, patterns []string) int {
	if len(design) == 0 {
		return 1
	}
	if val, ok := cache[design]; ok {
		return val
	}
	var result int
	for _, p := range patterns {
		l := len(p)
		if l <= len(design) && design[:l] == p {
			result += countWays(design[l:], patterns)
		}
	}
	cache[design] = result
	return result
}

func part1(s string) int {
	var patterns, designs = parse(s)
	cache = make(map[string]int)
	var result int
	for i := range designs {
		if countWays(designs[i], patterns) != 0 {
			result++
		}
	}
	return result
}

func part2(s string) int {
	var patterns, designs = parse(s)
	//cache = make(map[string]int)
	var result int
	for i := range designs {
		result += countWays(designs[i], patterns)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day19/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day19/input.txt")

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
