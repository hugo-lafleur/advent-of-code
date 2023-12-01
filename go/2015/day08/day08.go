package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func format(s string) []string {
	return strings.Split(s, "\n")

}

func is_hexa(b byte) bool {
	tab := [16]string{"a", "b", "c", "d", "e", "f", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, str := range tab {
		if string(b) == str {
			return true
		}
	}
	return false
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
					fmt.Printf("%d : %q\n", i, c)
					m++
					i++
					break
				}
				if line[i+1] == '\\' {
					fmt.Printf("%d : %q\n", i, c)
					m++
					i++
					break
				}
				if line[i+1] == 'x' /*&& is_hexa(line[i+2]) && is_hexa(line[i+3])*/ {
					fmt.Printf("%d : %q\n", i, c)
					m++
					i = i + 3
					break
				}
			case '"':
			default:
				fmt.Printf("%d : %q\n", i, c)
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
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
