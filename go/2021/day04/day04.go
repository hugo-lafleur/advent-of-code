package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type board [][]int

func format(s string) ([]int, []board) {
	parts := strings.Split(s, "\n\n")
	boards := []board{}
	list := parts[0]
	listInt := []int{}
	listSplit := strings.Split(list, ",")
	for _, x := range listSplit {
		n, _ := strconv.Atoi(x)
		listInt = append(listInt, n)
	}
	for i := 1; i < len(parts); i++ {
		var b board
		boardLines := strings.Split(parts[i], "\n")
		for _, line := range boardLines {
			lineBoard := []int{}
			lineSplit := strings.Fields(line)
			for _, x := range lineSplit {
				n, _ := strconv.Atoi(x)
				lineBoard = append(lineBoard, n)
			}
			b = append(b, lineBoard)
		}
		boards = append(boards, b)
	}
	return listInt, boards
}

func checkColumn(b board, i int) bool {
	for j := 0; j < len(b); j++ {
		if b[j][i] != -1 {
			return false
		}
	}
	return true
}

func checkRow(b board, i int) bool {
	for j := 0; j < len(b); j++ {
		if b[i][j] != -1 {
			return false
		}
	}
	return true
}

func check(b board) bool {
	for i := 0; i < len(b); i++ {
		if checkColumn(b, i) || checkRow(b, i) {
			return true
		}
	}
	return false
}

func score(b board) int {
	res := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] != -1 {
				res += b[i][j]
			}
		}
	}
	return res
}

func part1(s string) int {
	list, boards := format(s)
	for _, n := range list {
		for index, b := range boards {
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(b[i]); j++ {
					if b[i][j] == n {
						boards[index][i][j] = -1
					}
				}
			}
		}
		for _, b := range boards {
			if check(b) {
				return n * score(b)
			}
		}
	}
	return 0
}

func part2(s string) int {
	list, boards := format(s)
	var res int
	done := make(map[int]bool)
	for _, n := range list {
		for index, b := range boards {
			for i := 0; i < len(b); i++ {
				for j := 0; j < len(b[i]); j++ {
					if b[i][j] == n {
						boards[index][i][j] = -1
					}
				}
			}
		}
		for i, b := range boards {
			_, ok := done[i]
			if check(b) && !ok {
				res = n * score(b)
				done[i] = true
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day04/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day04/input.txt")

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
