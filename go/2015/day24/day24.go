package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []int {
	lines := strings.Split(s, "\n")
	tab := []int{}
	for _, line := range lines {
		n, _ := strconv.Atoi(line)
		tab = append(tab, n)
	}
	return tab
}

func multi(tab []int) int {
	res := 1
	for _, x := range tab {
		res *= x
	}
	return res
}

func sum(tab []int) int {
	res := 0
	for _, x := range tab {
		res += x
	}
	return res
}

func bin(i int) [28]bool {
	res := [28]bool{}
	j := 27
	for i > 0 {
		if i%2 == 0 {
			res[j] = false
		} else {
			res[j] = true
		}
		i /= 2
		j--
	}
	return res
}

func exp(x int, n int) int {
	if n == 1 {
		return x
	} else {
		return x * exp(x, n-1)
	}
}

func part1(s string) int {
	tab := format(s)
	i := 0
	min := 0
	len_group := 0
	l := len(tab)
	n := sum(tab)
	for i < exp(2, l) {
		group1 := []int{}
		bin := bin(i)
		//fmt.Println(i)
		for j := range bin {
			if j < l && bin[27-j] {
				group1 = append(group1, tab[j])
			}
		}
		if sum(group1) == n/3 {
			if min == 0 {
				min = multi(group1)
				len_group = len(group1)
			}
			if len(group1) < len_group {
				min = multi(group1)
				len_group = len(group1)
			}
			if len_group == len(group1) && multi(group1) < min {
				min = multi(group1)
			}
		}
		i++
	}
	return min
}

func part2(s string) int {
	tab := format(s)
	i := 0
	min := 0
	len_group := 0
	l := len(tab)
	n := sum(tab)
	for i < exp(2, l) {
		group1 := []int{}
		bin := bin(i)
		//fmt.Println(i)
		for j := range bin {
			if j < l && bin[27-j] {
				group1 = append(group1, tab[j])
			}
		}
		if sum(group1) == n/4 {
			if min == 0 {
				min = multi(group1)
				len_group = len(group1)
			}
			if len(group1) < len_group {
				min = multi(group1)
				len_group = len(group1)
			}
			if len_group == len(group1) && multi(group1) < min {
				min = multi(group1)
			}
		}
		i++
	}
	return min
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
