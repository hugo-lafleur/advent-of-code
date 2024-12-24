package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

func parse(s string) (map[string]int, [][4]string) {
	var wires, gates = map[string]int{}, [][4]string{}
	var lines = strings.Split(s, "\n")
	var i int
	for lines[i] != "" {
		var name string
		var value int
		fmt.Sscanf(lines[i], "%s %d", &name, &value)
		wires[name[:3]] = value
		i++
	}
	i++
	for i < len(lines) {
		var wire1, gate, wire2, wire3 string
		fmt.Sscanf(lines[i], "%s %s %s -> %s", &wire1, &gate, &wire2, &wire3)
		gates = append(gates, [4]string{wire1, gate, wire2, wire3})
		i++
	}
	return wires, gates
}

func part1(s string) int {
	var wires, gates = parse(s)
	var dq deque.Deque[[4]string]
	for i := range gates {
		dq.PushBack(gates[i])
	}
	for dq.Len() != 0 {
		curr := dq.PopFront()
		_, ok1 := wires[curr[0]]
		_, ok2 := wires[curr[2]]
		if ok1 && ok2 {
			switch curr[1] {
			case "AND":
				wires[curr[3]] = wires[curr[0]] & wires[curr[2]]
			case "OR":
				wires[curr[3]] = wires[curr[0]] | wires[curr[2]]
			case "XOR":
				wires[curr[3]] = wires[curr[0]] ^ wires[curr[2]]
			}
		} else {
			dq.PushBack(curr)
		}
	}
	var Z []string
	for key := range wires {
		if key[0] == 'z' {
			Z = append(Z, key)
		}
	}
	slices.Sort(Z)
	var result string
	for i := range Z {
		result = strconv.Itoa(wires[Z[i]]) + result
	}
	n, _ := strconv.ParseInt(result, 2, 64)
	return int(n)
}

func part2(s string) string {
	var _, gates = parse(s)
	var result []string
	for _, g := range gates {
		if g[1] == "XOR" {
			if (g[0][0] == 'x' && g[2][0] == 'y') || (g[0][0] == 'y' && g[2][0] == 'x') {
				if g[3][0] == 'z' && g[3] != "z00" {
					result = append(result, g[3])
				}
				for _, g2 := range gates {
					if g2[1] == "XOR" && (g2[0] == g[3] || g2[2] == g[3]) {
						if g2[3][0] != 'z' {
							result = append(result, g2[3])
						}
					}
					if g2[1] == "OR" && (g2[0] == g[3] || g2[2] == g[3]) {
						result = append(result, g[3])
					}
				}
			} else if g[3][0] != 'z' {
				result = append(result, g[3])
			}
		}
		if g[1] == "AND" {
			if (g[0][0] == 'x' && g[2][0] == 'y') || (g[0][0] == 'y' && g[2][0] == 'x') {
				if g[0] != "x00" && g[0] != "y00" {
					for _, g2 := range gates {
						if g2[1] != "OR" && (g2[0] == g[3] || g2[2] == g[3]) {
							result = append(result, g[3])
						}
					}
				}
			}
		}
		if g[3][0] == 'z' && g[3] != "z45" && g[1] != "XOR" {
			result = append(result, g[3])
		}
	}
	slices.Sort(result)
	result = slices.Compact(result)
	/*var g = graph.New(graph.StringHash, graph.Directed())
	for _, gate := range gates {
		g.AddVertex(gate[0])
		g.AddVertex(gate[2])
		g.AddVertex(gate[3])
		g.AddEdge(gate[0], gate[3])
		g.AddEdge(gate[2], gate[3])
	}

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g, file)*/
	return strings.Join(result, ",")
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day24/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day24/input.txt")

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
