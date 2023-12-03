package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	return tab
}

func part1(s string) int {
	tab := format(s)
	res := 0
	tmp := []int{}
	reg := 1
	for _, line := range tab {
		if line[0] == "noop" {
			tmp = append(tmp, 0)
		} else {
			n, _ := strconv.Atoi(line[1])
			tmp = append(tmp, 0)
			tmp = append(tmp, n)
		}

	}
	i := 1
	for i < len(tmp) {
		reg += tmp[i-1]
		y := i + 1
		if y == 20 || y == 60 || y == 100 || y == 140 || y == 180 || y == 220 {
			res += y * reg
		}
		i++
	}
	return res
}

func part2(s string) [6][40]string {
	tab := format(s)
	tmp := []int{}
	reg := 1
	crt := [6][40]string{}
	for _, line := range tab {
		if line[0] == "noop" {
			tmp = append(tmp, 0)
		} else {
			n, _ := strconv.Atoi(line[1])
			tmp = append(tmp, 0)
			tmp = append(tmp, n)
		}

	}
	j := 0
	i := 1
	for i < len(tmp) {
		y := i
		if y == 41 || y == 81 || y == 121 || y == 161 || y == 201 {
			j++
		}
		a := reg - 1
		b := reg + 1
		x := i - 1 - 40*j
		if x == a || x == b || x == reg {

			crt[j][x] = "#"
		} else {
			crt[j][x] = "."
		}
		reg += tmp[i-1]
		i++
	}
	return crt
}

func printRes(res [6][40]string) {
	for _, line := range res {
		fmt.Printf("%v\n", line)
	}
}

func main() {
	content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 :\n")
	printRes(part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 :\n")
	printRes(part2(string(content)))
	fmt.Println(time.Since(start))
}
