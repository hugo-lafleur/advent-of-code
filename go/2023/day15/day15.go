package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type lens map[string]int

type box []lens

type boxes []box

func format(s string) []string {
	return strings.Split(s, ",")
}

func HASH(s string) int {
	res := 0
	for _, r := range s {
		res += int(r)
		res *= 17
		res = res % 256
	}
	return res
}

func lensInBox(label string, boxesList boxes, i int) bool {
	box := boxesList[i]
	for _, lens := range box {
		_, ok := lens[label]
		if ok {
			return ok
		}
	}
	return false
}

func (boxesList boxes) removeLens(label string, i int) {
	j := 0
	for j < len(boxesList[i]) {
		_, ok := boxesList[i][j][label]
		if ok {
			boxesList[i] = append(boxesList[i][:j], boxesList[i][j+1:]...)
			break
		}
		j++
	}
}

func (boxesList boxes) addLens(label string, focalLength int, i int) {
	j := 0
	done := false
	for j < len(boxesList[i]) {
		_, ok := boxesList[i][j][label]
		if ok {
			boxesList[i][j][label] = focalLength
			done = true
		}
		j++
	}
	if !done {
		boxesList[i] = append(boxesList[i], lens(map[string]int{label: focalLength}))
	}
}

func part1(s string) int {
	c := 0
	list := format(s)
	for _, str := range list {
		c += HASH(str)
	}
	return c
}

func part2(s string) int {
	c := 0
	stepList := format(s)
	boxesList := make(boxes, 256)
	for _, step := range stepList {
		l := len(step)
		if step[l-1] == '-' {
			label := step[:l-1]
			b := HASH(label)
			if lensInBox(label, boxesList, b) {
				boxesList.removeLens(label, b)
			}
		}
		if step[l-2] == '=' {
			instr := strings.Split(step, "=")
			n, _ := strconv.Atoi(instr[1])
			b := HASH(instr[0])
			boxesList.addLens(instr[0], n, b)

		}
	}
	for i, box := range boxesList {
		for j, lens := range box {
			for _, value := range lens {
				c += (i + 1) * (j + 1) * value
			}
		}
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
