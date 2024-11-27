package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type movement struct {
	write    int
	slot     int
	newState string
}

type state struct {
	name string
	zero movement
	one  movement
}

func Split(r rune) bool {
	return r == '.' || r == ' ' || r == '	' || r == ':'
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func checksum(tape map[int]int) int {
	r := 0
	for _, v := range tape {
		if v == 1 {
			r++
		}
	}
	return r
}

func part1(s string) int {
	input := format(s)
	n, _ := strconv.Atoi(input[1][5])
	rules := make(map[state]bool)
	for i := 3; i < len(input); i += 6 {
		name := input[i][2]
		write, _ := strconv.Atoi(input[i+2][4])
		direction := input[i+3][6]
		var slot int
		if direction == "left" {
			slot = -1
		} else {
			slot = 1
		}
		newState := input[i+4][4]
		zero := movement{write, slot, newState}
		i += 4
		write, _ = strconv.Atoi(input[i+2][4])
		direction = input[i+3][6]
		if direction == "left" {
			slot = -1
		} else {
			slot = 1
		}
		newState = input[i+4][4]
		one := movement{write, slot, newState}
		rules[state{name, zero, one}] = true
	}
	steps := 0
	currentState := input[0][3]
	cursor := 0
	tape := make(map[int]int)
	for steps < n {
		var currentRule state
		for rule := range rules {
			if rule.name == currentState {
				currentRule = rule
				break
			}
		}
		n, ok := tape[cursor]
		if !ok || n == 0 {
			tape[cursor] = currentRule.zero.write
			cursor += currentRule.zero.slot
			currentState = currentRule.zero.newState
		} else {
			tape[cursor] = currentRule.one.write
			cursor += currentRule.one.slot
			currentState = currentRule.one.newState
		}
		steps++
	}
	return checksum(tape)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day25/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
