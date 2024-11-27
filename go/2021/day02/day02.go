package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type command struct {
	dir string
	n   int
}

type submarine struct {
	pos complex128
	aim float64
}

func format(s string) []command {
	lines := strings.Split(s, "\n")
	res := []command{}
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		n, _ := strconv.Atoi(lineSplit[1])
		res = append(res, command{dir: lineSplit[0], n: n})
	}
	return res
}

func part1(s string) int {
	list := format(s)
	pos := complex(0, 0)
	for _, com := range list {
		switch com.dir {
		case "forward":
			pos += complex(float64(com.n), 0)
		case "down":
			pos += complex(0, float64(com.n))
		case "up":
			pos += complex(0, -float64(com.n))
		}
	}
	return int(real(pos)) * int(imag(pos))
}

func part2(s string) int {
	list := format(s)
	curr := submarine{}
	for _, com := range list {
		switch com.dir {
		case "forward":
			curr.pos += complex(float64(com.n), 0)
			curr.pos += complex(0, curr.aim*float64(com.n))
		case "down":
			curr.aim += float64(com.n)
		case "up":
			curr.aim -= float64(com.n)
		}
	}
	return int(real(curr.pos)) * int(imag(curr.pos))
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day02/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day02/input.txt")

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
