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
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func isExecuted(instr []string, registers map[string]int) bool {
	n, _ := strconv.Atoi(instr[6])
	x := registers[instr[4]]
	switch instr[5] {
	case ">":
		return x > n
	case "<":
		return x < n
	case ">=":
		return x >= n
	case "<=":
		return x <= n
	case "==":
		return x == n
	case "!=":
		return x != n
	}
	return false
}

func part1(s string) int {
	instrList := format(s)
	registers := make(map[string]int)
	for _, instr := range instrList {
		registers[instr[0]] = 0
		registers[instr[4]] = 0
	}
	for _, instr := range instrList {
		if isExecuted(instr, registers) {
			n, _ := strconv.Atoi(instr[2])
			switch instr[1] {
			case "inc":
				registers[instr[0]] += n
			case "dec":
				registers[instr[0]] -= n
			}
		}
	}
	var max int
	for _, value := range registers {
		max = value
	}
	for _, value := range registers {
		if value > max {
			max = value
		}
	}
	return max
}

func part2(s string) int {
	instrList := format(s)
	registers := make(map[string]int)
	max := 0
	for _, instr := range instrList {
		registers[instr[0]] = 0
		registers[instr[4]] = 0
	}
	for _, instr := range instrList {
		if isExecuted(instr, registers) {
			n, _ := strconv.Atoi(instr[2])
			switch instr[1] {
			case "inc":
				registers[instr[0]] += n
			case "dec":
				registers[instr[0]] -= n
			}
			if registers[instr[0]] > max {
				max = registers[instr[0]]
			}
		}
	}
	return max
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day08/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2017/day08/input.txt")

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
