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

func isIn(n int, l []int) bool {
	for _, x := range l {
		if x == n {
			return true
		}
	}
	return false
}

func solve(s string) []int {
	reg := registers([]int{0, 0, 0, 0, 0, 0})
	pointerRegister, instructionList := format(s)
	ip := 0
	halt := []int{}
	for ip < len(instructionList) && ip >= 0 {
		instr := instructionList[ip]
		reg[pointerRegister] = ip
		if ip < len(instructionList)-2 {
			next := instructionList[ip+1]
			nextNext := instructionList[ip+2]
			if instr.operator == "addi" && next.operator == "muli" && nextNext.operator == "gtrr" && instr.B == 1 && instr.C == next.A && next.A == next.C && next.A == nextNext.A {
				reg[instr.A] = (reg[nextNext.B] / next.B)
			}
		}
		reg = nextRegisters(reg, instr)
		if instr.operator == "eqrr" && (instr.A == 0 || instr.B == 0) {
			var new int
			if instr.A == 0 {
				new = reg[instr.B]
			} else {
				new = reg[instr.A]
			}
			if isIn(new, halt) {
				fmt.Println(new)
				return halt
			}
			halt = append(halt, new)
		}
		if ip >= len(instructionList) && ip < 0 {
			break
		}
		ip = reg[pointerRegister] + 1
	}
	return halt
}

func part1(s string) int {
	return solve(s)[0]
}

func part2(s string) int {
	res := solve(s)
	return res[len(res)-1]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day21/input.txt")

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
