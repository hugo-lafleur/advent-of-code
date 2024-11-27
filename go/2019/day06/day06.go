package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ")"))
	}
	return res
}

func numberOfOrbits(s string, orbits map[string]string, memory map[string]int) (int, map[string]int) {
	direct, ok := orbits[s]
	res := 0
	if !ok {
		res = 0
		memory[s] = res
		return res, memory
	}
	directNumber, ok2 := memory[direct]
	if ok2 {
		res = 1 + directNumber
		memory[s] = res
		return res, memory
	}
	res, _ = numberOfOrbits(direct, orbits, memory)
	res++
	memory[s] = res
	return res, memory
}

func part1(s string) int {
	c := 0
	var n int
	orbits := make(map[string]string)
	memory := make(map[string]int)
	l := format(s)
	for _, line := range l {
		orbits[line[1]] = line[0]
	}
	for key := range orbits {
		n, memory = numberOfOrbits(key, orbits, memory)
		c += n
	}
	return c
}

func part2(s string) int {
	g := dijkstra.NewGraph()
	l := format(s)
	for _, line := range l {
		g.AddMappedVertex(line[0])
		g.AddMappedVertex(line[1])
		g.AddMappedArc(line[0], line[1], 1)
		g.AddMappedArc(line[1], line[0], 1)
	}
	idYOU, _ := g.GetMapping("YOU")
	idSAN, _ := g.GetMapping("SAN")
	best, _ := g.Shortest(idYOU, idSAN)
	return int(best.Distance) - 2
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day06/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day06/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day06/input.txt")

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
