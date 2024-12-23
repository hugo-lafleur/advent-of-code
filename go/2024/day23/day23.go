package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
	"time"
)

func parse(s string) map[string]map[string]bool {
	var result = make(map[string]map[string]bool)
	var lines = strings.Split(s, "\n")
	for i := range lines {
		parts := strings.Split(lines[i], "-")
		if result[parts[0]] == nil {
			result[parts[0]] = map[string]bool{}
		}
		if result[parts[1]] == nil {
			result[parts[1]] = map[string]bool{}
		}
		result[parts[0]][parts[1]] = true
		result[parts[1]][parts[0]] = true
	}
	return result
}

func mapToSlices(m map[string]bool) []string {
	var result []string
	for key := range m {
		result = append(result, key)
	}
	return result
}

func part1(s string) int {
	var graph = parse(s)
	var result = map[[3]string]bool{}
	for key := range graph {
		if key[0] == 't' {
			for neigh1 := range graph[key] {
				for neigh2 := range graph[key] {
					if graph[neigh1][neigh2] {
						var arr = []string{key, neigh1, neigh2}
						slices.Sort(arr)
						result[[3]string{arr[0], arr[1], arr[2]}] = true
					}
				}
			}
		}
	}
	return len(result)
}

func part2(s string) string {
	var graph = parse(s)
	var V = make(map[string]bool)
	for key := range graph {
		V[key] = true
	}
	var result []string
	var bronKerbosch func(R, P, X map[string]bool)
	bronKerbosch = func(R, P, X map[string]bool) {
		if len(P) == 0 && len(X) == 0 {
			clique := mapToSlices(R)
			if len(clique) > len(result) {
				result = clique
			}
		}
		for v := range P {
			var newR = maps.Clone(R)
			var newP = maps.Clone(P)
			var newX = maps.Clone(X)
			newR[v] = true
			maps.DeleteFunc(newP, func(key string, val bool) bool { return !graph[v][key] })
			maps.DeleteFunc(newX, func(key string, val bool) bool { return !graph[v][key] })
			bronKerbosch(newR, newP, newX)
			delete(P, v)
			X[v] = true
		}
	}
	bronKerbosch(make(map[string]bool), V, make(map[string]bool))
	slices.Sort(result)
	return strings.Join(result, ",")
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day23/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2024/day23/input.txt")

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
