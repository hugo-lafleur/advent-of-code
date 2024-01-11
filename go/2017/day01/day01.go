package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []int {
	strs := strings.Split(s, "")
	res := []int{}
	for _, s := range strs {
		n, _ := strconv.Atoi(s)
		res = append(res, n)
	}
	return res
}

func part1(s string) int {
	c := 0
	list := format(s)
	for i := 0; i < len(list); i++ {
		if list[i] == list[(i+1)%(len(list))] {
			c += list[i]
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for i := 0; i < len(list); i++ {
		if list[i] == list[(i+(len(list)/2))%(len(list))] {
			c += list[i]
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day01/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day01/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day01/input.data")

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
