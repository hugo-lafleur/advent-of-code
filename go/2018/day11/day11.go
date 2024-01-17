package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type point struct {
	x, y int
}

type square struct {
	x, y, n int
}

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func hundredsDigit(n int) int {
	res := n % 1000
	res = res / 100
	return res
}

func power3x3(p point, powerLevel map[point]int) int {
	if p.x+2 > 300 || p.y+2 > 300 {
		return 0
	}
	res := 0
	for _, i := range []int{0, 1, 2} {
		for _, j := range []int{0, 1, 2} {
			res += powerLevel[point{p.x + i, p.y + j}]
		}
	}
	return res
}

func powerNxN(p point, powerLevel map[point]int, n int) int {
	if p.x+n-1 > 300 || p.y+n-1 > 300 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += powerLevel[point{p.x + i, p.y + j}]
		}
	}
	return res
}

func maximum(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func part1(s string) point {
	max := 0
	var res point
	serialNumber := format(s)
	powerLevel := make(map[point]int)
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			p := point{i, j}
			rackID := i + 10
			power := rackID * j
			power += serialNumber
			power *= rackID
			power = hundredsDigit(power)
			power -= 5
			powerLevel[p] = power
		}
	}
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			p := point{i, j}
			totalPower := power3x3(p, powerLevel)
			if totalPower > max {
				max = totalPower
				res = p
			}
		}
	}
	return res
}

func part2(s string) square {
	max := 0
	var res square
	serialNumber := format(s)
	powerLevel := make(map[point]int)
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			p := point{i, j}
			rackID := i + 10
			power := rackID * j
			power += serialNumber
			power *= rackID
			power = hundredsDigit(power)
			power -= 5
			powerLevel[p] = power
		}
	}
	for i := 1; i < 301; i++ {
		for j := 1; j < 301; j++ {
			for n := 4; maximum(i, j)+n < 301; n++ {
				p := point{i, j}
				totalPower := powerNxN(p, powerLevel, n)
				if totalPower < 0 {
					break
				}
				if totalPower > max {
					max = totalPower
					res = square{i, j, n}
				}
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day11/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day11/input.data")

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
