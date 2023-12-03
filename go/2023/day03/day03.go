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
		tab = append(tab, strings.Split(line, ""))
	}
	return tab
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isSym(s string) bool {
	return s == "*" || s == "/" || s == "$" || s == "#" || s == "%" || s == "-" || s == "+" || s == "=" || s == "@" || s == "&"
}

func symNear(i int, j int, tab [][]string) bool {
	res := false
	if j > 0 {
		if isSym(tab[i][j-1]) {
			res = true
		}
	}
	if j > 0 && i > 0 {
		if isSym(tab[i-1][j-1]) {
			res = true
		}
	}
	if i > 0 {
		if isSym(tab[i-1][j]) {
			res = true
		}
	}
	if j < len(tab[0])-1 && i > 0 {
		if isSym(tab[i-1][j+1]) {
			res = true
		}
	}
	if j < len(tab[0])-1 {
		if isSym(tab[i][j+1]) {
			res = true
		}
	}
	if j < len(tab[0])-1 && i < len(tab)-1 {
		if isSym(tab[i+1][j+1]) {
			res = true
		}
	}
	if i < len(tab)-1 {
		if isSym(tab[i+1][j]) {
			res = true
		}
	}
	if j > 0 && i < len(tab)-1 {
		if isSym(tab[i+1][j-1]) {
			res = true
		}
	}
	return res
}

func part1(s string) int {
	c := 0
	tab := format(s)
	l := len(tab[0])
	for i, line := range tab {
		j := 0
		for j < l {
			char := line[j]
			if isNumber(char) {
				sym := false
				if symNear(i, j, tab) {
					sym = true
				}
				k := j + 1
				n := char
				for k < l && isNumber(line[k]) {
					if symNear(i, k, tab) {
						sym = true
					}
					n += line[k]
					k += 1
				}
				j = k
				if sym {
					number, _ := strconv.Atoi(n)
					c += number
				}
			}
			j += 1
		}
	}
	return c
}

func isGear(s string) bool {
	return s == "*"
}

func gearNear(i int, j int, tab [][]string) (bool, int, int) {
	if j > 0 {
		if isGear(tab[i][j-1]) {
			return true, i, j - 1
		}
	}
	if j > 0 && i > 0 {
		if isGear(tab[i-1][j-1]) {
			return true, i - 1, j - 1
		}
	}
	if i > 0 {
		if isGear(tab[i-1][j]) {
			return true, i - 1, j
		}
	}
	if j < len(tab[0])-1 && i > 0 {
		if isGear(tab[i-1][j+1]) {
			return true, i - 1, j + 1
		}
	}
	if j < len(tab[0])-1 {
		if isGear(tab[i][j+1]) {
			return true, i, j + 1
		}
	}
	if j < len(tab[0])-1 && i < len(tab)-1 {
		if isGear(tab[i+1][j+1]) {
			return true, i + 1, j + 1
		}
	}
	if i < len(tab)-1 {
		if isGear(tab[i+1][j]) {
			return true, i + 1, j
		}
	}
	if j > 0 && i < len(tab)-1 {
		if isGear(tab[i+1][j-1]) {
			return true, i + 1, j - 1
		}
	}
	return false, i, j
}

func isIn(tab []int, x int) (bool, int) {
	for i, y := range tab {
		if x == y {
			return true, i
		}
	}
	return false, 0
}

func part2(s string) int {
	c := 0
	tab := format(s)
	l := len(tab[0])
	gears := []int{}
	first := []int{}
	for i, line := range tab {
		j := 0
		for j < l {
			char := line[j]
			if isNumber(char) {
				gear, a, b := gearNear(i, j, tab)
				k := j + 1
				n := char
				for k < l && isNumber(line[k]) {
					if !gear {
						gear, a, b = gearNear(i, k, tab)
					}
					n += line[k]
					k += 1
				}
				j = k
				number, _ := strconv.Atoi(n)
				if gear {
					m := a*l + b
					isIn, p := isIn(gears, m)
					if isIn {
						c += number * first[p]
					} else {
						first = append(first, number)
						gears = append(gears, m)
					}
				}
			}
			j += 1
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}

	test, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test 1 : %d\n", part1(string(test)))
	fmt.Printf("Test 2 : %d\n", part2(string(test)))

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
