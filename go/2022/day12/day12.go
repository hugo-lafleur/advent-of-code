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
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ""))
	}
	return tab
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func ascii(s string) int {
	return int([]rune(s)[0])

}

func part1(s string) int {
	graph := dijkstra.NewGraph()
	tab := format(s)
	n := 0
	l := len(tab[0])
	dict := make(map[int]string)
	start := 0
	end := 0
	for i := range tab {
		for j := range tab[i] {
			graph.AddVertex(i*l + j)
			dict[i*l+j] = tab[i][j]
		}
	}
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			if str == "S" {
				tab[i][j] = "a"
				start = i*l + j
			}
			if str == "E" {
				tab[i][j] = "z"
				end = i*l + j
			}
		}
	}
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			n = ascii(str)
			if i > 0 {
				if abs(n-ascii(tab[i-1][j])) < 2 || ascii(tab[i-1][j]) < n {
					graph.AddArc(i*l+j, (i-1)*l+j, 1)
				}
			}
			if i < len(tab)-1 {
				if abs(n-ascii(tab[i+1][j])) < 2 || ascii(tab[i+1][j]) < n {
					graph.AddArc(i*l+j, (i+1)*l+j, 1)
				}
			}
			if j > 0 {
				if abs(n-ascii(tab[i][j-1])) < 2 || ascii(tab[i][j-1]) < n {
					graph.AddArc(i*l+j, i*l+j-1, 1)
				}
			}
			if j < len(tab[i])-1 {
				if abs(n-ascii(tab[i][j+1])) < 2 || ascii(tab[i][j+1]) < n {
					graph.AddArc(i*l+j, i*l+j+1, 1)
				}
			}
		}
	}
	best, err := graph.Shortest(start, end)
	if err != nil {
		log.Fatal(err)
	}
	return int(best.Distance)
}

func part2(s string) int {
	graph := dijkstra.NewGraph()
	tab := format(s)
	n := 0
	l := len(tab[0])
	dict := make(map[int]string)
	end := 0
	for i := range tab {
		for j := range tab[i] {
			graph.AddVertex(i*l + j)
			dict[i*l+j] = tab[i][j]
		}
	}
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			if str == "E" {
				tab[i][j] = "z"
				end = i*l + j
			}
		}
	}
	for i := range tab {
		for j := range tab[i] {
			str := tab[i][j]
			n = ascii(str)
			if i > 0 {
				if abs(n-ascii(tab[i-1][j])) < 2 || ascii(tab[i-1][j]) < n {
					graph.AddArc(i*l+j, (i-1)*l+j, 1)
				}
			}
			if i < len(tab)-1 {
				if abs(n-ascii(tab[i+1][j])) < 2 || ascii(tab[i+1][j]) < n {
					graph.AddArc(i*l+j, (i+1)*l+j, 1)
				}
			}
			if j > 0 {
				if abs(n-ascii(tab[i][j-1])) < 2 || ascii(tab[i][j-1]) < n {
					graph.AddArc(i*l+j, i*l+j-1, 1)
				}
			}
			if j < len(tab[i])-1 {
				if abs(n-ascii(tab[i][j+1])) < 2 || ascii(tab[i][j+1]) < n {
					graph.AddArc(i*l+j, i*l+j+1, 1)
				}
			}
		}
	}
	res := 5000
	for i := range tab {
		for j := range tab[i] {
			if dict[i*l+j] == "a" {
				best, _ := graph.Shortest(i*l+j, end)
				if int(best.Distance) < res && int(best.Distance) != 0 {
					res = int(best.Distance)
				}
			}
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day12/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day12/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
