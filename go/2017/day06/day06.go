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
	return r == ' ' || r == '	'
}

func format(s string) []int {
	temp := strings.FieldsFunc(s, Split)
	res := []int{}
	for _, s := range temp {
		n, _ := strconv.Atoi(s)
		res = append(res, n)
	}
	return res
}

func banksToString(b []int) string {
	res := ""
	for _, x := range b {
		res += strconv.Itoa(x) + ";"
	}
	return res
}

func maxBlocks(b []int) int {
	index := 0
	max := 0
	for i, x := range b {
		if x > max {
			max = x
			index = i
		}
	}
	return index
}

func part1(s string) int {
	c := 0
	banks := format(s)
	seen := make(map[string]bool)
	for {
		_, ok := seen[banksToString(banks)]
		if ok {
			return c
		}
		seen[banksToString(banks)] = true
		i := maxBlocks(banks)
		toDistribute := banks[i]
		banks[i] = 0
		for j := 1; j <= toDistribute; j++ {
			banks[(i+j)%len(banks)]++
		}
		c++
	}
}

func part2(s string) int {
	c := 0
	banks := format(s)
	seen := make(map[string]int)
	for {
		loop, ok := seen[banksToString(banks)]
		if ok {
			return c - loop
		}
		seen[banksToString(banks)] = c
		i := maxBlocks(banks)
		toDistribute := banks[i]
		banks[i] = 0
		for j := 1; j <= toDistribute; j++ {
			banks[(i+j)%len(banks)]++
		}
		c++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day06/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day06/input.data")

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
