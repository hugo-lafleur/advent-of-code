package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	name string
	n    int
}

func format(s string) []instruction {
	lines := strings.Split(s, "\n")
	res := []instruction{}
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		name := lineSplit[0]
		n, _ := strconv.Atoi(lineSplit[1])
		res = append(res, instruction{name: name, n: n})
	}
	return res
}

func solve(instructions []instruction) (bool, int) {
	memory := make(map[int]int)
	i := 0
	acc := 0
	for i < len(instructions) {
		instr := instructions[i]
		if memory[i] == 1 {
			return false, acc
		}
		memory[i]++
		switch instr.name {
		case "nop":
			i++
		case "acc":
			acc += instr.n
			i++
		case "jmp":
			i += instr.n
		}
	}
	return true, acc
}

func part1(s string) int {
	instructions := format(s)
	_, n := solve(instructions)
	return n
}

func part2(s string) int {
	baseInstructions := format(s)
	for i := 0; i < len(baseInstructions); i++ {
		if baseInstructions[i].name == "acc" {
			continue
		}
		instructions := make([]instruction, len(baseInstructions))
		copy(instructions, baseInstructions)
		switch baseInstructions[i].name {
		case "nop":
			instructions[i].name = "jmp"
		case "jmp":
			instructions[i].name = "nop"
		}
		ok, acc := solve(instructions)
		if ok {
			return acc
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day08/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day08/input.data")

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
