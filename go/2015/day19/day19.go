package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func format(s string) ([][]string, []string) {
	lines := strings.Split(s, "\n")
	l := len(lines)
	last := strings.Split(lines[l-1], "")
	mol := []string{}
	i := 0
	tmp := ""
	for i < len(last) {
		if int(lines[l-1][i]) < 91 {
			if tmp != "" {
				mol = append(mol, tmp)
			}
			tmp = last[i]
		} else {
			tmp = tmp + last[i]
			mol = append(mol, tmp)
			tmp = ""
		}
		i++
	}
	if tmp != "" {
		mol = append(mol, tmp)
	}

	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.Split(x, " "))
	}
	rules := [][]string{}
	for _, x := range tab {
		if len(x) > 1 && x[1] == "=>" {
			new := []string{}
			new = append(new, x[0])
			new = append(new, "=>")
			str := strings.Split(x[2], "")
			n := len(str)
			i := 0
			tmp := ""
			for i < n {
				if int(x[2][i]) < 91 {
					if tmp != "" {
						new = append(new, tmp)
					}
					tmp = str[i]
				} else {
					tmp = tmp + str[i]
					new = append(new, tmp)
					tmp = ""
				}
				i++
			}
			if tmp != "" {
				new = append(new, tmp)
			}
			rules = append(rules, new)
		}
	}
	return rules, mol
}

func add(tab []string, j int, add []string) []string {
	new := []string{}
	for i, x := range tab {
		if i == j {
			new = append(new, add...)
		} else {
			new = append(new, x)
		}
	}
	return new
}

func equal(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func is_in(tab [][]string, s []string) bool {
	for _, x := range tab {
		if equal(x, s) {
			return true
		}
	}
	return false
}

func remove(mol []string, j int, rule []string) []string {
	new := []string{}
	i := 0
	for i < len(mol) {
		x := mol[i]
		if i == j {
			new = append(new, rule[0])
			i = i + len(rule) - 2
		} else {
			new = append(new, x)
			i++
		}
	}
	return new
}

func in_rule(mol []string, j int, rule []string) bool {
	i := 2
	for i < len(rule) {
		if j >= len(mol) {
			return false
		}
		if mol[j] != rule[i] {
			return false
		}
		i++
		j++
	}
	return true
}

func minimze(tab [][]string) [][]string {
	res := tab[0]
	for _, x := range tab {
		if len(x) < len(res) {
			res = x
		}
	}
	return [][]string{res}
}

func part1(s string) int {
	rules, mol := format(s)
	list := [][]string{}
	for _, x := range rules {
		start := x[0]
		ends := []string{}
		i := 2
		for i < len(x) {
			ends = append(ends, x[i])
			i++
		}
		for i, y := range mol {
			if y == start {
				new := add(mol, i, ends)
				if !(is_in(list, new)) {
					list = append(list, new)
				}
			}
			i++
		}
	}
	return len(list)
}

func part2(s string) int {
	rules, molas := format(s)
	start := [][]string{molas}
	cpt := 0
	for {
		new := [][]string{}
		for _, mol := range start {
			for i := range mol {
				for _, x := range rules {
					if in_rule(mol, i, x) {
						add := remove(mol, i, x)
						if !is_in(new, add) && (len(add) < len(mol) || equal(add, []string{"e"})) {
							new = append(new, add)
						}
					}
				}
			}
		}
		cpt++
		for _, x := range new {
			if equal(x, []string{"e"}) {
				return cpt
			}
		}
		start = minimze(new)
	}
}

func main() {
	/*content, err := os.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))*/

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
