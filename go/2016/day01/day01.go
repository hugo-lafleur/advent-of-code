package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func format(s string) []string {
	line := strings.Split(s, "")
	res := []string{}
	i := 0
	for i < len(line) {
		char := line[i]
		if char == "R" || char == "L" {
			res = append(res, char)
		}
		if isNumber(char) {
			n := char
			k := i + 1
			for k < len(line) && isNumber(line[k]) {
				n += line[k]
				k++
			}
			res = append(res, n)
			i = k
		}
		i++
	}
	return res
}

func part1(s string) int {
	i := 0
	d := "n"
	path := format(s)
	a, b := 0, 0
	for i < len(path)-1 {
		n, _ := strconv.Atoi(path[i+1])
		if path[i] == "R" {
			switch d {
			case "n":
				d = "e"
			case "e":
				d = "s"
			case "s":
				d = "w"
			case "w":
				d = "n"
			}

		}
		if path[i] == "L" {
			switch d {
			case "n":
				d = "w"
			case "e":
				d = "n"
			case "s":
				d = "e"
			case "w":
				d = "s"
			}

		}
		switch d {
		case "n":
			b += n
		case "e":
			a += n
		case "s":
			b -= n
		case "w":
			a -= n
		}
		i += 2

	}
	return abs(a) + abs(b)
}

func part2(s string) int {
	i := 0
	d := "n"
	path := format(s)
	a, b := 0, 0
	knownA := []int{}
	knownB := []int{}
	for i < len(path)-1 {
		n, _ := strconv.Atoi(path[i+1])
		if path[i] == "R" {
			switch d {
			case "n":
				d = "e"
			case "e":
				d = "s"
			case "s":
				d = "w"
			case "w":
				d = "n"
			}

		}
		if path[i] == "L" {
			switch d {
			case "n":
				d = "w"
			case "e":
				d = "n"
			case "s":
				d = "e"
			case "w":
				d = "s"
			}

		}
		switch d {
		case "n":
			k := 0
			for k < n {
				b++
				knownA = append(knownA, a)
				knownB = append(knownB, b)
				j := 0
				for j < len(knownA)-1 {
					if knownA[j] == a && knownB[j] == b {
						return abs(a) + abs(b)
					}
					j++
				}
				k++
			}
		case "e":
			k := 0
			for k < n {
				a++
				knownA = append(knownA, a)
				knownB = append(knownB, b)
				j := 0
				for j < len(knownA)-1 {
					if knownA[j] == a && knownB[j] == b {
						return abs(a) + abs(b)
					}
					j++
				}
				k++
			}
		case "s":
			k := 0
			for k < n {
				b--
				knownA = append(knownA, a)
				knownB = append(knownB, b)
				j := 0
				for j < len(knownA)-1 {
					if knownA[j] == a && knownB[j] == b {
						return abs(a) + abs(b)
					}
					j++
				}
				k++
			}
		case "w":
			k := 0
			for k < n {
				a--
				knownA = append(knownA, a)
				knownB = append(knownB, b)
				j := 0
				for j < len(knownA)-1 {
					if knownA[j] == a && knownB[j] == b {
						return abs(a) + abs(b)
					}
					j++
				}
				k++
			}
		}
		i += 2

	}
	return 0
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
