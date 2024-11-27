package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
)

type point struct {
	x, y, z, t int
}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '	'
}

func format(s string) []point {
	res := []point{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		splitLines := strings.FieldsFunc(line, Split)
		x, _ := strconv.Atoi(splitLines[0])
		y, _ := strconv.Atoi(splitLines[1])
		z, _ := strconv.Atoi(splitLines[2])
		t, _ := strconv.Atoi(splitLines[3])
		res = append(res, point{x, y, z, t})
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y) + abs(p1.z-p2.z) + abs(p1.t-p2.t)
}

func pointHash(p point) string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y) + "," + strconv.Itoa(p.z) + "," + strconv.Itoa(p.t)
}

func part1(s string) int {
	l := format(s)
	g := graph.New(pointHash, graph.Directed())
	for _, p := range l {
		g.AddVertex(p)
	}
	for _, p1 := range l {
		for _, p2 := range l {
			if p1 != p2 && distance(p1, p2) <= 3 {
				g.AddEdge(pointHash(p1), pointHash(p2))
			}
		}
	}
	scc, _ := graph.StronglyConnectedComponents(g)
	return len(scc)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day25/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
