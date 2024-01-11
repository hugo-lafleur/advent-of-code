package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type rule struct {
	part       string
	comparator string
	number     int
	next       string
}

type defaultRule string

type workflow struct {
	name  string
	rules []rule
	dR    defaultRule
}

type tree struct {
	tag     map[string]int
	infTree *tree
	supTree *tree
}

type interval struct {
	min, max int
}

var listRanges [][]map[string]int

func split(r rune) bool {
	return r == '{' || r == '}' || r == ':' || r == ',' || r == '='
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, split))
	}
	return res
}

func createWorkflows(tab [][]string) map[string]workflow {
	workflows := make(map[string]workflow)
	for _, line := range tab {
		if len(line) == 0 {
			break
		}
		rules := []rule{}
		for i := 1; i < len(line)-1; i = i + 2 {
			part := string(line[i][0])
			comparator := string(line[i][1])
			number := 0
			if comparator == ">" {
				number, _ = strconv.Atoi(strings.Split(line[i], ">")[1])
			} else {
				number, _ = strconv.Atoi(strings.Split(line[i], "<")[1])
			}
			rules = append(rules, rule{part, comparator, number, line[i+1]})
		}
		workflow := workflow{line[0], rules, defaultRule(line[len(line)-1])}
		workflows[line[0]] = workflow
	}
	return workflows
}

func length(i interval) int {
	return i.max - i.min + 1
}

func createTree(start string, workflows map[string]workflow, i int) tree {
	if start == "A" || start == "R" {
		tag := make(map[string]int)
		tag[start] = -1
		return tree{tag, nil, nil}
	}
	workflow := workflows[start]
	rules := workflow.rules
	if i == len(rules) {
		return createTree(string(workflow.dR), workflows, 0)
	}
	rule := rules[i]
	tag := make(map[string]int)
	tag[rule.part] = rule.number
	if rule.comparator == "<" {
		infTree := createTree(rule.next, workflows, 0)
		supTree := createTree(start, workflows, i+1)
		return tree{tag, &infTree, &supTree}
	} else {
		tag[rule.part]++
		supTree := createTree(rule.next, workflows, 0)
		infTree := createTree(start, workflows, i+1)
		return tree{tag, &infTree, &supTree}
	}
}

func searchTree(t tree, visited []map[string]int) {
	visited = append(visited, t.tag)
	if t.infTree != nil && t.supTree != nil {
		searchTree(*t.infTree, append(visited, map[string]int{"<": 0}))
		searchTree(*t.supTree, append(visited, map[string]int{">": 0}))
	} else {
		copy := []map[string]int{}
		copy = append(copy, visited...)
		listRanges = append(listRanges, copy)
	}
}

func part1(s string) int {
	c := 0
	tab := format(s)
	workflows := make(map[string]workflow)
	k := 0
	for _, line := range tab {
		if len(line) == 0 {
			break
		}
		rules := []rule{}
		for i := 1; i < len(line)-1; i = i + 2 {
			part := string(line[i][0])
			comparator := string(line[i][1])
			number := 0
			if comparator == ">" {
				number, _ = strconv.Atoi(strings.Split(line[i], ">")[1])
			} else {
				number, _ = strconv.Atoi(strings.Split(line[i], "<")[1])
			}
			rules = append(rules, rule{part, comparator, number, line[i+1]})
		}
		workflow := workflow{line[0], rules, defaultRule(line[len(line)-1])}
		workflows[line[0]] = workflow
		k++
	}
	k++
	for k < len(tab) {
		line := tab[k]
		dict := make(map[string]int)
		for i := 0; i < 7; i++ {
			n, _ := strconv.Atoi(line[i+1])
			dict[line[i]] = n
		}
		now := "in"
		for now != "A" && now != "R" {
			workflow := workflows[now]
			rules := workflow.rules
			for i := 0; i < len(rules)+1; i++ {
				if i == len(rules) {
					now = string(workflow.dR)
					break
				}
				rule := rules[i]
				if rule.comparator == ">" {
					if dict[rule.part] > rule.number {
						now = rule.next
						break
					}
				} else {
					if dict[rule.part] < rule.number {
						now = rule.next
						break
					}
				}
			}
		}
		if now == "A" {
			c += dict["x"] + dict["m"] + dict["a"] + dict["s"]
		}
		k++
	}
	return c
}

func part2(s string) int {
	c := 0
	tab := format(s)
	workflows := createWorkflows(tab)
	tree := createTree("in", workflows, 0)
	visited := []map[string]int{}
	listRanges = [][]map[string]int{}
	searchTree(tree, visited)
	fmt.Println("")
	for _, path := range listRanges {
		intervals := make(map[string]interval)
		intervals["x"] = interval{1, 4000}
		intervals["m"] = interval{1, 4000}
		intervals["a"] = interval{1, 4000}
		intervals["s"] = interval{1, 4000}
		_, ok := path[len(path)-1]["A"]
		if ok {
			for i := 0; i < len(path)-1; i++ {
				_, sign := path[i+1]["<"]
				if sign {
					for key, value := range path[i] {
						max := intervals[key].max
						if value < max {
							temp := intervals[key]
							temp.max = value - 1
							intervals[key] = temp
						}
					}
				} else {
					for key, value := range path[i] {
						min := intervals[key].min
						if value > min {
							temp := intervals[key]
							temp.min = value
							intervals[key] = temp
						}
					}
				}
			}
			c += length(intervals["x"]) * length(intervals["m"]) * length(intervals["a"]) * length(intervals["s"])
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day19/test.data")

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

	content, err = os.ReadFile("../../../inputs/2023/day19/input.data")

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
