package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	l := len(lines)
	graph := make([][]int, l)
	for i := range graph {
		graph[i] = make([]int, 3)
	}
	for i, line := range tab {
		a, _ := strconv.Atoi(line[3])
		b, _ := strconv.Atoi(line[6])
		c, _ := strconv.Atoi(line[13])
		graph[i][0] = a
		graph[i][1] = b
		graph[i][2] = c
	}
	return graph
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
	cpt := []int{}
	res := []int{}
	b := []bool{}
	for _, x := range graph {
		cpt = append(cpt, x[1])
		res = append(res, 0)
		b = append(b, true)
	}
	n := 2503
	i := 0
	for i < n {
		for j, x := range graph {
			if cpt[j] > 0 {
				res[j] = res[j] + x[0]
				cpt[j]--
			}
			if cpt[j] < 0 {
				cpt[j]++
			}
			if cpt[j] == 0 {
				if b[j] {
					cpt[j] = -x[2]
					b[j] = false
				} else {
					cpt[j] = x[1]
					b[j] = true
				}
			}
		}
		i++
	}
	return max(res)
}

func i_max(tab []int) []int {
	m := max(tab)
	res := []int{}
	for i, x := range tab {
		if x == m {
			res = append(res, i)
		}
	}
	return res
}

func part2(s string) int {
	graph := format(s)
	cpt := []int{}
	res := []int{}
	b := []bool{}
	points := []int{}
	for _, x := range graph {
		cpt = append(cpt, x[1])
		res = append(res, 0)
		b = append(b, true)
		points = append(points, 0)
	}
	n := 2503
	i := 0
	for i < n {
		for j, x := range graph {
			if cpt[j] > 0 {
				res[j] = res[j] + x[0]
				cpt[j]--
			}
			if cpt[j] < 0 {
				cpt[j]++
			}
			if cpt[j] == 0 {
				if b[j] {
					cpt[j] = -x[2]
					b[j] = false
				} else {
					cpt[j] = x[1]
					b[j] = true
				}
			}
		}
		for _, x := range i_max(res) {
			points[x]++
		}
		i++
	}
	return max(points)
}

func main() {
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
