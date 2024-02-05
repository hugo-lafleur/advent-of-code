package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type policy struct {
	min      int
	max      int
	letter   string
	password string
}

func Split(r rune) bool {
	return r == '-' || r == ':' || r == ' '
}

func format(s string) []policy {
	lines := strings.Split(s, "\n")
	res := []policy{}
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		a, _ := strconv.Atoi(lineSplit[0])
		b, _ := strconv.Atoi(lineSplit[1])
		pol := policy{a, b, lineSplit[2], lineSplit[3]}
		res = append(res, pol)
	}
	return res
}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, pol := range list {
		n := strings.Count(pol.password, pol.letter)
		if n >= pol.min && n <= pol.max {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for _, pol := range list {
		k := 0
		for _, i := range []int{pol.min, pol.max} {
			if string(pol.password[i-1]) == pol.letter {
				k++
			}
		}
		if k == 1 {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day02/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day02/input.data")

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
