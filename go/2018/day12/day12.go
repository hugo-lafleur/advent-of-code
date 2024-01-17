package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type mem struct {
	sum, g int
}

func format(s string) (string, [][]string) {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Fields(line))
	}
	return res[0][2], res[2:]
}

func position(s string) int {
	for i, r := range s {
		if r == '#' {
			return i
		}
	}
	return 0
}

func count(s string) int {
	c := 0
	for _, r := range s {
		if r == '#' {
			c++
		}
	}
	return c
}

func part1(s string) int {
	c := 0
	state, rulesStr := format(s)
	rules := make(map[string]string)
	for _, line := range rulesStr {
		rules[line[0]] = line[2]
	}
	generation := 0
	offset := 0
	for generation < 20 {
		copy := "...." + strings.Clone(state) + "...."
		state = ""
		for i := 2; i < len(copy)-2; i++ {
			res, ok := rules[copy[i-2:i+3]]
			if ok {
				state += res
				continue
			}
			state += "."
		}
		offset += 2
		generation++
	}
	for i, r := range state {
		if r == '#' {
			c += (i - offset)
		}
	}
	return c
}

func part2(s string) int {
	memory := make(map[string]mem)
	state, rulesStr := format(s)
	rules := make(map[string]string)
	var infinite int
	for _, line := range rulesStr {
		rules[line[0]] = line[2]
		if strings.Trim(line[0], ".") == "#" && line[2] == "#" {
			infinite = 2 - position(line[0])
		}
	}
	generation := 0
	offset := 0
	for {
		generation++
		copy := "...." + strings.Clone(state) + "...."
		state = ""
		for i := 2; i < len(copy)-2; i++ {
			res, ok := rules[copy[i-2:i+3]]
			if ok {
				state += res
				continue
			}
			state += "."
		}
		offset += 2
		last, ok := memory[strings.Trim(state, ".")]
		if ok {
			return last.sum + count(strings.Trim(state, "."))*(50000000000-last.g)*infinite
		}
		c := 0
		for i, r := range state {
			if r == '#' {
				c += (i - offset)
			}
		}
		memory[strings.Trim(state, ".")] = mem{c, generation}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day12/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day12/input.data")

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
