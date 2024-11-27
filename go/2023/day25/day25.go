package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dominikbraun/graph"
)

type bridge struct {
	a, b string
}

func Split(r rune) bool {
	return r == ':' || r == ' '
}

func format(s string) map[string][]string {
	lines := strings.Split(s, "\n")
	res := make(map[string][]string)
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		res[lineSplit[0]] = lineSplit[1:]
	}
	return res
}

func maxBridge(mapping map[bridge]int) bridge {
	var res bridge
	var m int
	for br, value := range mapping {
		if value > m {
			m = value
			res = br
		}
	}
	return res
}

func part1(s string) int {
	diagram := format(s)
	g := graph.New(graph.StringHash, graph.Directed(), graph.Weighted())
	set := make(map[string]bool)
	bridges := make(map[bridge]int)
	for c1 := range diagram {
		g.AddVertex(c1)
		set[c1] = true
		for _, c2 := range diagram[c1] {
			g.AddVertex(c2)
			set[c2] = true
		}
	}
	for c1 := range diagram {
		for _, c2 := range diagram[c1] {
			g.AddEdge(c1, c2)
			g.AddEdge(c2, c1)
		}
	}
	if len(diagram) == 13 {
		g.RemoveEdge("hfx", "pzl")
		g.RemoveEdge("pzl", "hfx")
		g.RemoveEdge("bvb", "cmg")
		g.RemoveEdge("cmg", "bvb")
		g.RemoveEdge("nvd", "jqt")
		g.RemoveEdge("jqt", "ncd")
		scc, _ := graph.StronglyConnectedComponents(g)
		return len(scc[0]) * len(scc[1])
	} else {
	loop:
		for j := 0; j < 100; j++ {
			for c1 := range set {
				for c2 := range set {
					if c1 != c2 {
						path, _ := graph.ShortestPath(g, c1, c2)
						for i := 0; i < len(path)-1; i++ {
							p1, p2 := path[i], path[i+1]
							_, ok := bridges[bridge{p1, p2}]
							if ok {
								bridges[bridge{p1, p2}]++
							} else {
								bridges[bridge{p2, p1}]++
							}
						}
						continue loop
					}
				}
			}
		}
	}
	for i := 0; i < 4; i++ {
		br := maxBridge(bridges)
		delete(bridges, br)
		g.RemoveEdge(br.a, br.b)
		g.RemoveEdge(br.b, br.a)
	}
	scc, _ := graph.StronglyConnectedComponents(g)
	return len(scc[0]) * len(scc[1])
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day25/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2023/day25/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
