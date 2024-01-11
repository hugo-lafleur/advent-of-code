package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

type point struct {
	x, y int
}

func format(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func build(p point, f int) string {
	c := 0
	c += p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y
	c += f
	ones := 0
	for c != 0 {
		ones += c & 1
		c >>= 1
	}
	if ones%2 == 0 {
		return "open"
	} else {
		return "wall"
	}
}

func isValid(p point, d int) bool {
	return p.x > -1 && p.y > -1 && p.x < d && p.y < d
}

func part1(s string) int {
	f := format(s)
	graph := dijkstra.NewGraph()
	layout := make(map[point]string)
	d := 100
	var x, y int
	if f == 10 {
		x = 7
		y = 4
	} else {
		x = 31
		y = 39
	}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			graph.AddVertex(i*d + j)
			layout[point{i, j}] = build(point{i, j}, f)
		}
	}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			p := point{i, j}
			if layout[p] == "open" {
				if isValid(point{i + 1, j}, d) && layout[point{i + 1, j}] == "open" {
					graph.AddArc(i*d+j, (i+1)*d+j, 1)
				}
				if isValid(point{i, j + 1}, d) && layout[point{i, j + 1}] == "open" {
					graph.AddArc(i*d+j, i*d+j+1, 1)
				}
				if isValid(point{i - 1, j}, d) && layout[point{i - 1, j}] == "open" {
					graph.AddArc(i*d+j, (i-1)*d+j, 1)
				}
				if isValid(point{i, j - 1}, d) && layout[point{i, j - 1}] == "open" {
					graph.AddArc(i*d+j, i*d+j-1, 1)
				}
			}
		}
	}
	best, _ := graph.Shortest(1*d+1, x*d+y)
	return int(best.Distance)
}

func part2(s string) int {
	c := 0
	f := format(s)
	graph := dijkstra.NewGraph()
	layout := make(map[point]string)
	d := 60
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			graph.AddVertex(i*d + j)
			layout[point{i, j}] = build(point{i, j}, f)
		}
	}
	for i := 0; i < d; i++ {
		for j := 0; j < d; j++ {
			p := point{i, j}
			if layout[p] == "open" {
				if isValid(point{i + 1, j}, d) && layout[point{i + 1, j}] == "open" {
					graph.AddArc(i*d+j, (i+1)*d+j, 1)
				}
				if isValid(point{i, j + 1}, d) && layout[point{i, j + 1}] == "open" {
					graph.AddArc(i*d+j, i*d+j+1, 1)
				}
				if isValid(point{i - 1, j}, d) && layout[point{i - 1, j}] == "open" {
					graph.AddArc(i*d+j, (i-1)*d+j, 1)
				}
				if isValid(point{i, j - 1}, d) && layout[point{i, j - 1}] == "open" {
					graph.AddArc(i*d+j, i*d+j-1, 1)
				}
			}
		}
	}
	for x := 0; x < d; x++ {
		for y := 0; y < d; y++ {
			best, err := graph.Shortest(1*d+1, x*d+y)
			if err == nil {
				if best.Distance < 51 {
					c++
				}
			}
		}
	}
	return c + 1 //+1 is for (1,1) (empty path)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day13/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day13/input.data")

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
