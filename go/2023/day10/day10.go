package main

import (
	"fmt"
	"log"
	"os"
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
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func next(current point, previous point, tab [][]string) point {
	s := tab[current.x][current.y]
	if s == "|" {
		if current.x == previous.x+1 {
			return point{current.x + 1, current.y}
		} else {
			return point{current.x - 1, current.y}
		}
	}
	if s == "-" {
		if current.y == previous.y+1 {
			return point{current.x, current.y + 1}
		} else {
			return point{current.x, current.y - 1}
		}
	}
	if s == "L" {
		if current.x == previous.x+1 && current.y == previous.y {
			return point{current.x, current.y + 1}
		} else {
			return point{current.x - 1, current.y}
		}
	}
	if s == "J" {
		if current.x == previous.x+1 && current.y == previous.y {
			return point{current.x, current.y - 1}
		} else {
			return point{current.x - 1, current.y}
		}
	}
	if s == "7" {
		if current.x == previous.x-1 && current.y == previous.y {
			return point{current.x, current.y - 1}
		} else {
			return point{current.x + 1, current.y}
		}
	}
	if s == "F" {
		if current.x == previous.x-1 && current.y == previous.y {
			return point{current.x, current.y + 1}
		} else {
			return point{current.x + 1, current.y}
		}
	}
	return point{0, 0}
}

func part1(s string) int {
	tab := format(s)
	var source point
	for i, line := range tab {
		for j, s := range line {
			if s == "S" {
				source.x = i
				source.y = j
			}
		}
	}
	var current point
	var previous point
	previous.x = source.x
	previous.y = source.y
	if source.x < len(tab)-1 && (tab[source.x+1][source.y] == "|" || tab[source.x+1][source.y] == "L" || tab[source.x+1][source.y] == "J") {
		current.x = source.x + 1
		current.y = source.y
	} else {
		if source.x > 0 && (tab[source.x-1][source.y] == "|" || tab[source.x-1][source.y] == "F" || tab[source.x-1][source.y] == "7") {
			current.x = source.x - 1
			current.y = source.y
		} else {
			if source.y > 0 && (tab[source.x][source.y-1] == "-" || tab[source.x][source.y-1] == "L" || tab[source.x][source.y-1] == "F") {
				current.x = source.x
				current.y = source.y - 1
			} else {
				if source.y < len(tab[0])-1 && (tab[source.x][source.y+1] == "-" || tab[source.x][source.y+1] == "J" || tab[source.x][source.y-1] == "7") {
					current.x = source.x
					current.y = source.y + 1
				}
			}
		}
	}
	i := 1
	for !(current.x == source.x && current.y == source.y) {
		temp := current
		current = next(current, previous, tab)
		previous = temp
		i++
	}
	return i / 2
}

func isIn(x point, l []point) bool {
	for _, y := range l {
		if x.x == y.x && x.y == y.y {
			return true
		}
	}
	return false
}

func part2(s string) int {
	c := 0
	tab := format(s)
	var source point
	for i, line := range tab {
		for j, s := range line {
			if s == "S" {
				source.x = i
				source.y = j
			}
		}
	}
	var current point
	var previous point
	previous.x = source.x
	previous.y = source.y
	if source.x < len(tab)-1 && (tab[source.x+1][source.y] == "|" || tab[source.x+1][source.y] == "L" || tab[source.x+1][source.y] == "J") {
		current.x = source.x + 1
		current.y = source.y
	} else {
		if source.x > 0 && (tab[source.x-1][source.y] == "|" || tab[source.x-1][source.y] == "F" || tab[source.x-1][source.y] == "7") {
			current.x = source.x - 1
			current.y = source.y
		} else {
			if source.y > 0 && (tab[source.x][source.y-1] == "-" || tab[source.x][source.y-1] == "L" || tab[source.x][source.y-1] == "F") {
				current.x = source.x
				current.y = source.y - 1
			} else {
				if source.y < len(tab[0])-1 && (tab[source.x][source.y+1] == "-" || tab[source.x][source.y+1] == "J" || tab[source.x][source.y-1] == "7") {
					current.x = source.x
					current.y = source.y + 1
				}
			}
		}
	}
	polygon := []point{}
	polygon = append(polygon, source, current)
	for !(current.x == source.x && current.y == source.y) {
		temp := current
		current = next(current, previous, tab)
		previous = temp
		polygon = append(polygon, current)
	}
	i := 0
	scale2Polygon := []point{}
	scale2Polygon = append(scale2Polygon, point{2 * source.x, 2 * source.y})
	for i < len(polygon)-1 {
		current = polygon[i]
		nextPoint := polygon[i+1]
		if current.x == nextPoint.x {
			if current.y < nextPoint.y {
				scale2Polygon = append(scale2Polygon, point{2 * current.x, 2 * current.y}, point{2 * current.x, 2*current.y + 1}, point{2 * nextPoint.x, 2 * nextPoint.y})
			} else {
				scale2Polygon = append(scale2Polygon, point{2 * current.x, 2 * current.y}, point{2 * current.x, 2*current.y - 1}, point{2 * nextPoint.x, 2 * nextPoint.y})
			}
		} else {
			if current.x < nextPoint.x {
				scale2Polygon = append(scale2Polygon, point{2 * current.x, 2 * current.y}, point{2*current.x + 1, 2 * current.y}, point{2 * nextPoint.x, 2 * nextPoint.y})
			} else {
				scale2Polygon = append(scale2Polygon, point{2 * current.x, 2 * current.y}, point{2*current.x - 1, 2 * current.y}, point{2 * nextPoint.x, 2 * nextPoint.y})
			}
		}
		i++
	}
	outside := []point{}
	n := 2 * len(tab)
	m := 2 * len(tab[0])
	x := 0
	y := 0
	for x < n-1 {
		p := point{x, 0}
		if !isIn(p, scale2Polygon) && !isIn(p, outside) {
			outside = append(outside, p)
		}
		p = point{x, m - 2}
		if !isIn(p, scale2Polygon) && !isIn(p, outside) {
			outside = append(outside, p)
		}
		x++
	}
	for y < m-1 {
		p := point{0, y}
		if !isIn(p, scale2Polygon) && !isIn(p, outside) {
			outside = append(outside, p)
		}
		p = point{n - 2, y}
		if !isIn(p, scale2Polygon) && !isIn(p, outside) {
			outside = append(outside, p)
		}
		y++
	}
	i = 0
	for i < len(outside) {
		pnt := outside[i]
		x := pnt.x
		y := pnt.y
		p := point{x + 1, y}
		if x < 2*len(tab)-2 && !isIn(p, outside) && !isIn(p, scale2Polygon) {
			outside = append(outside, p)
		}
		p = point{x - 1, y}
		if x > 0 && !isIn(p, outside) && !isIn(p, scale2Polygon) {
			outside = append(outside, p)
		}
		p = point{x, y + 1}
		if y < 2*len(tab[0])-2 && !isIn(p, outside) && !isIn(p, scale2Polygon) {
			outside = append(outside, p)
		}
		p = point{x, y - 1}
		if y > 0 && !isIn(p, outside) && !isIn(p, scale2Polygon) {
			outside = append(outside, p)
		}
		i++
	}
	//print the graphical view of the output
	for i = 0; i < len(tab); i++ {
		for j := 0; j < len(tab[0]); j++ {
			/*if isIn(point{2 * i, 2 * j}, outside) {
				fmt.Print("O")
			}
			if isIn(point{2 * i, 2 * j}, scale2Polygon) {
				fmt.Printf("*")
			}*/
			if !isIn(point{2 * i, 2 * j}, scale2Polygon) && !isIn(point{2 * i, 2 * j}, outside) {
				//fmt.Printf("I")
				c++
			}
		}
		//fmt.Print("\n")
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day10/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2023/day10/test2.data")

	if err != nil {
		log.Fatal(err)
	}

	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2023/day10/input.data")

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
