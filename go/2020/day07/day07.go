package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '.'
}

func format(s string) map[string]map[string]int {
	lines := strings.Split(s, "\n")
	res := make(map[string]map[string]int)
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		name := lineSplit[0] + lineSplit[1]
		if lineSplit[4] == "no" {
			res[name] = make(map[string]int)
			continue
		}
		rules := make(map[string]int)
		for i := 4; i < len(lineSplit); i += 4 {
			n, _ := strconv.Atoi(lineSplit[i])
			rules[lineSplit[i+1]+lineSplit[i+2]] = n
		}
		res[name] = rules
	}
	return res
}

func canContainsShinyGold(bag string, bags map[string]map[string]int) bool {
	_, ok := bags[bag]["shinygold"]
	if ok {
		return true
	}
	for key := range bags[bag] {
		if canContainsShinyGold(key, bags) {
			return true
		}
	}
	return false
}

func numberOfBags(bag string, bags map[string]map[string]int) int {
	res := 1
	for key, value := range bags[bag] {
		res += value * numberOfBags(key, bags)
	}
	return res
}

func part1(s string) int {
	c := 0
	bags := format(s)
	for key := range bags {
		if canContainsShinyGold(key, bags) {
			c++
		}
	}
	return c
}

func part2(s string) int {
	bags := format(s)
	return numberOfBags("shinygold", bags) - 1
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day07/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day07/input.data")

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
