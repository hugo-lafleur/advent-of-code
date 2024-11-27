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

func createTest(stackStr string) []Stack {
	lines := strings.Split(stackStr, "\n")
	input := [][]string{}
	for _, x := range lines {
		input = append(input, strings.Split(x, ""))
	}
	var stack1 Stack
	var stack2 Stack
	var stack3 Stack
	i := len(input) - 2
	for i > -1 {
		k := 0
		for k < len(input[0]) {
			if (k%4) == 1 && input[i][k] != " " {
				if k == 1 {
					stack1.Push(input[i][k])
				}
				if k == 5 {
					stack2.Push(input[i][k])
				}
				if k == 9 {
					stack3.Push(input[i][k])
				}
			}
			k++
		}
		i--
	}
	res := []Stack{}
	res = append(res, stack1)
	res = append(res, stack2)
	res = append(res, stack3)
	return res
}

func createInput(stackStr string) []Stack {
	lines := strings.Split(stackStr, "\n")
	input := [][]string{}
	for _, x := range lines {
		input = append(input, strings.Split(x, ""))
	}
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
	i := len(input) - 2
	for i > -1 {
		k := 0
		for k < len(input[0]) {
			if (k%4) == 1 && input[i][k] != " " {
				if k == 1 {
					stack1.Push(input[i][k])
				}
				if k == 5 {
					stack2.Push(input[i][k])
				}
				if k == 9 {
					stack3.Push(input[i][k])
				}
				if k == 13 {
					stack4.Push(input[i][k])
				}
				if k == 17 {
					stack5.Push(input[i][k])
				}
				if k == 21 {
					stack6.Push(input[i][k])
				}
				if k == 25 {
					stack7.Push(input[i][k])
				}
				if k == 29 {
					stack8.Push(input[i][k])
				}
				if k == 33 {
					stack9.Push(input[i][k])
				}
			}
			k++
		}
		i--
	}
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

func part1(s string, stackStr string) []string {
	tab := format(s)
	stacks := []Stack{}
	if len(tab) == 4 {
		stacks = createTest(stackStr)
	}
	if len(tab) > 4 {
		stacks = createInput(stackStr)
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

func part2(s string, stackStr string) []string {
	tab := format(s)
	stacks := []Stack{}
	if len(tab) == 4 {
		stacks = createTest(stackStr)
	}
	if len(tab) > 4 {
		stacks = createInput(stackStr)
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
	content, err := os.ReadFile("../../../inputs/2022/day05/test.txt")

	if err != nil {
		log.Fatal(err)
	}

	contentStack, err := os.ReadFile("../../../inputs/2022/day05/testStack.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content), string(contentStack)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content), string(contentStack)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day05/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	contentStack, err = os.ReadFile("../../../inputs/2022/day05/inputStack.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content), string(contentStack)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content), string(contentStack)))
	fmt.Println(time.Since(start))
}
