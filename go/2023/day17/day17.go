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

type state struct {
	i, j   int
	d      string
	blocks int
}

var tab [][]string

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func uniqueCode(s state) string {
	res := ""
	res += strconv.Itoa(s.i)
	res += ";"
	res += strconv.Itoa(s.j)
	res += ";"
	res += s.d
	res += ";"
	res += strconv.Itoa(s.blocks)
	return res
}

func part1(s string) int {
	tab = format(s)
	n := len(tab)
	m := len(tab[0])
	graph := dijkstra.NewGraph()
	dict := make(map[int]state)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range []string{"U", "R", "D", "L"} {
				for b := 0; b < 4; b++ {
					s := state{i, j, d, b}
					dict[graph.AddMappedVertex(uniqueCode(s))] = s
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range []string{"U", "R", "D", "L"} {
				for b := 0; b < 4; b++ {
					s := state{i, j, d, b}
					if j < m-1 && d != "L" {
						n, _ := strconv.Atoi(tab[i][j+1])
						if !(d == "R") {
							nextS := state{i, j + 1, "R", 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
						if d == "R" && b < 3 {
							nextS := state{i, j + 1, "R", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
					if j > 0 && d != "R" {
						n, _ := strconv.Atoi(tab[i][j-1])
						if !(d == "L") {
							nextS := state{i, j - 1, "L", 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
						if d == "L" && b < 3 {
							nextS := state{i, j - 1, "L", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
					if i < n-1 && d != "U" {
						n, _ := strconv.Atoi(tab[i+1][j])
						if !(d == "D") {
							nextS := state{i + 1, j, "D", 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
						if d == "D" && b < 3 {
							nextS := state{i + 1, j, "D", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
					if i > 0 && d != "D" {
						n, _ := strconv.Atoi(tab[i-1][j])
						if !(d == "U") {
							nextS := state{i - 1, j, "U", 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
						if d == "U" && b < 3 {
							nextS := state{i - 1, j, "U", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
				}
			}
		}
	}
	for i := 0; i < 4; i++ {
		for _, d := range []string{"R", "D", "L", "U"} {
			s := state{n - 1, m - 1, d, i}
			end := state{-1, -1, "E", 0}
			graph.AddMappedArc(uniqueCode(s), uniqueCode(end), 0)
		}
	}
	idSource := graph.AddMappedVertex(uniqueCode(state{0, 0, "R", 0}))
	idDest := graph.AddMappedVertex(uniqueCode(state{-1, -1, "E", 0}))
	best, _ := graph.Shortest(idSource, idDest)
	return int(best.Distance)
}

func part2(s string) int {
	tab = format(s)
	n := len(tab)
	m := len(tab[0])
	graph := dijkstra.NewGraph()
	dict := make(map[int]state)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range []string{"U", "R", "D", "L"} {
				for b := 0; b < 10; b++ {
					s := state{i, j, d, b}
					dict[graph.AddMappedVertex(uniqueCode(s))] = s
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, d := range []string{"U", "R", "D", "L"} {
				for b := 0; b < 11; b++ {
					s := state{i, j, d, b}
					if j < m-1 && d != "L" {
						if (!(d == "R") && j < m-4) || (i == 0 && j == 0) {
							n1, _ := strconv.Atoi(tab[i][j+1])
							n2, _ := strconv.Atoi(tab[i][j+2])
							n3, _ := strconv.Atoi(tab[i][j+3])
							n4, _ := strconv.Atoi(tab[i][j+4])
							nextS := state{i, j + 4, "R", 4}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n1+n2+n3+n4))
						}
						if d == "R" && b < 10 && b > 3 {
							n, _ := strconv.Atoi(tab[i][j+1])
							nextS := state{i, j + 1, "R", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}

					}
					if j > 0 && d != "R" {
						if !(d == "L") && j > 3 {
							n1, _ := strconv.Atoi(tab[i][j-1])
							n2, _ := strconv.Atoi(tab[i][j-2])
							n3, _ := strconv.Atoi(tab[i][j-3])
							n4, _ := strconv.Atoi(tab[i][j-4])
							nextS := state{i, j - 4, "L", 4}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n1+n2+n3+n4))
						}
						if d == "L" && b < 10 && b > 3 {
							n, _ := strconv.Atoi(tab[i][j-1])
							nextS := state{i, j - 1, "L", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
					if i < n-1 && d != "U" {
						if !(d == "D") && i < n-4 {
							n1, _ := strconv.Atoi(tab[i+1][j])
							n2, _ := strconv.Atoi(tab[i+2][j])
							n3, _ := strconv.Atoi(tab[i+3][j])
							n4, _ := strconv.Atoi(tab[i+4][j])
							nextS := state{i + 4, j, "D", 4}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n1+n2+n3+n4))
						}
						if d == "D" && b < 10 && b > 3 {
							n, _ := strconv.Atoi(tab[i+1][j])
							nextS := state{i + 1, j, "D", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
					if i > 0 && d != "D" {
						if !(d == "U") && i > 3 {
							n1, _ := strconv.Atoi(tab[i-1][j])
							n2, _ := strconv.Atoi(tab[i-2][j])
							n3, _ := strconv.Atoi(tab[i-3][j])
							n4, _ := strconv.Atoi(tab[i-4][j])
							nextS := state{i - 4, j, "U", 4}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n1+n2+n3+n4))
						}
						if d == "U" && b < 10 && b > 3 {
							n, _ := strconv.Atoi(tab[i-1][j])
							nextS := state{i - 1, j, "U", b + 1}
							graph.AddMappedArc(uniqueCode(s), uniqueCode(nextS), int64(n))
						}
					}
				}
			}
		}
	}
	for i := 0; i < 11; i++ {
		for _, d := range []string{"R", "D", "L", "U"} {
			s := state{n - 1, m - 1, d, i}
			end := state{-1, -1, "E", 0}
			graph.AddMappedArc(uniqueCode(s), uniqueCode(end), 0)
		}
	}
	idSource := graph.AddMappedVertex(uniqueCode(state{0, 0, "R", 0}))
	idDest := graph.AddMappedVertex(uniqueCode(state{-1, -1, "E", 0}))
	best, _ := graph.Shortest(idSource, idDest)
	return int(best.Distance)
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
