package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func format(s string) []int {
	tab_s := strings.Split(s, "\n")
	tab := []int{}
	for _, x := range tab_s {
		n, _ := strconv.Atoi(x)
		tab = append(tab, n)
	}
	return tab
}

func next_5(b [5]bool) [5]bool {
	i := 0
	for b[i] {
		i++
	}
	b[i] = true
	j := i - 1
	for j > -1 {
		b[j] = false
		j--
	}
	return b
}

func next_20(b [20]bool) [20]bool {
	i := 0
	for b[i] {
		i++
	}
	b[i] = true
	j := i - 1
	for j > -1 {
		b[j] = false
		j--
	}
	return b
}

func cmpt_5(b [5]bool) int {
	i := 0
	res := 0
	for i < len(b) {
		if b[i] {
			res++
		}
		i++
	}
	return res
}
func cmpt_20(b [20]bool) int {
	i := 0
	res := 0
	for i < len(b) {
		if b[i] {
			res++
		}
		i++
	}
	return res
}

func min(tab []int) int {
	res := tab[0]
	for _, x := range tab {
		if x < res {
			res = x
		}
	}
	return res
}
func part1(s string) int {
	tab := format(s)
	res := 0
	n := len(tab)
	if n == 5 {
		b := [5]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [5]bool{true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 5 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 25 {
				res++
			}
			b = next_5(b)
		}
	}
	if n == 20 {
		b := [20]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [20]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 20 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 150 {
				res++
			}
			b = next_20(b)
		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	res := 0
	n := len(tab)
	cpt := []int{}
	if n == 5 {
		b := [5]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [5]bool{true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 5 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 25 {
				cpt = append(cpt, cmpt_5(b))
				res++
			}
			b = next_5(b)
		}
	}
	if n == 20 {
		b := [20]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [20]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 20 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 150 {
				cpt = append(cpt, cmpt_20(b))
				res++
			}
			b = next_20(b)
		}
	}
	res2 := 0
	if n == 5 {
		b := [5]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [5]bool{true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 5 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 25 && cmpt_5(b) == min(cpt) {
				res2++
			}
			b = next_5(b)
		}
	}
	if n == 20 {
		b := [20]bool{}
		for i := range b {
			b[i] = false
		}
		for b != [20]bool{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true} {
			tmp := 0
			i := 0
			for i < 20 {
				if b[i] {
					tmp += tab[i]
				}
				i++
			}
			if tmp == 150 && cmpt_20(b) == min(cpt) {
				res2++
			}
			b = next_20(b)
		}
	}
	return res2
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
