package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type program struct {
	name     string
	weight   int
	children []string
}

func Split(r rune) bool {
	return r == ' ' || r == '(' || r == ')' || r == '-' || r == ',' || r == '>'
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func calculateWeight(p program, l []program) int {
	res := p.weight
	for _, children := range p.children {
		for _, x := range l {
			if children == x.name {
				res += calculateWeight(x, l)
			}
		}
	}
	return res
}

func allSame(l []int) bool {
	c := l[0]
	for _, x := range l {
		if x != c {
			return false
		}
	}
	return true
}

func indexDifferent(l []int) int {
	if l[0] == l[1] {
		for i, x := range l {
			if x != l[0] {
				return i
			}
		}
	}
	if l[0] == l[2] {
		return 1
	}
	return 0
}

func part1(s string) string {
	tab := format(s)
	programList := []program{}
	for _, line := range tab {
		name := line[0]
		n, _ := strconv.Atoi(line[1])
		children := []string{}
		for i := 2; i < len(line); i++ {
			children = append(children, line[i])
		}
		programList = append(programList, program{name, n, children})
	}
	isChildren := make(map[string]bool)
	for _, program := range programList {
		isChildren[program.name] = false
	}
	for _, program := range programList {
		for _, children := range program.children {
			isChildren[children] = true
		}
	}
	for key, value := range isChildren {
		if !value {
			return key
		}
	}
	return ""
}

func part2(s string) int {
	tab := format(s)
	programList := []program{}
	weights := make(map[string]int)
	for _, line := range tab {
		name := line[0]
		n, _ := strconv.Atoi(line[1])
		children := []string{}
		for i := 2; i < len(line); i++ {
			children = append(children, line[i])
		}
		weights[name] = n
		programList = append(programList, program{name, n, children})
	}
	for _, p := range programList {
		weights[p.name] = calculateWeight(p, programList)
	}
	candidates := []program{}
	for _, p := range programList {
		childrenWeights := []int{}
		for _, children := range p.children {
			childrenWeights = append(childrenWeights, weights[children])
		}
		if len(childrenWeights) > 1 && !allSame(childrenWeights) {
			candidates = append(candidates, p)
		}
	}
	var toBalance program
loop:
	for _, c := range candidates {
		for _, children := range c.children {
			for _, test := range candidates {
				if test.name == children {
					continue loop
				}
			}
		}
		toBalance = c
	}
	childrenWeights := []int{}
	for _, children := range toBalance.children {
		childrenWeights = append(childrenWeights, weights[children])
	}
	unbanlancedChildren := toBalance.children[indexDifferent(childrenWeights)]
	var baseWeight int
	for _, p := range programList {
		if p.name == unbanlancedChildren {
			baseWeight = p.weight
		}
	}
	for _, w := range childrenWeights {
		if weights[unbanlancedChildren] != w {
			return (w - weights[unbanlancedChildren]) + baseWeight
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day07/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day07/input.data")

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
