package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func mul(s string) int {
	s = s[4 : len(s)-1]
	parts := strings.Split(s, ",")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a * b
}

func part1(s string) int {
	var result int
	var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, str := range re.FindAllString(s, -1) {
		result += mul(str)
	}
	return result
}

func part2(s string) int {
	var result int
	var enabled = 1
	var re = regexp.MustCompile(`(mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`)
	for _, str := range re.FindAllString(s, -1) {
		if str == "do()" {
			enabled = 1
		} else if str == "don't()" {
			enabled = 0
		} else {
			result += mul(str) * enabled
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day03/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day03/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day03/input.txt")

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
