package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aclements/go-z3/z3"
	"github.com/gammazero/deque"
)

type machine struct {
	diagram []bool
	buttons [][]int
	joltage []int
}

func format(s string) []machine {
	lines := strings.Split(s, "\n")
	result := []machine{}
	for _, line := range lines {
		diagram := []bool{}
		buttons := [][]int{}
		joltage := []int{}
		parts := strings.Split(line, " ")
		for i := range parts[0] {
			if parts[0][i] != '[' && parts[0][i] != ']' {
				if parts[0][i] == '#' {
					diagram = append(diagram, true)
				} else {
					diagram = append(diagram, false)
				}
			}
		}
		for i := 1; i < len(parts)-1; i++ {
			button := []int{}
			for _, str := range strings.FieldsFunc(parts[i], func(r rune) bool {
				return r == '(' || r == ')' || r == ','
			}) {
				n, _ := strconv.Atoi(str)
				button = append(button, n)
			}
			buttons = append(buttons, button)
		}
		for _, str := range strings.FieldsFunc(parts[len(parts)-1], func(r rune) bool {
			return r == ',' || r == '{' || r == '}'
		}) {
			n, _ := strconv.Atoi(str)
			joltage = append(joltage, n)
		}
		result = append(result, machine{diagram, buttons, joltage})
	}
	return result
}

type state struct {
	lights  int
	presses int
}

func fewestPressesLight(m machine) int {
	n := len(m.diagram)
	target := 0
	for i := range m.diagram {
		target <<= 1
		if m.diagram[i] {
			target++
		}
	}
	var dq deque.Deque[state]
	dq.PushBack(state{0, 0})
	visited := map[int]bool{}
	visited[0] = true
	for dq.Len() != 0 {
		curr := dq.PopFront()
		if curr.lights == target {
			return curr.presses
		}
		for _, b := range m.buttons {
			nextLights := curr.lights
			for _, i := range b {
				nextLights ^= (1 << (n - 1 - i))
			}
			if !visited[nextLights] {
				visited[nextLights] = true
				dq.PushBack(state{nextLights, curr.presses + 1})
			}
		}
	}
	return -1
}

func part1(s string) int {
	machines := format(s)
	result := 0
	for _, m := range machines {
		result += fewestPressesLight(m)
	}
	return result
}

func fewestPressesJoltage(m machine, ctx *z3.Context) int {
	n := len(m.joltage)
	s := z3.NewSolver(ctx)
	vars := []z3.Int{}
	total := ctx.IntConst("total")
	for i := range m.buttons {
		vars = append(vars, ctx.IntConst(fmt.Sprintf("%d", i)))
		s.Assert(vars[i].GE(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
	}

	s.Assert(total.GE(ctx.FromInt(0, ctx.IntSort()).(z3.Int)))
	s.Assert(total.Eq(vars[0].Add(vars[1:]...)))
	joltageVars := make([][]z3.Int, n)
	for i := range n {
		joltageVars[i] = []z3.Int{}
	}
	for i := range m.buttons {
		for _, j := range m.buttons[i] {
			joltageVars[j] = append(joltageVars[j], vars[i])
		}
	}
	for i := range m.joltage {
		sum := joltageVars[i][0].Add(joltageVars[i][1:]...)
		j := ctx.FromInt(int64(m.joltage[i]), ctx.IntSort()).(z3.Int)
		s.Assert(sum.Eq(j))
	}
	result := -1
	for {
		sat, err := s.Check()
		if !sat || err != nil {
			break
		}
		model := s.Model()
		val, _, _ := model.Eval(ctx.IntConst("total"), true).(z3.Int).AsInt64()
		result = int(val)
		s.Assert(total.LT(ctx.FromInt(val, ctx.IntSort()).(z3.Int)))
	}
	return result
}

func part2(s string) int {
	machines := format(s)
	result := 0
	ctx := z3.NewContext(nil)
	for i, m := range machines {
		fmt.Println(i+1, "/", len(machines))
		result += fewestPressesJoltage(m, ctx)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day10/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2025/day10/input.txt")

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
