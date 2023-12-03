package main

import (
	"fmt"
	"log"
	"os"
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

func shape(c string) int {
	switch c {
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}
	return 0

}

func outcome(a string, b string) int {
	switch a {
	case "A":
		switch b {
		case "X":
			return 3
		case "Y":
			return 6
		case "Z":
			return 0
		}
	case "B":
		switch b {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
	case "C":
		switch b {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
		}
	}
	return 0
}

func outcome_2(c string) int {
	switch c {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	}
	return 0
}

func shape_2(a string, b string) int {
	switch a {
	case "A":
		switch b {
		case "X":
			return 3
		case "Y":
			return 1
		case "Z":
			return 2
		}
	case "B":
		switch b {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		}
	case "C":
		switch b {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		}
	}
	return 0
}

func part1(s string) int {
	tab := format(s)
	sum := 0
	for _, line := range tab {
		sum += shape(line[1]) + outcome(line[0], line[1])
	}
	return sum
}

func part2(s string) int {
	tab := format(s)
	sum := 0
	for _, line := range tab {
		sum += shape_2(line[0], line[1]) + outcome_2(line[1])
	}
	return sum
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
