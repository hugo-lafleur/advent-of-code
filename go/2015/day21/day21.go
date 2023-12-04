package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var hitPoints int
var baseDamage int
var baseArmor int

func win(damage int, armor int) bool {
	boss := hitPoints
	player := 100
	for {
		boss = boss - damage + baseArmor
		//fmt.Println(boss)
		if boss < 1 {
			return true
		}
		player = player - baseDamage + armor
		if player < 1 {
			return false
		}
	}
}

func min(tab []int) int {
	m := tab[0]
	for _, x := range tab {
		if x < m {
			m = x
		}
	}
	return m
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

func part1(s string) int {
	stats := strings.Split(s, "\n")
	for i, line := range stats {
		tab := strings.Split(line, " ")
		l := len(tab)
		switch i {
		case 0:
			n, _ := strconv.Atoi(tab[l-1])
			hitPoints = n
		case 1:
			n, _ := strconv.Atoi(tab[l-1])
			baseDamage = n
		case 2:
			n, _ := strconv.Atoi(tab[l-1])
			baseArmor = n
		}

	}
	var damage_weapon = make(map[int]int)
	damage_weapon[0] = 4
	damage_weapon[1] = 5
	damage_weapon[2] = 6
	damage_weapon[3] = 7
	damage_weapon[4] = 8
	var cost_weapon = make(map[int]int)
	cost_weapon[0] = 8
	cost_weapon[1] = 10
	cost_weapon[2] = 25
	cost_weapon[3] = 40
	cost_weapon[4] = 74
	var armor_armor = make(map[int]int)
	armor_armor[0] = 1
	armor_armor[1] = 2
	armor_armor[2] = 3
	armor_armor[3] = 4
	armor_armor[4] = 5
	armor_armor[5] = 0
	var cost_armor = make(map[int]int)
	cost_armor[0] = 13
	cost_armor[1] = 31
	cost_armor[2] = 53
	cost_armor[3] = 75
	cost_armor[4] = 102
	cost_armor[5] = 0
	var damage_rings = make(map[int]int)
	damage_rings[0] = 1
	damage_rings[1] = 2
	damage_rings[2] = 3
	damage_rings[3] = 0
	damage_rings[4] = 0
	damage_rings[5] = 0
	damage_rings[6] = 0
	damage_rings[7] = 0
	var armor_rings = make(map[int]int)
	armor_rings[0] = 0
	armor_rings[1] = 0
	armor_rings[2] = 0
	armor_rings[3] = 1
	armor_rings[4] = 2
	armor_rings[5] = 3
	armor_rings[6] = 0
	armor_rings[7] = 0
	var cost_rings = make(map[int]int)
	cost_rings[0] = 25
	cost_rings[1] = 50
	cost_rings[2] = 100
	cost_rings[3] = 20
	cost_rings[4] = 40
	cost_rings[5] = 80
	cost_rings[6] = 0
	cost_rings[7] = 0
	res := []int{}
	armor := 0
	for armor < 6 {
		weapon := 0
		for weapon < 5 {
			i := 0
			for i < 8 {
				j := 0
				for j < 8 {
					if i != j {
						if win(damage_weapon[weapon]+damage_rings[i]+damage_rings[j], armor_armor[armor]+armor_rings[i]+armor_rings[j]) {
							res = append(res, cost_armor[armor]+cost_weapon[weapon]+cost_rings[i]+cost_rings[j])
						}
					}
					j++
				}
				i++
			}
			weapon++
		}
		armor++
	}
	return min(res)
}

func part2(s string) int {
	stats := strings.Split(s, "\n")
	for i, line := range stats {
		tab := strings.Split(line, " ")
		l := len(tab)
		switch i {
		case 0:
			n, _ := strconv.Atoi(tab[l-1])
			hitPoints = n
		case 1:
			n, _ := strconv.Atoi(tab[l-1])
			baseDamage = n
		case 2:
			n, _ := strconv.Atoi(tab[l-1])
			baseArmor = n
		}

	}
	var damage_weapon = make(map[int]int)
	damage_weapon[0] = 4
	damage_weapon[1] = 5
	damage_weapon[2] = 6
	damage_weapon[3] = 7
	damage_weapon[4] = 8
	var cost_weapon = make(map[int]int)
	cost_weapon[0] = 8
	cost_weapon[1] = 10
	cost_weapon[2] = 25
	cost_weapon[3] = 40
	cost_weapon[4] = 74
	var armor_armor = make(map[int]int)
	armor_armor[0] = 1
	armor_armor[1] = 2
	armor_armor[2] = 3
	armor_armor[3] = 4
	armor_armor[4] = 5
	armor_armor[5] = 0
	var cost_armor = make(map[int]int)
	cost_armor[0] = 13
	cost_armor[1] = 31
	cost_armor[2] = 53
	cost_armor[3] = 75
	cost_armor[4] = 102
	cost_armor[5] = 0
	var damage_rings = make(map[int]int)
	damage_rings[0] = 1
	damage_rings[1] = 2
	damage_rings[2] = 3
	damage_rings[3] = 0
	damage_rings[4] = 0
	damage_rings[5] = 0
	damage_rings[6] = 0
	damage_rings[7] = 0
	var armor_rings = make(map[int]int)
	armor_rings[0] = 0
	armor_rings[1] = 0
	armor_rings[2] = 0
	armor_rings[3] = 1
	armor_rings[4] = 2
	armor_rings[5] = 3
	armor_rings[6] = 0
	armor_rings[7] = 0
	var cost_rings = make(map[int]int)
	cost_rings[0] = 25
	cost_rings[1] = 50
	cost_rings[2] = 100
	cost_rings[3] = 20
	cost_rings[4] = 40
	cost_rings[5] = 80
	cost_rings[6] = 0
	cost_rings[7] = 0
	res := []int{}
	armor := 0
	for armor < 6 {
		weapon := 0
		for weapon < 5 {
			i := 0
			for i < 8 {
				j := 0
				for j < 8 {
					if i != j {
						if !win(damage_weapon[weapon]+damage_rings[i]+damage_rings[j], armor_armor[armor]+armor_rings[i]+armor_rings[j]) {
							res = append(res, cost_armor[armor]+cost_weapon[weapon]+cost_rings[i]+cost_rings[j])
						}
					}
					j++
				}
				i++
			}
			weapon++
		}
		armor++
	}
	return max(res)
}

func main() {
	content, err := os.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
