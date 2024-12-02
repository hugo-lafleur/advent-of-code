package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parse(s string) ([]int, []int) {
	var left, right []int
	var lines = strings.Split(s, "\n")
	for i := range lines {
		parts := strings.Split(lines[i], " ")
		a, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		b, _ := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
		left = append(left, a)
		right = append(right, b)
	}
	return left, right
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(s string) int {
	var result int
	var left, right = parse(s)
	slices.Sort(left)
	slices.Sort(right)
	for i := range left {
		result += abs(left[i] - right[i])
	}
	return result
}

func part2(s string) int {
	var result int
	var left, right = parse(s)
	var count = make(map[int]int)
	for i := range right {
		count[right[i]]++
	}
	for i := range left {
		result += left[i] * count[left[i]]
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day01/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day01/input.txt")

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
