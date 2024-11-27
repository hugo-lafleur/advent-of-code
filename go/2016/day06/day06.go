package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func maxDict(dict map[string]int) string {
	max := 0
	var maxStr string
	for key, value := range dict {
		if value > max {
			max = value
			maxStr = key
		}
	}
	return maxStr
}

func minDict(dict map[string]int) string {
	min := 100000000
	var minStr string
	for key, value := range dict {
		if value < min {
			min = value
			minStr = key
		}
	}
	return minStr
}

func part1(s string) string {
	res := ""
	tab := format(s)
	for j := 0; j < len(tab[0]); j++ {
		occ := make(map[string]int)
		for i := 0; i < len(tab); i++ {
			occ[tab[i][j]]++
		}
		res += maxDict(occ)
	}
	return res
}

func part2(s string) string {
	res := ""
	tab := format(s)
	for j := 0; j < len(tab[0]); j++ {
		occ := make(map[string]int)
		for i := 0; i < len(tab); i++ {
			occ[tab[i][j]]++
		}
		res += minDict(occ)
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2016/day06/input.txt")

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
