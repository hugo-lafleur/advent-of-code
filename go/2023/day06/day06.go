package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func clean(l []string) []string {
	res := []string{}
	for _, x := range l {
		if !(x == "") {
			res = append(res, x)
		}
	}
	return res
}

func format(s string) []int {
	lines := strings.Split(s, "\n")
	time := clean(strings.Split(lines[0], " "))
	distance := clean(strings.Split(lines[1], " "))
	res := []int{}
	i := 1
	for i < len(time) && i < len(distance) {
		if !(time[i] == " ") {
			t, _ := strconv.Atoi(time[i])
			d, _ := strconv.Atoi(distance[i])
			res = append(res, t, d)
		}
		i++
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	i := 0
	res := 1
	for i < len(tab)-1 {
		ways := 0
		time := tab[i]
		distance := tab[i+1]
		k := 1
		for k < time {
			race := k * (time - k)
			if race > distance {
				ways++
			}
			k++
		}
		res *= ways
		i = i + 2
	}
	return res
}

func format2(s string) []int {
	tab := format(s)
	i := 0
	time := ""
	distance := ""
	for i < len(tab)-1 {
		time += strconv.Itoa(tab[i])
		distance += strconv.Itoa(tab[i+1])
		i = i + 2
	}
	res := []int{}
	a, _ := strconv.Atoi(time)
	b, _ := strconv.Atoi(distance)
	res = append(res, a, b)
	return res
}

func part2(s string) int {
	tab := format2(s)
	ways := 0
	time := tab[0]
	distance := tab[1]
	k := 1
	for k < time {
		race := k * (time - k)
		if race > distance {
			ways++
		}
		k++
	}
	return ways
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day06/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2023/day06/input.txt")

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
