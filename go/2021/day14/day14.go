package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) (string, map[string]string) {
	lines := strings.Split(s, "\n")
	rules := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		lineSplit := strings.Split(lines[i], " -> ")
		rules[lineSplit[0]] = lineSplit[1]
	}
	return lines[0], rules
}

func diffMaxMin(m map[string]int) int {
	var min, max int
	for _, value := range m {
		min = value
		max = value
	}
	for _, value := range m {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}
	return max - min
}

func part1(s string) int {
	str, rules := format(s)
	polymer := make(map[string]int)
	total := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		polymer[str[i:(i+2)]]++
	}
	for i := 0; i < len(str); i++ {
		total[str[i:(i+1)]]++
	}
	steps := 0
	for steps < 10 {
		newMap := make(map[string]int)
		for key, value := range polymer {
			insert := rules[key]
			newMap[key[0:1]+insert] += value
			total[insert] += value
			newMap[insert+key[1:2]] += value
		}
		steps++
		polymer = newMap
	}
	return diffMaxMin(total)
}

func part2(s string) int {
	str, rules := format(s)
	polymer := make(map[string]int)
	total := make(map[string]int)
	for i := 0; i < len(str)-1; i++ {
		polymer[str[i:(i+2)]]++
	}
	for i := 0; i < len(str); i++ {
		total[str[i:(i+1)]]++
	}
	steps := 0
	for steps < 40 {
		newMap := make(map[string]int)
		for key, value := range polymer {
			insert := rules[key]
			newMap[key[0:1]+insert] += value
			total[insert] += value
			newMap[insert+key[1:2]] += value
		}
		steps++
		polymer = newMap
	}
	return diffMaxMin(total)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day14/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day14/input.data")

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
