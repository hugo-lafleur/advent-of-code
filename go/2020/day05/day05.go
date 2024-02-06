package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
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

func part1(s string) int {
	var max int
	list := format(s)
	for _, seat := range list {
		F, B := 0, 127
		for i := 0; i < 7; i++ {
			m := (F + B) / 2
			if seat[i] == "F" {
				B = m
			} else {
				F = m + 1
			}
		}
		row := F
		L, R := 0, 7
		for i := 7; i < len(seat); i++ {
			m := (L + R) / 2
			if seat[i] == "L" {
				R = m
			} else {
				L = m + 1
			}
		}
		colum := L
		ID := row*8 + colum
		if ID > max {
			max = ID
		}
	}
	return max
}

func part2(s string) int {
	IDs := []int{}
	list := format(s)
	for _, seat := range list {
		F, B := 0, 127
		for i := 0; i < 7; i++ {
			m := (F + B) / 2
			if seat[i] == "F" {
				B = m
			} else {
				F = m + 1
			}
		}
		row := F
		L, R := 0, 7
		for i := 7; i < len(seat); i++ {
			m := (L + R) / 2
			if seat[i] == "L" {
				R = m
			} else {
				L = m + 1
			}
		}
		colum := L
		ID := row*8 + colum
		IDs = append(IDs, ID)
	}
	sort.Ints(IDs)
	i := IDs[0]
	for {
		if !slices.Contains(IDs, i) {
			return i
		}
		i++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day05/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day05/input.data")

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
