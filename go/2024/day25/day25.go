package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func parse(s string) ([][5]int, [][5]int) {
	var locks, keys [][5]int
	var parts = strings.Split(s, "\n\n")
	for _, p := range parts {
		var columns [5]int
		for i := range columns {
			columns[i] = -1
		}
		var lines = strings.Split(p, "\n")
		for _, line := range lines {
			for j := range line {
				if line[j] == '#' {
					columns[j]++
				}
			}
		}
		if p[0] == '#' {
			locks = append(locks, columns)
		} else {
			keys = append(keys, columns)
		}
	}
	return locks, keys
}

func fit(lock, key [5]int) bool {
	for i := range 5 {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func part1(s string) int {
	var locks, keys = parse(s)
	var result int
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				result++
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day25/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
