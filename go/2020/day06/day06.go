package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][][]string {
	res := [][][]string{}
	groups := strings.Split(s, "\n\n")
	for _, group := range groups {
		groupRes := [][]string{}
		groupSplit := strings.Split(group, "")
		i := 0
		groupRes = append(groupRes, []string{})
		for _, answer := range groupSplit {
			if answer != "\n" {
				groupRes[i] = append(groupRes[i], answer)
			} else {
				i++
				groupRes = append(groupRes, []string{})
			}
		}
		res = append(res, groupRes)
	}
	return res
}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, group := range list {
		m := make(map[string]bool)
		for _, person := range group {
			for _, answer := range person {
				m[answer] = true
			}
		}
		c += len(m)
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for _, group := range list {
		m := make(map[string]int)
		for _, person := range group {
			for _, answer := range person {
				m[answer]++
			}
		}
		for _, value := range m {
			if value == len(group) {
				c++
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day06/input.txt")

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
