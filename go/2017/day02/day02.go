package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == ' ' || r == '	'
}

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	resString := [][]string{}
	for _, line := range lines {
		resString = append(resString, strings.FieldsFunc(line, Split))
	}
	res := [][]int{}
	for _, line := range resString {
		lineInt := []int{}
		for _, x := range line {
			n, _ := strconv.Atoi(x)
			lineInt = append(lineInt, n)
		}
		res = append(res, lineInt)
	}
	return res
}

func part1(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		sort.Ints(line)
		c += line[len(line)-1] - line[0]
	}
	return c
}

func part2(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		for _, a := range line {
			for _, b := range line {
				if b > a && b%a == 0 && a != 1 {
					c += b / a
				}
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day02/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day02/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day02/input.txt")

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
