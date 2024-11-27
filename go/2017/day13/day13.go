package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type layer struct {
	depth     int
	length    int
	scanner   int
	direction int
}

func Split(r rune) bool {
	return r == ' ' || r == ':'
}

func positionScanner(length int, time int) int {
	n := time % ((length - 1) * 2)
	if n < length {
		return n
	}
	return length - 2 - (n - length)
}

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	res := [][]int{}
	for _, line := range lines {
		intLine := []int{}
		strs := strings.FieldsFunc(line, Split)
		for _, x := range strs {
			n, _ := strconv.Atoi(x)
			intLine = append(intLine, n)
		}
		res = append(res, intLine)
	}
	return res
}
func part1(s string) int {
	c := 0
	input := format(s)
	layers := []layer{}
	for _, line := range input {
		layers = append(layers, layer{line[0], line[1], 0, 1})
	}
	maxDepth := input[len(input)-1][0]
	time := 0
	for time < maxDepth+1 {
		for _, layer := range layers {
			if time == layer.depth && layer.scanner == 0 {
				c += time * layer.length
			}
		}
		for i, layer := range layers {
			switch layer.direction {
			case 1:
				if layer.scanner == layer.length-1 {
					layers[i].scanner--
					layers[i].direction = -1
				} else {
					layers[i].scanner++
				}
			case -1:
				if layer.scanner == 0 {
					layers[i].scanner++
					layers[i].direction = 1
				} else {
					layers[i].scanner--
				}
			}
		}
		time++
	}
	return c
}

func part2(s string) int {
	input := format(s)
	layers := []layer{}
	for _, line := range input {
		layers = append(layers, layer{line[0], line[1], 0, 1})
	}
	delay := 0
loop:
	for {
		for _, layer := range layers {
			if positionScanner(layer.length, delay+layer.depth) == 0 {
				delay++
				continue loop
			}
		}
		return delay
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day13/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2017/day13/input.txt")

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
