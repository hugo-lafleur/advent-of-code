package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type couple struct {
	a, b int
}

func format(s string) (int, int) {
	lines := strings.Split(s, "\n")
	a, _ := strconv.Atoi(strings.Split(lines[0], " ")[4])
	b, _ := strconv.Atoi(strings.Split(lines[1], " ")[4])
	return a, b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(pos1, score1, pos2, score2, turn int, cache map[string]couple) (int, int) {
	key := fmt.Sprint(pos1, score1, pos2, score2, turn)
	_, ok := cache[key]
	if ok {
		return cache[key].a, cache[key].b
	}
	var a, b int
	var tempA, tempB int
	if turn == 0 {
		for _, roll := range []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9} {
			new_pos := pos1
			new_pos += roll
			for new_pos > 10 {
				new_pos -= 10
			}
			new_score := new_pos + score1
			if new_score >= 21 {
				a++
				continue
			}
			tempA, tempB = solve(new_pos, new_score, pos2, score2, 1, cache)
			a += tempA
			b += tempB
		}
	}
	if turn == 1 {
		for _, roll := range []int{3, 4, 5, 4, 5, 6, 5, 6, 7, 4, 5, 6, 5, 6, 7, 6, 7, 8, 5, 6, 7, 6, 7, 8, 7, 8, 9} {
			new_pos := pos2
			new_pos += roll
			for new_pos > 10 {
				new_pos -= 10
			}
			new_score := new_pos + score2
			if new_score >= 21 {
				b++
				continue
			}
			tempA, tempB = solve(pos1, score1, new_pos, new_score, 0, cache)
			a += tempA
			b += tempB
		}
	}
	cache[key] = couple{a, b}
	return a, b
}

func part1(s string) int {
	pos1, pos2 := format(s)
	score1, score2 := 0, 0
	dice := 1
	for score1 < 1000 && score2 < 1000 {
		roll := 0
		for i := 0; i < 3; i++ {
			roll += dice
			dice++

		}
		pos1 += roll
		for pos1 > 10 {
			pos1 -= 10
		}
		score1 += pos1
		if score1 >= 1000 {
			break
		}
		roll = 0
		for i := 0; i < 3; i++ {
			roll += dice
			dice++

		}
		pos2 += roll
		for pos2 > 10 {
			pos2 -= 10
		}
		score2 += pos2
	}
	return min(score1, score2) * (dice - 1)
}

func part2(s string) int {
	pos1, pos2 := format(s)
	a, b := solve(pos1, 0, pos2, 0, 0, make(map[string]couple))
	return max(a, b)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day21/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day21/input.data")

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
