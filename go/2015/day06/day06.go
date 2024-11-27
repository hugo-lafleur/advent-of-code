package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	tab := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	return tab

}

func part1(s string) int {
	cmd := format(s)
	lights := [1000][1000]bool{}
	for _, instruction := range cmd {
		if instruction[0] == "turn" {
			start := strings.Split(instruction[2], ",")
			a, _ := strconv.Atoi(start[0])
			b, _ := strconv.Atoi(start[1])
			end := strings.Split(instruction[4], ",")
			c, _ := strconv.Atoi(end[0])
			d, _ := strconv.Atoi(end[1])
			if instruction[1] == "on" {
				i := a
				j := b
				for i < c+1 {
					j = b
					for j < d+1 {
						lights[i][j] = true
						j++
					}
					i++
				}
			}
			if instruction[1] == "off" {
				i := a
				j := b
				for i < c+1 {
					j = b
					for j < d+1 {
						lights[i][j] = false
						j++
					}
					i++
				}
			}
		}
		if instruction[0] == "toggle" {
			start := strings.Split(instruction[1], ",")
			a, _ := strconv.Atoi(start[0])
			b, _ := strconv.Atoi(start[1])
			end := strings.Split(instruction[3], ",")
			c, _ := strconv.Atoi(end[0])
			d, _ := strconv.Atoi(end[1])
			i := a
			j := b
			for i < c+1 {
				j = b
				for j < d+1 {
					lights[i][j] = !lights[i][j]
					j++
				}
				i++

			}
		}
	}
	c := 0
	i := 0
	j := 0
	for i < 1000 {
		j = 0
		for j < 1000 {
			if lights[i][j] {
				c++
			}
			j++
		}
		i++
	}
	return c
}

func part2(s string) int {
	cmd := format(s)
	bright := [1000][1000]int{}
	for _, instruction := range cmd {
		if instruction[0] == "turn" {
			start := strings.Split(instruction[2], ",")
			a, _ := strconv.Atoi(start[0])
			b, _ := strconv.Atoi(start[1])
			end := strings.Split(instruction[4], ",")
			c, _ := strconv.Atoi(end[0])
			d, _ := strconv.Atoi(end[1])
			if instruction[1] == "on" {
				i := a
				j := b
				for i < c+1 {
					j = b
					for j < d+1 {
						bright[i][j]++
						j++
					}
					i++
				}
			}
			if instruction[1] == "off" {
				i := a
				j := b
				for i < c+1 {
					j = b
					for j < d+1 {
						bright[i][j]--
						if bright[i][j] == -1 {
							bright[i][j] = 0
						}
						j++
					}
					i++
				}
			}
		}
		if instruction[0] == "toggle" {
			start := strings.Split(instruction[1], ",")
			a, _ := strconv.Atoi(start[0])
			b, _ := strconv.Atoi(start[1])
			end := strings.Split(instruction[3], ",")
			c, _ := strconv.Atoi(end[0])
			d, _ := strconv.Atoi(end[1])
			i := a
			j := b
			for i < c+1 {
				j = b
				for j < d+1 {
					bright[i][j] += 2
					j++
				}
				i++

			}
		}
	}
	c := 0
	i := 0
	j := 0
	for i < 1000 {
		j = 0
		for j < 1000 {
			c += bright[i][j]
			j++
		}
		i++
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2015/day06/input.txt")

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
