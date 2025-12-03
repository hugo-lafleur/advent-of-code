package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	result := make([][]int, len(lines))
	for i := range lines {
		result[i] = []int{}
		for j := range lines[i] {
			result[i] = append(result[i], int(lines[i][j]-'0'))
		}
	}
	return result
}

func part1(s string) int {
	result := 0
	lines := format(s)
	for _, line := range lines {
		mx := 0
		largestJoltage := 0
		for i := range len(line) - 1 {
			largestJoltage = max(largestJoltage, 10*mx+line[i])
			mx = max(mx, line[i])
		}
		largestJoltage = max(largestJoltage, 10*mx+line[len(line)-1])
		result += largestJoltage
	}
	return result
}

func part2(s string) int {
	result := 0
	lines := format(s)
	for _, line := range lines {
		maxJoltage := make([]int, 12)
		for i := range line {
			for j := max(0, 12-(len(line)-i)); j < 12; j++ {
				if line[i] > maxJoltage[j] {
					maxJoltage[j] = line[i]
					for k := j + 1; k < 12; k++ {
						maxJoltage[k] = 0
					}
					break
				}
			}
		}
		joltage := 0
		for _, n := range maxJoltage {
			joltage *= 10
			joltage += n
		}
		result += joltage
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day03/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day03/input.txt")

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
