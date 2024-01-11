package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == ';' || r == '='
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.FieldsFunc(x, Split))
	}
	return tab
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

/*func is_in(tab []int, x int) bool {
	for _, y := range tab {
		if x == y {
			return true
		}
	}
	return false
}*/

func sum(tab []int) int {
	s := 0
	for _, x := range tab {
		s += x
	}
	return s
}

func max(tab []int) int {
	m := tab[0]
	for _, x := range tab {
		if x > m {
			m = x
		}
	}
	return m
}

/*func max_map(dict map[int]int) int {
	m := 0
	res := 0
	for x := range dict {
		if dict[x] > m {
			m = dict[x]
			res = x
		}
	}
	return res
}*/

func is_in(tab []int, x int) bool {
	for _, y := range tab {
		if x == y {
			return true
		}
	}
	return false
}

func part1(s string) int {
	tab := format(s)
	dict := make(map[string]int)
	valves := []int{}
	res := []int{}
	for i := range tab {
		dict[tab[i][1]] = i
	}
	graph := dijkstra.NewGraph()
	for i := range tab {
		graph.AddVertex(i)
	}
	for _, line := range tab {
		src := dict[line[1]]
		j := 10
		for j < len(line) {
			graph.AddArc(src, dict[line[j]], 1)
			j++
		}
	}
	for i, line := range tab {
		n, _ := strconv.Atoi(line[5])
		if n != 0 {
			valves = append(valves, i)
		}
	}
	allPath := [][]int{}
	i := 0
	flow := make(map[int]int)
	for _, x := range valves {
		n, _ := strconv.Atoi(tab[x][5])
		flow[x] = n
	}
	for i < 100000 {
		done := []int{0}
		path := []int{0}
		for len(done) <= len(valves) {
			next := 0
			for is_in(done, next) {
				next = valves[rand.Intn(len(valves))]
			}
			path = append(path, next)
			done = append(done, next)
		}
		allPath = append(allPath, path)
		i++
	}
	//fmt.Println(allPath)
	for _, path := range allPath {
		tmp := []int{0}
		i := 0
		for i < len(path)-1 {
			best, _ := graph.Shortest(path[i], path[i+1])
			tmp = append(tmp, best.Path[1:]...)
			i++
		}
		fullpath := tmp
		open := path
		i = 0
		time := 0
		pressure := 0
		add := []int{}
		for i < len(fullpath) && time < 30 {
			pressure += sum(add)
			time++
			if fullpath[i] == open[0] {
				n, _ := strconv.Atoi(tab[fullpath[i]][5])
				add = append(add, n)
				pressure += sum(add)
				open = open[1:]
				time++
			}
			i++
		}
		//fmt.Println(time)
		if time < 30 {
			pressure += sum(add) * (30 - time + 1)
			res = append(res, pressure)
		}
	}
	sort.Ints(res)
	//fmt.Println(res)
	return max(res)
}

func part2(s string) int {
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day16/test.data")

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

	content, err = os.ReadFile("../../../inputs/2022/day16/input.data")

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
