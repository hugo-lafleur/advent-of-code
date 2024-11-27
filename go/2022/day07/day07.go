package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Stack []string

func stack_to_string(s Stack) string {
	res := ""
	for !s.IsEmpty() {
		str, _ := s.Pop()
		res = str + "/" + res
	}
	if res == "//" {
		return "/"
	}
	return res[1 : len(res)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(i string) {
	*s = append(*s, i)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		i := len(*s) - 1
		x := (*s)[i]
		*s = (*s)[:i]
		return x, true
	}
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	return tab
}

func is_in(s string, tab []string) bool {
	for _, x := range tab {
		if s == x {
			return true
		}
	}
	return false
}

func total_size(s string, parent map[string]string, size map[string]int, dir []string, total map[string]int) int {
	res, b := total[s]
	sum := 0
	if b {
		return res
	} else {
		for x := range parent {
			if parent[x] == s {
				if is_in(x, dir) {
					sum += total_size(x, parent, size, dir, total)
				} else {
					sum += size[x]
				}
			}
		}
	}
	return sum
}

func part1(s string) int {
	tab := format(s)
	var total = make(map[string]int)
	var parent = make(map[string]string)
	var size = make(map[string]int)
	var dir = []string{}
	var current string
	var stack Stack
	i := 0
	for i < len(tab) {
		line := tab[i]
		if line[0] == "$" && line[1] == "ls" {
			i++
			line = tab[i]
			for line[0] != "$" {
				if line[0] != "dir" {
					stack.Push(line[1])
					str := stack_to_string(stack)
					stack.Pop()
					parent[str] = stack_to_string(stack)
					n, _ := strconv.Atoi(line[0])
					size[str] = n
				} else {
					stack.Push(line[1])
					str := stack_to_string(stack)
					stack.Pop()
					parent[str] = stack_to_string(stack)
				}
				i++
				if i == len(tab) {
					for _, x := range dir {
						total[x] = total_size(x, parent, size, dir, total)
					}
					res := 0
					for x := range total {
						if total[x] <= 100000 {
							res += total[x]
						}
					}
					return res
				}
				line = tab[i]
			}
		}
		if line[0] == "$" && line[1] == "cd" && line[2] != ".." {
			current = line[2]
			stack.Push(current)
			dir = append(dir, stack_to_string(stack))
		}
		if line[0] == "$" && line[1] == "cd" && line[2] == ".." {
			current, _ = stack.Pop()
		}
		i++
	}
	return 0
}

func part2(s string) int {
	tab := format(s)
	var total = make(map[string]int)
	var parent = make(map[string]string)
	var size = make(map[string]int)
	var dir = []string{}
	var current string
	var stack Stack
	i := 0
	for i < len(tab) {
		line := tab[i]
		if line[0] == "$" && line[1] == "ls" {
			i++
			line = tab[i]
			for line[0] != "$" {
				if line[0] != "dir" {
					stack.Push(line[1])
					str := stack_to_string(stack)
					stack.Pop()
					parent[str] = stack_to_string(stack)
					n, _ := strconv.Atoi(line[0])
					size[str] = n
				} else {
					stack.Push(line[1])
					str := stack_to_string(stack)
					stack.Pop()
					parent[str] = stack_to_string(stack)
				}
				i++
				if i == len(tab) {
					for _, x := range dir {
						total[x] = total_size(x, parent, size, dir, total)
					}
					free := 70000000 - total["/"]
					need := 30000000 - free
					min := 70000000
					for _, x := range dir {
						n := total[x]
						if n-30000000 < min-30000000 && n > need {
							min = n
						}
					}
					return min
				}
				line = tab[i]
			}
		}
		if line[0] == "$" && line[1] == "cd" && line[2] != ".." {
			current = line[2]
			stack.Push(current)
			dir = append(dir, stack_to_string(stack))
		}
		if line[0] == "$" && line[1] == "cd" && line[2] == ".." {
			current, _ = stack.Pop()
		}
		i++
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day07/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2022/day07/input.txt")

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
