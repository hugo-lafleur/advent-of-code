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

type state struct {
	p    point
	tool string
}

func Split(r rune) bool {
	return r == '\n' || r == ' ' || r == ','
}

func format(s string) (int, point) {
	strs := strings.FieldsFunc(s, Split)
	depth, _ := strconv.Atoi(strs[1])
	x, _ := strconv.Atoi(strs[3])
	y, _ := strconv.Atoi(strs[4])
	return depth, point{x, y}
}

func geologicIndex(p point, erosionLevel map[point]int) int {
	if p.y == 0 {
		return p.x * 16807
	}
	if p.x == 0 {
		return p.y * 48271
	}
	return erosionLevel[point{p.x - 1, p.y}] * erosionLevel[point{p.x, p.y - 1}]
}

func hashState(st state) string {
	return strconv.Itoa(st.p.x) + "," + strconv.Itoa(st.p.y) + "," + st.tool
}

func part1(s string) int {
	c := 0
	depth, target := format(s)
	erosionLevel := make(map[point]int)
	for i := 0; i <= target.x; i++ {
		for j := 0; j <= target.y; j++ {
			p := point{i, j}
			var index int
			if (i == 0 && j == 0) || (i == target.x && j == target.y) {
				index = 0
			} else {
				index = geologicIndex(p, erosionLevel)
			}
			erosionLevel[p] = (index + depth) % 20183
		}
	}
	for _, value := range erosionLevel {
		c += (value % 3)
	}
	return c
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areAdjacent(p1, p2 point) bool {
	if abs(p1.x-p2.x) == 0 {
		return abs(p1.y-p2.y) == 1
	}
	if abs(p1.y-p2.y) == 0 {
		return abs(p1.x-p2.x) == 1
	}
	return false
}

func part2(s string) int {
	depth, target := format(s)
	erosionLevel := make(map[point]int)
	nature := make(map[point]string)
	correction := 30 // you may want to change that value based on your input : allow you to go beyond target point but it should not be too big as it increase run time
	for i := 0; i <= target.x+correction; i++ {
		for j := 0; j <= target.y+correction; j++ {
			p := point{i, j}
			var index int
			if (i == 0 && j == 0) || (i == target.x && j == target.y) {
				index = 0
			} else {
				index = geologicIndex(p, erosionLevel)
			}
			erosionLevel[p] = (index + depth) % 20183
		}
	}
	for key, value := range erosionLevel {
		switch value % 3 {
		case 0:
			nature[key] = "r"
		case 1:
			nature[key] = "w"
		case 2:
			nature[key] = "n"
		}
	}
	dijkstra := dijkstra.NewGraph()
	statesPossibles := []state{}
	for key, value := range nature {
		var st state
		switch value {
		case "r":
			st = state{key, "gear"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
			st = state{key, "torch"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
		case "w":
			st = state{key, "gear"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
			st = state{key, "none"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
		case "n":
			st = state{key, "none"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
			st = state{key, "torch"}
			dijkstra.AddMappedVertex(hashState(st))
			statesPossibles = append(statesPossibles, st)
		}
	}

	for _, st1 := range statesPossibles {
		for _, st2 := range statesPossibles {
			if areAdjacent(st1.p, st2.p) && st1.tool == st2.tool {
				dijkstra.AddMappedArc(hashState(st1), hashState(st2), 1)
			}
			if st1.p == st2.p && st1.tool != st2.tool {
				dijkstra.AddMappedArc(hashState(st1), hashState(st2), 7)
			}
		}
	}

	sourceID, _ := dijkstra.GetMapping(hashState(state{point{0, 0}, "torch"}))
	destID, _ := dijkstra.GetMapping(hashState(state{target, "torch"}))

	best, err := dijkstra.Shortest(sourceID, destID)
	if err == nil {
		return int(best.Distance)
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day22/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day22/input.data")

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
