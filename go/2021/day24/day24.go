package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type couple struct {
	a, b int
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func part1(s string) int {
	instr := format(s)
	array := make([]string, 14)
	var q deque.Deque[couple]
	for i := 0; i < 14; i++ {
		a, _ := strconv.Atoi(instr[18*i+4][2])
		b, _ := strconv.Atoi(instr[18*i+5][2])
		c, _ := strconv.Atoi(instr[18*i+15][2])
		if a == 1 {
			q.PushBack(couple{i, c})
		} else {
			c := q.PopBack()
			r := c.b + b
			t := 9
			for t+r > 9 {
				t--
			}
			for t+r < 1 {
				t++
			}
			array[c.a] = strconv.Itoa(t)
			array[i] = strconv.Itoa(t + r)
		}
	}
	n, _ := strconv.Atoi(strings.Join(array, ""))
	return n
}

func part2(s string) int {
	instr := format(s)
	array := make([]string, 14)
	var q deque.Deque[couple]
	for i := 0; i < 14; i++ {
		a, _ := strconv.Atoi(instr[18*i+4][2])
		b, _ := strconv.Atoi(instr[18*i+5][2])
		c, _ := strconv.Atoi(instr[18*i+15][2])
		if a == 1 {
			q.PushBack(couple{i, c})
		} else {
			c := q.PopBack()
			r := c.b + b
			t := 1
			for t+r > 9 {
				t--
			}
			for t+r < 1 {
				t++
			}
			array[c.a] = strconv.Itoa(t)
			array[i] = strconv.Itoa(t + r)
		}
	}
	n, _ := strconv.Atoi(strings.Join(array, ""))
	return n
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day24/input.txt")

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
