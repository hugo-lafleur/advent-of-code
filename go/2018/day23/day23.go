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
	x, y, z int
}

type nanobot struct {
	p      point
	radius int
}

type interval struct {
	a, b int
}

func Split(r rune) bool {
	return r == '<' || r == '>' || r == ',' || r == '='
}

func format(s string) []nanobot {
	lines := strings.Split(s, "\n")
	res := []nanobot{}
	for _, line := range lines {
		strs := strings.FieldsFunc(line, Split)
		x, _ := strconv.Atoi(strs[1])
		y, _ := strconv.Atoi(strs[2])
		z, _ := strconv.Atoi(strs[3])
		r, _ := strconv.Atoi(strs[5])
		res = append(res, nanobot{point{x, y, z}, r})
	}
	return res
}

func maxRadius(nanobotList []nanobot) nanobot {
	r := nanobotList[0].radius
	index := 0
	for i, nanobot := range nanobotList {
		if nanobot.radius > r {
			r = nanobot.radius
			index = i
		}
	}
	return nanobotList[index]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(n1, n2 point) int {
	return abs(n1.x-n2.x) + abs(n1.y-n2.y) + abs(n1.z-n2.z)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func nanobotToInterval(n nanobot) interval {
	return interval{max(0, distance(point{0, 0, 0}, n.p)-n.radius), distance(point{0, 0, 0}, n.p) + n.radius}
}

func numberOfIntervals(a int, l []interval) int {
	c := 0
	for _, i := range l {
		if i.a <= a && a <= i.b {
			c++
		}
	}
	return c
}

func part1(s string) int {
	c := 0
	nanobotList := format(s)
	maxRadiusNanobot := maxRadius(nanobotList)
	for _, n := range nanobotList {
		if distance(maxRadiusNanobot.p, n.p) <= maxRadiusNanobot.radius {
			c++
		}
	}
	return c
}

func part2(s string) int {
	nanobotList := format(s)
	var intervals []interval
	for _, n := range nanobotList {
		intervals = append(intervals, nanobotToInterval(n))
	}
	distanceToNumberOfIntervals := make(map[int]int)
	for _, inter := range intervals {
		distanceToNumberOfIntervals[inter.a] = numberOfIntervals(inter.a, intervals)
	}
	var m int
	for _, value := range distanceToNumberOfIntervals {
		m = value
	}
	for _, value := range distanceToNumberOfIntervals {
		if value > m {
			m = value
		}
	}
	var shortest int
	for key, value := range distanceToNumberOfIntervals {
		if value == m {
			shortest = key
		}
	}
	for key, value := range distanceToNumberOfIntervals {
		if value == m && key < shortest {
			shortest = key
		}
	}
	return shortest
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day23/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day23/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day23/input.data")

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
