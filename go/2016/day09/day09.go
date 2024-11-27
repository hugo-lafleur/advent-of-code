package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "")
}

func eval(input []string) ([]string, int) {
	i := 1
	a := input[i]
	i++
	for input[i] != "x" {
		a += input[i]
		i++
	}
	i++
	b := ""
	for input[i] != ")" {
		b += input[i]
		i++
	}
	n, _ := strconv.Atoi(a)
	m, _ := strconv.Atoi(b)
	i++
	str := []string{}
	for j := 0; j < n; j++ {
		str = append(str, input[i+j])
	}
	res := []string{}
	for j := 0; j < m; j++ {
		res = append(res, str...)
	}
	for j := i + n; j < len(input); j++ {
		res = append(res, input[j])
	}
	return res, n*m - 1
}

func sum(t []int) int {
	res := 0
	for _, x := range t {
		res += x
	}
	return res
}

func part1(s string) int {
	input := format(s)
	res := []string{}
	i := 0
	for i < len(input) {
		c := input[i]
		if c != "(" {
			res = append(res, c)
		}
		if c == "(" {
			input, i = eval(input[i:])
			res = append(res, input[:i+1]...)
		}
		i++
	}
	return len(res)
}

func part2(s string) int {
	input := format(s)
	res := []int{}
	for i := 0; i < len(input); i++ {
		res = append(res, 1)
	}
	i := 0
	for i < len(input) {
		c := input[i]
		if c == "(" {
			i++
			a := input[i]
			res[i] = 0
			res[i-1] = 0
			i++
			for input[i] != "x" {
				a += input[i]
				res[i] = 0
				i++
			}
			i++
			res[i] = 0
			res[i-1] = 0
			b := ""
			for input[i] != ")" {
				b += input[i]
				res[i] = 0
				i++
			}
			res[i] = 0
			n, _ := strconv.Atoi(a)
			m, _ := strconv.Atoi(b)
			for j := 1; j < n+1; j++ {
				res[i+j] = res[i+j] * m
			}
		}
		i++
	}
	return sum(res)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day09/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2016/day09/input.txt")

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
