package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type group struct {
	units      int
	hitPoints  int
	immune     []string
	weak       []string
	damage     int
	damageType string
	initiative int
}

type groupList []group

func (g group) AttackPower() int {
	return g.damage * g.units
}

func (g groupList) Len() int {
	return len(g)
}

func (g groupList) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]

}

func (g groupList) Less(i, j int) bool {
	if g[i].AttackPower() == g[j].AttackPower() {
		return g[i].initiative > g[j].initiative
	}
	return g[i].AttackPower() > g[j].AttackPower()
}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == ';' || r == '(' || r == ')'
}

func format(s string) []groupList {
	dStr := strings.Split(s, "\n\n")
	immuneStr := strings.Split(dStr[0], "\n")
	infectionStr := strings.Split(dStr[1], "\n")
	var immuneList groupList
	for _, line := range immuneStr[1:] {
		strs := strings.FieldsFunc(line, Split)
		var units, hitPoints, damage, initiative int
		var weak, immune []string
		var damageType string
		for i, word := range strs {
			if word == "units" {
				n, _ := strconv.Atoi(strs[i-1])
				units = n
			}
			if word == "hit" {
				n, _ := strconv.Atoi(strs[i-1])
				hitPoints = n
			}
			if word == "damage" {
				n, _ := strconv.Atoi(strs[i-2])
				damage = n
				damageType = strs[i-1]
			}
			if word == "immune" {
				i++
				for strs[i] != "with" && strs[i] != "weak" {
					if strs[i] != "to" {
						immune = append(immune, strs[i])
					}
					i++
				}
				i--
			}
			if word == "weak" {
				i++
				for strs[i] != "with" && strs[i] != "immune" {
					if strs[i] != "to" {
						weak = append(weak, strs[i])
					}
					i++
				}
				i--
			}
			if word == "initiative" {
				n, _ := strconv.Atoi(strs[i+1])
				initiative = n
				break
			}
		}
		immuneList = append(immuneList, group{units, hitPoints, immune, weak, damage, damageType, initiative})
	}
	var infectionList groupList
	for _, line := range infectionStr[1:] {
		strs := strings.FieldsFunc(line, Split)
		var units, hitPoints, damage, initiative int
		var weak, immune []string
		var damageType string
		for i, word := range strs {
			if word == "units" {
				n, _ := strconv.Atoi(strs[i-1])
				units = n
			}
			if word == "hit" {
				n, _ := strconv.Atoi(strs[i-1])
				hitPoints = n
			}
			if word == "damage" {
				n, _ := strconv.Atoi(strs[i-2])
				damage = n
				damageType = strs[i-1]
			}
			if word == "immune" {
				i++
				for strs[i] != "with" && strs[i] != "weak" {
					if strs[i] != "to" {
						immune = append(immune, strs[i])
					}
					i++
				}
				i--
			}
			if word == "weak" {
				i++
				for strs[i] != "with" && strs[i] != "immune" {
					if strs[i] != "to" {
						weak = append(weak, strs[i])
					}
					i++
				}
				i--
			}
			if word == "initiative" {
				n, _ := strconv.Atoi(strs[i+1])
				initiative = n
				break
			}
		}
		infectionList = append(infectionList, group{units, hitPoints, immune, weak, damage, damageType, initiative})
	}
	return []groupList{immuneList, infectionList}
}

func isIn(l []string, s string) bool {
	for _, x := range l {
		if s == x {
			return true
		}
	}
	return false
}

func maxKeyMap(m map[int]int) []int {
	var max int
	for _, value := range m {
		max = value
	}
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	res := []int{}
	for key, value := range m {
		if value == max {
			res = append(res, key)
		}
	}
	return res
}

func allZero(m map[int]int) bool {
	for _, value := range m {
		if value != 0 {
			return false
		}
	}
	return true
}

func wouldDamage(g group, target group) int {
	if isIn(target.weak, g.damageType) {
		return 2 * g.AttackPower()
	}
	if isIn(target.immune, g.damageType) {
		return 0
	}
	return g.AttackPower()
}

