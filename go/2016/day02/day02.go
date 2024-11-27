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
	return strings.Split(s, "\n")
}

func next(i int, r rune) int {
	switch r {
	case 'U':
		if i > 3 {
			return i - 3
		} else {
			return i
		}
	case 'D':
		if i > 6 {
			return i
		} else {
			return i + 3
		}
	case 'L':
		if (i % 3) == 1 {
			return i
		} else {
			return i - 1
		}
	case 'R':
		if (i % 3) == 0 {
			return i
		} else {
			return i + 1
		}
	}
	return 0
}

func part1(s string) int {
	lines := format(s)
	str := ""
	n := 5
	for _, line := range lines {
		for _, char := range line {
			n = next(n, char)
		}
		str += strconv.Itoa(n)
	}
	r, _ := strconv.Atoi(str)
	return r
}

func next2(s string, r rune) string {
	switch s {
	case "1":
		if r == 'D' {
			return "3"
		} else {
			return "1"
		}
	case "2":
		if r == 'D' {
			return "6"
		}
		if r == 'R' {
			return "3"
		} else {
			return "2"
		}
	case "3":
		switch r {
		case 'U':
			return "1"
		case 'D':
			return "7"
		case 'R':
			return "4"
		case 'L':
			return "1"
		}
	case "4":
		if r == 'L' {
			return "3"
		}
		if r == 'D' {
			return "8"
		} else {
			return "4"
		}
	case "5":
		if r == 'R' {
			return "6"
		} else {
			return "5"
		}
	case "6":
		switch r {
		case 'U':
			return "2"
		case 'D':
			return "A"
		case 'R':
			return "7"
		case 'L':
			return "5"
		}
	case "7":
		switch r {
		case 'U':
			return "3"
		case 'D':
			return "B"
		case 'R':
			return "8"
		case 'L':
			return "6"
		}
	case "8":
		switch r {
		case 'U':
			return "4"
		case 'D':
			return "C"
		case 'R':
			return "9"
		case 'L':
			return "7"
		}
	case "9":
		if r == 'L' {
			return "8"
		} else {
			return "9"
		}
	case "D":
		if r == 'U' {
			return "B"
		} else {
			return "D"
		}
	case "A":
		if r == 'U' {
			return "6"
		}
		if r == 'R' {
			return "B"
		} else {
			return "A"
		}
	case "B":
		switch r {
		case 'U':
			return "7"
		case 'D':
			return "D"
		case 'R':
			return "C"
		case 'L':
			return "A"
		}
	case "C":
		if r == 'L' {
			return "B"
		}
		if r == 'U' {
			return "8"
		} else {
			return "C"
		}
	}
	return "0"
}

func part2(s string) string {
	lines := format(s)
	str := ""
	n := "5"
	for _, line := range lines {
		for _, char := range line {
			n = next2(n, char)
		}
		str += n
	}
	return str
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day02/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2016/day02/input.txt")

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
