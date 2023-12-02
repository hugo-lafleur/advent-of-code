package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func format(s string) []string {
	return strings.Split(s, "\n")
}

func formatLine(s string) []string {
	chars := strings.Split(s, "")
	res := []string{}
	for i, char := range chars {
		if i > 7 && i < len(chars)-1 {
			_, err1 := strconv.Atoi(char)
			_, err2 := strconv.Atoi(chars[i+1])
			if err1 == nil && err2 == nil {
				res = append(res, char+chars[i+1])
				res = append(res, chars[i+3])
				i += 2
			}
			if err1 == nil && err2 != nil {
				res = append(res, char)
				res = append(res, chars[i+2])
			}
			if char == ";" {
				res = append(res, ";")
			}
		}
	}
	return res
}

func maxColor(line []string) [3]int {
	res := [3]int{}
	for i := range line {
		if line[i] == "r" {
			n, _ := strconv.Atoi(line[i-1])
			if n > res[0] {
				res[0] = n
			}
		}
		if line[i] == "g" {
			n, _ := strconv.Atoi(line[i-1])
			if n > res[1] {
				res[1] = n
			}
		}
		if line[i] == "b" {
			n, _ := strconv.Atoi(line[i-1])
			if n > res[2] {
				res[2] = n
			}
		}
	}
	return res
}

func part1(s string) int {
	c := 0
	lines := format(s)
	for i, line := range lines {
		lineFormat := formatLine(line)
		max := maxColor(lineFormat)
		if max[0] < 13 && max[1] < 14 && max[2] < 15 {
			c += i + 1
		}
	}
	return c
}

func power(max [3]int) int {
	return max[0] * max[1] * max[2]
}

func part2(s string) int {
	c := 0
	lines := format(s)
	for _, line := range lines {
		lineFormat := formatLine(line)
		max := maxColor(lineFormat)
		c += power(max)
	}
	return c
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}

	test, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test 1 : %d\n", part1(string(test)))
	fmt.Printf("Test 2 : %d\n", part2(string(test)))

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
