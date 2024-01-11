package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		n := len(parts)
		name := ""
		for i := 0; i < len(parts)-1; i++ {
			name += parts[i]
		}
		last := strings.Split(parts[n-1], "[")
		formatLine := []string{}
		formatLine = append(formatLine, name, last[0], last[1][:len(last[1])-1])
		res = append(res, formatLine)
	}
	return res
}

func createDict(name string) map[rune]int {
	res := make(map[rune]int)
	for _, r := range name {
		res[r]++
	}
	return res
}

func maxDict(dict map[rune]int) rune {
	max := 0
	var maxRune rune
	for key, value := range dict {
		if value > max {
			max = value
			maxRune = key
		}
		if value == max {
			if int(key) < int(maxRune) {
				max = value
				maxRune = key
			}
		}
	}
	return maxRune
}

func runeInString(r rune, s string) bool {
	for _, c := range s {
		if r == c {
			return true
		}
	}
	return false
}

func rotation(r rune, k int) rune {
	n := int(r)
	new := ((n+k)-97)%26 + 97
	return rune(new)
}

func part1(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		dict := createDict(line[0])
		real := true
		for i := 0; i < 5; i++ {
			max := maxDict(dict)
			delete(dict, max)
			if !(runeInString(max, line[2])) {
				real = false
				break
			}
		}
		if real {
			n, _ := strconv.Atoi(line[1])
			c += n
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	tab := format(s)
	for _, line := range tab {
		n, _ := strconv.Atoi(line[1])
		s := ""
		for _, c := range line[0] {
			s += string(rotation(c, n))
		}
		if s == "northpoleobjectstorage" {
			c = n
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day04/test.data")

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

	content, err = os.ReadFile("../../../inputs/2016/day04/input.data")

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
