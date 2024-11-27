package main

import (
	"fmt"
	"log"
	"os"
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

func part1(s string) int {
	fish := format(s)
	day := 0
	for day < 80 {
		for i := range fish {
			fish[i]--
			if fish[i] == -1 {
				fish[i] = 6
				fish = append(fish, 8)
			}
		}
		day++
	}
	return len(fish)
}

func part2(s string) int {
	c := 0
	fish := format(s)
	m := make(map[int]int)
	for _, x := range fish {
		m[x]++
	}
	day := 0
	for day < 256 {
		newMap := make(map[int]int)
		for key, value := range m {
			if key == 0 {
				newMap[8] += value
				newMap[6] += value
			} else {
				newMap[key-1] += value
			}
		}
		m = newMap
		day++
	}
	for _, value := range m {
		c += value
	}
	return c
}
func main() {
	content, err := os.ReadFile("../../../inputs/2021/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day06/input.txt")

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
