package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func is_in(s string, tab []string) bool {
	for _, x := range tab {
		if s == x {
			return true
		}
	}
	return false
}

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	cities := []string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	for _, line := range tab {
		if !(is_in(line[0], cities)) {
			cities = append(cities, line[0])
		}
	}
	cities = append(cities, tab[len(tab)-1][2])
	l := len(cities)
	graph := make([][]int, l)
	for i := range graph {
		graph[i] = make([]int, l)
	}
	for _, line := range tab {
		for i, src := range cities {
			for j, dest := range cities {
				if src == line[0] && dest == line[2] {
					n, _ := strconv.Atoi(line[4])
					graph[i][j] = n
					graph[j][i] = n
				}
			}
		}
	}
	return graph
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func min(tab []int) int {
	m := tab[0]
	for _, x := range tab {
		if x < m {
			m = x
		}
	}
	return m
}

func max(tab []int) int {
	m := tab[0]
	for _, x := range tab {
		if x > m {
			m = x
		}
	}
	return m
}

func part1(s string) int {
	graph := format(s)
	res := []int{}
	l := []int{}
	for i := range graph {
		l = append(l, i)
	}
	p := permutations(l)
	n := 0
	for _, x := range p {
		n = 0
		for i := 0; i < len(x)-1; i++ {
			n = n + graph[x[i]][x[i+1]]
		}
		res = append(res, n)
	}
	return min(res)
}

func part2(s string) int {
	graph := format(s)
	res := []int{}
	l := []int{}
	for i := range graph {
		l = append(l, i)
	}
	p := permutations(l)
	n := 0
	for _, x := range p {
		n = 0
		for i := 0; i < len(x)-1; i++ {
			n = n + graph[x[i]][x[i+1]]
		}
		res = append(res, n)
	}
	return max(res)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day09/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2015/day09/input.txt")

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
