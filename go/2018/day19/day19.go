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
	operator string
	A, B, C  int
}

type registers []int

func format(s string) (int, []instruction) {
	lines := strings.Split(s, "\n")
	strs := [][]string{}
	for _, line := range lines {
		strs = append(strs, strings.Split(line, " "))
	}
	instrList := []instruction{}
	pointerRegister, _ := strconv.Atoi(strs[0][1])
	for i := 1; i < len(strs); i++ {
		a, _ := strconv.Atoi(strs[i][1])
		b, _ := strconv.Atoi(strs[i][2])
		c, _ := strconv.Atoi(strs[i][3])
		instrList = append(instrList, instruction{strs[i][0], a, b, c})
	}
	return pointerRegister, instrList
}

func nextRegisters(before registers, instr instruction) registers {
	after := make(registers, len(before))
	copy(after, before)
	operator := instr.operator
	A := instr.A
	B := instr.B
	C := instr.C
	switch operator {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(s string) int {
	reg := registers([]int{0, 0, 0, 0, 0, 0})
	pointerRegister, instructionList := format(s)
	ip := 0
	for ip < len(instructionList) && ip >= 0 {
		instr := instructionList[ip]
		reg[pointerRegister] = ip
		reg = nextRegisters(reg, instr)
		if ip >= len(instructionList) && ip < 0 {
			break
		}
		ip = reg[pointerRegister] + 1
	}
	return reg[0]
}

func part2(s string) int {
	reg := registers([]int{1, 0, 0, 0, 0, 0})
	pointerRegister, instructionList := format(s)
	limit := 0
	ip := 0
	for ip < len(instructionList) && ip >= 0 {
		instr := instructionList[ip]
		reg[pointerRegister] = ip
		if ip < len(instructionList)-1 {
			nextInstr := instructionList[ip+1]
			if instr.operator == "mulr" && nextInstr.operator == "eqrr" {
				limit = max(reg[nextInstr.A], reg[nextInstr.B])
				break
			}
		}
		reg = nextRegisters(reg, instr)
		if ip >= len(instructionList) && ip < 0 {
			break
		}
		ip = reg[pointerRegister] + 1

	}
	res := 0
	for i := 1; i <= limit; i++ {
		if (10551364 % i) == 0 {
			res += i
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day19/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day19/input.data")

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
