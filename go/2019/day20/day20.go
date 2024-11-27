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

type point3d struct {
	x, y  int
	level int
}

func hashPoint(p point) string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y)
}

func hashPoint3d(p point3d) string {
	return strconv.Itoa(p.x) + "," + strconv.Itoa(p.y) + "," + strconv.Itoa(p.level)
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func addPoint(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func addPoint3d(p1, p2 point3d) point3d {
	return point3d{p1.x + p2.x, p1.y + p2.y, p1.level}
}

func isCloseToPortal(p point, tab [][]string) (bool, string, string) {
	for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		firstPortal := addPoint(p, offset)
		i, j := firstPortal.x, firstPortal.y
		if tab[i][j][0] >= 'A' {
			for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				secondPortal := addPoint(firstPortal, offset)
				i, j := secondPortal.x, secondPortal.y
				if tab[i][j][0] >= 'A' {
					return true, tab[firstPortal.x][firstPortal.y], tab[i][j]
				}
			}
		}
	}
	return false, "", ""
}

func isOutside(p point, tab [][]string) bool {
	x, y := p.x, p.y
	n, m := len(tab), len(tab[0])
	return x == 2 || y == 2 || x == n-3 || y == m-3
}

func part1(s string) int {
	tab := format(s)
	graph := dijkstra.NewGraph()
	portals := make(map[string][]point)
	for i := range tab {
		for j := range tab[i] {
			s := tab[i][j]
			p := point{i, j}
			if s == "." {
				graph.AddMappedVertex(hashPoint(p))
				ok, A, B := isCloseToPortal(p, tab)
				if ok {
					_, exist := portals[A+B]
					if exist {
						portals[A+B] = append(portals[A+B], p)
					} else {
						portals[B+A] = append(portals[B+A], p)
					}
				}
			}
		}
	}
	for i := range tab {
		for j := range tab[i] {
			s := tab[i][j]
			p := point{i, j}
			if s == "." {
				for _, offset := range []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					next := addPoint(p, offset)
					if tab[next.x][next.y] == "." {
						graph.AddMappedArc(hashPoint(p), hashPoint(next), 1)
					}
				}
			}
		}
	}
	for _, value := range portals {
		if len(value) == 2 {
			graph.AddMappedArc(hashPoint(value[0]), hashPoint(value[1]), 1)
			graph.AddMappedArc(hashPoint(value[1]), hashPoint(value[0]), 1)
		}
	}
	srcID, _ := graph.GetMapping(hashPoint(portals["AA"][0]))
	destID, _ := graph.GetMapping(hashPoint(portals["ZZ"][0]))
	best, _ := graph.Shortest(srcID, destID)
	return int(best.Distance)
}

func part2(s string) int {
	tab := format(s)
	graph := dijkstra.NewGraph()
	portals := make(map[string][]point)
	maxLevel := 50
	for level := 0; level < maxLevel; level++ {
		for i := range tab {
			for j := range tab[i] {
				s := tab[i][j]
				p := point3d{i, j, level}
				if s == "." {
					graph.AddMappedVertex(hashPoint3d(p))
					ok, A, B := isCloseToPortal(point{i, j}, tab)
					if ok && level == 0 {
						_, exist := portals[A+B]
						if exist {
							portals[A+B] = append(portals[A+B], point{i, j})
						} else {
							portals[B+A] = append(portals[B+A], point{i, j})
						}
					}
				}
			}
		}
	}
	for level := 0; level < maxLevel; level++ {
		for i := range tab {
			for j := range tab[i] {
				s := tab[i][j]
				p := point3d{i, j, level}
				if s == "." {
					for _, offset := range []point3d{{1, 0, level}, {-1, 0, level}, {0, 1, level}, {0, -1, level}} {
						next := addPoint3d(p, offset)
						if tab[next.x][next.y] == "." {
							graph.AddMappedArc(hashPoint3d(p), hashPoint3d(next), 1)
						}
					}
				}
			}
		}
	}
	for _, value := range portals {
		if len(value) == 2 {
			p1 := value[0]
			p2 := value[1]
			if isOutside(p1, tab) {
				for level := 1; level < maxLevel; level++ {
					graph.AddMappedArc(hashPoint3d(point3d{p1.x, p1.y, level}), hashPoint3d(point3d{p2.x, p2.y, level - 1}), 1)
					graph.AddMappedArc(hashPoint3d(point3d{p2.x, p2.y, level - 1}), hashPoint3d(point3d{p1.x, p1.y, level}), 1)
				}
			}
			if isOutside(p2, tab) {
				for level := 0; level < maxLevel-1; level++ {
					graph.AddMappedArc(hashPoint3d(point3d{p1.x, p1.y, level}), hashPoint3d(point3d{p2.x, p2.y, level + 1}), 1)
					graph.AddMappedArc(hashPoint3d(point3d{p2.x, p2.y, level + 1}), hashPoint3d(point3d{p1.x, p1.y, level}), 1)
				}
			}
		}
	}
	srcID, _ := graph.GetMapping(hashPoint3d(point3d{portals["AA"][0].x, portals["AA"][0].y, 0}))
	destID, _ := graph.GetMapping(hashPoint3d(point3d{portals["ZZ"][0].x, portals["ZZ"][0].y, 0}))
	best, _ := graph.Shortest(srcID, destID)
	return int(best.Distance)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day20/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day20/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day20/input.txt")

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
