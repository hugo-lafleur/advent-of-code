package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(strings.Replace(line, " ", "", -1), ""))
	}
	return res
}

func solve1(expr []string) int {
	curr := make([]string, len(expr))
	copy(curr, expr)
	for len(curr) != 1 {
		new := []string{}
		for i := 0; i < len(curr)-2; i++ {
			a, err1 := strconv.Atoi(curr[i])
			b, err2 := strconv.Atoi(curr[i+2])
			if err1 == nil && err2 == nil {
				switch curr[i+1] {
				case "+":
					new = append(new, strconv.Itoa(a+b))
				case "*":
					new = append(new, strconv.Itoa(a*b))
				}
				new = append(new, curr[i+3:]...)
				break
			}
			if curr[i] == "(" && curr[i+2] == ")" {
				new = append(new, curr[i+1])
				new = append(new, curr[i+3:]...)
				break
			}
			new = append(new, curr[i])
		}
		curr = new
	}
	n, _ := strconv.Atoi(curr[0])
	return n
}

func hasAddition(expr []string, index int) bool {
	k := 0
	for i := index; i < len(expr); i++ {
		if expr[i] == "+" && k == 0 {
			return true
		}
		if expr[i] == "(" {
			k++
		}
		if expr[i] == ")" {
			k--
		}
	}
	k = 0
	for i := index; i >= 0; i-- {
		if expr[i] == "+" && k == 0 {
			return true
		}
		if expr[i] == "(" {
			k++
		}
		if expr[i] == ")" {
			k--
		}
	}
	return false
}

func solve2(expr []string) int {
	curr := make([]string, len(expr))
	copy(curr, expr)
	for len(curr) != 1 {
		new := []string{}
		for i := 0; i < len(curr)-2; i++ {
			a, err1 := strconv.Atoi(curr[i])
			b, err2 := strconv.Atoi(curr[i+2])
			if err1 == nil && err2 == nil && curr[i+1] == "+" {
				new = append(new, strconv.Itoa(a+b))
				new = append(new, curr[i+3:]...)
				break
			}
			if err1 == nil && err2 == nil && curr[i+1] == "*" && !hasAddition(curr, i+1) {
				new = append(new, strconv.Itoa(a*b))
				new = append(new, curr[i+3:]...)
				break
			}
			if curr[i] == "(" && curr[i+2] == ")" {
				new = append(new, curr[i+1])
				new = append(new, curr[i+3:]...)
				break
			}
			new = append(new, curr[i])
		}
		curr = new
	}
	n, _ := strconv.Atoi(curr[0])
	return n
}

func part1(s string) int {
	c := 0
	l := format(s)
	for _, expr := range l {
		c += solve1(expr)
	}
	return c
}

func part2(s string) int {
	c := 0
	l := format(s)
	for _, expr := range l {
		c += solve2(expr)
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day18/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day18/input.txt")

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
