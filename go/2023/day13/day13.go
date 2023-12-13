package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "\n\n")
}

func testRows(r1 int, r2 int, input [][]string) bool {
	l := len(input[0])
	i := 0
	for r1-i > -1 && r2+i < len(input) {
		j := 0
		for j < l {
			if input[r1-i][j] != input[r2+i][j] {
				return false
			}
			j++
		}
		i++
	}
	return true
}

func testColumns(c1 int, c2 int, input [][]string) bool {
	l := len(input)
	j := 0
	for c1-j > -1 && c2+j < len(input[0]) {
		i := 0
		for i < l {
			if input[i][c1-j] != input[i][c2+j] {
				return false
			}
			i++
		}
		j++
	}
	return true
}

func part1(s string) int {
	c := 0
	tab := format(s)
	for _, input := range tab {
		lines := strings.Split(input, "\n")
		realInput := [][]string{}
		for _, line := range lines {
			realInput = append(realInput, strings.Split(line, ""))
		}
		m := len(realInput)
		n := len(realInput[0])
		//test rows
		i := 0
		for i < m-1 {
			if testRows(i, i+1, realInput) {
				c += 100 * (i + 1)
			}
			i++
		}
		//test columns
		j := 0
		for j < n-1 {
			if testColumns(j, j+1, realInput) {
				c += j + 1
			}
			j++
		}
	}
	return c
}

func testRows2(r1 int, r2 int, input [][]string) bool {
	l := len(input[0])
	i := 0
	smudge := false
	for r1-i > -1 && r2+i < len(input) {
		j := 0
		for j < l {
			if input[r1-i][j] != input[r2+i][j] {
				if smudge {
					return false
				} else {
					smudge = true
				}
			}
			j++
		}
		i++
	}
	return smudge
}

func testColumns2(c1 int, c2 int, input [][]string) bool {
	l := len(input)
	j := 0
	smudge := false
	for c1-j > -1 && c2+j < len(input[0]) {
		i := 0
		for i < l {
			if input[i][c1-j] != input[i][c2+j] {
				if smudge {
					return false
				} else {
					smudge = true
				}
			}
			i++
		}
		j++
	}
	return smudge
}

func part2(s string) int {
	c := 0
	tab := format(s)
	for _, input := range tab {
		lines := strings.Split(input, "\n")
		realInput := [][]string{}
		for _, line := range lines {
			realInput = append(realInput, strings.Split(line, ""))
		}
		m := len(realInput)
		n := len(realInput[0])
		//test rows
		i := 0
		for i < m-1 {
			if testRows2(i, i+1, realInput) {
				c += 100 * (i + 1)
			}
			i++
		}
		//test columns
		j := 0
		for j < n-1 {
			if testColumns2(j, j+1, realInput) {
				c += j + 1
			}
			j++
		}
	}
	return c
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
