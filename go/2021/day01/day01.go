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
	res := []int{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	c := 0
	list := format(s)
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for i := 2; i < len(list)-1; i++ {
		if list[i]+list[i+1]+list[i-1] > list[i-1]+list[i-2]+list[i] {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day01/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day01/input.data")

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
