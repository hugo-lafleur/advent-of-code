package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"slices"
	"strings"
	"time"

	"gonum.org/v1/gonum/stat"
)

type Point struct {
	x, y int
}
type Robot struct {
	initial Point
	vx, vy  int
}

func parse(s string) []Robot {
	var lines = strings.Split(s, "\n")
	var result []Robot
	for i := range lines {
		var px, py, vx, vy int
		fmt.Sscanf(lines[i], "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		result = append(result, Robot{initial: Point{px, py}, vx: vx, vy: vy})
	}
	return result
}

func position(r Robot, t int) Point {
	return Point{r.initial.x + t*r.vx, r.initial.y + t*r.vy}
}

func part1(s string) int {
	var robots = parse(s)
	var m, n = 103, 101
	var t = 100
	if len(robots) == 12 {
		m, n = 7, 11
	}
	var grid = make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}
	for _, r := range robots {
		p := position(r, t)
		p.x = ((p.x % n) + n) % n
		p.y = ((p.y % m) + m) % m
		grid[p.y][p.x]++
	}
	var quadrants = make([]int, 4)
	for i := range grid {
		for j := range grid[i] {
			if i == m/2 || j == n/2 {
				continue
			}
			quadrants[2*((i-1)/(m/2))+((j-1)/(n/2))] += grid[i][j]
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func part2(s string) int { //using variance to find most "organized" time
	var robots = parse(s)
	var m, n = 103, 101
	if len(robots) == 12 {
		m, n = 7, 11
	}
	var bx, by int //
	var minVarx, minVary = math.MaxFloat64, math.MaxFloat64
	for t := range max(m, n) {
		var listx = make([]float64, len(robots))
		var listy = make([]float64, len(robots))
		for i, r := range robots {
			p := position(r, t)
			p.x = ((p.x % n) + n) % n
			p.y = ((p.y % m) + m) % m
			listx[i] = float64(p.x)
			listy[i] = float64(p.y)
		}
		varx := stat.Variance(listx, nil)
		if varx < minVarx {
			minVarx = varx
			bx = t
		}
		vary := stat.Variance(listy, nil)
		if vary < minVary {
			minVary = vary
			by = t
		}
	}
	var Z big.Int
	var G = big.NewInt(int64(m))
	var N = big.NewInt(int64(n))
	return by + (int(Z.ModInverse(G, N).Int64())*(bx-by)%n)*m
}

func part2_old(s string) int { //bruteforce : finding a long straight line of robots
	var robots = parse(s)
	var m, n = 103, 101
	if len(robots) == 12 {
		m, n = 7, 11
	}
	for t := 0; ; t++ {
		var grid = make([][]byte, m)
		for i := range grid {
			grid[i] = slices.Repeat([]byte{' '}, n)
		}
		for _, r := range robots {
			p := position(r, t)
			p.x = ((p.x % n) + n) % n
			p.y = ((p.y % m) + m) % m
			grid[p.y][p.x] = '*'
		}
		for i := range grid {
			if strings.Contains(string(grid[i]), "*******************************") {
				return t
			}
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day14/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day14/input.txt")

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
