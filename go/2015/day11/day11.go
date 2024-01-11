package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func format(s string) []int {
	res := []int{}
	for _, c := range s {
		res = append(res, int(c))
	}
	return res
}

func straight(tab []int, i int) bool {
	if tab[i+1] == tab[i]+1 && tab[i+2] == tab[i]+2 {
		return true
	}
	return false
}

func three(tab []int) bool {
	n := len(tab) - 2
	for i := 0; i < n; i++ {
		if straight(tab, i) {
			return true
		}
	}
	return false
}

func is_in(tab []int, n int) bool {
	for _, x := range tab {
		if x == n {
			return true
		}
	}
	return false
}

func forbidden(tab []int) bool {
	for _, x := range tab {
		if x == 105 || x == 111 || x == 108 {
			return false
		}
	}
	return true
}

func pairs(tab []int) bool {
	i := 0
	c := 0
	done := []int{}
	for i < len(tab) {
		if i == len(tab)-1 {
			return c > 1
		}
		if tab[i] == tab[i+1] && !is_in(done, tab[i]) {
			c++
			done = append(done, tab[i])

		}
		i++
	}
	return c > 1
}

func next(tab []int) []int {
	i := len(tab) - 1
	for tab[i] == 122 {
		tab[i] = 97
		i--
	}
	tab[i]++
	return tab
}

func part1(s string) []string {
	tab := format(s)
	for !(three(tab) && forbidden(tab) && pairs(tab)) {
		//fmt.Println(tab)
		tab = next(tab)
	}
	res := []string{}
	for _, x := range tab {
		res = append(res, string(rune(x)))
	}
	return res
}

func part2(s string) []string {
	tab := format(s)
	for !(three(tab) && forbidden(tab) && pairs(tab)) {
		//fmt.Println(tab)
		tab = next(tab)
	}
	tab = next(tab)
	for !(three(tab) && forbidden(tab) && pairs(tab)) {
		//fmt.Println(tab)
		tab = next(tab)
	}
	res := []string{}
	for _, x := range tab {
		res = append(res, string(rune(x)))
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day11/test.data")

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

	content, err = os.ReadFile("../../../inputs/2015/day11/input.data")

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
