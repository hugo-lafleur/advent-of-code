package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/RyanCarrier/dijkstra"
	"github.com/gammazero/deque"
)

type valve_description struct {
	flow int
	near []string
}

type state struct {
	pos      string
	opened   []string
	pressure int
	time     int
}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == ';' || r == '='
}

func Split2(r rune) bool {
	return r == '[' || r == ']'
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.FieldsFunc(x, Split))
	}
	return tab
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func hash(s state) string {
	return fmt.Sprint(s.pos) + fmt.Sprint(s.opened) + fmt.Sprint(s.time)
}

func hashToOpened(s string) []string {
	parts := strings.FieldsFunc(s, Split2)
	if len(parts) == 1 {
		return []string{}
	}
	return strings.Split(parts[1], " ")
}

func solve(s string, max_time int) (int, map[string]int) {
	tab := format(s)
	var res int
	valves := make(map[string]valve_description)
	var q deque.Deque[state]
	var usefulValves []string
	paths := make(map[string]map[string]int)
	graph := dijkstra.NewGraph()
	for _, line := range tab {
		graph.AddMappedVertex(line[1])
	}
	for _, line := range tab {
		flow, _ := strconv.Atoi(line[5])
		near := []string{}
		near = append(near, line[1])
		for i := 10; i < len(line); i++ {
			graph.AddMappedArc(line[1], line[i], 1)
			near = append(near, line[i])
		}
		desc := valve_description{flow: flow, near: near}
		if flow > 0 {
			usefulValves = append(usefulValves, line[1])
		}
		valves[line[1]] = desc
	}
	for _, line := range tab {
		valve := line[1]
		paths[valve] = make(map[string]int)
		for _, usefulValve := range usefulValves {
			if valve != usefulValve {
				srcID, _ := graph.GetMapping(valve)
				destID, _ := graph.GetMapping(usefulValve)
				best, _ := graph.Shortest(srcID, destID)
				paths[valve][usefulValve] = int(best.Distance)
			}
		}
	}
	visited := make(map[string]int)
	start := state{pos: "AA", opened: []string{}, pressure: 0, time: 0}
	visited[hash(start)] = 0
	q.PushBack(start)
	for q.Len() != 0 {
		curr := q.PopFront()
		pressure := visited[hash(curr)]
		res = max(res, pressure)
		if curr.time == max_time {
			continue
		}
		new_pressure := pressure
		for _, valve := range curr.opened {
			new_pressure += valves[valve].flow
		}
		states_to_add := []state{}
		if slices.Contains(curr.opened, curr.pos) {
			new_opened := []string{}
			new_opened = append(new_opened, curr.opened...)
			slices.Sort(new_opened)
			new_state := state{pos: curr.pos, opened: new_opened, pressure: pressure + (new_pressure-pressure)*(max_time-curr.time), time: max_time}
			states_to_add = append(states_to_add, new_state)
		}
		if slices.Contains(usefulValves, curr.pos) && !slices.Contains(curr.opened, curr.pos) {
			new_opened := []string{}
			new_opened = append(new_opened, curr.opened...)
			new_opened = append(new_opened, curr.pos)
			slices.Sort(new_opened)
			new_state := state{pos: curr.pos, opened: new_opened, pressure: new_pressure, time: curr.time + 1}
			states_to_add = append(states_to_add, new_state)
		} else {
			if len(curr.opened) != len(usefulValves) {
				for _, usefulValve := range usefulValves {
					if !slices.Contains(curr.opened, usefulValve) {
						new_opened := []string{}
						new_opened = append(new_opened, curr.opened...)
						slices.Sort(new_opened)
						new_state := state{pos: usefulValve, opened: new_opened, pressure: pressure + (new_pressure-pressure)*(paths[curr.pos][usefulValve]), time: curr.time + paths[curr.pos][usefulValve]}
						if new_state.time > max_time {
							new_state.pressure = pressure + (new_pressure-pressure)*(max_time-curr.time)
							new_state.time = max_time
						}
						states_to_add = append(states_to_add, new_state)
					}
				}
			}
		}
		for _, new_state := range states_to_add {
			h := hash(new_state)
			old, seen := visited[h]
			if (!seen || old < new_state.pressure) && new_state.time <= max_time {
				q.PushBack(new_state)
				visited[h] = new_state.pressure
			}
		}
	}
	return res, visited
}

func part1(s string) int {
	res, _ := solve(s, 30)
	return res
}

func part2(s string) int {
	res := 0
	_, mapping := solve(s, 26)
	pressure := make(map[string]int)
	paths := []string{}
	for key := range mapping {
		array := hashToOpened(key)
		str := strings.Join(array, " ")
		_, added := pressure[str]
		if !added {
			paths = append(paths, str)
		}
		pressure[str] = max(pressure[str], mapping[key])
	}
	for i := 0; i < len(paths); i++ {
		fmt.Println(i)
	next:
		for j := i + 1; j < len(paths); j++ {
			path1 := paths[i]
			path2 := paths[j]
			opened1 := strings.Split(path1, " ")
			opened2 := strings.Split(path2, " ")
			for _, e := range opened1 {
				if slices.Contains(opened2, e) {
					continue next
				}
			}
			for _, e := range opened2 {
				if slices.Contains(opened1, e) {
					continue next
				}
			}
			res = max(res, pressure[path1]+pressure[path2])
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day16/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day16/input.txt")

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
