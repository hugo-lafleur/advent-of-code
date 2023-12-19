package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type chips []int

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(s string) int {
	instr := format(s)
	bots := make(map[int]chips)
	outputs := make(map[int]chips)
	for _, line := range instr {
		if line[0] == "value" {
			value, _ := strconv.Atoi(line[1])
			bot, _ := strconv.Atoi(line[5])
			bots[bot] = append(bots[bot], value)
		}
	}
	for i := 0; i < len(instr); i++ {
		for _, line := range instr {
			if line[0] == "bot" {
				bot, _ := strconv.Atoi(line[1])
				if len(bots[bot]) == 2 {
					lowTo, _ := strconv.Atoi(line[6])
					highTo, _ := strconv.Atoi(line[11])
					a, b := bots[bot][0], bots[bot][1]
					if min(a, b) == 17 && max(a, b) == 61 {
						return bot
					}
					if line[5] == "bot" {
						bots[lowTo] = append(bots[lowTo], min(a, b))
					} else {
						outputs[lowTo] = append(outputs[lowTo], min(a, b))
					}
					if line[10] == "bot" {
						bots[highTo] = append(bots[highTo], max(a, b))
					} else {
						outputs[highTo] = append(outputs[highTo], max(a, b))
					}
					bots[bot] = []int{}
				}
			}
		}
	}
	return 0
}

func part2(s string) int {
	instr := format(s)
	bots := make(map[int]chips)
	outputs := make(map[int]chips)
	for _, line := range instr {
		if line[0] == "value" {
			value, _ := strconv.Atoi(line[1])
			bot, _ := strconv.Atoi(line[5])
			bots[bot] = append(bots[bot], value)
		}
	}
	for i := 0; i < len(instr); i++ {
		for _, line := range instr {
			if line[0] == "bot" {
				bot, _ := strconv.Atoi(line[1])
				if len(bots[bot]) == 2 {
					lowTo, _ := strconv.Atoi(line[6])
					highTo, _ := strconv.Atoi(line[11])
					a, b := bots[bot][0], bots[bot][1]
					if line[5] == "bot" {
						bots[lowTo] = append(bots[lowTo], min(a, b))
					} else {
						outputs[lowTo] = append(outputs[lowTo], min(a, b))
					}
					if line[10] == "bot" {
						bots[highTo] = append(bots[highTo], max(a, b))
					} else {
						outputs[highTo] = append(outputs[highTo], max(a, b))
					}
					bots[bot] = []int{}
				}
			}
		}
	}
	return outputs[0][0] * outputs[1][0] * outputs[2][0]
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
