package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) (int, int) {
	lines := strings.Split(s, "\n")
	a, _ := strconv.Atoi(lines[0])
	b, _ := strconv.Atoi(lines[1])
	return a, b
}

func transform(subject, loopSize int) int {
	curr := 1
	for i := 0; i < loopSize; i++ {
		curr = curr * subject
		curr = (curr % 20201227)
	}
	return curr
}

func part1(s string) int {
	cardKey, doorKey := format(s)
	curr := 1
	var cardLoop, doorLoop int
	for loop := 0; cardLoop == 0 || doorLoop == 0; loop++ {
		if curr == cardKey {
			cardLoop = loop
		}
		if curr == doorKey {
			doorLoop = loop
		}
		curr = curr * 7
		curr = (curr % 20201227)
	}
	return transform(cardKey, doorLoop)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day25/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
