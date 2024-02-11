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

func format(s string) (map[point]bool, [][]string) {
	mapping := make(map[point]bool)
	instr := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if strings.Contains(line, ",") {
			points := strings.Split(line, ",")
			a, _ := strconv.Atoi(points[0])
			b, _ := strconv.Atoi(points[1])
			mapping[point{a, b}] = true
		}
		if strings.Contains(line, "fold") {
			instr = append(instr, strings.Split(line, "="))
		}
	}
	return mapping, instr
}

func mappingToString(m map[point]bool) string {
	var maxX, maxY int
	for key := range m {
		maxX = key.x
		maxY = key.y
	}
	for key := range m {
		if key.x > maxX {
			maxX = key.x
		}
		if key.y > maxY {
			maxY = key.y
		}
	}
	res := "\n"
	for i := 0; i <= maxY; i++ {
		for j := 0; j <= maxX; j++ {
			if m[point{j, i}] {
				res += "# "
			} else {
				res += "  "
			}
		}
		res += "\n"
	}
	return res
}

func part1(s string) int {
	mapping, instrs := format(s)
	instr := instrs[0]
	newMap := make(map[point]bool)
	if strings.Contains(instr[0], "x") {
		n, _ := strconv.Atoi(instr[1])
		for key := range mapping {
			if key.x < n {
				newMap[key] = true
			} else {
				newMap[point{2*n - key.x, key.y}] = true
			}
		}
	}
	if strings.Contains(instr[0], "y") {
		n, _ := strconv.Atoi(instr[1])
		for key := range mapping {
			if key.y < n {
				newMap[key] = true
			} else {
				newMap[point{key.x, 2*n - key.y}] = true
			}
		}
	}
	return len(newMap)
}

func part2(s string) string {
	mapping, instrs := format(s)
	for _, instr := range instrs {
		newMap := make(map[point]bool)
		if strings.Contains(instr[0], "x") {
			n, _ := strconv.Atoi(instr[1])
			for key := range mapping {
				if key.x < n {
					newMap[key] = true
				} else {
					newMap[point{2*n - key.x, key.y}] = true
				}
			}
		}
		if strings.Contains(instr[0], "y") {
			n, _ := strconv.Atoi(instr[1])
			for key := range mapping {
				if key.y < n {
					newMap[key] = true
				} else {
					newMap[point{key.x, 2*n - key.y}] = true
				}
			}
		}
		mapping = newMap
	}
	return mappingToString(mapping)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day13/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	content, err = os.ReadFile("../../../inputs/2021/day13/input.data")

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
