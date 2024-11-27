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
	str := strings.Split(s, "")
	tab := []int{}
	for _, x := range str {
		n, _ := strconv.Atoi(x)
		tab = append(tab, n)
	}
	return tab
}

func next(tab []int) []int {
	c := 1
	res := []int{}
	for i, x := range tab {
		if i == len(tab)-1 {
			if c > 9 {
				s := strconv.Itoa(c)
				l := strings.Split(s, "")
				for _, x := range l {
					n, _ := strconv.Atoi(x)
					tab = append(tab, n)
				}
			} else {
				res = append(res, c)
			}
			res = append(res, x)
			return res
		}
		if tab[i+1] == x {
			c++
		}
		if tab[i+1] != x {
			if c > 9 {
				s := strconv.Itoa(c)
				l := strings.Split(s, "")
				for _, x := range l {
					n, _ := strconv.Atoi(x)
					tab = append(tab, n)
				}
			} else {
				res = append(res, c)
			}
			res = append(res, x)
			c = 1
		}
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	for i := 0; i < 40; i++ {
		tab = next(tab)
	}
	return len(tab)
}

func part2(s string) int {
	tab := format(s)
	for i := 0; i < 50; i++ {
		tab = next(tab)
	}
	return len(tab)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day10/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2015/day10/input.txt")

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
