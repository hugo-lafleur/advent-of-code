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

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func exp(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * exp(x, n-1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2aux(x int) point {
	k := 1
	for exp(k, 2) < x {
		k = k + 2
	}
	start := exp(k, 2)
	i := k / 2
	j := -k / 2
	if x <= start && x >= start-k+1 {
		return point{(i + x - start), j}
	}
	i = i - k + 1
	start = start - k + 1
	if x <= start && x >= start-k+1 {
		return point{i, (j - x + start)}
	}
	j = j + k - 1
	start = start - k + 1
	if x <= start && x >= start-k+1 {
		return point{(i - x + start), j}
	}
	i = i + k - 1
	start = start - k + 1
	return point{i, (j + x - start)}
}

func part1(s string) int {
	x := format(s)
	k := 1
	for exp(k, 2) < x {
		k = k + 2
	}
	start := exp(k, 2)
	i := k / 2
	j := -k / 2
	if x <= start && x >= start-k+1 {
		return abs(j) + abs(i+x-start)
	}
	i = i - k + 1
	start = start - k + 1
	if x <= start && x >= start-k+1 {
		return abs(i) + abs(j-x+start)
	}
	j = j + k - 1
	start = start - k + 1
	if x <= start && x >= start-k+1 {
		return abs(j) + abs(i-x+start)
	}
	i = i + k - 1
	start = start - k + 1
	return abs(i) + abs(j+x-start)

}

func part2(s string) int {
	n := format(s)
	intToPoint := make(map[int]point)
	pointToInt := make(map[point]int)
	intToValue := make(map[int]int)
	intToPoint[1] = point{0, 0}
	intToValue[1] = 1
	pointToInt[point{0, 0}] = 1
	k := 2
	for {
		p := part2aux(k)
		intToPoint[k] = p
		pointToInt[p] = k
		v := 0
		for _, i := range []int{p.x - 1, p.x + 1, p.x} {
			for _, j := range []int{p.y - 1, p.y + 1, p.y} {
				s, ok := pointToInt[point{i, j}]
				if ok {
					v += intToValue[s]
				}
			}
		}
		intToValue[k] = v
		if v > n {
			return v
		}
		k++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day03/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day03/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day03/input.data")

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
