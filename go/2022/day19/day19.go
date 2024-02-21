package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type state struct {
	aor, ac, aob, ag int
	ror, rc, rob, rg int
	time             int
}

func format(s string) []map[string]map[string]int {
	lines := strings.Split(s, "\n")
	res := make([]map[string]map[string]int, len(lines))
	for i, line := range lines {
		lineSplit := strings.Split(line, " ")
		res[i] = make(map[string]map[string]int)
		res[i]["ore"] = make(map[string]int)
		res[i]["clay"] = make(map[string]int)
		res[i]["obsidian"] = make(map[string]int)
		res[i]["geode"] = make(map[string]int)
		n, _ := strconv.Atoi(lineSplit[6])
		res[i]["ore"]["ore"] = n
		n, _ = strconv.Atoi(lineSplit[12])
		res[i]["clay"]["ore"] = n
		n, _ = strconv.Atoi(lineSplit[18])
		res[i]["obsidian"]["ore"] = n
		n, _ = strconv.Atoi(lineSplit[21])
		res[i]["obsidian"]["clay"] = n
		n, _ = strconv.Atoi(lineSplit[27])
		res[i]["geode"]["ore"] = n
		n, _ = strconv.Atoi(lineSplit[30])
		res[i]["geode"]["obsidian"] = n
	}
	return res
}

func maximums(blueprint map[string]map[string]int) map[string]int {
	res := make(map[string]int)
	for robot := range blueprint {
		for ressource := range blueprint[robot] {
			res[ressource] = max(res[ressource], blueprint[robot][ressource])
		}
	}
	return res
}

func solve(blueprint map[string]map[string]int, maxTime int) int {
	var res int
	var triangle []int
	for i := 0; i <= maxTime; i++ {
		triangle = append(triangle, ((i-1)*i)/2)
	}
	limit := maximums(blueprint)
	lor, lc, lob := limit["ore"], limit["clay"], limit["obsidian"]
	start := state{aor: 0, ac: 0, aob: 0, ag: 0, ror: 1, rc: 0, rob: 0, rg: 0, time: maxTime}
	oror, cor, obor, obc, gor, gob := blueprint["ore"]["ore"], blueprint["clay"]["ore"], blueprint["obsidian"]["ore"], blueprint["obsidian"]["clay"], blueprint["geode"]["ore"], blueprint["geode"]["obsidian"]
	var dfs func(state)
	dfs = func(curr state) {
		res = max(res, curr.ag)
		if curr.ag+curr.time*curr.rg+triangle[curr.time] <= res {
			return
		}
		if curr.time == 0 {
			return
		}
		for _, robotToBuy := range []string{"geode", "obsidian", "clay", "ore"} {
			newState := state{aor: curr.aor + curr.ror, ac: curr.ac + curr.rc, aob: curr.aob + curr.rob, ag: curr.ag + curr.rg, ror: curr.ror, rc: curr.rc, rob: curr.rob, rg: curr.rg, time: curr.time - 1}
			if robotToBuy == "ore" {
				if curr.aor >= oror && newState.ror < lor {
					newState.aor -= oror
					newState.ror++
					dfs(newState)
				}
			}
			if robotToBuy == "clay" {
				if curr.aor >= cor && newState.rc < lc {
					newState.aor -= cor
					newState.rc++
					dfs(newState)
				}
			}
			if robotToBuy == "obsidian" && curr.rc > 0 {
				if curr.aor >= obor && curr.ac >= obc && newState.rob < lob {
					newState.aor -= obor
					newState.ac -= obc
					newState.rob++
					dfs(newState)
				}
			}
			if robotToBuy == "geode" && curr.rob > 1 {
				if curr.aor >= gor && curr.aob >= gob {
					newState.aor -= gor
					newState.aob -= gob
					newState.rg++
					dfs(newState)
				}
			}
		}
		newState := state{aor: curr.aor + curr.ror, ac: curr.ac + curr.rc, aob: curr.aob + curr.rob, ag: curr.ag + curr.rg, ror: curr.ror, rc: curr.rc, rob: curr.rob, rg: curr.rg, time: curr.time - 1}
		dfs(newState)
	}
	dfs(start)
	return res
}

func part1(s string) int {
	c := 0
	blueprints := format(s)
	for i, b := range blueprints {
		c += (i + 1) * solve(b, 24)
	}
	return c
}

func part2(s string) int {
	c := 1
	blueprints := format(s)
	if len(blueprints) > 2 {
		blueprints = blueprints[:3]
	}
	for _, b := range blueprints {
		c = c * solve(b, 32)
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day19/test.data")

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

	content, err = os.ReadFile("../../../inputs/2022/day19/input.data")

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
