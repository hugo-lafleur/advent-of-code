package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
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

func format(s string) ([]int, []int) {
	ranges := []int{}
	ids := []int{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		} else if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			a, _ := strconv.Atoi(parts[0])
			b, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, a)
			ranges = append(ranges, -b)
		} else {
			n, _ := strconv.Atoi(line)
			ids = append(ids, n)
		}
	}
	slices.SortFunc(ranges, func(a, b int) int {
		return cmp.Or(cmp.Compare(abs(a), abs(b)), cmp.Compare(b, a))
	})
	return ranges, ids
}

func part1(s string) int {
	ranges, ids := format(s)
	result := 0
	for _, id := range ids {
		curr := 0
		for _, r := range ranges {
			if id <= abs(r) {
				if curr > 0 {
					result++
				}
				break
			}
			if r < 0 {
				curr--
			} else {
				curr++
			}
		}
	}
	return result
}

func part2(s string) int {
	ranges, _ := format(s)
	result := 0
	curr := 0
	last := 0
	for _, r := range ranges {
		if r > 0 {
			if curr == 0 {
				last = r
			}
			curr++
		} else {
			curr--
			if curr == 0 {
				result += -r - last + 1
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day05/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day05/input.txt")

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
