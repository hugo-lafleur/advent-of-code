package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func format(s string) []string {
	return strings.Split(s, "\n")
}

func calibrationValue(line string) int {
	values := []int{}
	for _, char := range line {
		n := (int(char)) - int('0')
		if n > -1 && n < 10 {
			values = append(values, n)
		}
	}
	i := values[0]
	j := values[len(values)-1]
	a := strconv.Itoa(i)
	b := strconv.Itoa(j)
	ab := a + b
	r, _ := strconv.Atoi(ab)
	return r
}

func part1(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		c += calibrationValue(line)
	}
	return c
}

func calibrationValue2(s string) int {
	l := len(s)
	values := []int{}
	for i, char := range s {
		n := int(char) - int('0')
		if n > -1 && n < 10 {
			values = append(values, n)
		}
		if i < l-2 && char == 'o' && s[i+1] == 'n' && s[i+2] == 'e' {
			values = append(values, 1)
		}
		if i < l-2 && char == 't' && s[i+1] == 'w' && s[i+2] == 'o' {
			values = append(values, 2)
		}
		if i < l-4 && char == 't' && s[i+1] == 'h' && s[i+2] == 'r' && s[i+3] == 'e' && s[i+4] == 'e' {
			values = append(values, 3)
		}
		if i < l-3 && char == 'f' && s[i+1] == 'o' && s[i+2] == 'u' && s[i+3] == 'r' {
			values = append(values, 4)
		}
		if i < l-3 && char == 'f' && s[i+1] == 'i' && s[i+2] == 'v' && s[i+3] == 'e' {
			values = append(values, 5)
		}
		if i < l-2 && char == 's' && s[i+1] == 'i' && s[i+2] == 'x' {
			values = append(values, 6)
		}
		if i < l-4 && char == 's' && s[i+1] == 'e' && s[i+2] == 'v' && s[i+3] == 'e' && s[i+4] == 'n' {
			values = append(values, 7)
		}
		if i < l-4 && char == 'e' && s[i+1] == 'i' && s[i+2] == 'g' && s[i+3] == 'h' && s[i+4] == 't' {
			values = append(values, 8)
		}
		if i < l-3 && char == 'n' && s[i+1] == 'i' && s[i+2] == 'n' && s[i+3] == 'e' {
			values = append(values, 9)
		}
	}
	i := values[0]
	j := values[len(values)-1]
	a := strconv.Itoa(i)
	b := strconv.Itoa(j)
	ab := a + b
	r, _ := strconv.Atoi(ab)
	return r
}

func part2(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		c += calibrationValue2(line)
	}
	return c
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}

	test1, err := os.ReadFile("test1.data")

	if err != nil {
		log.Fatal(err)
	}

	test2, err := os.ReadFile("test2.data")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test 1 : %d\n", part1(string(test1)))
	fmt.Printf("Test 2 : %d\n", part2(string(test2)))

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
