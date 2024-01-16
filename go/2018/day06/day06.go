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

func Split(r rune) bool {
	return r == '\n' || r == ',' || r == ' '
}

func format(s string) []point {
	strs := strings.FieldsFunc(s, Split)
	res := []point{}
	for i := 0; i < len(strs); i += 2 {
		x, _ := strconv.Atoi(strs[i])
		y, _ := strconv.Atoi(strs[i+1])
		res = append(res, point{x, y})
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func closest(p point, l []point) []int {
	d := distance(p, l[0])
	for _, x := range l {
		if distance(p, x) < d {
			d = distance(p, x)
		}
	}
	res := []int{}
	for i, x := range l {
		if distance(p, x) == d {
			res = append(res, i)
		}
	}
	return res
}

func max(m map[int]int) int {
	var max int
	for _, v := range m {
		max = v
	}
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}

func maxCoord(l []point) int {
	m := l[0].x
	for _, p := range l {
		if p.x > m {
			m = p.x
		}
		if p.y > m {
			m = p.y
		}
	}
	return m
}

func part1(s string) int {
	points := format(s)
	edge := maxCoord(points) + 1
	firstAreaSize := make(map[int]int)
	secondAreaSize := make(map[int]int)
	grid := make(map[point]int)
	for i := 0; i < edge; i++ {
		for j := 0; j < edge; j++ {
			for _, x := range []int{i, -i} {
				for _, y := range []int{j, -j} {
					p := point{x, y}
					listClosest := closest(p, points)
					if len(listClosest) == 1 {
						grid[p] = listClosest[0]
						firstAreaSize[listClosest[0]]++
					}
				}
			}
		}
	}
	edge++
	for i := -edge; i < edge; i++ {
		j := edge
		for _, x := range []int{i, -i} {
			for _, y := range []int{j, -j} {
				p := point{x, y}
				listClosest := closest(p, points)
				if len(listClosest) == 1 {
					delete(firstAreaSize, listClosest[0])
				}
			}
		}
	}
	i := edge
	for j := -edge; j < edge; j++ {
		for _, x := range []int{i, -i} {
			for _, y := range []int{j, -j} {
				p := point{x, y}
				listClosest := closest(p, points)
				if len(listClosest) == 1 {
					delete(firstAreaSize, listClosest[0])
				}
			}
		}
	}
	for key := range secondAreaSize {
		delete(firstAreaSize, key)
	}
	return max(firstAreaSize)
}

func part2(s string) int {
	c := 0
	points := format(s)
	edge := maxCoord(points) + 1
	var limit int
	if len(points) == 6 {
		limit = 32
	} else {
		limit = 10000
	}
	for i := 0; i < edge; i++ {
		for j := 0; j < edge; j++ {
			for _, x := range []int{i, -i} {
				for _, y := range []int{j, -j} {
					p := point{x, y}
					d := 0
					for _, pt := range points {
						d += distance(p, pt)
						if d > limit {
							break
						}
					}
					if d < limit {
						c++
					}
				}
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day06/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day06/input.data")

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
