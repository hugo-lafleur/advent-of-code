package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][2]int {
	ranges := strings.Split(s, ",")
	result := [][2]int{}
	for _, r := range ranges {
		ids := strings.Split(r, "-")
		a, _ := strconv.Atoi(ids[0])
		b, _ := strconv.Atoi(ids[1])
		result = append(result, [2]int{a, b})
	}
	return result
}

func invalidID1(n int) bool {
	str := strconv.Itoa(n)
	l := len(str) / 2
	return str[:l] == str[l:]
}

func invalidID2(n int) bool {
	str := strconv.Itoa(n)
	for l := 1; l <= len(str)/2; l++ {
		if len(str)%l == 0 {
			pattern := str[:l]
			if str == strings.Repeat(pattern, len(str)/l) {
				return true
			}
		}
	}
	return false
}

func part1(s string) int {
	result := 0
	ranges := format(s)
	for _, ids := range ranges {
		for i := ids[0]; i <= ids[1]; i++ {
			if invalidID1(i) {
				result += i
			}
		}
	}
	return result
}

func part2(s string) int {
	result := 0
	ranges := format(s)
	for _, ids := range ranges {
		for i := ids[0]; i <= ids[1]; i++ {
			if invalidID2(i) {
				result += i
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day02/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day02/input.txt")

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
