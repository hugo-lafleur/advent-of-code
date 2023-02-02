package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func min(tab [10000000]int, n int) int {
	for j, x := range tab {
		if x > n {
			return j
		}
	}
	return 0
}

func part1(s string) int {
	n := format(s)
	i := 0
	if n == 29000000 {
		i = 665270
	}
	present := []int{}
	for len(present) == 0 || present[len(present)-1] < n {
		s := 0
		j := 1
		for j < i+1 {
			if i%j == 0 {
				s += 10 * j
			}
			j++
		}
		present = append(present, s)
		i++
	}
	return i - 1
}

func part2(s string) int {
	n := format(s)
	i := 1
	max := 0
	if n == 29000000 {
		max = 3255840
	}
	present := [10000000]int{}
	for i < max {
		j := 1
		c := 0
		for c < 50 && i*j < max {
			present[j*i] += 11 * i
			c++
			j++
		}
		i++
	}
	return min(present, n)
}

func main() {
	content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	fmt.Printf("Part 1 : %d\n", part1(string(content)))

	content, err = os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
