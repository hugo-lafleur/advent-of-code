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

func part1(s string) int {
	c := 0
	instrs := format(s)
	registers := make(map[string]int)
	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]
		switch instr[0] {
		case "set":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] = n
			} else {
				registers[instr[1]] = registers[instr[2]]
			}
		case "sub":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] -= n
			} else {
				registers[instr[1]] -= registers[instr[2]]
			}
		case "mul":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] *= n
			} else {
				registers[instr[1]] *= registers[instr[2]]
			}
			c++
		case "jnz":
			n, err1 := strconv.Atoi(instr[1])
			m, err2 := strconv.Atoi(instr[2])
			if err1 != nil {
				n = registers[instr[1]]
			}
			if err2 != nil {
				m = registers[instr[2]]
			}
			if n != 0 {
				i = i + m - 1
			}
		}
	}
	return c
}

func part2(s string) int {
	instrs := format(s)
	registers := make(map[string]int)
	registers["a"] = 1
	res := 0
	var b, c, step int
	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]
		switch instr[0] {
		case "set":
			n, err := strconv.Atoi(instr[2])
			if instr[1] == "f" && instr[2] == "1" {
				b = registers["b"]
				c = registers["c"]
				goto out
			}
			if err == nil {
				registers[instr[1]] = n
			} else {
				registers[instr[1]] = registers[instr[2]]
			}
		case "sub":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] -= n
			} else {
				registers[instr[1]] -= registers[instr[2]]
			}
		case "mul":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] *= n
			} else {
				registers[instr[1]] *= registers[instr[2]]
			}
		case "jnz":
			n, err1 := strconv.Atoi(instr[1])
			m, err2 := strconv.Atoi(instr[2])
			if err1 != nil {
				n = registers[instr[1]]
			}
			if err2 != nil {
				m = registers[instr[2]]
			}
			if n != 0 {
				i = i + m - 1
			}
		}
	}
out:
	for i := len(instrs) - 1; i >= 0; i-- {
		if instrs[i][0] == "sub" && instrs[i][1] == "b" {
			step, _ = strconv.Atoi(instrs[i][2])
			break
		}
	}
	step = -step
	for ; b <= c; b += step {
		for d := 2; d*d < b; d++ {
			if b%d == 0 {
				res++
				break
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day23/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
