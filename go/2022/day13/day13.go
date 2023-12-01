package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func format(s string) []string {
	lines := strings.Split(s, "\n")
	return lines
}

func test(a any, b any) int {
	atab, aok := a.([]any)
	btab, bok := b.([]any)
	if !aok && !bok {
		return int(a.(float64) - b.(float64))
	}
	if !aok {
		atab = []any{a}
	}
	if !bok {
		btab = []any{b}
	}
	i := 0
	for i < len(btab) && i < len(atab) {
		c := test(atab[i], btab[i])
		if c != 0 {
			return c
		}
		i++
	}
	return len(atab) - len(btab)
}

func part1(s string) int {
	tab := format(s)
	res := 0
	i := 0
	for i < len(tab)-2 {
		var a, b any
		json.Unmarshal([]byte(tab[i]), &a)
		json.Unmarshal([]byte(tab[i+1]), &b)
		if test(a, b) < 0 {
			res += i/3 + 1
		}
		i = i + 3
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	list := []any{}
	var a, b any
	res := 1
	i := 0
	for i < len(tab) {
		json.Unmarshal([]byte(tab[i]), &a)
		json.Unmarshal([]byte(tab[i+1]), &b)
		list = append(list, a)
		list = append(list, b)
		i = i + 3
	}
	str := "[[2]]"
	str2 := "[[6]]"
	json.Unmarshal([]byte(str), &a)
	json.Unmarshal([]byte(str2), &b)
	list = append(list, a)
	list = append(list, b)
	sort.Slice(list, func(i, j int) bool { return test(list[i], list[j]) < 0 })
	for i, x := range list {
		if test(a, x) == 0 || test(x, b) == 0 {
			res *= (i + 1)
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
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	start = time.Now()
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
