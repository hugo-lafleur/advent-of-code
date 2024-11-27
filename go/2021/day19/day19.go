package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y, z int
}

type couple struct {
	a, b point
}

func format(s string) [][]point {
	res := [][]point{}
	lines := strings.Split(s, "\n")
	var list []point
	for _, line := range lines {
		if strings.Contains(line, ",") {
			lineSplit := strings.Split(line, ",")
			a, _ := strconv.Atoi(lineSplit[0])
			b, _ := strconv.Atoi(lineSplit[1])
			c, _ := strconv.Atoi(lineSplit[2])
			list = append(list, point{a, b, c})
		}
		if line == "" {
			res = append(res, list)
			list = []point{}
		}
	}
	res = append(res, list)
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dist(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z)
}

func plus(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y, p1.z + p2.z}
}

func minus(p1, p2 point) point {
	return point{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}
}

func rotations(p point) []point {
	temp := []int{p.x, p.y, p.z}
	res := []point{}
	for _, x := range []int{0, 1, 2} {
		for _, y := range []int{0, 1, 2} {
			for _, z := range []int{0, 1, 2} {
				if x != y && x != z && y != z {
					res = append(res, point{temp[x], temp[y], temp[z]})
					res = append(res, point{temp[x], temp[y], -temp[z]})
					res = append(res, point{temp[x], -temp[y], temp[z]})
					res = append(res, point{temp[x], -temp[y], -temp[z]})
					res = append(res, point{-temp[x], temp[y], temp[z]})
					res = append(res, point{-temp[x], temp[y], -temp[z]})
					res = append(res, point{-temp[x], -temp[y], temp[z]})
					res = append(res, point{-temp[x], -temp[y], -temp[z]})
				}
			}
		}
	}
	return res
}

func listDist(p point, l []point) []int {
	res := []int{}
	for _, x := range l {
		res = append(res, dist(p, x))
	}
	return res
}

func inter(l1, l2 []int) []int {
	res := []int{}
	for _, x := range l1 {
		if slices.Contains(l2, x) {
			res = append(res, x)
		}
	}
	return res
}

func align(s1, s2 []point) (bool, []point, point) {
	matches := []couple{}
	for _, p1 := range s1 {
		for _, p2 := range s2 {
			if len(inter(listDist(p1, s1), listDist(p2, s2))) >= 12 {
				matches = append(matches, couple{p1, p2})
			}
		}
	}
	if len(matches) == 0 {
		return false, []point{}, point{}
	}
	rots := make(map[point][]point)
	for _, p2 := range s2 {
		rots[p2] = rotations(p2)
	}
	var rotation int
	for i := 0; i < 48; i++ {
		set := make(map[point]bool)
		for _, c := range matches {
			set[minus(c.a, rots[c.b][i])] = true
		}
		if len(set) == 1 {
			rotation = i
			break
		}
	}
	scanner_pos := minus(matches[0].a, rots[matches[0].b][rotation])
	res := []point{}
	for _, p2 := range s2 {
		res = append(res, plus(rots[p2][rotation], scanner_pos))
	}
	return true, res, scanner_pos
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func part1(s string) int {
	scanners := format(s)
	aligned_scanners := [][]point{scanners[0]}
	other_scanners := scanners[1:]
next_scanner:
	for len(other_scanners) != 0 {
		scan := other_scanners[0]
		other_scanners = other_scanners[1:]
		for _, refS := range aligned_scanners {
			res, temp, _ := align(refS, scan)
			if res {
				aligned_scanners = append(aligned_scanners, temp)
				continue next_scanner
			}
		}
		other_scanners = append(other_scanners, scan)
	}
	set := make(map[point]bool)
	for _, scan := range aligned_scanners {
		for _, p := range scan {
			set[p] = true
		}
	}
	return len(set)
}

func part2(s string) int {
	scanners := format(s)
	aligned_scanners := [][]point{scanners[0]}
	other_scanners := scanners[1:]
	scanners_positions := []point{{0, 0, 0}}
next_scanner:
	for len(other_scanners) != 0 {
		scan := other_scanners[0]
		other_scanners = other_scanners[1:]
		for _, refS := range aligned_scanners {
			res, temp, scan_pos := align(refS, scan)
			if res {
				aligned_scanners = append(aligned_scanners, temp)
				scanners_positions = append(scanners_positions, scan_pos)
				continue next_scanner
			}
		}
		other_scanners = append(other_scanners, scan)
	}
	var res int
	for i := 0; i < len(scanners_positions); i++ {
		for j := i + 1; j < len(scanners_positions); j++ {
			res = max(res, dist(scanners_positions[i], scanners_positions[j]))
		}
	}
	return res
}
func main() {
	content, err := os.ReadFile("../../../inputs/2021/day19/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day19/input.txt")

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
