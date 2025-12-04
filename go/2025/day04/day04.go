package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	result := [][]string{}
	for _, line := range lines {
		result = append(result, strings.Split(line, ""))
	}
	return result
}

func removableList(tab [][]string) [][2]int {
	dirs := [][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}}
	result := [][2]int{}
	for r := range tab {
		for c := range tab[r] {
			if tab[r][c] == "@" {
				count := 0
				for _, dir := range dirs {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nc >= 0 && nr < len(tab) && nc < len(tab[0]) && tab[nr][nc] == "@" {
						count++
					}
				}
				if count < 4 {
					result = append(result, [2]int{r, c})
				}
			}
		}
	}
	return result
}

func part1(s string) int {
	tab := format(s)
	return len(removableList(tab))
}

func part2(s string) int {
	tab := format(s)
	result := 0
	for {
		list := removableList(tab)
		if len(list) == 0 {
			return result
		}
		result += len(list)
		for _, p := range list {
			tab[p[0]][p[1]] = "."
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day04/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day04/input.txt")

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
