package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []map[string]int {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.Split(x, " "))
	}
	graph := []map[string]int{}
	for _, x := range tab {
		sue := make(map[string]int)
		a, _ := strconv.Atoi(x[3][:len(x[3])-1])
		b, _ := strconv.Atoi(x[5][:len(x[5])-1])
		c, _ := strconv.Atoi(x[7])
		sue[x[2][:len(x[2])-1]] = a
		sue[x[4][:len(x[4])-1]] = b
		sue[x[6][:len(x[6])-1]] = c
		graph = append(graph, sue)
	}
	return graph
}

func part1(s string) []int {
	graph := format(s)
	res := []int{}
	for i, sue := range graph {
		b := true
		if val, ok := sue["children"]; ok {
			if val != 3 {
				b = false
			}
		}
		if val, ok := sue["cats"]; ok {
			if val != 7 {
				b = false
			}
		}
		if val, ok := sue["samoyeds"]; ok {
			if val != 2 {
				b = false
			}
		}
		if val, ok := sue["pomeranians"]; ok {
			if val != 3 {
				b = false
			}
		}
		if val, ok := sue["akitas"]; ok {
			if val != 0 {
				b = false
			}
		}
		if val, ok := sue["vizslas"]; ok {
			if val != 0 {
				b = false
			}
		}
		if val, ok := sue["goldfish"]; ok {
			if val != 5 {
				b = false
			}
		}
		if val, ok := sue["trees"]; ok {
			if val != 3 {
				b = false
			}
		}
		if val, ok := sue["cars"]; ok {
			if val != 2 {
				b = false
			}
		}
		if val, ok := sue["perfumes"]; ok {
			if val != 1 {
				b = false
			}
		}
		if b {
			res = append(res, i+1)
		}
	}
	return res
}

func part2(s string) []int {
	graph := format(s)
	res := []int{}
	for i, sue := range graph {
		b := true
		if val, ok := sue["children"]; ok {
			if val != 3 {
				b = false
			}
		}
		if val, ok := sue["cats"]; ok {
			if val <= 7 {
				b = false
			}
		}
		if val, ok := sue["samoyeds"]; ok {
			if val != 2 {
				b = false
			}
		}
		if val, ok := sue["pomeranians"]; ok {
			if val >= 3 {
				b = false
			}
		}
		if val, ok := sue["akitas"]; ok {
			if val != 0 {
				b = false
			}
		}
		if val, ok := sue["vizslas"]; ok {
			if val != 0 {
				b = false
			}
		}
		if val, ok := sue["goldfish"]; ok {
			if val >= 5 {
				b = false
			}
		}
		if val, ok := sue["trees"]; ok {
			if val <= 3 {
				b = false
			}
		}
		if val, ok := sue["cars"]; ok {
			if val != 2 {
				b = false
			}
		}
		if val, ok := sue["perfumes"]; ok {
			if val != 1 {
				b = false
			}
		}
		if b {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day16/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
