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
	return strings.Split(s, "\n")
}

func winningNumbers(card []string) []int {
	res := []int{}
	for _, s := range card {
		if s == "|" {
			return res
		}
		n, err := strconv.Atoi(s)
		if err == nil && n != 0 {
			res = append(res, n)
		}
	}
	return res
}

func numbers(card []string) []int {
	res := []int{}
	win := true
	for _, s := range card {
		if !win {
			n, _ := strconv.Atoi(s)
			if n != 0 {
				res = append(res, n)
			}
		}
		if s == "|" {
			win = false
		}
	}
	return res
}

func isIn(x int, list []int) bool {
	for _, y := range list {
		if x == y {
			return true
		}
	}
	return false
}

func matches(winningNumbers []int, numbers []int) int {
	c := 0
	for _, w := range winningNumbers {
		if isIn(w, numbers) {
			c++
		}
	}
	return c
}

func points(i int) int {
	if i < 2 {
		return i
	} else {
		return 2 * points(i-1)
	}
}

func part1(s string) int {
	c := 0
	lines := format(s)
	for _, line := range lines {
		card := strings.Split(line, " ")
		winningNumbers := winningNumbers(card)
		numbers := numbers(card)
		matches := matches(winningNumbers, numbers)
		c += points(matches)
	}
	return c
}

func sum(l []int) int {
	s := 0
	for _, x := range l {
		s += x
	}
	return s
}

func part2(s string) int {
	lines := format(s)
	cards := []int{}
	l := len(lines)
	i := 0
	for i < l {
		cards = append(cards, 1)
		i++
	}
	for i, line := range lines {
		card := strings.Split(line, " ")
		winningNumbers := winningNumbers(card)
		numbers := numbers(card)
		matches := matches(winningNumbers, numbers)
		k := 1
		for k < matches+1 {
			cards[i+k] += cards[i]
			k++
		}
	}
	return sum(cards)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day04/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day04/input.txt")

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
