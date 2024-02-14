package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func format(s string) (map[point]string, string) {
	mapping := make(map[point]string)
	enhancement := ""
	parts := strings.Split(s, "\n\n")
	enhancement = parts[0]
	lines := strings.Split(parts[1], "\n")
	for i, line := range lines {
		for j, s := range strings.Split(line, "") {
			mapping[point{i, j}] = s
		}
	}
	return mapping, enhancement
}

func neighbors(p point) []point {
	res := []point{}
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			res = append(res, point{p.x + x, p.y + y})
		}
	}
	return res
}

func countLits(mapping map[point]string) int {
	c := 0
	for _, value := range mapping {
		if value == "#" {
			c++
		}
	}
	return c
}

func part1(s string) int {
	mapping, enhancement := format(s)
	steps := 0
	for steps < 2 {
		new_mapping := make(map[point]string)
		for pixel := range mapping {
			for _, neigh := range neighbors(pixel) {
				_, done := new_mapping[neigh]
				if !done {
					str := ""
					for _, new_neigh := range neighbors(neigh) {
						c, ok := mapping[new_neigh]
						if !ok {
							if enhancement[0:1] == "." {
								str += "0"
							}
							if enhancement[1:2] == "#" && enhancement[(len(enhancement)-1):] == "." {
								if steps%2 == 0 {
									str += "0"
								} else {
									str += "1"
								}
							}
						} else {
							if c == "#" {
								str += "1"
							} else {
								str += "0"
							}
						}
					}
					index, _ := strconv.ParseInt(str, 2, 64)
					new_mapping[neigh] = string(enhancement[index])
				}
			}
		}
		mapping = new_mapping
		steps++
	}
	return countLits(mapping)
}

func part2(s string) int {
	mapping, enhancement := format(s)
	steps := 0
	for steps < 50 {
		new_mapping := make(map[point]string)
		for pixel := range mapping {
			for _, neigh := range neighbors(pixel) {
				_, done := new_mapping[neigh]
				if !done {
					str := ""
					for _, new_neigh := range neighbors(neigh) {
						c, ok := mapping[new_neigh]
						if !ok {
							if enhancement[0:1] == "." {
								str += "0"
							}
							if enhancement[1:2] == "#" && enhancement[(len(enhancement)-1):] == "." {
								if steps%2 == 0 {
									str += "0"
								} else {
									str += "1"
								}
							}
						} else {
							if c == "#" {
								str += "1"
							} else {
								str += "0"
							}
						}
					}
					index, _ := strconv.ParseInt(str, 2, 64)
					new_mapping[neigh] = string(enhancement[index])
				}
			}
		}
		mapping = new_mapping
		steps++
	}
	return countLits(mapping)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day20/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day20/input.data")

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
