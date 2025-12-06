package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format1(s string) ([][]int, []string) {
	lines := strings.Split(s, "\n")
	tab := [][]int{}
	for i := range len(lines) - 1 {
		parts := strings.Fields(lines[i])
		lineInt := []int{}
		for _, str := range parts {
			n, _ := strconv.Atoi(str)
			lineInt = append(lineInt, n)
		}
		tab = append(tab, lineInt)
	}
	return tab, strings.Fields(lines[len(lines)-1])
}

func format2(s string) ([][]int, []string) {
	lines := strings.Split(s, "\n")
	n := len(lines)
	symbols := []string{}
	tab := [][]int{}
	lineTab := []int{}
	for i, char := range lines[n-1] {
		x := 0
		for j := range len(lines) - 1 {
			if lines[j][i] == ' ' {
				continue
			}
			x *= 10
			x += int(lines[j][i] - '0')
		}
		if x != 0 {
			lineTab = append(lineTab, x)
		}
		if i == len(lines[n-1])-1 || lines[n-1][i+1] == '+' || lines[n-1][i+1] == '*' {
			tab = append(tab, lineTab)
			lineTab = []int{}
			continue
		}
		if char == '*' || char == '+' {
			symbols = append(symbols, string(char))
		}

	}
	return tab, symbols
}

func part1(s string) int {
	tab, symbols := format1(s)
	result := 0
	for i, sym := range symbols {
		problem := 0
		if sym == "*" {
			problem = 1
			for j := range tab {
				problem *= tab[j][i]
			}
			result += problem
		} else {
			for j := range tab {
				problem += tab[j][i]
			}
			result += problem
		}
	}
	return result
}

func part2(s string) int {
	tab, symbols := format2(s)
	result := 0
	for i, sym := range symbols {
		problem := 0
		if sym == "*" {
			problem = 1
			for j := range tab[i] {
				problem *= tab[i][j]
			}
			result += problem
		} else {
			for j := range tab[i] {
				problem += tab[i][j]
			}
			result += problem
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day06/input.txt")

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
