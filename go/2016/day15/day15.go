package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type disc struct {
	index     int
	positions int
	start     int
}

func Split(r rune) bool {
	return r == ' ' || r == '.'
}

func format(s string) []disc {
	lines := strings.Split(s, "\n")
	res := []disc{}
	for i, line := range lines {
		infos := strings.FieldsFunc(line, Split)
		d := disc{}
		d.index = i + 1
		a, _ := strconv.Atoi(infos[3])
		d.positions = a
		b, _ := strconv.Atoi(infos[11])
		d.start = b
		res = append(res, d)
	}
	return res
}

func verifySlot(slot map[int]bool) bool {
	for _, value := range slot {
		if !value {
			return false
		}
	}
	return true
}

func part1(s string) int {
	time := 0
	discs := format(s)
	for {
		slot := make(map[int]bool)
		for _, disc := range discs {
			if (time+disc.index+disc.start)%disc.positions == 0 {
				slot[disc.index] = true
			} else {
				slot[disc.index] = false
			}
		}
		if verifySlot(slot) {
			return time
		}
		time++
	}
}

func part2(s string) int {
	time := 0
	discs := format(s)
	discs = append(discs, disc{7, 11, 0})
	for {
		slot := make(map[int]bool)
		for _, disc := range discs {
			if (time+disc.index+disc.start)%disc.positions == 0 {
				slot[disc.index] = true
			} else {
				slot[disc.index] = false
			}
		}
		if verifySlot(slot) {
			return time
		}
		time++
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day15/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day15/input.txt")

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
