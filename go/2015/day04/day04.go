package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func part1(s string) int {
	i := 1
	for true {
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
	for true {
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
	content, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 : %d\n", part1(string(content)))
	fmt.Printf("Part 2 : %d\n", part2(string(content)))
}
