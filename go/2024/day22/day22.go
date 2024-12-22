package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Pair struct {
	bananas int
	change  int
}

func parse(s string) []int {
	var lines = strings.Split(s, "\n")
	var result = make([]int, len(lines))
	for i := range result {
		n, _ := strconv.Atoi(lines[i])
		result[i] = n
	}
	return result
}

func deltaList(n int) []Pair {
	var mod = 16777216
	var last = n
	var result []Pair
	for range 2000 {
		n = ((n * 64) ^ n) % mod
		n = ((n / 32) ^ n) % mod
		n = ((n * 2048) ^ n) % mod
		result = append(result, Pair{n % 10, n%10 - last%10})
		last = n
	}
	return result
}

func part1(s string) int {
	var result int
	var mod = 16777216
	var list = parse(s)
	for _, n := range list {
		for range 2000 {
			n = ((n * 64) ^ n) % mod
			n = ((n / 32) ^ n) % mod
			n = ((n * 2048) ^ n) % mod
		}
		result += n
	}
	return result
}

func part2(s string) int {
	var list = parse(s)
	var bananas = make([]int, 130321)
	for _, n := range list {
		var delta = deltaList(n)
		var seen = make([]bool, 130321)
		var seq int
		for i := range 3 {
			seq *= 19
			seq += delta[i].change + 9
		}
		for i := 3; i < len(delta); i++ {
			seq *= 19
			seq = seq % 130321
			seq += delta[i].change + 9
			if !seen[seq] {
				bananas[seq] += delta[i].bananas
				seen[seq] = true
			}
		}
	}
	return slices.Max(bananas)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day22/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day22/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day22/input.txt")

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
