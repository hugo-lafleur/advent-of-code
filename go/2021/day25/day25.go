package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	tab := [][]string{}
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		tab = append(tab, strings.Split(line, ""))
	}
	return tab
}

func mod(a int, b int) int {
	m := a % b
	if m < 0 {
		m = m + b
	}
	return m
}

func check(tab [][]string) [][]bool {
	n := len(tab)
	m := len(tab[0])
	check := make([][]bool, n)
	for i := range check {
		check[i] = make([]bool, m)
	}
	for i, line := range tab {
		for j, x := range line {
			if (x == ">" && line[mod(j+1, m)] == ".") || (x == "v" && tab[mod(i+1, n)][j] == ".") {
				check[i][j] = true
			}
		}
	}
	return check
}

func move(tab [][]string) ([][]string, bool) {
	v := false
	n := len(tab)
	m := len(tab[0])
	b := check(tab)
	for i, line := range tab {
		for j, x := range line {
			if x == ">" && b[i][j] {
				line[mod(j+1, m)] = ">"
				line[j] = "."
				v = true
			}
		}
	}
	b = check(tab)
	for i, line := range tab {
		for j, x := range line {
			if x == "v" && b[i][j] {
				tab[mod(i+1, n)][j] = "v"
				tab[i][j] = "."
				v = true
			}
		}
	}
	return tab, !v
}

func part1(s string) int {
	c := 0
	tab := format(s)
	var done bool
	for {
		c += 1
		if tab, done = move(tab); done {
			break
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day25/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2021/day25/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
