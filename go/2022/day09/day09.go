package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	motions := []string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	for _, line := range tab {
		n, _ := strconv.Atoi(line[1])
		i := 0
		for i < n {
			motions = append(motions, line[0])
			i++
		}
	}
	return motions
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func near(h [2]int, t [2]int) bool {
	x := h[0]
	y := h[1]
	a := t[0]
	b := t[1]
	return (abs(x-a) < 2) && (abs(y-b) < 2)
}

func cpt(done [1000][1000]bool) int {
	res := 0
	for i := range done {
		for j := range done[i] {
			if done[i][j] {
				res++
			}
		}
	}
	return res
}

func update(rope [20]int) [20]int {
	i := 2
	for i < 20 {
		a := rope[i-2]
		b := rope[i-1]
		x := rope[i]
		y := rope[i+1]
		if abs(a-x) > 1 || abs(b-y) > 1 {
			if a != x {
				if x > a {
					rope[i]--
				} else {
					rope[i]++
				}
			}
			if b != y {
				if y > b {
					rope[i+1]--
				} else {
					rope[i+1]++
				}
			}
		}
		i = i + 2
	}
	return rope
}

func part1(s string) int {
	done := [1000][1000]bool{}
	h := [2]int{500, 500}
	prev := [2]int{500, 500}
	t := [2]int{500, 500}
	motions := format(s)
	for _, x := range motions {
		prev[0] = h[0]
		prev[1] = h[1]
		switch x {
		case "R":
			h[1]++
		case "L":
			h[1]--
		case "U":
			h[0]--
		case "D":
			h[0]++
		}
		if !(near(h, t)) {
			t[0] = prev[0]
			t[1] = prev[1]
		}
		done[t[0]][t[1]] = true
	}
	return cpt(done)
}

func part2(s string) int {
	done := [1000][1000]bool{}
	rope := [20]int{}
	i := 0
	for i < 20 {
		rope[i] = 500
		i++
	}
	motions := format(s)
	for _, x := range motions {
		switch x {
		case "R":
			rope[1]++
		case "L":
			rope[1]--
		case "U":
			rope[0]--
		case "D":
			rope[0]++
		}
		rope = update(rope)
		done[rope[18]][rope[19]] = true
	}

	return cpt(done)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day09/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day09/test2.data")

	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day09/input.data")

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
