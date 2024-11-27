package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func isValid(i, j int, tab [][]string) bool {
	n, m := len(tab), len(tab[0])
	return i >= 0 && j >= 0 && i < n && j < m
}

func adjacent(tab [][]string, i, j int) map[string]int {
	res := make(map[string]int)
	for _, x := range []int{1, 0, -1} {
		for _, y := range []int{1, 0, -1} {
			if (x != 0 || y != 0) && isValid(i+x, j+y, tab) {
				res[tab[i+x][j+y]]++
			}
		}
	}
	return res
}

func resourceValue(tab [][]string) int {
	acres := 0
	lumber := 0
	for _, line := range tab {
		for _, s := range line {
			if s == "|" {
				acres++
			}
			if s == "#" {
				lumber++
			}
		}
	}
	return acres * lumber
}

func next(tab [][]string) [][]string {
	n := len(tab)
	m := len(tab[0])
	res := make([][]string, n)
	for i := range res {
		res[i] = make([]string, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			adj := adjacent(tab, i, j)
			switch tab[i][j] {
			case ".":
				if adj["|"] >= 3 {
					res[i][j] = "|"
				} else {
					res[i][j] = "."
				}
			case "|":
				if adj["#"] >= 3 {
					res[i][j] = "#"
				} else {
					res[i][j] = "|"
				}
			case "#":
				if adj["#"] > 0 && adj["|"] > 0 {
					res[i][j] = "#"
				} else {
					res[i][j] = "."
				}
			}
		}
	}
	return res
}

func printTab(tab [][]string) {
	for _, line := range tab {
		fmt.Println(line)
	}
	fmt.Println()
}

func hash(tab [][]string) string {
	res := ""
	for _, line := range tab {
		for _, s := range line {
			res += s
		}
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	minute := 0
	for minute < 10 {
		tab = next(tab)
		minute++
	}
	return resourceValue(tab)
}

func part2(s string) int {
	tab := format(s)
	memory := make(map[string]int)
	minute := 0
	for minute < 1000000000 {
		tab = next(tab)
		minute++
		cycle, alreadySeen := memory[hash(tab)]
		cycleLength := minute - cycle
		if alreadySeen {
			for minute+cycleLength < 1000000000 {
				minute += cycleLength
			}
		}
		memory[hash(tab)] = minute
	}
	return resourceValue(tab)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day18/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day18/input.txt")

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
