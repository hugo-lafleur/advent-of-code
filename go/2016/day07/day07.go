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
	res := [][]string{}
	for _, line := range lines {
		firstPart := strings.Split(line, "[")
		secondPart := []string{}
		for _, part := range firstPart {
			secondPart = append(secondPart, strings.Split(part, "]")...)
		}
		formatLine := []string{}
		formatLine = append(formatLine, secondPart...)
		res = append(res, formatLine)
	}
	return res
}

func TLS(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

func SSL(s string) []string {
	res := []string{}
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			res = append(res, s[i:i+3])
		}
	}
	return res
}

func ABAtoBAB(ABA []string, BAB []string) bool {
	for _, aba := range ABA {
		for _, bab := range BAB {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}
	return false
}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, line := range list {
		res := false
		for i := 0; i < len(line); i++ {
			if i%2 == 0 {
				res = res || TLS(line[i])
			} else {
				if TLS(line[i]) {
					res = false
					break
				}
			}
		}
		if res {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for _, line := range list {
		ABA := []string{}
		BAB := []string{}
		for i := 0; i < len(line); i++ {
			if i%2 == 0 {
				ABA = append(ABA, SSL(line[i])...)
			} else {
				BAB = append(BAB, SSL(line[i])...)
			}
		}
		if ABAtoBAB(ABA, BAB) {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day07/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))

	content, err = os.ReadFile("../../../inputs/2016/day07/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day07/input.txt")

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
