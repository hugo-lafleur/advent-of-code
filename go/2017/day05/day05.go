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
	res := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	c := 0
	offsets := format(s)
	i := 0
	for i >= 0 && i < len(offsets) {
		offsets[i]++
		i += offsets[i] - 1
		c++
	}
	return c
}

func part2(s string) int {
	c := 0
	offsets := format(s)
	i := 0
	for i >= 0 && i < len(offsets) {
		j := offsets[i]
		if j >= 3 {
			offsets[i]--
		} else {
			offsets[i]++
		}
		i += j
		c++

	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day05/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2017/day05/input.txt")

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
