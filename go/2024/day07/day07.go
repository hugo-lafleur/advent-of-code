package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse(s string) [][]int {
	var lines = strings.Split(s, "\n")
	var result = make([][]int, len(lines))
	for i := range lines {
		parts := strings.FieldsFunc(lines[i], func(r rune) bool { return r == ' ' || r == ':' })
		for j := range parts {
			n, _ := strconv.Atoi(parts[j])
			result[i] = append(result[i], n)
		}
	}
	return result
}

func concatenation(x, y int) int {
	return x*int(math.Pow10(int(math.Log10(float64(y)))+1)) + y
}

func testEquation(goal int, curr int, remaining []int, p2 bool) bool {
	if curr > goal || len(remaining) == 0 {
		return goal == curr
	}
	var result = testEquation(goal, curr+remaining[0], remaining[1:], p2)
	result = result || testEquation(goal, curr*remaining[0], remaining[1:], p2)
	if p2 {
		result = result || testEquation(goal, concatenation(curr, remaining[0]), remaining[1:], p2)
	}
	return result
}

func part1(s string) int {
	var list = parse(s)
	var result int
	for i := range list {
		if testEquation(list[i][0], list[i][1], list[i][2:], false) {
			result += list[i][0]
		}
	}
	return result
}

func part2(s string) int {
	var list = parse(s)
	var result int
	for i := range list {
		if testEquation(list[i][0], list[i][1], list[i][2:], true) {
			result += list[i][0]
		}
	}
	return result
}
func main() {
	content, err := os.ReadFile("../../../inputs/2024/day07/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day07/input.txt")

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
