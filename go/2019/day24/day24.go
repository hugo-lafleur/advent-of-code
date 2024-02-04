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

func addPoint(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func isValid(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < 5 && p.y < 5
}

func next(tab [][]string) ([][]string, string) {
	res := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			p := point{i, j}
			s := tab[i][j]
			c := 0
			for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				newP := addPoint(p, offset)
				if isValid(newP) && tab[newP.x][newP.y] == "#" {
					c++
				}
			}
			if s == "#" {
				if c == 1 {
					res += "#"
				} else {
					res += "."
				}
			}
			if s == "." {
				if c == 1 || c == 2 {
					res += "#"
				} else {
					res += "."
				}
			}
		}
		res += "\n"
	}
	res = strings.TrimSuffix(res, "\n")
	return format(res), res
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func rating(s string) int {
	res := 0
	power := 1
	for _, r := range s {
		if r == '#' {
			res += power
		}
		if r != '\n' {
			power *= 2
		}
	}
	return res
}

func countInside(down [][]string, i, j int) int {
	res := 0
	if i == 1 {
		for _, s := range down[0] {
			if s == "#" {
				res++
			}
		}
	}
	if i == 3 {
		for _, s := range down[4] {
			if s == "#" {
				res++
			}
		}
	}
	if j == 1 {
		for _, line := range down {
			if line[0] == "#" {
				res++
			}
		}
	}
	if j == 3 {
		for _, line := range down {
			if line[4] == "#" {
				res++
			}
		}
	}
	return res
}

func countOutside(up [][]string, i, j int) int {
	res := 0
	if i == 0 && up[1][2] == "#" {
		res++
	}
	if i == 4 && up[3][2] == "#" {
		res++
	}
	if j == 0 && up[2][1] == "#" {
		res++
	}
	if j == 4 && up[2][3] == "#" {
		res++
	}
	return res
}

func next2(tab [][]string, up [][]string, down [][]string) string {
	res := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i != 2 || j != 2 {
				p := point{i, j}
				s := tab[i][j]
				c := 0
				for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					newP := addPoint(p, offset)
					if isValid(newP) && tab[newP.x][newP.y] == "#" {
						c++
					}
				}
				if (i+j == 3 && (i == 1 || j == 1)) || (i+j == 5 && (i == 2 || j == 2)) {
					c += countInside(down, i, j)
				} else {
					c += countOutside(up, i, j)
				}
				if s == "#" {
					if c == 1 {
						res += "#"
					} else {
						res += "."
					}
				}
				if s == "." {
					if c == 1 || c == 2 {
						res += "#"
					} else {
						res += "."
					}
				}
			} else {
				res += "?"
			}
		}
		res += "\n"
	}
	res = strings.TrimSuffix(res, "\n")
	return res
}

func bugsNumber(levels map[int]string) int {
	res := 0
	for _, value := range levels {
		for _, r := range value {
			if r == '#' {
				res++
			}
		}
	}
	return res
}

func part1(s string) int {
	str := s
	tab := format(s)
	memory := make(map[string]bool)
	memory[str] = true
	for {
		tab, str = next(tab)
		_, ok := memory[str]
		if ok {
			return rating(str)
		}
		memory[str] = true
	}
}

func part2(s string) int {
	var minutes int
	levels := make(map[int]string)
	levels[0] = s
	if s == "....#\n#..#.\n#..##\n..#..\n#...." {
		minutes = 10
	} else {
		minutes = 200
	}
	for i := 1; i <= (minutes/2)+1; i++ {
		levels[i] = ".....\n.....\n.....\n.....\n....."
		levels[-i] = ".....\n.....\n.....\n.....\n....."
	}
	for m := 0; m < minutes; m++ {
		newLevels := make(map[int]string)
		for key, value := range levels {
			newLevels[key] = value
		}
		for index := -(minutes / 2); index <= (minutes / 2); index++ {
			newLevels[index] = next2(format(levels[index]), format(levels[index-1]), format(levels[index+1]))
		}
		levels = newLevels
	}
	return bugsNumber(levels)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day24/test.data")

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

	content, err = os.ReadFile("../../../inputs/2019/day24/input.data")

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
