package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	tab := [][]int{}
	add := []int{}
	for _, line := range lines {
		if line == "" {
			tab = append(tab, add)
			add = []int{}
		} else {
			n, _ := strconv.Atoi(line)
			add = append(add, n)
		}
	}
	tab = append(tab, add)
	return tab
}

func part1(s string) int {
	tab := format(s)
	sum := 0
	m := 0
	for _, elf := range tab {
		sum = 0
		for _, item := range elf {
			sum += item
		}
		if sum > m {
			m = sum
		}
	}
	return m
}

func part2(s string) int {
	tab := format(s)
	sum := 0
	res := []int{}
	for _, elf := range tab {
		sum = 0
		for _, item := range elf {
			sum += item
		}
		res = append(res, sum)
	}
	sort.Ints(res)
	l := len(res)
	return res[l-1] + res[l-2] + res[l-3]
}

func main() {
	content, err := ioutil.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))

	content, err = ioutil.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
