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

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func shoelace(l []point) int {
	res := 0
	for i := 0; i < len(l)-1; i++ {
		p1 := l[i]
		p2 := l[i+1]
		res += p1.x*p2.y - p1.y*p2.x
	}
	return res / 2
}

func part1(s string) int {
	list := format(s)
	trench := []point{}
	i := 0
	j := 0
	trench = append(trench, point{i, j})
	length := 0
	for _, line := range list {
		x, _ := strconv.Atoi(line[1])
		length += x
		switch line[0] {
		case "R":
			j = j + x
		case "L":
			j = j - x
		case "U":
			i = i - x
		case "D":
			i = i + x
		}
		trench = append(trench, point{i, j})
	}
	return abs(shoelace(trench)) + 1 + length/2
}

func part2(s string) int {
	list := format(s)
	trench := []point{}
	i := 0
	j := 0
	trench = append(trench, point{i, j})
	length := 0
	for _, line := range list {
		codes := strings.Split(line[2], "")
		y, _ := strconv.ParseInt(codes[2]+codes[3]+codes[4]+codes[5]+codes[6], 16, 64)
		x := int(y)
		length += x
		switch line[2][7] {
		case '0':
			j = j + x
		case '2':
			j = j - x
		case '3':
			i = i - x
		case '1':
			i = i + x

		}
		trench = append(trench, point{i, j})
	}
	return abs(shoelace(trench)) + 1 + length/2
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day18/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day18/input.txt")

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
