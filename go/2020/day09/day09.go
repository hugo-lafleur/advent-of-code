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

func isSum(n int, l []int) bool {
	for i := 0; i < len(l); i++ {
		for j := i + 1; j < len(l); j++ {
			if l[i]+l[j] == n {
				return true
			}
		}
	}
	return false
}

func sum(l []int) int {
	res := 0
	for _, x := range l {
		res += x
	}
	return res
}

func min(l []int) int {
	min := l[0]
	for _, x := range l {
		if x < min {
			min = x
		}
	}
	return min
}

func max(l []int) int {
	max := l[0]
	for _, x := range l {
		if x > max {
			max = x
		}
	}
	return max
}

func part1(s string) int {
	list := format(s)
	var preambleSize int
	if len(list) == 20 {
		preambleSize = 5
	} else {
		preambleSize = 25
	}
	for i := preambleSize; i < len(list); i++ {
		if !isSum(list[i], list[i-preambleSize:i]) {
			return list[i]
		}
	}
	return 0
}

func part2(s string) int {
	list := format(s)
	var preambleSize int
	var end int
	if len(list) == 20 {
		preambleSize = 5
	} else {
		preambleSize = 25
	}
	for i := preambleSize; i < len(list); i++ {
		if !isSum(list[i], list[i-preambleSize:i]) {
			end = i
			break
		}
	}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			l := list[i : j+1]
			if sum(l) == list[end] {
				return min(l) + max(l)
			}
			if sum(l) > list[end] {
				break
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day09/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day09/input.data")

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
