package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func format(s string) []string {
	return strings.Split(s, "")
}

func is_in(tab []string, x string) bool {
	for _, y := range tab {
		if x == y {
			return true
		}
	}
	return false
}

func no_different(tab []string) bool {
	i := 0
	b := true
	for i < len(tab) {
		//fmt.Println(tab[:i], tab[i+1:], tab[i])
		if is_in(tab[:i], tab[i]) || is_in(tab[i+1:], tab[i]) {
			b = false
		}
		i++
	}
	return b
}

func part1(s string) int {
	tab := format(s)
	i := 3
	for i < len(s) {
		a := tab[i]
		b := tab[i-1]
		c := tab[i-2]
		d := tab[i-3]
		if a != b && a != c && a != d && b != c && b != d && c != d {
			return i + 1
		}
		i++
	}
	return 0
}

func part2(s string) int {
	tab := format(s)
	i := 14
	for i < len(s) {
		if no_different(tab[i-14 : i]) {
			return i
		}
		i++
	}
	return 0
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
