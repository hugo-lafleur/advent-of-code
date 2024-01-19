package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

func Split(r rune) bool {
	return r == '=' || r == ',' || r == '.' || r == ' '
}

func format(s string) map[point]string {
	res := make(map[point]string)
	lines := strings.Split(s, "\n")
	strs := [][]string{}
	for _, line := range lines {
		strs = append(strs, strings.FieldsFunc(line, Split))
	}
	for _, line := range strs {
		a, _ := strconv.Atoi(line[1])
		b, _ := strconv.Atoi(line[3])
		c, _ := strconv.Atoi(line[4])
		if line[0] == "x" {
			for i := b; i <= c; i++ {
				res[point{a, i}] = "#"
			}
			continue
		}
		if line[0] == "y" {
			for i := b; i <= c; i++ {
				res[point{i, a}] = "#"
			}
			continue
		}
	}
	return res
}

func flooded(tab map[point]string, p point) bool {
	if tab[p] == "#" {
		return true
	}
	if tab[p] != "~" {
		return false
	}
	left := point{p.x, p.y}
	for tab[left] == "~" {
		left.x--
	}
	right := point{p.x, p.y}
	for tab[right] == "~" {
		right.x++
	}
	return tab[left] == "#" && tab[right] == "#"
}

func canFall(tab map[point]string, p point) bool {
	if tab[p] != "|" {
		return false
	}
	left := point{p.x - 1, p.y}
	for tab[left] != "#" {
		if tab[point{left.x, left.y + 1}] == "" {
			return true
		}
		left.x--
	}
	right := point{p.x + 1, p.y}
	for tab[right] != "#" {
		if tab[point{right.x, right.y + 1}] == "" {
			return true
		}
		right.x++
	}
	return false
}

func calculateTiles(tab map[point]string, max, min int) int {
	c := 0
	for key, value := range tab {
		if value != "#" && key.y <= max && key.y >= min {
			c++
		}
	}
	return c
}

func calculateRetained(tab map[point]string, max, min int) int {
	c := 0
	for key, value := range tab {
		if value == "~" && key.y <= max && key.y >= min {
			c++
		}
	}
	return c
}

func resolve(s string) (int, int, map[point]string) {
	tab := format(s)
	var fallingWater deque.Deque[point]
	var flatWater deque.Deque[point]
	maximumYValue := 0
	minYvalue := 100000
	for key := range tab {
		if key.y > maximumYValue {
			maximumYValue = key.y
		}
		if key.y < minYvalue {
			minYvalue = key.y
		}
	}
	tab[point{500, 0}] = "|"
	fallingWater.PushBack(point{500, 0})
	for {
		for flatWater.Len() != 0 { // flood the current region if water is not falling
			key := flatWater.PopBack()
			belowPoint := point{key.x, key.y + 1}
			rightPoint := point{key.x + 1, key.y}
			leftPoint := point{key.x - 1, key.y}
			right := tab[rightPoint]
			left := tab[leftPoint]
			below := tab[belowPoint]
			if key.y >= maximumYValue {
				continue
			}
			if below == "|" || below == "" {
				tab[belowPoint] = "~"
				flatWater.PushBack(belowPoint)
			}
			if right == "|" || right == "" {
				tab[rightPoint] = "~"
				flatWater.PushBack(rightPoint)
			}
			if left == "|" || left == "" {
				tab[leftPoint] = "~"
				flatWater.PushBack(leftPoint)
			}
		}
		if fallingWater.Len() == 0 {
			return maximumYValue, minYvalue, tab
		}
		key := fallingWater.Back()
		belowPoint := point{key.x, key.y + 1}
		rightPoint := point{key.x + 1, key.y}
		leftPoint := point{key.x - 1, key.y}
		right := tab[rightPoint]
		left := tab[leftPoint]
		below := tab[belowPoint]
		if key.y >= maximumYValue { // if off-limit go back
			fallingWater.PopBack()
			continue
		}
		if below == "|" { // if going back go back further
			fallingWater.PopBack()
			continue
		}
		if below == "" { // fall
			tab[belowPoint] = "|"
			fallingWater.PushBack(belowPoint)
			continue
		}
		if below == "#" && left != "|" && right != "|" { // if falling water just hit ground
			tab[key] = "~"
			flatWater.PushBack(key)
			fallingWater.PopBack()
			continue
		}
		if below == "~" && flooded(tab, belowPoint) && !canFall(tab, key) { // if water is risng up
			tab[key] = "~"
			flatWater.PushBack(key)
			fallingWater.PopBack()
			continue
		}
		if flooded(tab, belowPoint) && canFall(tab, key) { // if water overflow
			fallingWater.PopBack()
			if right == "" {
				tab[rightPoint] = "|"
				fallingWater.PushBack(rightPoint)
			}
			if left == "" {
				tab[leftPoint] = "|"
				fallingWater.PushBack(leftPoint)
			}
			continue
		}
	}
}

func part1(s string) int {
	max, min, tab := resolve(s)
	return calculateTiles(tab, max, min)
}

func part2(s string) int {
	max, min, tab := resolve(s)
	return calculateRetained(tab, max, min)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day17/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day17/input.data")

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
