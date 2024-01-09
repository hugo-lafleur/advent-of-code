package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func dragonCurve(s string) string {
	a := s
	b := []byte{}
	for i := len(a) - 1; i > -1; i-- {
		c := a[i]
		if c == '1' {
			b = append(b, byte('0'))
		} else {
			b = append(b, byte('1'))
		}
	}
	return a + "0" + string(b)
}

func pairs(s string) string {
	res := []byte{}
	for i := 0; i < len(s)-1; i = i + 2 {
		if s[i] == s[i+1] {
			res = append(res, byte('1'))
		} else {
			res = append(res, byte('0'))
		}
	}
	return string(res)
}

func part1(s string) string {
	var length int
	var base string
	if len(s) == 5 {
		length = 20
	} else {
		length = 272
	}
	base = s
	for len(base) < length {
		base = dragonCurve(base)
	}
	base = base[:length]
	checksum := pairs(base)
	for len(checksum)%2 == 0 {
		checksum = pairs(checksum)
	}
	return checksum
}

func part2(s string) string {
	var length int
	var base string
	length = 35651584
	base = s
	for len(base) < length {
		base = dragonCurve(base)
	}
	base = base[:length]
	checksum := pairs(base)
	for len(checksum)%2 == 0 {
		checksum = pairs(checksum)
	}
	return checksum
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
