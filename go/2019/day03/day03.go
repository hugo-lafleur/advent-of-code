package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type segment struct {
	start, end complex128
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ","))
	}
	return res
}

func next(curr complex128, s string) (complex128, segment) {
	var end complex128
	d := string(s[0])
	n, _ := strconv.Atoi(s[1:])
	switch d {
	case "U":
		end = curr + complex(float64(n), 0)*1i
	case "D":
		end = curr + complex(float64(n), 0)*(-1i)
	case "L":
		end = curr - complex(float64(n), 0)
	case "R":
		end = curr + complex(float64(n), 0)
	}
	return end, segment{curr, end}
}

func inInterval(x int, a, b int) bool {
	if a <= b {
		return x >= a && x <= b
	}
	return inInterval(x, b, a)
}

func intersection(seg1, seg2 segment) (bool, complex128) {
	a, b, c, d := seg1.start, seg1.end, seg2.start, seg2.end
	d1 := b - a
	d2 := d - c
	if imag(d1*d2) == 0 && real(d1*d2) != 0 {
		return false, 0
	}
	if imag(d1) == 0 {
		if inInterval(int(imag(a)), int(imag(c)), int(imag(d))) && inInterval(int(real(c)), int(real(a)), int(real(b))) {
			return true, complex(real(c), imag(a))
		}
		return false, 0
	}
	return intersection(seg2, seg1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(c1, c2 complex128) int {
	return abs(int(real(c1))-int(real(c2))) + abs(int(imag(c1))-int(imag(c2)))
}

func minMap(m map[complex128]int) int {
	var min int
	for _, value := range m {
		min = value
	}
	for _, value := range m {
		if value < min {
			min = value
		}
	}
	return min
}

func part1(s string) int {
	wires := format(s)
	wire1, wire2 := wires[0], wires[1]
	curr := complex128(0)
	var seg segment
	wire1List := []segment{}
	wire2List := []segment{}
	intersections := make(map[complex128]int)
	for _, w := range wire1 {
		curr, seg = next(curr, w)
		wire1List = append(wire1List, seg)

	}
	curr = complex128(0)
	for _, w := range wire2 {
		curr, seg = next(curr, w)
		wire2List = append(wire2List, seg)
	}
	for _, seg1 := range wire1List {
		for _, seg2 := range wire2List {
			ok, p := intersection(seg1, seg2)
			if ok && p != 0 {
				intersections[p] = distance(p, 0)
			}
		}
	}
	return minMap(intersections)
}

func part2(s string) int {
	wires := format(s)
	wire1, wire2 := wires[0], wires[1]
	curr := complex128(0)
	var seg segment
	wire1List := []segment{}
	wire2List := []segment{}
	intersections := make(map[complex128]int)
	steps := make(map[complex128]int)
	for _, w := range wire1 {
		curr, seg = next(curr, w)
		wire1List = append(wire1List, seg)

	}
	curr = complex128(0)
	for _, w := range wire2 {
		curr, seg = next(curr, w)
		wire2List = append(wire2List, seg)
	}
	for _, seg1 := range wire1List {
		for _, seg2 := range wire2List {
			ok, p := intersection(seg1, seg2)
			if ok && p != 0 {
				intersections[p] = distance(p, 0)
			}
		}
	}
	for key := range intersections {
		l := 0
		for _, seg := range wire1List {
			ok, _ := intersection(seg, segment{key, key})
			if ok {
				steps[key] += l + distance(key, seg.start)
				break
			}
			l += distance(seg.end, seg.start)
		}
		l = 0
		for _, seg := range wire2List {
			ok, _ := intersection(seg, segment{key, key})
			if ok {
				steps[key] += l + distance(key, seg.start)
				break
			}
			l += distance(seg.end, seg.start)
		}
	}
	return minMap(steps)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day03/test.data")

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

	content, err = os.ReadFile("../../../inputs/2019/day03/input.data")

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
