package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type galaxy struct {
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

func isIn(x int, l []int) bool {
	for _, y := range l {
		if x == y {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func part1(s string) int {
	c := 0
	image := format(s)
	emptyLines := []int{}
	emptyColumns := []int{}
	i := 0
	for i < len(image) {
		j := 0
		empty := true
		for j < len(image[0]) {
			if image[i][j] == "#" {
				empty = false
				break
			}
			j++
		}
		if empty {
			emptyLines = append(emptyLines, i)
		}
		i++
	}
	j := 0
	for j < len(image[0]) {
		i := 0
		empty := true
		for i < len(image) {
			if image[i][j] == "#" {
				empty = false
				break
			}
			i++
		}
		if empty {
			emptyColumns = append(emptyColumns, j)
		}
		j++
	}
	i = 0
	newImage := [][]string{}
	for i < len(image) {
		if isIn(i, emptyLines) {
			newLine := []string{}
			for k := 0; k < len(image[0])+len(emptyColumns); k++ {
				newLine = append(newLine, ".")
			}
			newImage = append(newImage, newLine, newLine)
		} else {
			newLine := []string{}
			j = 0
			for j < len(image[0]) {
				if isIn(j, emptyColumns) {
					newLine = append(newLine, ".")
				}
				if image[i][j] == "#" {
					newLine = append(newLine, "#")
				}
				if image[i][j] == "." {
					newLine = append(newLine, ".")
				}
				j++
			}
			newImage = append(newImage, newLine)
		}
		i++
	}
	galaxies := []galaxy{}
	i = 0
	for i < len(newImage) {
		j = 0
		for j < len(newImage[0]) {
			if newImage[i][j] == "#" {
				galaxies = append(galaxies, galaxy{i, j})
			}
			j++
		}
		i++
	}
	i = 0
	for i < len(galaxies) {
		j = i + 1
		for j < len(galaxies) {
			c += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
			j++
		}
		i++
	}
	return c
}

func howHigher(x int, l []int) int {
	for i, y := range l {
		if x < y {
			return i
		}
	}
	return len(l)
}

func part2(s string) int {
	c := 0
	d := 0
	image := format(s)
	if len(image) == 10 {
		d = 100
	} else {
		d = 1000000
	}
	emptyLines := []int{}
	emptyColumns := []int{}
	i := 0
	for i < len(image) {
		j := 0
		empty := true
		for j < len(image[0]) {
			if image[i][j] == "#" {
				empty = false
				break
			}
			j++
		}
		if empty {
			emptyLines = append(emptyLines, i)
		}
		i++
	}
	j := 0
	for j < len(image[0]) {
		i := 0
		empty := true
		for i < len(image) {
			if image[i][j] == "#" {
				empty = false
				break
			}
			i++
		}
		if empty {
			emptyColumns = append(emptyColumns, j)
		}
		j++
	}
	galaxies := []galaxy{}
	i = 0
	for i < len(image) {
		j = 0
		for j < len(image[0]) {
			if image[i][j] == "#" {
				galaxies = append(galaxies, galaxy{i, j})
			}
			j++
		}
		i++
	}
	for i, galaxy := range galaxies {
		galaxies[i].x = galaxy.x + (d-1)*howHigher(galaxy.x, emptyLines)
		galaxies[i].y = galaxy.y + (d-1)*howHigher(galaxy.y, emptyColumns)
	}
	i = 0
	for i < len(galaxies) {
		j = i + 1
		for j < len(galaxies) {
			c += abs(galaxies[i].x-galaxies[j].x) + abs(galaxies[i].y-galaxies[j].y)
			j++
		}
		i++
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day11/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day11/input.txt")

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
