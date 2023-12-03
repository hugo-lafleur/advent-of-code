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
	return strings.Split(s, "")
}

func is_number(s string) bool {
	if s == "1" || s == "2" || s == "3" || s == "4" || s == "5" || s == "6" || s == "7" || s == "8" || s == "9" || s == "0" {
		return true
	}
	return false
}

func part1(s string) int {
	res := 0
	tab := format(s)
	l := len(tab)
	i := 0
	neg := false
	for i < l {
		c := 0
		if tab[i] == "-" {
			neg = true
			i++
		}
		for is_number(tab[i]) {
			n, _ := strconv.Atoi(tab[i])
			c = 10 * c
			c = c + n
			i++
			if i == l {
				if neg {
					return res - c
				}
				return res + c
			}

		}
		if neg {
			res = res - c
			neg = false
		} else {
			res = res + c
		}
		i++
	}
	return res
}

func find_next(tab []string, i int) int {
	j := i
	c := 1
	for c != 0 {
		if tab[j] == "}" {
			c = c - 1
		}
		if tab[j] == "{" {
			c = c + 1
		}
		j++
	}
	return j - 1
}

type Stack []int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i int) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		i := len(*s) - 1
		x := (*s)[i]
		*s = (*s)[:i]
		return x, true
	}
}

func part2(s string) int {
	var stack Stack
	res := 0
	tab := format(s)
	l := len(tab)
	i := 0
	neg := false
	tmp := 0
	for i < l {
		c := 0
		for is_number(tab[i]) {
			n, _ := strconv.Atoi(tab[i])
			c = 10 * c
			c = c + n
			i++
		}
		if neg {
			tmp = tmp - c
			neg = false
		} else {
			tmp = tmp + c
		}
		if tab[i] == "r" && tab[i+1] == "e" && tab[i+2] == "d" && tab[i-2] == ":" {
			j := find_next(tab, i)
			i = j
			tmp = 0
		}
		if tab[i] == "{" {
			stack.Push(tmp)
			tmp = 0
		}
		if tab[i] == "}" {
			stack.Push(tmp)
			x, b := stack.Pop()
			if b {
				y, b := stack.Pop()
				if b {
					tmp = x + y
				}
			}
		}
		if tab[i] == "-" {
			neg = true
		}
		i++
	}
	if tmp != 0 {
		res = res + tmp
	}
	return res
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
