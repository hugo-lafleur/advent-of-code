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
	lines := strings.Split(s, ",")
	res := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	list := format(s)
	m := make(map[int]int)
	var last, lastIndex int
	for i, n := range list {
		m[n] = i + 1
		last = n
		lastIndex = i + 1
	}
	for i := len(list) + 1; i <= 2020; i++ {
		n := (i - 1) - lastIndex
		lastIndex = m[n]
		if lastIndex == 0 {
			lastIndex = i
		}
		last = n
		m[n] = i
	}
	return last
}

func part2(s string) int {
	list := format(s)
	m := make(map[int]int)
	var last, lastIndex int
	for i, n := range list {
		m[n] = i + 1
		last = n
		lastIndex = i + 1
	}
	for i := len(list) + 1; i <= 30000000; i++ {
		n := (i - 1) - lastIndex
		lastIndex = m[n]
		if lastIndex == 0 {
			lastIndex = i
		}
		last = n
		m[n] = i
	}
	return last
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day15/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day15/input.txt")

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
