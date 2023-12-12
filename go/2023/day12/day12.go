package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type combination struct {
	arr, goal string
	inc       int
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

var cache = make(map[combination]int)

func mainFunc(arr []string, goal []string, inc int) int {
	arrStr := strings.Join(arr, "")
	goalStr := strings.Join(goal, ",")
	comb := combination{arrStr, goalStr, inc}
	if len(arr) == 0 {
		if (len(goal) == 1 && strconv.Itoa(inc) == goal[0]) || (inc == 0 && len(goal) == 0) {
			return 1
		} else {
			return 0
		}
	}
	char := arr[0]
	v, ok := cache[comb]
	if ok {
		return v
	}
	if len(goal) == 0 {
		if char == "." {
			res := mainFunc(arr[1:], goal, inc)
			cache[comb] = res
			return res
		}
	}
	if char == "." {
		if inc == 0 {
			res := mainFunc(arr[1:], goal, 0)
			cache[comb] = res
			return res
		} else {
			n, _ := strconv.Atoi(goal[0])
			if n == inc {
				res := mainFunc(arr[1:], goal[1:], 0)
				cache[comb] = res
				return res
			} else {
				return 0
			}
		}
	}
	if char == "#" {
		return mainFunc(arr[1:], goal, inc+1)
	} else {
		arr1 := make([]string, len(arr))
		arr2 := make([]string, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		arr1[0] = "#"
		arr2[0] = "."
		res := mainFunc(arr1, goal, inc) + mainFunc(arr2, goal, inc)
		cache[comb] = res
		return res
	}

}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, line := range list {
		arr := strings.Split(line[0], "")
		goal := strings.Split(line[1], ",")
		c += mainFunc(arr, goal, 0)
	}
	return c
}

func part2(s string) int {
	c := 0
	list := format(s)
	for _, line := range list {
		baseArr := strings.Clone(line[0])
		baseGoal := strings.Clone(line[1])
		for i := 0; i < 4; i++ {
			line[0] += "?"
			line[0] += baseArr
			line[1] += ","
			line[1] += baseGoal
		}
		arr := strings.Split(line[0], "")
		goal := strings.Split(line[1], ",")
		//fmt.Println(i)
		c += mainFunc(arr, goal, 0)
		//fmt.Println(c)
	}
	return c
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
