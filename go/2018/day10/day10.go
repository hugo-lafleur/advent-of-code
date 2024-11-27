package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector struct {
	x, y int
}

type point struct {
	position, velocity vector
}

func Split(r rune) bool {
	return r == '>' || r == ',' || r == '<' || r == ' '
}

func format(s string) []point {
	lines := strings.Split(s, "\n")
	res := []point{}
	for _, line := range lines {
		strs := strings.FieldsFunc(line, Split)
		a, _ := strconv.Atoi(strs[1])
		b, _ := strconv.Atoi(strs[2])
		c, _ := strconv.Atoi(strs[4])
		d, _ := strconv.Atoi(strs[5])
		res = append(res, point{vector{a, b}, vector{c, d}})
	}
	return res
}

func size(points []point) int {
	minX := points[0].position.x
	maxX := points[0].position.x
	minY := points[0].position.y
	maxY := points[0].position.y
	for _, p := range points {
		if p.position.x > maxX {
			maxX = p.position.x
		}
		if p.position.x < minX {
			minX = p.position.x
		}
		if p.position.y > maxY {
			maxY = p.position.y
		}
		if p.position.y < minY {
			minY = p.position.y
		}
	}
	return (maxX - minX) * (maxY - minY)
}

func pointsToString(points []point) string {
	minX := points[0].position.x
	maxX := points[0].position.x
	minY := points[0].position.y
	maxY := points[0].position.y
	for _, p := range points {
		if p.position.x > maxX {
			maxX = p.position.x
		}
		if p.position.x < minX {
			minX = p.position.x
		}
		if p.position.y > maxY {
			maxY = p.position.y
		}
		if p.position.y < minY {
			minY = p.position.y
		}
	}
	tab := [][]string{}
	for i := 0; i < maxY-minY+1; i++ {
		line := []string{}
		for j := 0; j < maxX-minX+1; j++ {
			line = append(line, ".")
		}
		tab = append(tab, line)
	}
	for _, p := range points {
		tab[p.position.y-minY][p.position.x-minX] = "#"
	}
	res := ""
	for _, line := range tab {
		for _, s := range line {
			if s == "." {
				res += "  "
				continue
			}
			res += s + " "
		}
		res += "\n"
	}
	return res
}

func part1(s string) string {
	points := format(s)
	res := append([]point{}, points...)
	min := size(points)
	for {
		for i, p := range points {
			points[i].position.x += p.velocity.x
			points[i].position.y += p.velocity.y
		}
		if size(points) < min {
			min = size(points)
			res = append([]point{}, points...)
		}
		if size(points) > min {
			return pointsToString(res)
		}
	}
}

func part2(s string) int {
	c := 0
	points := format(s)
	min := size(points)
	for {
		for i, p := range points {
			points[i].position.x += p.velocity.x
			points[i].position.y += p.velocity.y
		}
		if size(points) < min {
			min = size(points)
		}
		if size(points) > min {
			return c
		}
		c++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day10/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : \n%v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day10/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : \n%v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
