package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func parse(s string) ([7]int, []int) {
	var lines = strings.Split(s, "\n")
	var A, B, C int
	fmt.Sscanf(lines[0], "Register A: %d", &A)
	fmt.Sscanf(lines[1], "Register B: %d", &B)
	fmt.Sscanf(lines[2], "Register C: %d", &C)
	var program []int
	for _, str := range strings.Split(lines[4][9:], ",") {
		n, _ := strconv.Atoi(str)
		program = append(program, n)
	}
	return [7]int{0, 1, 2, 3, A, B, C}, program
}

func solve(registers [7]int, program []int) []int {
	var i int
	var result []int
	for i < len(program) {
		var literal = program[i+1]
		var combo int
		if literal < 7 {
			combo = registers[program[i+1]]
		}
		switch program[i] {
		case 0:
			registers[4] = registers[4] / (1 << combo)
		case 1:
			registers[5] = registers[5] ^ literal
		case 2:
			registers[5] = combo % 8
		case 3:
			if registers[4] != 0 {
				i = literal
				continue
			}
		case 4:
			registers[5] = registers[5] ^ registers[6]
		case 5:
			result = append(result, combo%8)
		case 6:
			registers[5] = registers[4] / (1 << combo)
		case 7:
			registers[6] = registers[4] / (1 << combo)
		}
		i += 2
	}
	return result
}

func part1(s string) string {
	var registers, program = parse(s)
	var result []string
	var output = solve(registers, program)
	for i := range output {
		result = append(result, strconv.Itoa(output[i]))
	}
	return strings.Join(result, ",")
}

func part2(s string) int {
	var registers, program = parse(s)
	var result int
	var mul = 1
	var i int
	for i < len(program) {
		for offset := range 8 {
			registers[4] = result + offset*mul
			if solve(registers, program)[0] == program[len(program)-1-i] {
				result = registers[4] << 3
				i++
				break
			}
			if offset == 7 {
				mul *= 8
			}
		}
	}
	return result / 8
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day17/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day17/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day17/input.txt")

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
