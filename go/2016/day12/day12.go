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
	instr := format(s)
	registers := make(map[string]int)
	for _, str := range []string{"a", "b", "c", "d"} {
		registers[str] = 0
	}
	for i := 0; i < len(instr); i++ {
		line := instr[i]
		switch line[0] {
		case "cpy":
			n, err := strconv.Atoi(line[1])
			if err == nil {
				registers[line[2]] = n
			} else {
				registers[line[2]] = registers[line[1]]
			}
		case "inc":
			registers[line[1]]++
		case "dec":
			registers[line[1]]--
		case "jnz":
			c, err := strconv.Atoi(line[1])
			if registers[line[1]] != 0 || (err == nil && c != 0) {
				n, _ := strconv.Atoi(line[2])
				i += n - 1
			}
		}

	}
	return registers["a"]
}

func part2(s string) int {
	instr := format(s)
	registers := make(map[string]int)
	for _, str := range []string{"a", "b", "d"} {
		registers[str] = 0
	}
	registers["c"] = 1
	for i := 0; i < len(instr); i++ {
		line := instr[i]
		switch line[0] {
		case "cpy":
			n, err := strconv.Atoi(line[1])
			if err == nil {
				registers[line[2]] = n
			} else {
				registers[line[2]] = registers[line[1]]
			}
		case "inc":
			registers[line[1]]++
		case "dec":
			registers[line[1]]--
		case "jnz":
			c, err := strconv.Atoi(line[1])
			if registers[line[1]] != 0 || (err == nil && c != 0) {
				n, _ := strconv.Atoi(line[2])
				i += n - 1
			}
		}

	}
	return registers["a"]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day12/test.data")

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

	content, err = os.ReadFile("../../../inputs/2016/day12/input.data")

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
