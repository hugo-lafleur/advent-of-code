package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == '/' || r == ','
}

func format(s string) [][]string {
	instrs := strings.Split(s, ",")
	res := [][]string{}
	for _, instr := range instrs {
		switch instr[0] {
		case byte('s'):
			res = append(res, strings.SplitAfter(instr, "s"))
		case byte('x'):
			temp := []string{"x"}
			temp = append(temp, strings.Split(instr[1:], "/")...)
			res = append(res, temp)
		case byte('p'):
			temp := []string{"p"}
			temp = append(temp, strings.Split(instr[1:], "/")...)
			res = append(res, temp)
		}
	}
	return res
}

func spin(s string, n int) string {
	if n == 0 {
		return s
	}
	l := len(s)
	return spin(string(s[l-1])+s[:(l-1)], n-1)
}

func exchange(s string, a, b string) string {
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	res := ""
	for i, r := range s {
		switch i {
		case n:
			res += string(s[m])
		case m:
			res += string(s[n])
		default:
			res += string(r)
		}
	}
	return res
}

func partner(s string, a, b string) string {
	res := ""
	for _, r := range s {
		switch string(r) {
		case a:
			res += b
		case b:
			res += a
		default:
			res += string(r)
		}
	}
	return res
}

func part1(s string) string {
	instrs := format(s)
	var curr string
	if len(instrs) == 3 {
		curr = "abcde"
	} else {
		curr = "abcdefghijklmnop"
	}
	for _, instr := range instrs {
		switch instr[0] {
		case "s":
			n, _ := strconv.Atoi(instr[1])
			curr = spin(curr, n)
		case "x":
			curr = exchange(curr, instr[1], instr[2])
		case "p":
			curr = partner(curr, instr[1], instr[2])
		}
	}
	return curr
}

func part2(s string) string {
	instrs := format(s)
	var curr string
	memory := make(map[string]int)
	k := -1
	if len(instrs) == 3 {
		curr = "abcde"
	} else {
		curr = "abcdefghijklmnop"
	}
	memory[curr] = 0
	for i := 1; true; i++ {
		for _, instr := range instrs {
			switch instr[0] {
			case "s":
				n, _ := strconv.Atoi(instr[1])
				curr = spin(curr, n)
			case "x":
				curr = exchange(curr, instr[1], instr[2])
			case "p":
				curr = partner(curr, instr[1], instr[2])
			}
		}
		for key, value := range memory {
			if key == curr {
				c := (i - value)
				k = (1000000000 % c)
				break
			}
		}
		for key, value := range memory {
			if value == k {
				return key
			}
		}
		memory[curr] = i
	}
	return curr
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day16/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day16/input.txt")

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
