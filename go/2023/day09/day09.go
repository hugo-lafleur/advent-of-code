package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "\n")
}

func isZero(l []int) bool {
	for _, x := range l {
		if x != 0 {
			return false
		}
	}
	return true
}

func history(numbers []int) int {
	h := 0
	list := [][]int{}
	list = append(list, numbers)
	i := 0
	for !(isZero(list[i])) {
		new := []int{}
		curr := list[i]
		j := 0
		for j < len(curr)-1 {
			new = append(new, curr[j+1]-curr[j])
			j++
		}
		list = append(list, new)
		i++
	}
	i = 0
	for i < len(list) {
		h += list[i][len(list[i])-1]
		i++
	}
	return h
}

func part1(s string) int {
	c := 0
	lines := format(s)
	for _, line := range lines {
		chars := strings.Split(line, " ")
		numbers := []int{}
		for _, x := range chars {
			n, _ := strconv.Atoi(x)
			numbers = append(numbers, n)
		}
		c += history(numbers)
	}
	return c
}

func powersMinusOne(n int) int {
	if n%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func history2(numbers []int) int {
	h := 0
	list := [][]int{}
	list = append(list, numbers)
	i := 0
	for !(isZero(list[i])) {
		new := []int{}
		curr := list[i]
		j := 0
		for j < len(curr)-1 {
			new = append(new, curr[j+1]-curr[j])
			j++
		}
		list = append(list, new)
		i++
	}
	i = 0
	for i < len(list) {
		h += list[i][0] * powersMinusOne(i)
		i++
	}
	return h
}

func part2(s string) int {
	c := 0
	lines := format(s)
	for _, line := range lines {
		chars := strings.Split(line, " ")
		numbers := []int{}
		for _, x := range chars {
			n, _ := strconv.Atoi(x)
			numbers = append(numbers, n)
		}
		c += history2(numbers)
	}
	return c
}

func main() {
	content, err := os.ReadFile("test.data")

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
