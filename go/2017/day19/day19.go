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

func tabToString(t []string) string {
	res := ""
	for _, x := range t {
		res += x
	}
	return res
}

func next(i, j int, tab [][]string, d string, path []string) (int, int, string, []string) {
	c := tab[i][j]
	if c != "|" && c != "-" && c != "+" && c != " " {
		path = append(path, c)
	}
	if c == " " {
		return -1, -1, "", path
	}
	switch d {
	case "s":
		i++
		n := tab[i][j]
		switch n {
		case "-":
			return i + 1, j, "s", path
		case "+":
			if tab[i][j+1] != " " && tab[i][j+1] != "|" {
				return i, j + 1, "e", path
			} else {
				return i, j - 1, "w", path
			}
		default:
			return i, j, "s", path
		}
	case "n":
		i--
		n := tab[i][j]
		switch n {
		case "-":
			return i - 1, j, "n", path
		case "+":
			if tab[i][j+1] != " " && tab[i][j+1] != "|" {
				return i, j + 1, "e", path
			} else {
				return i, j - 1, "w", path
			}
		default:
			return i, j, "n", path
		}
	case "w":
		j--
		n := tab[i][j]
		switch n {
		case "|":
			return i, j - 1, "w", path
		case "+":
			if (i+1) < len(tab) && tab[i+1][j] != " " && tab[i+1][j] != "-" {
				return i + 1, j, "s", path
			} else {
				return i - 1, j, "n", path
			}
		default:
			return i, j, "w", path
		}
	case "e":
		j++
		n := tab[i][j]
		switch n {
		case "|":
			return i, j + 1, "e", path
		case "+":
			if (i+1) < len(tab) && tab[i+1][j] != " " && tab[i+1][j] != "-" {
				return i + 1, j, "s", path
			} else {
				return i - 1, j, "n", path
			}
		default:
			return i, j, "e", path
		}
	}
	return 0, 0, "", path
}

func next2(i, j int, tab [][]string, d string, step int) (int, int, string, int) {
	c := tab[i][j]
	if c == " " {
		return -1, -1, "", step
	}
	switch d {
	case "s":
		i++
		n := tab[i][j]
		switch n {
		case "-":
			return i + 1, j, "s", step + 2
		case "+":
			if tab[i][j+1] != " " && tab[i][j+1] != "|" {
				return i, j + 1, "e", step + 2
			} else {
				return i, j - 1, "w", step + 2
			}
		default:
			return i, j, "s", step + 1
		}
	case "n":
		i--
		n := tab[i][j]
		switch n {
		case "-":
			return i - 1, j, "n", step + 2
		case "+":
			if tab[i][j+1] != " " && tab[i][j+1] != "|" {
				return i, j + 1, "e", step + 2
			} else {
				return i, j - 1, "w", step + 2
			}
		default:
			return i, j, "n", step + 1
		}
	case "w":
		j--
		n := tab[i][j]
		switch n {
		case "|":
			return i, j - 1, "w", step + 2
		case "+":
			if (i+1) < len(tab) && tab[i+1][j] != " " && tab[i+1][j] != "-" {
				return i + 1, j, "s", step + 2
			} else {
				return i - 1, j, "n", step + 2
			}
		default:
			return i, j, "w", step + 1
		}
	case "e":
		j++
		n := tab[i][j]
		switch n {
		case "|":
			return i, j + 1, "e", step + 2
		case "+":
			if (i+1) < len(tab) && tab[i+1][j] != " " && tab[i+1][j] != "-" {
				return i + 1, j, "s", step + 2
			} else {
				return i - 1, j, "n", step + 2
			}
		default:
			return i, j, "e", step + 1
		}
	}
	return 0, 0, "", 0
}

func part1(s string) string {
	tab := format(s)
	var x, y int
	x = 0
	d := "s"
	path := []string{}
	for j, x := range tab[0] {
		if x == "|" {
			y = j
		}
	}
	for {
		x, y, d, path = next(x, y, tab, d, path)
		if x == -1 && y == -1 {
			return tabToString(path)
		}
	}
}

func part2(s string) int {
	tab := format(s)
	var x, y int
	x = 0
	d := "s"
	steps := 0
	for j, x := range tab[0] {
		if x == "|" {
			y = j
		}
	}
	for {
		x, y, d, steps = next2(x, y, tab, d, steps)
		if x == -1 && y == -1 {
			return steps
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day19/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2017/day19/input.txt")

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
