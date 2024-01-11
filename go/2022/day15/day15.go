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

type sector struct {
	center   point
	distance int
	top      point
	bottom   point
	left     point
	right    point
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func dist(p point, q point) int {
	return abs(p.x-q.x) + abs(p.y-q.y)
}

func Split(r rune) bool {
	return r == ' ' || r == '=' || r == ':' || r == ','
}

func edges(s sector) []point {
	p := s.top
	p.x--
	res := []point{}
	for p.x != s.right.x {
		res = append(res, p)
		p.x++
		p.y++
	}
	for p.y != s.bottom.y {
		res = append(res, p)
		p.x++
		p.y--
	}
	for p.x != s.left.x {
		res = append(res, p)
		p.x--
		p.y--
	}
	for p.x != s.top.x {
		res = append(res, p)
		p.x--
		p.y++
	}
	res = append(res, p)
	return res
}

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.FieldsFunc(x, Split))
	}
	res := [][]int{}
	for i := range tab {
		a, _ := strconv.Atoi(tab[i][3])
		b, _ := strconv.Atoi(tab[i][5])
		c, _ := strconv.Atoi(tab[i][11])
		d, _ := strconv.Atoi(tab[i][13])
		tmp := []int{}
		tmp = append(tmp, a, b, c, d)
		res = append(res, tmp)
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	grid := make(map[int]bool)
	res := 0
	line := 0
	if len(tab) == 14 {
		line = 10
	} else {
		line = 2000000
	}
	min_y := 0
	max_y := 0
	for _, x := range tab {
		a, _, e, _ := x[0], x[1], x[2], x[3]
		if a < min_y {
			min_y = a
		}
		if e < min_y {
			min_y = e
		}
		if a > max_y {
			max_y = a
		}
		if e > max_y {
			max_y = e
		}
	}
	min_y *= 2
	max_y *= 2
	k := min_y
	for k <= max_y {
		p := point{line, k}
		for _, x := range tab {
			a, b, e, f := x[0], x[1], x[2], x[3]
			sensor := point{b, a}
			beacon := point{f, e}
			d := dist(sensor, beacon)
			if dist(p, sensor) <= d && (p.x != beacon.x || p.y != beacon.y) {
				grid[k] = true
			}
		}
		k++
	}
	for x := range grid {
		if grid[x] {
			res++
		}
	}
	return res
}

func part2(s string) int {
	tab := format(s)
	max := 0
	edgesList := []point{}
	list := []sector{}
	if len(tab) == 14 {
		max = 20
	} else {
		max = 4000000
	}
	for _, x := range tab {
		a, b, e, f := x[0], x[1], x[2], x[3]
		sensor := point{b, a}
		beacon := point{f, e}
		d := dist(sensor, beacon)
		var s sector
		s.center = sensor
		s.distance = d
		s.top = point{b - d, a}
		s.bottom = point{b + d, a}
		s.left = point{b, a - d}
		s.right = point{b, a + d}
		edgesList = append(edgesList, edges(s)...)
		list = append(list, s)
	}
	for _, x := range edgesList {
		b := true
		for _, s := range list {
			if dist(x, s.center) <= s.distance {
				b = false
				break
			}
		}
		if b && x.x >= 0 && x.x <= max && x.y <= max && x.y >= 0 {
			return x.y*4000000 + x.x
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day15/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day15/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
