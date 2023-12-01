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
	str := [][]string{}
	for _, line := range lines {
		str = append(str, strings.Split(line, ""))
	}
	l := len(str)
	tab := make([][]int, l)
	for i := range str {
		tab[i] = make([]int, l)
		for j := range str[i] {
			n, _ := strconv.Atoi(str[i][j])
			tab[i][j] = n
		}
	}
	return tab
}

func visible_top(tab [][]int, i int, j int) bool {
	n := tab[i][j]
	x, y := i, j
	for x != 0 {
		x--
		if tab[x][y] >= n {
			return false
		}
	}
	return true
}

func visible_bottom(tab [][]int, i int, j int) bool {
	n := tab[i][j]
	x, y := i, j
	for x != len(tab)-1 {
		x++
		if tab[x][y] >= n {
			return false
		}
	}
	return true
}

func visible_right(tab [][]int, i int, j int) bool {
	n := tab[i][j]
	x, y := i, j
	for y != len(tab)-1 {
		y++
		if tab[x][y] >= n {
			return false
		}
	}
	return true
}

func visible_left(tab [][]int, i int, j int) bool {
	n := tab[i][j]
	x, y := i, j
	for y != 0 {
		y--
		if tab[x][y] >= n {
			return false
		}
	}
	return true
}

func score_top(tab [][]int, i int, j int) int {
	n := tab[i][j]
	x, y := i, j
	res := 1
	for x != 0 {
		x--
		if tab[x][y] >= n {
			return res
		}
		res++
	}
	return res - 1
}

func score_bottom(tab [][]int, i int, j int) int {
	n := tab[i][j]
	x, y := i, j
	res := 1
	for x != len(tab)-1 {
		x++
		if tab[x][y] >= n {
			return res
		}
		res++
	}
	return res - 1
}

func score_right(tab [][]int, i int, j int) int {
	n := tab[i][j]
	x, y := i, j
	res := 1
	for y != len(tab)-1 {
		y++
		if tab[x][y] >= n {
			return res
		}
		res++
	}
	return res - 1
}

func score_left(tab [][]int, i int, j int) int {
	n := tab[i][j]
	x, y := i, j
	res := 1
	for y != 0 {
		y--
		if tab[x][y] >= n {
			return res
		}
		res++
	}
	return res - 1
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
	tab := format(s)
	res := 0
	for i := range tab {
		for j := range tab[i] {
			if visible_bottom(tab, i, j) || visible_left(tab, i, j) || visible_right(tab, i, j) || visible_top(tab, i, j) {
				res += 1
			}

		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	score := []int{}
	for i := range tab {
		for j := range tab[i] {
			if i > 0 && i < len(tab)-1 && j > 0 && j < len(tab[i])-1 {
				score = append(score, score_bottom(tab, i, j)*score_left(tab, i, j)*score_right(tab, i, j)*score_top(tab, i, j))
			}
		}
	}
	return max(score)
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
