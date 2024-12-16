package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra/v2"
)

type Point struct {
	x, y int
}

type State struct {
	x, y int
	dir  int
}

func parse(s string) [][]byte {
	var lines = strings.Split(s, "\n")
	var grid = make([][]byte, len(lines))
	for i := range grid {
		grid[i] = []byte(lines[i])
	}
	return grid
}

func buildGraph(s string) (dijkstra.MappedGraph[State], State, State) {
	var grid = parse(s)
	var graph = dijkstra.NewMappedGraph[State]()
	var start, end State
	var dirs = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				start = State{i, j, 1}
				for k := range dirs {
					graph.AddEmptyVertex(State{i, j, k})
				}
			} else if grid[i][j] == 'E' {
				end = State{i, j, -1}
				graph.AddEmptyVertex(State{i, j, -1})
			} else if grid[i][j] != '#' {
				for k := range dirs {
					x, y := i+dirs[k][0], j+dirs[k][1]
					if grid[x][y] != '#' {
						graph.AddEmptyVertex(State{i, j, k})
					}
				}
			}
		}
	}
	for i := range dirs {
		if i != start.dir {
			graph.AddArc(start, State{start.x, start.y, i}, 1000)
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '#' {
				for k := range dirs {
					x, y := i+dirs[k][0], j+dirs[k][1]
					if grid[x][y] == 'E' {
						graph.AddArc(State{i, j, k}, State{x, y, -1}, 1)
					} else {
						if grid[x][y] != '#' {
							for l := range dirs {
								x2, y2 := x+dirs[l][0], y+dirs[l][1]
								if grid[x2][y2] != '#' {
									if k == l {
										graph.AddArc(State{i, j, k}, State{x, y, l}, 1)
									} else if k%2 != l%2 {
										graph.AddArc(State{i, j, k}, State{x, y, l}, 1001)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return graph, start, end
}

func part1(s string) int {
	var graph, start, end = buildGraph(s)
	best, _ := graph.Shortest(start, end)
	return int(best.Distance)
}

func part2(s string) int {
	var graph, start, end = buildGraph(s)
	var set = make(map[Point]bool)
	best, _ := graph.ShortestAll(start, end)
	for _, path := range best.Paths {
		for _, state := range path {
			set[Point{state.x, state.y}] = true
		}
	}
	return len(set)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day16/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day16/input.txt")

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
