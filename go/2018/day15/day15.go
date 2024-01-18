package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type point struct {
	x, y int
}

type unit struct {
	race   string
	health int
	attack int
	pos    point
}

type unitList []unit

type pointList []point

func (u unitList) Len() int {
	return len(u)
}

func (u unitList) Less(i, j int) bool {
	if u[i].pos.x == u[j].pos.x {
		return u[i].pos.y < u[j].pos.y
	}
	return u[i].pos.x < u[j].pos.x
}

func (u unitList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u pointList) Len() int {
	return len(u)
}

func (u pointList) Less(i, j int) bool {
	if u[i].x == u[j].x {
		return u[i].y < u[j].y
	}
	return u[i].x < u[j].x
}
func (u pointList) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func format(s string) (int, map[point]string) {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ""))
	}
	res := make(map[point]string)
	for i, line := range tab {
		for j, s := range line {
			res[point{i, j}] = s
		}
	}
	return len(lines), res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nextTo(p1, p2 point) bool {
	return (p1.x == p2.x && abs(p1.y-p2.y) == 1) || (p1.y == p2.y && abs(p1.x-p2.x) == 1)
}

func removeUnit(l unitList, index int) unitList {
	var res unitList
	for i, u := range l {
		if i != index {
			res = append(res, u)
		}
	}
	return res
}

func sumHealth(units unitList) int {
	res := 0
	for _, u := range units {
		res += u.health
	}
	return res
}

func IsIn(l pointList, p point) bool {
	for _, x := range l {
		if x == p {
			return true
		}
	}
	return false
}

func sumPoint(p1, p2 point) point {
	return point{p1.x + p2.x, p1.y + p2.y}
}

func solve(s string, power int, breakIfElfDie bool) (bool, int) {
	rounds := 0
	_, field := format(s)
	var units unitList
	for key, value := range field {
		if value == "E" {
			units = append(units, unit{value, 200, power, key})
		}
		if value == "G" {
			units = append(units, unit{value, 200, 3, key})
		}
	}
	for {
		sort.Sort(units)

		for k := 0; k < len(units); k++ {
			unit := units[k]
			var targetList []int           // list of targets
			var inRangeList pointList      // in range points as described
			var queue deque.Deque[point]   // deque for BFS
			var seen pointList             // remember points seen
			var distToUnit map[point]int   // map to find distance to points
			var firstPoint map[point]point // first point in path to point
			var pointsToGo pointList       // list of points that are in range and reached by BFS
			var maxDistance int            // once an in range point is reached, do not go any further
			var pointToGo point            // chosen point
			var nextPoint point            //next point according to dijkstra
			for j, u := range units {
				if unit.race != u.race {
					targetList = append(targetList, j)
				}
			}
			if len(targetList) == 0 {
				return false, rounds * sumHealth(units)
			}
			for _, j := range targetList {
				if nextTo(unit.pos, units[j].pos) {
					goto attack
				}
			}
			inRangeList = []point{}
			for _, index := range targetList {
				for p, s := range field {
					if nextTo(units[index].pos, p) && s == "." {
						inRangeList = append(inRangeList, p)
					}
				}
			}
			queue.PushBack(unit.pos)
			seen = append(seen, unit.pos)
			distToUnit = make(map[point]int)
			firstPoint = make(map[point]point)
			distToUnit[unit.pos] = 0
			maxDistance = 1000
			for queue.Len() != 0 {
				curr := queue.PopFront()
				if distToUnit[curr] > maxDistance {
					break
				}
				if IsIn(inRangeList, curr) {
					pointsToGo = append(pointsToGo, curr)
					maxDistance = distToUnit[curr]
				}
				for _, offset := range []point{{-1, 0}, {0, -1}, {0, +1}, {+1, 0}} {
					p := sumPoint(curr, offset)
					if field[p] == "." && !IsIn(seen, p) {
						queue.PushBack(p)
						seen = append(seen, p)
						distToUnit[p] = distToUnit[curr] + 1
						if distToUnit[p] == 1 {
							firstPoint[p] = p
						} else {
							firstPoint[p] = firstPoint[curr]
						}
					}
				}
			}
			if len(pointsToGo) == 0 {
				continue
			}
			sort.Sort(pointsToGo)
			pointToGo = pointsToGo[0]
			nextPoint = firstPoint[pointToGo]
			field[unit.pos] = "."
			field[nextPoint] = unit.race
			units[k].pos = nextPoint
		attack:
			unit = units[k]
			var minHealth int
			var targetInRangeList []int
			var minHealthTargetSet map[point]int
			var minHealthTargetPositionList pointList
			var indexTarget int
			for _, index := range targetList {
				if nextTo(units[index].pos, unit.pos) {
					targetInRangeList = append(targetInRangeList, index)
				}
			}
			for _, index := range targetInRangeList {
				minHealth = units[index].health
			}
			for _, index := range targetInRangeList {
				if units[index].health < minHealth {
					minHealth = units[index].health
				}
			}
			minHealthTargetSet = make(map[point]int)
			for _, index := range targetInRangeList {
				if minHealth == units[index].health {
					minHealthTargetSet[units[index].pos] = index
					minHealthTargetPositionList = append(minHealthTargetPositionList, units[index].pos)
				}
			}
			if len(minHealthTargetPositionList) == 0 {
				continue
			}
			sort.Sort(minHealthTargetPositionList)
			indexTarget = minHealthTargetSet[minHealthTargetPositionList[0]]
			units[indexTarget].health -= units[k].attack
			if units[indexTarget].health <= 0 {
				if units[indexTarget].race == "E" && breakIfElfDie {
					return true, 0
				}
				field[units[indexTarget].pos] = "."
				units = removeUnit(units, indexTarget)
				if indexTarget < k {
					k--
				}
			}
		}
		rounds++
	}
}

func part1(s string) int {
	_, res := solve(s, 3, false)
	return res
}

func part2(s string) int {
	for power := 4; ; power++ {
		elfKilled, res := solve(s, power, true)
		if !elfKilled {
			return res
		}
	}

}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day15/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day15/input.data")

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
