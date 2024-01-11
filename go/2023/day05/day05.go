package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "\n\n")
}

func min(tab []int) int {
	min := tab[0]
	for _, n := range tab {
		if n < min {
			min = n
		}
	}
	return min
}

func mapped(n int, tab []string) int {
	j := 1
	for j < 8 {
		maps := strings.Split(tab[j], "\n")
		k := 1
		for k < len(maps) {
			line := strings.Split(maps[k], " ")
			a, _ := strconv.Atoi(line[0])
			b, _ := strconv.Atoi(line[1])
			c, _ := strconv.Atoi(line[2])
			if n > b-1 && n < b+c {
				n = a + (n - b)
				k = len(maps)
			}
			k++
		}
		j++
	}
	return n

}

func part1(s string) int {
	locations := []int{}
	tab := format(s)
	seeds := strings.Split(tab[0], " ")
	i := 1
	for i < len(seeds) {
		seed, _ := strconv.Atoi(seeds[i])
		locations = append(locations, mapped(seed, tab))
		i++
	}
	return min(locations)
}

func unmapped(loc int, tab []string) int {
	i := 7
	n := loc
	for i > 0 {
		maps := strings.Split(tab[i], "\n")
		k := 1
		for k < len(maps) {
			line := strings.Split(maps[k], " ")
			a, _ := strconv.Atoi(line[0])
			b, _ := strconv.Atoi(line[1])
			c, _ := strconv.Atoi(line[2])
			if n > a-1 && n < a+c {
				n = b + (n - a)
				k = len(maps)
			}
			k++
		}
		i--
	}
	return n
}

func isSeed(loc int, tab []string) bool {
	seed := unmapped(loc, tab)
	seeds := strings.Split(tab[0], " ")
	k := 1
	for k < len(seeds)-1 {
		a, _ := strconv.Atoi(seeds[k])
		b, _ := strconv.Atoi(seeds[k+1])
		if seed > a-1 && seed < a+b {
			return true
		}
		k = k + 2
	}
	return false
}

func minSeeds(tab []string) int {
	seeds := strings.Split(tab[0], " ")
	min := 10000000000
	for _, x := range seeds {
		n, err := strconv.Atoi(x)
		if err == nil && n < min {
			min = n
		}
	}
	return min
}

func part2(s string) int {
	tab := format(s)
	i := 0
	step := minSeeds(tab)
	grow := true
	for step > 1 {
		if grow {
			for !isSeed(i, tab) {
				i = i + step
			}
			grow = false
			step = step / 10
		} else {
			for isSeed(i, tab) {
				i = i - step
			}
			grow = true
			step = step / 10
		}
	}
	if isSeed(i, tab) {
		for isSeed(i, tab) {
			i--
		}
		return i + 1
	}
	if !isSeed(i, tab) {
		for !isSeed(i, tab) {
			i++
		}
		return i
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day05/test.data")

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

	content, err = os.ReadFile("../../../inputs/2023/day05/input.data")

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
