package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Stack []string

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

func createTest() []Stack {
	res := []Stack{}
	var stack1 Stack
	var stack2 Stack
	var stack3 Stack
	stack1.Push("Z")
	stack1.Push("N")
	stack2.Push("M")
	stack2.Push("C")
	stack2.Push("D")
	stack3.Push("P")
	res = append(res, stack1)
	res = append(res, stack2)
	res = append(res, stack3)
	return res
}

func createInput() []Stack {
	res := []Stack{}
	var stack1 Stack
	var stack2 Stack
	var stack3 Stack
	var stack4 Stack
	var stack5 Stack
	var stack6 Stack
	var stack7 Stack
	var stack8 Stack
	var stack9 Stack
	stack1.Push("S")
	stack1.Push("L")
	stack1.Push("W")
	stack2.Push("J")
	stack2.Push("T")
	stack2.Push("N")
	stack2.Push("Q")
	stack3.Push("S")
	stack3.Push("C")
	stack3.Push("H")
	stack3.Push("F")
	stack3.Push("J")
	stack4.Push("T")
	stack4.Push("R")
	stack4.Push("M")
	stack4.Push("W")
	stack4.Push("N")
	stack4.Push("G")
	stack4.Push("B")
	stack5.Push("T")
	stack5.Push("R")
	stack5.Push("L")
	stack5.Push("S")
	stack5.Push("D")
	stack5.Push("H")
	stack5.Push("Q")
	stack5.Push("B")
	stack6.Push("M")
	stack6.Push("J")
	stack6.Push("B")
	stack6.Push("V")
	stack6.Push("F")
	stack6.Push("H")
	stack6.Push("R")
	stack6.Push("L")
	stack7.Push("D")
	stack7.Push("W")
	stack7.Push("R")
	stack7.Push("N")
	stack7.Push("J")
	stack7.Push("M")
	stack8.Push("B")
	stack8.Push("Z")
	stack8.Push("T")
	stack8.Push("F")
	stack8.Push("H")
	stack8.Push("N")
	stack8.Push("D")
	stack8.Push("J")
	stack9.Push("H")
	stack9.Push("L")
	stack9.Push("Q")
	stack9.Push("N")
	stack9.Push("B")
	stack9.Push("F")
	stack9.Push("T")
	res = append(res, stack1)
	res = append(res, stack2)
	res = append(res, stack3)
	res = append(res, stack4)
	res = append(res, stack5)
	res = append(res, stack6)
	res = append(res, stack7)
	res = append(res, stack8)
	res = append(res, stack9)
	return res
}

func part1(s string) []string {
	tab := format(s)
	stacks := []Stack{}
	if len(tab) == 4 {
		stacks = createTest()
	}
	if len(tab) > 4 {
		stacks = createInput()
	}
	for _, line := range tab {
		n, _ := strconv.Atoi(line[1])
		i, _ := strconv.Atoi(line[3])
		j, _ := strconv.Atoi(line[5])
		k := 0
		for k < n {
			s, _ := stacks[i-1].Pop()
			stacks[j-1].Push(s)
			k++
		}
	}
	res := []string{}
	i := 0
	for i < len(stacks) {
		s, _ = stacks[i].Pop()
		res = append(res, s)
		i++
	}
	return res
}

func part2(s string) []string {
	tab := format(s)
	stacks := []Stack{}
	if len(tab) == 4 {
		stacks = createTest()
	}
	if len(tab) > 4 {
		stacks = createInput()
	}
	for _, line := range tab {
		n, _ := strconv.Atoi(line[1])
		i, _ := strconv.Atoi(line[3])
		j, _ := strconv.Atoi(line[5])
		k := 0
		var tmp Stack
		for k < n {
			s, _ := stacks[i-1].Pop()
			tmp.Push(s)
			k++
		}
		for !(tmp.IsEmpty()) {
			s, _ := tmp.Pop()
			stacks[j-1].Push(s)
		}
	}
	res := []string{}
	i := 0
	for i < len(stacks) {
		s, _ = stacks[i].Pop()
		res = append(res, s)
		i++
	}
	return res
}

func main() {
	content, err := ioutil.ReadFile("test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	fmt.Printf("Part 1 : %s\n", part1(string(content)))
	fmt.Printf("Part 2 : %s\n", part2(string(content)))

	content, err = ioutil.ReadFile("input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	fmt.Printf("Part 1 : %s\n", part1(string(content)))
	fmt.Printf("Part 2 : %s\n", part2(string(content)))
}
