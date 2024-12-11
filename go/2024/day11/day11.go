package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse(s string) map[int]int {
	var parts = strings.Split(s, " ")
	var result = make(map[int]int, len(parts))
	for _, str := range parts {
		n, _ := strconv.Atoi(str)
		result[n]++
	}
	return result
}

func split(stone int) (int, int) {
	var str = strconv.Itoa(stone)
	var n = len(str) / 2
	a, _ := strconv.Atoi(str[:n])
	b, _ := strconv.Atoi(str[n:])
	return a, b
}

func solve(s string, blinks int) int {
	var stones = parse(s)
	for range blinks {
		newStones := make(map[int]int, len(stones))
		for stone, val := range stones {
			if stone == 0 {
				newStones[1] += val
			} else if len(strconv.Itoa(stone))%2 == 0 {
				a, b := split(stone)
				newStones[a] += val
				newStones[b] += val
			} else {
				newStones[stone*2024] += val
			}
		}
		stones = newStones
	}
	var result int
	for _, val := range stones {
		result += val
	}
	return result
}

func part1(s string) int {
	return solve(s, 25)
}

func part2(s string) int {
	return solve(s, 75)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day11/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day11/input.txt")

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
