package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func part1(s string) int {
	i := 1
	for {
		txt := s + strconv.Itoa(i)
		b := []byte(txt)
		h := md5.Sum(b)
		if h[0] == 0 && h[1] == 0 && h[2] < 17 {
			break
		}
		i++
	}
	return i
}

func part2(s string) int {
	i := 1
	for {
		txt := s + strconv.Itoa(i)
		b := []byte(txt)
		h := md5.Sum(b)
		if h[0] == 0 && h[1] == 0 && h[2] == 0 {
			break
		}
		i++
	}
	return i
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day04/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2015/day04/input.txt")

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
