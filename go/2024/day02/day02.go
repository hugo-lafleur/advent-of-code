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

func parse(s string) [][]int {
	var lines = strings.Split(s, "\n")
	var result = make([][]int, len(lines))
	for i := range lines {
		list := strings.Split(lines[i], " ")
		for j := range list {
			n, _ := strconv.Atoi(list[j])
			result[i] = append(result[i], n)
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

func isUnsafe(report []int) bool {
	if !(slices.IsSortedFunc(report, func(a, b int) int { return a - b }) ||
		slices.IsSortedFunc(report, func(a, b int) int { return b - a })) {
		return true
	}
	for i := 0; i < len(report)-1; i++ {
		if abs(report[i]-report[i+1]) < 1 || abs(report[i]-report[i+1]) > 3 {
			return true
		}
	}
	return false
}

func removeIndex(s []int, i int) []int {
	var result []int
	result = append(result, s[:i]...)
	result = append(result, s[i+1:]...)
	return result
}

func part1(s string) int {
	var reports = parse(s)
	var unsafeCount int
	for _, r := range reports {
		if isUnsafe(r) {
			unsafeCount++
		}
	}
	return len(reports) - unsafeCount
}

func part2(s string) int {
	var reports = parse(s)
	var unsafeCount int
	for _, r := range reports {
		var unsafe = isUnsafe(r)
		for i := 0; i < len(r); i++ {
			unsafe = unsafe && isUnsafe(removeIndex(r, i))
		}
		if unsafe {
			unsafeCount++
		}
	}
	return len(reports) - unsafeCount
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day02/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day02/input.txt")

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
