package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type rangeIP struct {
	start, end int
}

func format(s string) []rangeIP {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, "-"))
	}
	ranges := []rangeIP{}
	for _, line := range res {
		a, _ := strconv.Atoi(line[0])
		b, _ := strconv.Atoi(line[1])
		ranges = append(ranges, rangeIP{a, b})
	}
	return ranges
}

func inRange(index int, r rangeIP) bool {
	return index >= r.start && index <= r.end
}

func part1(s string) int {
	c := 0
	ranges := format(s)
	for {
		blocked := false
		for _, rangeIP := range ranges {
			if inRange(c, rangeIP) {
				blocked = true
				c = rangeIP.end
				break
			}
		}
		if !blocked {
			return c
		}
		c++
	}
}

func part2(s string) int {
	c := 0
	ranges := format(s)
	i := 0
	for i < 4294967295+1 {
		blocked := false
		for _, rangeIP := range ranges {
			if inRange(i, rangeIP) {
				blocked = true
				i = rangeIP.end
				break
			}
		}
		if !blocked {
			c++
		}
		i++
	}
	return c
}

func main() {
	content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("input.data")

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
