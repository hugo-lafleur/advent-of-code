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

func count(w []string) map[string]int {
	res := make(map[string]int)
	for _, l := range w {
		res[l]++
	}
	return res
}

func hasDouble(c map[string]int) bool {
	for _, v := range c {
		if v == 2 {
			return true
		}
	}
	return false
}

func hasTriple(c map[string]int) bool {
	for _, v := range c {
		if v == 3 {
			return true
		}
	}
	return false
}

func differByOne(w1, w2 []string) bool {
	c := 0
	l := len(w1)
	for i := 0; i < l; i++ {
		if w1[i] != w2[i] {
			c++
		}
		if c > 1 {
			return false
		}
	}
	return c == 1
}

func removeDifference(w1, w2 []string) string {
	res := ""
	l := len(w1)
	for i := 0; i < l; i++ {
		if w1[i] != w2[i] {
			continue
		}
		res += w1[i]
	}
	return res
}

func part1(s string) int {
	words := format(s)
	d := 0
	t := 0
	for _, word := range words {
		c := count(word)
		if hasDouble(c) {
			d++
		}
		if hasTriple(c) {
			t++
		}
	}
	return d * t
}

func part2(s string) string {
	words := format(s)
	for _, word1 := range words {
		for _, word2 := range words {
			if differByOne(word1, word2) {
				return removeDifference(word1, word2)
			}
		}
	}
	return ""
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day02/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day02/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day02/input.data")

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
