package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	for i := range tab {
		for j := range tab[i] {
			if tab[i][j] == "a," {
				tab[i][j] = "a"
			}
			if tab[i][j] == "b," {
				tab[i][j] = "b"
			}
		}
	}
	return tab
}

func part1(s string) int {
	instr := format(s)
	l := len(instr)
	i := 0
	reg := make(map[string]int)
	reg["a"] = 0
	reg["b"] = 0
	for i < l {
		line := instr[i]
		if line[0] == "hlf" {
			reg[line[1]] /= 2
		}
		if line[0] == "tpl" {
			reg[line[1]] *= 3
		}
		if line[0] == "inc" {
			reg[line[1]] += 1
		}
		if line[0] == "jio" {
			if reg[line[1]] == 1 {
				n, _ := strconv.Atoi(line[2])
				i = i + n - 1
			}
		}
		if line[0] == "jie" {
			if reg[line[1]]%2 == 0 {
				n, _ := strconv.Atoi(line[2])
				i = i + n - 1
			}
		}
		if line[0] == "jmp" {
			n, _ := strconv.Atoi(line[1])
			i = i + n - 1
		}
		i++
	}
	return reg["b"]
}

func part2(s string) int {
	instr := format(s)
	l := len(instr)
	i := 0
	reg := make(map[string]int)
	reg["a"] = 1
	reg["b"] = 0
	for i < l {
		line := instr[i]
		if line[0] == "hlf" {
			reg[line[1]] /= 2
		}
		if line[0] == "tpl" {
			reg[line[1]] *= 3
		}
		if line[0] == "inc" {
			reg[line[1]] += 1
		}
		if line[0] == "jio" {
			if reg[line[1]] == 1 {
				n, _ := strconv.Atoi(line[2])
				i = i + n - 1
			}
		}
		if line[0] == "jie" {
			if reg[line[1]]%2 == 0 {
				n, _ := strconv.Atoi(line[2])
				i = i + n - 1
			}
		}
		if line[0] == "jmp" {
			n, _ := strconv.Atoi(line[1])
			i = i + n - 1
		}
		i++
	}
	return reg["b"]
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
