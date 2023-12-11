package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "\n")
}

func part1(s string) int {
	c := 0
	list := format(s)
	directions := list[0]
	nextLeft := make(map[string]string)
	nextRight := make(map[string]string)
	i := 2
	for i < len(list) {
		line := list[i]
		chars := strings.Split(line, "")
		source := chars[0] + chars[1] + chars[2]
		destLeft := chars[7] + chars[8] + chars[9]
		destRight := chars[12] + chars[13] + chars[14]
		nextLeft[source] = destLeft
		nextRight[source] = destRight
		i++
	}
	current := "AAA"
	j := 0
	for current != "ZZZ" {
		if j == len(directions) {
			j = 0
		}
		if directions[j] == 'L' {
			current = nextLeft[current]
		} else {
			current = nextRight[current]
		}
		j++
		c++
	}
	return c
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, list ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(list); i++ {
		res = LCM(res, list[i])
	}
	return res
}

func part2(s string) int {
	list := format(s)
	directions := list[0]
	nextLeft := make(map[string]string)
	nextRight := make(map[string]string)
	sources := []string{}
	i := 2
	for i < len(list) {
		line := list[i]
		chars := strings.Split(line, "")
		source := chars[0] + chars[1] + chars[2]
		destLeft := chars[7] + chars[8] + chars[9]
		destRight := chars[12] + chars[13] + chars[14]
		if chars[2] == "A" {
			sources = append(sources, source)
		}
		nextLeft[source] = destLeft
		nextRight[source] = destRight
		i++
	}
	ends := []int{}
	i = 0
	for i < len(sources) {
		current := sources[i]
		j := 0
		c := 0
		for current[2] != 'Z' {
			if j == len(directions) {
				j = 0
			}
			if directions[j] == 'L' {
				current = nextLeft[current]
			} else {
				current = nextRight[current]
			}
			j++
			c++
		}
		ends = append(ends, c)
		i++
	}
	return LCM(1, 1, ends...)
}

func main() {
	content, err := os.ReadFile("test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
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
