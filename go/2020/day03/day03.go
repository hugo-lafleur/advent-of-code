package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func part1(s string) int {
	c := 0
	tab := format(s)
	j := 0
	for i := 0; i < len(tab); i++ {
		if tab[i][j] == "#" {
			c++
		}
		j = (j + 3) % len(tab[0])
	}
	return c
}

func part2(s string) int {
	res := 1
	tab := format(s)
	c := 0
	j := 0
	for k := 1; k < 8; k += 2 {
		c = 0
		j = 0
		for i := 0; i < len(tab); i++ {
			if tab[i][j] == "#" {
				c++
			}
			j = (j + k) % len(tab[0])
		}
		//fmt.Println(k, c)
		res *= c
	}
	c = 0
	j = 0
	for i := 0; i < len(tab); i += 2 {
		if tab[i][j] == "#" {
			c++
		}
		j = (j + 1) % len(tab[0])
	}
	fmt.Println(c)
	res *= c
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day03/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day03/input.txt")

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
