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

func format(s string) []int {
	split := strings.Split(s, ",")
	res := []int{}
	for _, x := range split {
		n, _ := strconv.Atoi(x)
		res = append(res, n)
	}
	return res
}

func triangleSum(n int, cache map[int]int) int {
	_, ok := cache[n]
	if ok {
		return cache[n]
	}
	if n == 0 {
		return 0
	}
	cache[n] = n + triangleSum(n-1, cache)
	return cache[n]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumDiff(x int, list []int) int {
	res := 0
	for _, n := range list {
		res += abs(n - x)
	}
	return res
}

func sumDiffTriange(x int, list []int, cache map[int]int) int {
	res := 0
	for _, n := range list {
		res += triangleSum(abs(n-x), cache)
	}
	return res
}

func part1(s string) int {
	list := format(s)
	min, max := slices.Min(list), slices.Max(list)
	minFuel := sumDiff(min, list)
	for i := min; i <= max; i++ {
		if sumDiff(i, list) < minFuel {
			minFuel = sumDiff(i, list)
		}
	}
	return minFuel
}

func part2(s string) int {
	list := format(s)
	min, max := slices.Min(list), slices.Max(list)
	cache := make(map[int]int)
	minFuel := sumDiffTriange(min, list, cache)
	for i := min; i <= max; i++ {
		if sumDiffTriange(i, list, cache) < minFuel {
			minFuel = sumDiffTriange(i, list, cache)
		}
	}
	return minFuel
}
func main() {
	content, err := os.ReadFile("../../../inputs/2021/day07/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day07/input.txt")

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
