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
	names := []string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	for _, line := range tab {
		if !(is_in(line[0], names)) {
			names = append(names, line[0])
		}
	}
	l := len(names)
	graph := make([][]int, l)
	for i := range graph {
		graph[i] = make([]int, l)
	}
	for _, line := range tab {
		for i, scr := range names {
			for j, dest := range names {
				if line[0] == scr && line[10][:len(line[10])-1] == dest {
					if line[2] == "gain" {
						n, _ := strconv.Atoi(line[3])
						graph[i][j] = n
					} else {
						n, _ := strconv.Atoi(line[3])
						graph[i][j] = -n
					}
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
	list := []int{}
	for i := range graph {
		list = append(list, i)
	}
	p := permutations(list)
	res := []int{}
	for _, x := range p {
		l := len(x)
		i := 1
		tmp := 0
		for i < l-1 {
			tmp = tmp + graph[x[i]][x[i+1]] + graph[x[i]][x[i-1]]
			i++
		}
		tmp = tmp + graph[x[0]][x[1]] + graph[x[0]][x[l-1]]
		tmp = tmp + graph[x[l-1]][x[0]] + graph[x[l-1]][x[l-2]]
		res = append(res, tmp)
	}
	return max(res)
}

func part2(s string) int {
	graph := format(s)
	for i, x := range graph {
		graph[i] = append(x, 0)
	}
	me := []int{}
	for range graph {
		me = append(me, 0)
	}
	me = append(me, 0)
	graph = append(graph, me)
	list := []int{}
	for i := range graph {
		list = append(list, i)
	}
	p := permutations(list)
	res := []int{}
	for _, x := range p {
		l := len(x)
		i := 1
		tmp := 0
		for i < l-1 {
			tmp = tmp + graph[x[i]][x[i+1]] + graph[x[i]][x[i-1]]
			i++
		}
		tmp = tmp + graph[x[0]][x[1]] + graph[x[0]][x[l-1]]
		tmp = tmp + graph[x[l-1]][x[0]] + graph[x[l-1]][x[l-2]]
		res = append(res, tmp)
	}
	return max(res)
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
