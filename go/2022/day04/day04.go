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
		tmp1 := strings.Split(line, "-")    //tmp1 = [2 4/6 8]
		tmp2 := strings.Split(tmp1[1], ",") //tmp2 = [4 6]
		tmp3 := []string{}
		tmp3 = append(tmp3, tmp1[0])
		tmp3 = append(tmp3, tmp2[0])
		tmp3 = append(tmp3, tmp2[1])
		tmp3 = append(tmp3, tmp1[2])
		tab = append(tab, tmp3)
	}
	return tab
}
func check1(s []string) bool {
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	c, _ := strconv.Atoi(s[2])
	d, _ := strconv.Atoi(s[3])
	if c >= a && d <= b {
		return true
	}
	if a >= c && b <= d {
		return true
	}
	return false
}

func check2(s []string) bool {
	a, _ := strconv.Atoi(s[0])
	b, _ := strconv.Atoi(s[1])
	c, _ := strconv.Atoi(s[2])
	d, _ := strconv.Atoi(s[3])
	if b < c || d < a {
		return false
	}
	return true
}

func part1(s string) int {
	tab := format(s)
	c := 0
	for _, line := range tab {
		if check1(line) {
			c++
		}
	}
	return c
}

func part2(s string) int {
	tab := format(s)
	c := 0
	for _, line := range tab {
		if check2(line) {
			c++
		}
	}
	return c
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
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
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
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
