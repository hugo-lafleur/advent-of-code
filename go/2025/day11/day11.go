package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
)

func format(s string) map[string][]string {
	lines := strings.Split(s, "\n")
	result := map[string][]string{}
	for _, line := range lines {
		parts := strings.FieldsFunc(line, func(r rune) bool { return r == ' ' || r == ':' })
		result[parts[0]] = parts[1:]
	}
	return result
}

func part1(s string) int {
	list := format(s)
	g := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic())
	for key := range list {
		g.AddVertex(key)
	}
	g.AddVertex("out")
	for key := range list {
		for _, val := range list[key] {
			g.AddEdge(key, val)
		}
	}
	paths, _ := graph.AllPathsBetween(g, "you", "out")
	return len(paths)
}

func explore(adjMap map[string]map[string]graph.Edge[string], start, end string) int {
	cache := map[string]int{}
	var dfs func(node string) int
	dfs = func(node string) int {
		if val, ok := cache[node]; ok {
			return val
		}
		if node == end {
			return 1
		}
		total := 0
		for neigh := range adjMap[node] {
			total += dfs(neigh)
		}
		cache[node] = total
		return total
	}
	return dfs(start)
}

func part2(s string) int {
	list := format(s)
	g := graph.New(graph.StringHash, graph.Directed(), graph.Acyclic())
	for key := range list {
		g.AddVertex(key)
	}
	g.AddVertex("out")
	for key := range list {
		for _, val := range list[key] {
			g.AddEdge(key, val)
		}
	}
	adjMap, _ := g.AdjacencyMap()
	return explore(adjMap, "svr", "fft") * explore(adjMap, "fft", "dac") * explore(adjMap, "dac", "out")
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day11/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2025/day11/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2025/day11/input.txt")

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
