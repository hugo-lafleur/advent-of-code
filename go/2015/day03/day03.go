package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func is_diff(tab []complex128) int {
	done := []complex128{}
	b := true
	for _, z := range tab {
		b = true
		for _, d := range done {
			if z == d {
				b = false
			}
		}
		if b {
			done = append(done, z)
		}
	}
	return len(done)
}

func part1(s string) int {
	c := 0 + 0i
	tab := []complex128{}
	tab = append(tab, c)
	for _, char := range s {
		switch char {
		case 'v':
			c = c - 1i
		case '^':
			c = c + 1i
		case '<':
			c = c - 1
		case '>':
			c = c + 1
		}
		tab = append(tab, c)
	}

	return is_diff(tab)
}

func part2(s string) int {
	c1 := 0 + 0i
	c2 := 0 + 0i
	tab := []complex128{}
	tab = append(tab, c1)
	for i, char := range s {
		if i%2 == 0 {
			switch char {
			case 'v':
				c1 = c1 - 1i
			case '^':
				c1 = c1 + 1i
			case '<':
				c1 = c1 - 1
			case '>':
				c1 = c1 + 1
			}
			tab = append(tab, c1)
		}
		if i%2 == 1 {
			switch char {
			case 'v':
				c2 = c2 - 1i
			case '^':
				c2 = c2 + 1i
			case '<':
				c2 = c2 - 1
			case '>':
				c2 = c2 + 1
			}
			tab = append(tab, c2)
		}
	}
	return is_diff(tab)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
