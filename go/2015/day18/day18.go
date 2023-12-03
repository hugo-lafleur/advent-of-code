package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]bool {
	lines := strings.Split(s, "\n")
	l := len(lines)
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.Split(x, ""))
	}
	graph := make([][]bool, l)
	for i := range graph {
		graph[i] = make([]bool, l)
	}
	for i := range graph {
		for j := range graph[i] {
			if tab[i][j] == "." {
				graph[i][j] = false
			} else {
				graph[i][j] = true
			}
		}
	}
	return graph
}

func nb_on(graph [][]bool, i int, j int) int {
	n := len(graph)
	res := 0
	k := -1
	for k < 2 {
		l := -1
		for l < 2 {
			if (i+k) > -1 && (i+k) < n && (j+l) > -1 && (j+l) < n && graph[i+k][j+l] {
				if !(k == 0 && l == 0) {
					res++
				}
			}
			l++
		}
		k++
	}
	return res
}

func next(graph [][]bool) [][]bool {
	l := len(graph)
	new := make([][]bool, l)
	for i := range new {
		new[i] = make([]bool, l)
	}
	i := 0
	for i < l {
		j := 0
		for j < l {
			n := nb_on(graph, i, j)
			if graph[i][j] {
				if n == 2 || n == 3 {
					new[i][j] = true
				} else {
					new[i][j] = false
				}
			} else {
				if n == 3 {
					new[i][j] = true
				} else {
					new[i][j] = false
				}
			}
			j++
		}
		i++
	}
	return new
}

func next2(graph [][]bool) [][]bool {
	l := len(graph)
	new := make([][]bool, l)
	for i := range new {
		new[i] = make([]bool, l)
	}
	i := 0
	for i < l {
		j := 0
		for j < l {
			n := nb_on(graph, i, j)
			if graph[i][j] {
				if n == 2 || n == 3 {
					new[i][j] = true
				} else {
					new[i][j] = false
				}
			} else {
				if n == 3 {
					new[i][j] = true
				} else {
					new[i][j] = false
				}
			}
			j++
		}
		i++
	}
	new[0][0] = true
	new[0][l-1] = true
	new[l-1][0] = true
	new[l-1][l-1] = true
	return new
}

func cpt_on(graph [][]bool) int {
	l := len(graph)
	res := 0
	i := 0
	for i < l {
		j := 0
		for j < l {
			if graph[i][j] {
				res++
			}
			j++
		}
		i++
	}
	return res
}

func part1(s string) int {
	graph := format(s)
	l := len(graph)
	if l == 6 {
		i := 0
		for i < 4 {
			graph = next(graph)
			i++
		}
	}
	if l == 100 {
		i := 0
		for i < 100 {
			graph = next(graph)
			i++
		}
	}
	return cpt_on(graph)
}

func part2(s string) int {
	graph := format(s)
	l := len(graph)
	graph[0][0] = true
	graph[0][l-1] = true
	graph[l-1][0] = true
	graph[l-1][l-1] = true
	if l == 6 {
		i := 0
		for i < 5 {
			graph = next2(graph)
			i++
		}
	}
	if l == 100 {
		i := 0
		for i < 100 {
			graph = next2(graph)
			i++
		}
	}
	return cpt_on(graph)
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
