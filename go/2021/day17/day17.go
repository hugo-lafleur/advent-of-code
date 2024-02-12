package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector struct {
	x, y int
}

func Split(r rune) bool {
	return r == ' ' || r == '.' || r == '=' || r == ','
}

func format(s string) []int {
	split := strings.FieldsFunc(s, Split)
	a, _ := strconv.Atoi(split[3])
	b, _ := strconv.Atoi(split[4])
	c, _ := strconv.Atoi(split[6])
	d, _ := strconv.Atoi(split[7])
	return []int{a, b, c, d}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxHeight(velocity vector, target []int) (bool, int) {
	pos := vector{0, 0}
	maxY := 0
	vel := velocity
	for pos.x <= target[1] && pos.y >= target[2] {
		if pos.x >= target[0] && pos.y <= target[3] {
			return true, maxY
		}
		pos.x += vel.x
		pos.y += vel.y
		if vel.x > 0 {
			vel.x--
		}
		if vel.x < 0 {
			vel.x++
		}
		vel.y--
		maxY = max(maxY, pos.y)
	}
	return false, maxY
}

func part1(s string) int {
	var maxH int
	target := format(s)
	for i := 0; i <= target[1]+1; i++ {
		for j := -1000; j < 1000; j++ {
			velocity := vector{i, j}
			reached, maxY := maxHeight(velocity, target)
			if reached {
				maxH = max(maxH, maxY)
			}
		}
	}
	return maxH
}

func part2(s string) int {
	c := 0
	target := format(s)
	for i := 0; i <= target[1]+1; i++ {
		for j := -1000; j < 1000; j++ {
			velocity := vector{i, j}
			reached, _ := maxHeight(velocity, target)
			if reached {
				c++
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day17/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day17/input.data")

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
