package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "\n")

}

func part1(s string) int {
	tab := format(s)
	code := 0
	m := 0
	for _, line := range tab {
		n := len(line)
		code += n
		i := 0
		for i < n {
			c := line[i]
			switch c {
			case '\\':
				if line[i+1] == '"' {
					//fmt.Printf("%d : %q\n", i, c)
					m++
					i++
					break
				}
				if line[i+1] == '\\' {
					//fmt.Printf("%d : %q\n", i, c)
					m++
					i++
					break
				}
				if line[i+1] == 'x' /*&& is_hexa(line[i+2]) && is_hexa(line[i+3])*/ {
					//fmt.Printf("%d : %q\n", i, c)
					m++
					i = i + 3
				}
			case '"':
			default:
				//fmt.Printf("%d : %q\n", i, c)
				m++
			}
			i++
		}

	}
	return code - m
}

func part2(s string) int {
	tab := format(s)
	code := 0
	m := 0
	for _, line := range tab {
		m = m + 2
		n := len(line)
		code += n
		i := 0
		for i < n {
			c := line[i]
			switch c {
			case '"':
				m = m + 2
			case '\\':
				m = m + 2
			default:
				m++
			}
			i++
		}

	}
	return m - code
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day08/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2015/day08/input.txt")

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
