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

func pattern(l []int) bool {
	for i, x := range l {
		if i%2 == 0 && x != 0 {
			return false
		}
		if i%2 == 1 && x != 1 {
			return false
		}
	}
	return true
}

func part1(s string) int {
	instr := format(s)
	j := 0
	for {
		registers := make(map[string]int)
		clockSignal := []int{}
		for _, str := range []string{"a", "b", "c", "d"} {
			registers[str] = 0
		}
		registers["a"] = j
		limit := 0
		for i := 0; i < len(instr) && limit < 100; i++ {
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
			case "out":
				n, err := strconv.Atoi(line[1])
				limit++
				if err == nil {
					clockSignal = append(clockSignal, n)
				} else {
					clockSignal = append(clockSignal, registers[line[1]])
				}
				if !pattern(clockSignal) {
					limit = 100
				}
			}
		}
		if pattern(clockSignal) {
			return j
		}
		j++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
