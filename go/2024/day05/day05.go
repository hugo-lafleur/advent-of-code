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

func parse(s string) (map[int]map[int]bool, [][]int) {
	var rules = make(map[int]map[int]bool)
	var updates [][]int
	var lines = strings.Split(s, "\n")
	var i int
	for i < len(lines) {
		if lines[i] == "" {
			i++
			break
		}
		parts := strings.Split(lines[i], "|")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		if rules[a] == nil {
			rules[a] = make(map[int]bool)
		}
		if rules[b] == nil {
			rules[b] = make(map[int]bool)
		}
		rules[a][b] = false
		rules[b][a] = true
		i++
	}
	for i < len(lines) {
		parts := strings.Split(lines[i], ",")
		var temp []int
		for _, str := range parts {
			n, _ := strconv.Atoi(str)
			temp = append(temp, n)
		}
		updates = append(updates, temp)
		i++
	}
	return rules, updates
}

func correctlyOrderedUpdate(rules map[int]map[int]bool, update []int) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if rules[update[i]][update[j]] {
				return false
			}
		}
	}
	return true
}

func orderedUpdate(rules map[int]map[int]bool, update []int) []int {
	var result = make([]int, len(update))
	copy(result, update)
	var count = make(map[int]int)
	for _, x := range update {
		for _, y := range update {
			if !rules[x][y] {
				count[x]++
			}
		}
	}
	slices.SortFunc(result, func(a, b int) int {
		return count[b] - count[a]
	})
	return result
}

func part1(s string) int {
	var rules, updates = parse(s)
	var result int
	for _, upd := range updates {
		if correctlyOrderedUpdate(rules, upd) {
			result += upd[len(upd)/2]
		}
	}
	return result
}

func part2(s string) int {
	var rules, updates = parse(s)
	var result int
	for _, upd := range updates {
		if !correctlyOrderedUpdate(rules, upd) {
			result += orderedUpdate(rules, upd)[len(upd)/2]
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day05/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day05/input.txt")

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
