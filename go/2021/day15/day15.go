package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

type point struct {
	x, y int
}

func cycle(n int) int {
	if n > 9 {
		return cycle(n - 9)
	}
	return n
}

func format(s string) (map[point]int, int, int) {
	lines := strings.Split(s, "\n")
	res := make(map[point]int)
	for i, line := range lines {
		lineSplit := strings.Split(line, "")
		for j, x := range lineSplit {
			p := point{i, j}
			n, _ := strconv.Atoi(x)
			res[p] = n
		}
	}
	return res, len(lines), len(lines[0])
}

func format2(s string) (map[point]int, int, int) {
	lines := strings.Split(s, "\n")
	res := make(map[point]int)
	for i, line := range lines {
		lineSplit := strings.Split(line, "")
		for j, x := range lineSplit {
			for a := 0; a < 5; a++ {
				for b := 0; b < 5; b++ {
					p := point{(a * len(lines)) + i, (b * len(lines[i])) + j}
					n, _ := strconv.Atoi(x)
					res[p] = cycle(n + a + b)
				}
			}
		}
	}
	return res, 5 * len(lines), 5 * len(lines[0])
}

func neighs(p point) []point {
	return []point{{p.x + 1, p.y}, {p.x - 1, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}}
}

func hash(p point) string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func part1(s string) int {
	mapping, n, m := format(s)
	graph := dijkstra.NewGraph()
	for key := range mapping {
		graph.AddMappedVertex(hash(key))
	}
	for key := range mapping {
		for _, neigh := range neighs(key) {
			_, ok := mapping[neigh]
			if ok {
				graph.AddMappedArc(hash(key), hash(neigh), int64(mapping[neigh]))
			}
		}
	}
	srcID, _ := graph.GetMapping(hash(point{0, 0}))
	destID, _ := graph.GetMapping(hash(point{n - 1, m - 1}))
	best, _ := graph.Shortest(srcID, destID)
	return int(best.Distance)
}

func part2(s string) int {
	mapping, n, m := format2(s)
	graph := dijkstra.NewGraph()
	for key := range mapping {
		graph.AddMappedVertex(hash(key))
	}
	for key := range mapping {
		for _, neigh := range neighs(key) {
			_, ok := mapping[neigh]
			if ok {
				graph.AddMappedArc(hash(key), hash(neigh), int64(mapping[neigh]))
			}
		}
	}
	srcID, _ := graph.GetMapping(hash(point{0, 0}))
	destID, _ := graph.GetMapping(hash(point{n - 1, m - 1}))
	best, _ := graph.Shortest(srcID, destID)
	return int(best.Distance)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day15/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day15/input.txt")

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
