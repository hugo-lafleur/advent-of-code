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
	return r == ' ' || r == ':'
}

func format(s string) map[string][]string {
	lines := strings.Split(s, "\n")
	res := make(map[string][]string)
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		res[lineSplit[0]] = lineSplit[1:]
	}
	return res
}

func solve(s string, rules map[string][]string) int {
	if len(rules[s]) == 1 {
		n, _ := strconv.Atoi(rules[s][0])
		return n
	} else {
		switch rules[s][1] {
		case "+":
			return solve(rules[s][0], rules) + solve(rules[s][2], rules)
		case "*":
			return solve(rules[s][0], rules) * solve(rules[s][2], rules)
		case "-":
			return solve(rules[s][0], rules) - solve(rules[s][2], rules)
		case "/":
			return solve(rules[s][0], rules) / solve(rules[s][2], rules)
		case "=":
			a := solve(rules[s][0], rules)
			b := solve(rules[s][2], rules)
			return a - b
		}
	}
	return 0
}

func part1(s string) int {
	rules := format(s)
	return solve("root", rules)
}

func part2(s string) int {
	rules := format(s)
	rules["root"][1] = "="
	a, b := 0, 10000000000000000
	var f0, f10 int
	for i := 0; i <= 10; i += 10 {
		rules["humn"] = []string{fmt.Sprint(i)}
		res := solve("root", rules)
		if i == 0 {
			f0 = res
		}
		if i == 10 {
			f10 = res
		}
	}
	for {
		i := (a + b) / 2
		rules["humn"] = []string{fmt.Sprint(i)}
		res := solve("root", rules)
		if (res > 0 && f10 > f0) || (res < 0 && f10 < f0) {
			b = i
		}
		if (res < 0 && f10 > f0) || (res > 0 && f10 < f0) {
			a = i
		}
		if res == 0 {
			return i
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day21/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day21/input.txt")

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
