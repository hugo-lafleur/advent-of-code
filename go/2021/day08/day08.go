package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) ([]string, []string) {
	lines := strings.Split(s, "\n")
	pattern := []string{}
	output := []string{}
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		i := 0
		for lineSplit[i] != "|" {
			i++
		}
		pattern = append(pattern, lineSplit[:i]...)
		output = append(output, lineSplit[i+1:]...)
	}
	return pattern, output
}

func invSolve(mapping map[string]string, str string) int {
	x := str
	if len(x) == 2 {
		return 1
	}
	if len(x) == 4 {
		return 4
	}
	if len(x) == 3 {
		return 7
	}
	if len(x) == 7 {
		return 8
	}
	if len(x) == 6 && !strings.Contains(x, mapping["d"]) {
		return 0
	}
	if len(x) == 6 && !strings.Contains(x, mapping["c"]) {
		return 6
	}
	if len(x) == 6 && !strings.Contains(x, mapping["e"]) {
		return 9
	}
	if len(x) == 5 && !strings.Contains(x, mapping["f"]) {
		return 2
	}
	if len(x) == 5 && strings.Contains(x, mapping["c"]) && strings.Contains(x, mapping["f"]) {
		return 3
	}
	if len(x) == 5 && strings.Contains(x, mapping["b"]) {
		return 5
	}
	return -1
}

func solve(pattern []string, output []string) int {
	list := []string{"a", "b", "c", "d", "e", "f", "g"}
	mapping := make(map[string]string)
	patternMap := make(map[int]string)
	for _, x := range pattern {
		if len(x) == 2 {
			patternMap[1] = x
		}
		if len(x) == 4 {
			patternMap[4] = x
		}
		if len(x) == 3 {
			patternMap[7] = x
		}
		if len(x) == 7 {
			patternMap[8] = x
		}
	}
	for _, letter := range list {
		if strings.Contains(patternMap[7], letter) && !strings.Contains(patternMap[1], letter) {
			mapping["a"] = letter
		}
	}
	count := make(map[string]int)
	for _, letter := range list {
		for _, p := range pattern {
			if strings.Contains(p, letter) {
				count[letter]++
			}
		}
	}
	for key, value := range count {
		if value == 8 && key != mapping["a"] {
			mapping["c"] = key
		}
		if value == 6 {
			mapping["b"] = key
		}
		if value == 7 && strings.Contains(patternMap[4], key) {
			mapping["d"] = key
		}
		if value == 4 {
			mapping["e"] = key
		}
		if value == 9 {
			mapping["f"] = key
		}
		if value == 7 && !strings.Contains(patternMap[4], key) {
			mapping["g"] = key
		}
	}
	res := ""
	for _, o := range output {
		res += strconv.Itoa(invSolve(mapping, o))
	}
	n, _ := strconv.Atoi(res)
	return n
}

func part1(s string) int {
	c := 0
	_, output := format(s)
	for _, str := range output {
		if len(str) == 2 || len(str) == 4 || len(str) == 3 || len(str) == 7 {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	pattern, output := format(s)
	for i := 0; i < (len(pattern) / 10); i++ {
		currPattern := pattern[(i * 10):((i + 1) * 10)]
		currOutput := output[(i * 4):((i + 1) * 4)]
		c += solve(currPattern, currOutput)
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day08/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day08/input.data")

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
