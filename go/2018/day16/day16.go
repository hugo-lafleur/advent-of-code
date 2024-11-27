package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type registers []int

type instruction []int

type sample struct {
	before registers
	instr  instruction
	after  registers
}

type possibility struct {
	table [][]string
	index int
}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '[' || r == ']'
}

func format1(s string) ([]sample, []instruction) {
	res := []sample{}
	lines := strings.Split(s, "\n")
	strs := [][]string{}
	split := 0
	for i, line := range lines {
		if line == "" && lines[i+1] == "" {
			split = i + 3
			break
		}
		strs = append(strs, strings.FieldsFunc(line, Split))
	}
	for i := 0; i < len(strs)-2; i += 4 {
		var before registers
		var after registers
		var instr instruction
		for j := 1; j < 5; j++ {
			n, _ := strconv.Atoi(strs[i][j])
			before = append(before, n)
		}
		for j := 0; j < 4; j++ {
			n, _ := strconv.Atoi(strs[i+1][j])
			instr = append(instr, n)
		}
		for j := 1; j < 5; j++ {
			n, _ := strconv.Atoi(strs[i+2][j])
			after = append(after, n)
		}
		res = append(res, sample{before, instr, after})
	}
	instrList := []instruction{}
	for i := split; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		temp := []int{}
		for _, x := range line {
			n, _ := strconv.Atoi(x)
			temp = append(temp, n)
		}
		instrList = append(instrList, temp)
	}
	return res, instrList
}

func nextRegisters(before registers, instr instruction, opcode string) registers {
	after := make(registers, len(before))
	copy(after, before)
	A := instr[1]
	B := instr[2]
	C := instr[3]
	switch opcode {
	case "addr":
		after[C] = before[A] + before[B]
	case "addi":
		after[C] = before[A] + B
	case "mulr":
		after[C] = before[A] * before[B]
	case "muli":
		after[C] = before[A] * B
	case "banr":
		after[C] = before[A] & before[B]
	case "bani":
		after[C] = before[A] & B
	case "borr":
		after[C] = before[A] | before[B]
	case "bori":
		after[C] = before[A] | B
	case "setr":
		after[C] = before[A]
	case "seti":
		after[C] = A
	case "gtir":
		if A > before[B] {
			after[C] = 1
		} else {
			after[C] = 0
		}
	case "gtri":
		if before[A] > B {
			after[C] = 1
		} else {
			after[C] = 0
		}
	case "gtrr":
		if before[A] > before[B] {
			after[C] = 1
		} else {
			after[C] = 0
		}
	case "eqir":
		if A == before[B] {
			after[C] = 1
		} else {
			after[C] = 0
		}
	case "eqri":
		if before[A] == B {
			after[C] = 1
		} else {
			after[C] = 0
		}
	case "eqrr":
		if before[A] == before[B] {
			after[C] = 1
		} else {
			after[C] = 0
		}
	}
	return after
}

func registersEquality(before, after registers) bool {
	if len(before) != len(after) {
		return false
	}
	for i := 0; i < len(before); i++ {
		if before[i] != after[i] {
			return false
		}
	}
	return true
}

func removePossibility(l []string, opcode string) []string {
	res := []string{}
	for _, x := range l {
		if x != opcode {
			res = append(res, x)
		}
	}
	return res
}

func foundLanguage(p [][]string) bool {
	for _, opcode := range p {
		if len(opcode) != 1 {
			return false
		}
	}
	return true
}

func languageIsIncorrect(p [][]string) bool {
	for _, opcode := range p {
		if len(opcode) == 0 {
			return true
		}
	}
	return false
}

func part1(s string) int {
	c := 0
	samples, _ := format1(s)
	for _, sample := range samples {
		temp := 0
		for _, opcode := range []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "gtir", "gtrr", "gtri", "eqir", "eqri", "eqrr", "setr", "seti"} {
			result := nextRegisters(sample.before, sample.instr, opcode)
			if registersEquality(sample.after, result) {
				temp++
			}
		}
		if temp >= 3 {
			c++
		}
	}
	return c
}

func part2(s string) int {
	samples, instrList := format1(s)
	opcodeList := []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "gtir", "gtrr", "gtri", "eqir", "eqri", "eqrr", "setr", "seti"}
	possibilities := make([][]string, 16)
	var language [][]string
	var deque deque.Deque[possibility]
	for i := 0; i < 16; i++ {
		possibilities[i] = opcodeList
	}
	for _, sample := range samples {
		for _, opcode := range opcodeList {
			result := nextRegisters(sample.before, sample.instr, opcode)
			if !registersEquality(sample.after, result) {
				possibilities[sample.instr[0]] = removePossibility(possibilities[sample.instr[0]], opcode)
			}
		}
	}
	deque.PushBack(possibility{possibilities, 0})
	for deque.Len() != 0 {
		curr := deque.PopBack()
		if foundLanguage(curr.table) {
			language = curr.table
			break
		}
		i := curr.index
		for _, opcode := range curr.table[i] {
			next := possibility{make([][]string, 16), i + 1}
			next.table[i] = []string{opcode}
			for j := range next.table {
				if i != j {
					next.table[j] = append(next.table[j], curr.table[j]...)
					next.table[j] = removePossibility(next.table[j], opcode)
				}
			}
			if !languageIsIncorrect(next.table) {
				deque.PushBack(next)
			}
		}
	}
	registers := registers([]int{0, 0, 0, 0})
	for _, instr := range instrList {
		registers = nextRegisters(registers, instr, language[instr[0]][0])
	}
	return registers[0]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day16/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day16/input.txt")

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
