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
	l := format(s)
	for _, n := range l {
		c += (n / 3) - 2
	}
	return c
}

func part2(s string) int {
	c := 0
	l := format(s)
	i := 0
	for i < len(l) {
		n := l[i]
		f := (n / 3) - 2
		if f > 0 {
			l = append(l, f)
			c += f
		}
		i++
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day01/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2019/day01/input.txt")

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
