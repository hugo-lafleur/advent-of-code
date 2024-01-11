package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	res := [][]int{}
	for _, line := range lines {
		strs := strings.Split(line, " ")
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])
		c, _ := strconv.Atoi(strs[2])
		nbr := []int{}
		nbr = append(nbr, a, b, c)
		res = append(res, nbr)
	}
	return res
}

func part1(s string) int {
	c := 0
	lines := format(s)
	for _, line := range lines {
		if line[0]+line[1] > line[2] && line[2]+line[1] > line[0] && line[0]+line[2] > line[1] {
			c++
		}
	}
	return c
}

func part2(s string) int {
	res := 0
	lines := format(s)
	i := 0
	for i < len(lines)-2 {
		j := 0
		for j < 3 {
			a := lines[i][j]
			b := lines[i+1][j]
			c := lines[i+2][j]
			if a+b > c && a+c > b && b+c > a {
				res++
			}
			j++
		}
		i = i + 3
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day03/test.data")

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

	content, err = os.ReadFile("../../../inputs/2016/day03/input.data")

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
