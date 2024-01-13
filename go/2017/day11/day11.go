package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, ",")
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(s string) int {
	directions := format(s)
	i, j := 0, 0
	for _, d := range directions {
		switch d {
		case "n":
			j -= 2
		case "nw":
			j--
			i--
		case "ne":
			i++
			j--
		case "se":
			j++
			i++
		case "sw":
			j++
			i--
		case "s":
			j += 2
		}
	}
	return abs(i) + max(0, (abs(j)-abs(i))/2)
}

func part2(s string) int {
	directions := format(s)
	i, j := 0, 0
	res := 0
	for _, d := range directions {
		switch d {
		case "n":
			j -= 2
		case "nw":
			j--
			i--
		case "ne":
			i++
			j--
		case "se":
			j++
			i++
		case "sw":
			j++
			i--
		case "s":
			j += 2
		}
		distance := abs(i) + max(0, (abs(j)-abs(i))/2)
		res = max(res, distance)
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day11/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day11/input.data")

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
