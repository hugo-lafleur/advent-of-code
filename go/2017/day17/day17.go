package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func insert(index int, n int, l []int) []int {
	res := []int{}
	for i := 0; i < len(l); i++ {
		res = append(res, l[i])
		if i == index {
			res = append(res, n)
		}
	}
	return res
}

func part1(s string) int {
	steps := format(s)
	list := []int{0}
	index := 0
	for i := 1; i < 2018; i++ {
		l := len(list)
		index = (index + steps) % l
		list = insert(index, i, list)
		index++
	}
	return list[(index+1)%len(list)]
}

func part2(s string) int {
	steps := format(s)
	index := 0
	var indexOne int
	for i := 1; i < 50000000+1; i++ {
		index = (index+steps)%i + 1
		if index == 1 {
			indexOne = i
		}
	}
	return indexOne
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day17/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day17/input.txt")

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
