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

type state struct {
	pos point
	d   string
}

func format(s string) ([][]string, []string) {
	parts := strings.Split(s, "\n\n")
	res := [][]string{}
	lines := strings.Split(parts[0], "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	temp := ""
	for _, r := range parts[1] {
		if r == 'L' || r == 'R' {
			temp += " " + string(r) + " "
		} else {
			temp += string(r)
		}
	}
	return res, strings.Split(temp, " ")
}

func next(p point, d string) point {
	switch d {
	case "L":
		return point{p.x, p.y - 1}
	case "R":
		return point{p.x, p.y + 1}
	case "U":
		return point{p.x - 1, p.y}
	case "D":
		return point{p.x + 1, p.y}
	}
	return point{}
}

func invDir(d string) string {
	switch d {
	case "L":
		return "R"
	case "R":
		return "L"
	case "U":
		return "D"
	case "D":
		return "U"
	}
	return ""
}

func changeDir(d string, c string) string {
	switch d {
	case "L":
		switch c {
		case "L":
			return "D"
		case "R":
			return "U"
		}
	case "R":
		switch c {
		case "L":
			return "U"
		case "R":
			return "D"
		}
	case "U":
		switch c {
		case "L":
			return "L"
		case "R":
			return "R"
		}
	case "D":
		switch c {
		case "L":
			return "R"
		case "R":
			return "L"
		}
	}
	return ""
}

func facing(d string) int {
	switch d {
	case "L":
		return 2
	case "R":
		return 0
	case "U":
		return 3
	case "D":
		return 1
	}
	return -1
}

func nextCube4(pos point, d string, c int) (point, string) {
	x, y := pos.x, pos.y
	switch c {
	case 0:
		if y == 11 && d == "R" {
			return point{11 - x, 15}, "L"
		}
		if x == 0 && d == "U" {
			return point{4, 11 - y}, "D"
		}
		if y == 8 && d == "L" {
			return point{4, x + 4}, "D"
		}
	case 2:
		if x == 4 && d == "U" {
			return point{y - 4, 8}, "R"
		}
		if x == 7 && d == "D" {
			return point{15 - y, 8}, "R"
		}
	case 1:
		if x == 4 && d == "U" {
			return point{0, 11 - y}, "D"
		}
		if y == 0 && d == "L" {
			return point{11, 19 - x}, "U"
		}
		if x == 7 && d == "D" {
			return point{11, 11 - y}, "U"
		}
	case 4:
		if y == 8 && d == "L" {
			return point{7, 15 - x}, "U"
		}
		if x == 11 && d == "D" {
			return point{7, 11 - y}, "U"
		}
	case 5:
		if x == 11 && d == "D" {
			return point{19 - y, 0}, "R"
		}
		if y == 15 && d == "R" {
			return point{11 - x, 11}, "L"
		}
		if x == 8 && d == "U" {
			return point{19 - y, 11}, "L"
		}
	case 3:
		if y == 11 && d == "R" {
			return point{8, 19 - x}, "D"
		}
	}
	fmt.Println("error")
	return point{}, ""
}

func nextCube50(pos point, d string, c int) (point, string) {
	x, y := pos.x, pos.y
	switch c {
	case 0:
		if x == 0 && d == "U" {
			return point{y + 100, 0}, "R"
		}
		if y == 50 && d == "L" { //
			return point{149 - x, 0}, "R"
		}
	case 2:
		if y == 50 && d == "L" { //
			return point{100, x - 50}, "D"
		}
		if y == 99 && d == "R" {
			return point{49, x + 50}, "U"
		}
	case 4:
		if y == 99 && d == "R" {
			return point{149 - x, 149}, "L"
		}
		if x == 149 && d == "D" {
			return point{y + 100, 49}, "L"
		}
	case 5:
		if y == 49 && d == "R" {
			return point{149, x - 100}, "U"
		}
		if x == 199 && d == "D" {
			return point{0, y + 100}, "D"
		}
		if y == 0 && d == "L" {
			return point{0, x - 100}, "D"
		}
	case 3:
		if y == 0 && d == "L" {
			return point{149 - x, 50}, "R"
		}
		if x == 100 && d == "U" { //
			return point{y + 50, 50}, "R"
		}
	case 1:
		if x == 0 && d == "U" {
			return point{199, y - 100}, "U"
		}
		if y == 149 && d == "R" {
			return point{149 - x, 99}, "L"
		}
		if x == 49 && d == "D" {
			return point{y - 50, 99}, "L"
		}
	}
	fmt.Println("error")
	return point{}, ""
}

func part1(s string) int {
	tab, mov := format(s)
	curr := state{}
	curr.d = "R"
loop:
	for i, line := range tab {
		for j, c := range line {
			if c == "." {
				curr.pos = point{i, j}
				break loop
			}
		}
	}
	for _, instr := range mov {
		if instr == "L" || instr == "R" {
			curr.d = changeDir(curr.d, instr)
		} else {
			n, _ := strconv.Atoi(instr)
			pos := curr.pos
			for i := 0; i < n; i++ {
				pos = next(pos, curr.d)
				if pos.x < 0 || pos.y < 0 || pos.x >= len(tab) || pos.y >= len(tab[pos.x]) || (tab[pos.x][pos.y] != "." && tab[pos.x][pos.y] != "#") {
					oldPos := next(pos, invDir(curr.d))
					pos = next(pos, invDir(curr.d))
					for tab[pos.x][pos.y] == "." || tab[pos.x][pos.y] == "#" {
						pos = next(pos, invDir(curr.d))
						if pos.x < 0 || pos.y < 0 || pos.x >= len(tab) || pos.y >= len(tab[pos.x]) || (tab[pos.x][pos.y] != "." && tab[pos.x][pos.y] != "#") {
							break
						}
					}
					pos = next(pos, curr.d)
					if tab[pos.x][pos.y] == "#" {
						pos = oldPos
						break
					}
				}
				if tab[pos.x][pos.y] == "." {
					continue
				}
				if tab[pos.x][pos.y] == "#" {
					pos = next(pos, invDir(curr.d))
					break
				}
			}
			curr.pos = pos
		}
	}
	return 1000*(curr.pos.x+1) + 4*(curr.pos.y+1) + facing(curr.d)
}

func part2(s string) int {
	tab, mov := format(s)
	mapping3d := make(map[point]int)
	cubeSize := 50
	if len(tab) < 20 {
		cubeSize = 4
	}
	k := 0
	for i, line := range tab {
		for j, c := range line {
			if c == "." {
				_, ok := mapping3d[point{i / cubeSize, j / cubeSize}]
				if !ok {
					mapping3d[point{i / cubeSize, j / cubeSize}] = k
					k++
				}
			}
		}
	}
	curr := state{}
	curr.d = "R"
loop:
	for i, line := range tab {
		for j, c := range line {
			if c == "." {
				curr.pos = point{i, j}
				break loop
			}
		}
	}
	for _, instr := range mov {
		if instr == "L" || instr == "R" {
			curr.d = changeDir(curr.d, instr)
		} else {
			n, _ := strconv.Atoi(instr)
			pos := curr.pos
			for i := 0; i < n; i++ {
				pos = next(pos, curr.d)
				if pos.x < 0 || pos.y < 0 || pos.x >= len(tab) || pos.y >= len(tab[pos.x]) || (tab[pos.x][pos.y] != "." && tab[pos.x][pos.y] != "#") {
					oldPos := next(pos, invDir(curr.d))
					oldDir := curr.d
					if cubeSize == 4 {
						pos, curr.d = nextCube4(oldPos, curr.d, mapping3d[point{oldPos.x / cubeSize, oldPos.y / cubeSize}])
					}
					if cubeSize == 50 {
						pos, curr.d = nextCube50(oldPos, curr.d, mapping3d[point{oldPos.x / cubeSize, oldPos.y / cubeSize}])
					}
					if tab[pos.x][pos.y] == "#" {
						pos = oldPos
						curr.d = oldDir
						break
					}
				}
				if tab[pos.x][pos.y] == "." {
					continue
				}
				if tab[pos.x][pos.y] == "#" {
					pos = next(pos, invDir(curr.d))
					break
				}
			}
			curr.pos = pos
		}
	}
	return 1000*(curr.pos.x+1) + 4*(curr.pos.y+1) + facing(curr.d)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day22/test.data")

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

	content, err = os.ReadFile("../../../inputs/2022/day22/input.data")

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
