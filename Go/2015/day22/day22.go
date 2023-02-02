package main

import (
	"fmt"
	"math/rand"
)

func min(tab []int) int {
	m := tab[0]
	for _, x := range tab {
		if x < m {
			m = x
		}
	}
	return m
}

func part1() int {
	res := []int{}
	cost := make(map[int]int)
	cost[0] = 53
	cost[1] = 73
	cost[2] = 113
	cost[3] = 173
	cost[4] = 229
	i := 0
	for i < 5000000 {
		poison := 0
		shield := 0
		boss := 71
		player := 50
		mana := 500
		recharge := 0
		spent := 0
		for boss > 0 && player > 0 {
			if poison > 0 {
				poison--
				boss -= 3
			}
			if shield > 0 {
				shield--
			}
			if recharge > 0 {
				recharge--
				mana += 101
			}
			if boss <= 0 {
				res = append(res, spent)
				break
			}
			x := rand.Intn(5)
			if mana < 53 {
				break
			}
			for !((poison == 0 || (poison > 0 && x != 3)) && (shield == 0 || (shield > 0 && x != 2)) && (recharge == 0 || (recharge > 0 && x != 4))) {
				x = rand.Intn(5)
			}
			switch x {
			case 0:
				boss -= 4
				mana -= cost[0]
				spent += cost[0]
			case 1:
				mana -= cost[1]
				boss -= 2
				player += 2
				spent += cost[1]
			case 2:
				mana -= cost[2]
				shield = 6
				spent += cost[2]
			case 3:
				mana -= cost[3]
				poison = 6
				spent += cost[3]
			case 4:
				mana -= cost[4]
				recharge = 5
				spent += cost[4]
			}
			if poison > 0 {
				poison--
				boss -= 3
			}
			if shield > 0 {
				shield--
			}
			if recharge > 0 {
				recharge--
				mana += 101
			}
			if boss <= 0 {
				res = append(res, spent)
				break
			}
			if shield > 0 {
				player -= 3
			} else {
				player -= 10
			}
		}
		i++
	}
	return min(res)
}

func part2() int {
	res := []int{}
	cost := make(map[int]int)
	cost[0] = 53
	cost[1] = 73
	cost[2] = 113
	cost[3] = 173
	cost[4] = 229
	i := 0
	for i < 5000000 {
		poison := 0
		shield := 0
		boss := 71
		player := 50
		mana := 500
		recharge := 0
		spent := 0
		for boss > 0 && player > 0 {
			player--
			if player < 1 {
				break
			}
			if poison > 0 {
				poison--
				boss -= 3
			}
			if shield > 0 {
				shield--
			}
			if recharge > 0 {
				recharge--
				mana += 101
			}
			if boss <= 0 {
				res = append(res, spent)
				break
			}
			x := rand.Intn(5)
			if mana < 53 {
				break
			}
			for !((poison == 0 || (poison > 0 && x != 3)) && (shield == 0 || (shield > 0 && x != 2)) && (recharge == 0 || (recharge > 0 && x != 4))) {
				x = rand.Intn(5)
			}
			switch x {
			case 0:
				boss -= 4
				mana -= cost[0]
				spent += cost[0]
			case 1:
				mana -= cost[1]
				boss -= 2
				player += 2
				spent += cost[1]
			case 2:
				mana -= cost[2]
				shield = 6
				spent += cost[2]
			case 3:
				mana -= cost[3]
				poison = 6
				spent += cost[3]
			case 4:
				mana -= cost[4]
				recharge = 5
				spent += cost[4]
			}
			if poison > 0 {
				poison--
				boss -= 3
			}
			if shield > 0 {
				shield--
			}
			if recharge > 0 {
				recharge--
				mana += 101
			}
			if boss <= 0 {
				res = append(res, spent)
				break
			}
			if shield > 0 {
				player -= 3
			} else {
				player -= 10
			}
		}
		i++
	}
	return min(res)
}

func main() {
	fmt.Printf("Input :\n")
	fmt.Printf("Part 1 : %d\n", part1())
	fmt.Printf("Part 2 : %d\n", part2())
}
