package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solve(s string, maxSteps int) string {
	curr := []int{}
	for _, r := range s {
		curr = append(curr, int(r-'0'))
	}
	steps := 0
	for steps < maxSteps {
		next := []int{}
		l := len(curr)
		for i := 0; i < l; i++ {
			pattern := []int{}
			for len(pattern) < l+1 {
				for _, n := range []int{0, 1, 0, -1} {
					for j := 0; j < i+1; j++ {
						pattern = append(pattern, n)
					}
				}
			}
			pattern = pattern[1:(l + 1)]
			temp := 0
			for j := i; j < l; j++ {
				n := curr[j]
				temp += n * pattern[j]
			}
			next = append(next, abs(temp)%10)
		}
		curr = next
		steps++
	}
	res := ""
	for i := 0; i < 9; i++ {
		res += strconv.Itoa(curr[i])
	}
	return res
}

func solve2(s string, maxSteps int) string {
	curr := []int{}
	for _, r := range s {
		curr = append(curr, int(r-'0'))
	}
	steps := 0
	for steps < maxSteps {
		l := len(curr)
		temp := 0
		for j := l - 1; j >= 0; j-- {
			temp = (temp + curr[j]) % 10
			curr[j] = temp
		}
		steps++
	}
	res := ""
	for i := 0; i < 9; i++ {
		res += strconv.Itoa(curr[i])
	}
	return res
}

func part1(s string) string {
	return solve(s, 100)
}

func part2(s string) string {
	offset, _ := strconv.Atoi(s[0:7])
	str := strings.Repeat(s, 10000)
	str = str[offset:]
	return solve2(str, 100)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day16/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day16/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day16/input.data")

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
