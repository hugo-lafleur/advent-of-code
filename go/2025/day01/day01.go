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
	lines := strings.Split(s, "\n")
	result := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			result = append(result, -n)
		} else {
			result = append(result, n)
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(s string) int {
	instructions := format(s)
	current := 50
	result := 0
	for _, instr := range instructions {
		current = ((current + instr) + 100) % 100
		if current == 0 {
			result++
		}
	}
	return result
}

func part2(s string) int {
	instructions := format(s)
	current := 50
	result := 0
	for _, instr := range instructions {
		prev := current
		current += instr
		result += abs(current / 100)
		if current*prev < 0 || current == 0 {
			result++
		}
		current = ((current % 100) + 100) % 100
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day01/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day01/input.txt")

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
