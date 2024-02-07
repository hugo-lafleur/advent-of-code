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
	return r == ' ' || r == '[' || r == ']'
}

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func sumMap(m map[int]int) int {
	res := 0
	for _, value := range m {
		res += value
	}
	return res
}

func applyMask(n int, mask map[int]string) int {
	bin := fmt.Sprintf("%036b", n)
	newBin := ""
	for i := range bin {
		change, ok := mask[i]
		if ok {
			newBin += change
		} else {
			newBin += string(bin[i])
		}
	}
	n64, _ := strconv.ParseInt(newBin, 2, 64)
	return int(n64)
}

func intToBin(n int) string {
	return fmt.Sprintf("%036b", n)
}

func binToInt(s string) int {
	n64, _ := strconv.ParseInt(s, 2, 64)
	return int(n64)
}

func allAddresses(bin string, mask map[int]string) []string {
	res := []string{""}
	for i := 0; i < 36; i++ {
		change, ok := mask[i]
		if ok {
			for j := range res {
				if change == "0" {
					res[j] += string(bin[i])
				}
				if change == "1" {
					res[j] += "1"
				}
			}
		} else {
			for j := range res {
				res = append(res, res[j]+"1")
				res[j] += "0"
			}
		}
	}
	return res
}

func part1(s string) int {
	list := format(s)
	var mask map[int]string
	mem := make(map[int]int)
	for _, instr := range list {
		if instr[0] == "mask" {
			mask = make(map[int]string)
			for i, r := range instr[2] {
				if r == '0' {
					mask[i] = "0"
				}
				if r == '1' {
					mask[i] = "1"
				}
			}
		}
		if instr[0] == "mem" {
			n, _ := strconv.Atoi(instr[3])
			addr, _ := strconv.Atoi(instr[1])
			mem[addr] = int(applyMask(n, mask))
		}
	}
	return sumMap(mem)
}

func part2(s string) int {
	list := format(s)
	var mask map[int]string
	mem := make(map[int]int)
	for _, instr := range list {
		if instr[0] == "mask" {
			mask = make(map[int]string)
			for i, r := range instr[2] {
				if r == '0' {
					mask[i] = "0"
				}
				if r == '1' {
					mask[i] = "1"
				}
			}
		}
		if instr[0] == "mem" {
			addr, _ := strconv.Atoi(instr[1])
			bin := intToBin(addr)
			n, _ := strconv.Atoi(instr[3])
			for _, binary := range allAddresses(bin, mask) {
				mem[binToInt(binary)] = n
			}
		}
	}
	return sumMap(mem)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day14/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day14/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day14/input.data")

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