func chooseTarget(g group, l groupList, targets map[int]int) int {
	available := []int{}
	var possibleTargets []int
outerLoop:
	for i := range l {
		for _, value := range targets {
			if i == value {
				continue outerLoop
			}
		}
		available = append(available, i)
	}
	possibleDamage := make(map[int]int)
	for _, index := range available {
		possibleDamage[index] = wouldDamage(g, l[index])
	}
	if allZero(possibleDamage) {
		return -1
	}
	possibleTargets = maxKeyMap(possibleDamage)
	if len(possibleTargets) == 1 {
		return possibleTargets[0]
	}
	effectivePower := make(map[int]int)
	for _, index := range possibleTargets {
		effectivePower[index] = l[index].AttackPower()
	}

	possibleTargets = maxKeyMap(effectivePower)
	if len(possibleTargets) == 1 {
		return possibleTargets[0]
	}
	mostInitiative := make(map[int]int)
	for _, index := range possibleTargets {
		mostInitiative[index] = l[index].initiative
	}
	possibleTargets = maxKeyMap(mostInitiative)
	return possibleTargets[0]
}

func totalUnits(l groupList) int {
	c := 0
	for _, g := range l {
		c += g.units
	}
	return c
}

func solve(s string, boost int) (bool, int) {
	g := format(s)
	immuneSystem, infection := g[0], g[1]
	for i := range immuneSystem {
		immuneSystem[i].damage += boost
	}
	fights := 0
	for immuneSystem.Len() != 0 && infection.Len() != 0 {
		sort.Sort(immuneSystem)
		sort.Sort(infection)

		// Target Selection for Immune System

		targetForImmune := make(map[int]int)
		for i, g := range immuneSystem {
			j := chooseTarget(g, infection, targetForImmune)
			if j != -1 {
				targetForImmune[i] = j
			}
		}

		// Target Selection for Infection

		targetForInfection := make(map[int]int)
		for i, g := range infection {
			j := chooseTarget(g, immuneSystem, targetForInfection)
			if j != -1 {
				targetForInfection[i] = j
			}
		}

		// Attacking order creattion

		initiativeMap := make(map[int]int)
		for i, g := range immuneSystem {
			initiativeMap[i+1] = g.initiative
		}
		for i, g := range infection {
			initiativeMap[-i-1] = g.initiative
		}

		// Attacking Phase

		isThereAction := false

		for len(initiativeMap) != 0 {
			src := maxKeyMap(initiativeMap)[0]
			var ok bool
			var dest int
			if src > 0 {
				delete(initiativeMap, src)
				src--
				if immuneSystem[src].initiative == 0 {
					continue
				}
				dest, ok = targetForImmune[src]
				if ok { // check if src has a target
					damage := (wouldDamage(immuneSystem[src], infection[dest]) / infection[dest].hitPoints)
					infection[dest].units -= damage
					if damage > 0 {
						isThereAction = true
					}
					if infection[dest].units <= 0 {
						infection[dest].initiative = 0
					}

				}
			}
			if src < 0 {
				delete(initiativeMap, src)
				src++
				src = -src
				if infection[src].initiative == 0 {
					continue
				}
				dest, ok = targetForInfection[src]
				if ok {
					damage := (wouldDamage(infection[src], immuneSystem[dest]) / immuneSystem[dest].hitPoints)
					immuneSystem[dest].units -= damage
					if damage > 0 {
						isThereAction = true
					}
					if immuneSystem[dest].units <= 0 {
						immuneSystem[dest].initiative = 0
					}
				}
			}
		}
		var newImmune groupList
		var newInfection groupList
		for _, g := range immuneSystem {
			if g.initiative != 0 {
				newImmune = append(newImmune, g)
			}
		}
		for _, g := range infection {
			if g.initiative != 0 {
				newInfection = append(newInfection, g)
			}
		}
		if !isThereAction {
			return false, 0
		}
		immuneSystem = newImmune
		infection = newInfection
		fights++
	}
	if infection.Len() == 0 {
		return true, totalUnits(immuneSystem)
	} else {
		return false, totalUnits(infection)
	}
}

func part1(s string) int {
	_, res := solve(s, 0)
	return res
}

func part2(s string) int {
	for boost := 0; ; boost++ {
		ok, res := solve(s, boost)
		if ok {
			return res
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day24/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day24/input.data")

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
