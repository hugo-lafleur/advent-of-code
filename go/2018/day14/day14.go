package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func appear(l []int, s string) bool {
	if len(l) < len(s) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if strconv.Itoa(l[len(l)-1-i]) != string(s[len(s)-1-i]) {
			return false
		}
	}
	return true
}

func part1(s string) string {
	c := ""
	m := format(s)
	scoreboard := []int{3, 7}
	curr1 := 0
	curr2 := 1
	for len(scoreboard) < m+10 {
		r1 := scoreboard[curr1]
		r2 := scoreboard[curr2]
		newR := r1 + r2
		if newR > 9 {
			scoreboard = append(scoreboard, newR/10, newR%10)
		} else {
			scoreboard = append(scoreboard, newR)
		}
		curr1 = (1 + r1 + curr1) % len(scoreboard)
		curr2 = (1 + r2 + curr2) % len(scoreboard)
	}
	for i := m; i < m+10; i++ {
		c += strconv.Itoa(scoreboard[i])
	}
	return c
}

func part2(s string) int {
	c := 2
	scoreboard := []int{3, 7}
	curr1 := 0
	curr2 := 1
	for {
		r1 := scoreboard[curr1]
		r2 := scoreboard[curr2]
		newR := r1 + r2
		if newR > 9 {
			scoreboard = append(scoreboard, newR/10, newR%10)
			c = c + 2
		} else {
			scoreboard = append(scoreboard, newR)
			c++
		}
		//fmt.Println(scoreboard, c)
		curr1 = (1 + r1 + curr1) % len(scoreboard)
		curr2 = (1 + r2 + curr2) % len(scoreboard)
		if appear(scoreboard, s) {
			return c - len(s)
		}
		if appear(scoreboard[:len(scoreboard)-1], s) {
			return c - 1 - len(s)
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day14/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day14/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day14/input.data")

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
