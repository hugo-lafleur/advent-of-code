package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	tab := make([][]int, len(lines))
	for j, line := range lines {
		runes := []rune(line)
		res := []int{}
		for i := range runes {
			n := int(runes[i])
			if n < 96 {
				n = n - 38
			}
			if n > 96 {
				n = n - 96
			}
			res = append(res, n)
		}
		tab[j] = res
	}
	return tab
}

func is_in(tab []int, x int) bool {
	for _, y := range tab {
		if x == y {
			return true
		}
	}
	return false
}

func part1(s string) int {
	tab := format(s)
	res := 0
	for _, line := range tab {
		n := len(line) / 2
		for _, x := range line[0:n] {
			if is_in(line[n:], x) {
				res += x
				break
			}
		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	res := 0
	for i := 0; i < len(tab)-2; i = i + 3 {
		for _, x := range tab[i] {
			if is_in(tab[i+1], x) && is_in(tab[i+2], x) {
				res += x
				break
			}
		}
	}
	return res

}

func main() {
	content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
