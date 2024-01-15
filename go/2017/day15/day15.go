package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == ' ' || r == '\n'
}

func format(s string) (int, int) {
	strs := strings.FieldsFunc(s, Split)
	res := []int{}
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err == nil {
			res = append(res, n)
		}
	}
	return res[0], res[1]
}

func part1(s string) int {
	c := 0
	l := 40000000
	a, b := format(s)
	for i := 0; i < l; i++ {
		nextA := (a * 16807) % 2147483647
		nextB := (b * 48271) % 2147483647
		if nextA&65535 == nextB&65535 {
			c++
		}
		a, b = nextA, nextB
	}
	return c
}

func part2(s string) int {
	c := 0
	l := 5000000
	a, b := format(s)
	for i := 0; i < l; i++ {
		nextA := (a * 16807) % 2147483647
		nextB := (b * 48271) % 2147483647
		for (nextA % 4) != 0 {
			nextA = (nextA * 16807) % 2147483647
		}
		for (nextB % 8) != 0 {
			nextB = (nextB * 48271) % 2147483647
		}
		if nextA&65535 == nextB&65535 {
			c++
		}
		a, b = nextA, nextB
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day15/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day15/input.data")

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
