package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"time"
)

type food struct {
	ingredients []string
	allergens   []string
}

func Split(r rune) bool {
	return r == ' ' || r == '(' || r == ',' || r == ')'
}

func format(s string) []food {
	res := []food{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		i := 0
		var ing, all []string
		for lineSplit[i] != "contains" {
			ing = append(ing, lineSplit[i])
			i++
		}
		i++
		for i < len(lineSplit) {
			all = append(all, lineSplit[i])
			i++
		}
		res = append(res, food{ingredients: ing, allergens: all})
	}
	return res
}

func count(x string, l []string) int {
	res := 0
	for _, y := range l {
		if x == y {
			res++
		}
	}
	return res
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

func solve(perm []int, ings []string, alls []string, foods []food) (bool, map[string]string) {
	corres := make(map[string]string)
	for i := range ings {
		corres[alls[i]] = ings[perm[i]]
	}
	for _, fd := range foods {
		for _, all := range fd.allergens {
			if !slices.Contains(fd.ingredients, corres[all]) {
				return false, corres
			}
		}
	}
	return true, corres
}

func part1(s string) int {
	c := 0
	foods := format(s)
	possibilities := make(map[string]map[string]bool)
	for _, food := range foods {
		for _, ing := range food.ingredients {
			_, done := possibilities[ing]
			if !done {
				possibilities[ing] = make(map[string]bool)
			}
			for _, all := range food.allergens {
				possibilities[ing][all] = true
			}
		}
	}
	noAllergens := make(map[string]bool)
	for ing, allergens := range possibilities {
		k := 0
	allgs:
		for all := range allergens {
			for _, food := range foods {
				if slices.Contains(food.allergens, all) && !slices.Contains(food.ingredients, ing) {
					k++
					continue allgs
				}
			}
		}
		if k == len(allergens) {
			noAllergens[ing] = true
		}
	}
	for noAll := range noAllergens {
		for _, food := range foods {
			c += count(noAll, food.ingredients)
		}
	}
	return c
}

func part2(s string) string {
	foods := format(s)
	possibilities := make(map[string]map[string]bool)
	for _, food := range foods {
		for _, ing := range food.ingredients {
			_, done := possibilities[ing]
			if !done {
				possibilities[ing] = make(map[string]bool)
			}
			for _, all := range food.allergens {
				possibilities[ing][all] = true
			}
		}
	}
	noAllergens := make(map[string]bool)
	for ing, allergens := range possibilities {
		k := 0
	allgs:
		for all := range allergens {
			for _, food := range foods {
				if slices.Contains(food.allergens, all) && !slices.Contains(food.ingredients, ing) {
					k++
					continue allgs
				}
			}
		}
		if k == len(allergens) {
			noAllergens[ing] = true
		}
	}
	newFoods := []food{}
	ings := []string{}
	alls := []string{}
	for _, nwfood := range foods {
		newIngs := []string{}
		for _, ing := range nwfood.ingredients {
			_, noAll := noAllergens[ing]
			if !noAll {
				newIngs = append(newIngs, ing)
				if !slices.Contains(ings, ing) {
					ings = append(ings, ing)
				}
			}
			for _, all := range nwfood.allergens {
				if !slices.Contains(alls, all) {
					alls = append(alls, all)
				}
			}
		}
		newFoods = append(newFoods, food{newIngs, nwfood.allergens})
	}
	perm := []int{}
	for i := 0; i < len(ings); i++ {
		perm = append(perm, i)
	}
	allPerms := permutations(perm)
	i := 0
	var ok bool
	var corres map[string]string
	for ; !ok; i++ {
		ok, corres = solve(allPerms[i], ings, alls, newFoods)
	}
	sort.Strings(alls)
	res := ""
	for _, all := range alls {
		res += corres[all] + ","
	}
	return res[:len(res)-1]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day21/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day21/input.data")

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
