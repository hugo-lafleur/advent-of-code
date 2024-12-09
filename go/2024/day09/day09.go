package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func parse(s string) []int {
	var b = []byte(s)
	var result = []int{}
	var j int
	for i, char := range b {
		if i%2 == 0 {
			for range char - '0' {
				result = append(result, j)
			}
			j++
		} else {
			for range char - '0' {
				result = append(result, -1)
			}
		}
	}
	return result
}

func part1(s string) int {
	var blocks = parse(s)
	var i, j = 0, len(blocks) - 1
	for i < j {
		if blocks[i] == -1 {
			for blocks[j] == -1 {
				j--
			}
			blocks[i] = blocks[j]
			j--
		}
		i++
	}
	blocks = blocks[:j+1]
	var result int
	for i := range blocks {
		result += i * blocks[i]
	}
	return result
}

func part2(s string) int {
	var blocks = parse(s)
	var j1, j2 = len(blocks) - 1, len(blocks) - 1
	var start = 0
	for j1 >= 0 {
		for blocks[j2] == -1 {
			j2--
			j1--
		}
		for j1-1 >= 0 && blocks[j1-1] == blocks[j2] {
			j1--
		}
		if j1 <= 0 {
			break
		}
		var i1, i2 = start, start
		for i1 < len(blocks) {
			if blocks[i1] == -1 {
				for i2+1 < len(blocks) && blocks[i2+1] == -1 {
					i2++
				}
				if i1 < j1 && i2-i1 >= j2-j1 {
					for k := 0; k <= j2-j1; k++ {
						blocks[i1+k] = blocks[j1+k]
						blocks[j1+k] = -1
					}
					break
				} else {
					i1 = i2 + 1
					i2 = i1
				}
			} else {
				i1++
				i2++
			}
		}
		j2 = j1 - 1
		j1 = j2
	}
	var result int
	for i := range blocks {
		if blocks[i] != -1 {
			result += i * blocks[i]
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day09/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day09/input.txt")

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
