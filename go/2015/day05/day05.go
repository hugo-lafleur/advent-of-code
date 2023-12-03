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

func vowel(s string) bool {
	c := 0
	for _, char := range s {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			c++
		}
	}
	return c > 2
}

func dble(s string) bool {
	n := len(s)
	i := 0
	for i < (n - 1) {
		if s[i] == s[i+1] {
			return true
		}
		i++
	}
	return false
}

func no(s string) bool {
	n := len(s)
	i := 0
	for i < (n - 1) {
		if s[i] == 'a' && s[i+1] == 'b' {
			return false
		}
		if s[i] == 'c' && s[i+1] == 'd' {
			return false
		}
		if s[i] == 'p' && s[i+1] == 'q' {
			return false
		}
		if s[i] == 'x' && s[i+1] == 'y' {
			return false
		}
		i++
	}
	return true
}

func is_in(s string, tab []string) bool {
	for _, str := range tab {
		if s == str {
			return true
		}
	}
	return false
}

func twice(s string) bool {
	n := len(s)
	i := 0
	tab := []string{}
	for i < (n - 1) {
		str := string(s[i]) + string(s[i+1])
		if is_in(str, tab) {
			if s[i] == s[i+1] && s[i] == s[i-1] {
				if i > 1 && s[i-2] == s[i] {
					return true
				}

			} else {
				return true
			}
		}
		tab = append(tab, str)
		i++
	}
	return false
}

func xyx(s string) bool {
	n := len(s)
	i := 0
	for i < (n - 2) {
		if s[i] == s[i+2] {
			return true
		}
		i++
	}
	return false
}

func part1(s string) int {
	tab := format(s)
	c := 0
	for _, str := range tab {
		if vowel(str) && dble(str) && no(str) {
			c++
		}
	}
	return c
}

func part2(s string) int {
	tab := format(s)
	c := 0
	for _, str := range tab {
		if twice(str) && xyx(str) {
			c++
		}
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
