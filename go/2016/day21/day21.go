package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func reverse(n int) int {
	switch n {
	case 0:
		return 7
	case 1:
		return 7
	case 2:
		return 2
	case 3:
		return 6
	case 4:
		return 1
	case 5:
		return 5
	case 6:
		return 0
	case 7:
		return 4
	}
	return 0
}

func newString(str string, instr []string) string {
	newStr := ""
	switch instr[0] {
	case "swap":
		switch instr[1] {
		case "position":
			indexA, _ := strconv.Atoi(instr[2])
			indexB, _ := strconv.Atoi(instr[5])
			for i := range str {
				if i == indexA {
					newStr += string(str[indexB])
				}
				if i == indexB {
					newStr += string(str[indexA])
				}
				if i != indexA && i != indexB {
					newStr += string(str[i])
				}
			}
		case "letter":
			for _, r := range str {
				if string(r) == instr[2] {
					newStr += instr[5]
				}
				if string(r) == instr[5] {
					newStr += instr[2]
				}
				if string(r) != instr[2] && string(r) != instr[5] {
					newStr += string(r)
				}
			}
		}
	case "rotate":
		switch instr[3] {
		case "step", "steps":
			var direction int
			steps, _ := strconv.Atoi(instr[2])
			switch instr[1] {
			case "left":
				direction = 1
			case "right":
				direction = -1
			}
			for i := range str {
				k := i + direction*steps
				if k >= len(str) {
					k = k % len(str)
				}
				for k < 0 {
					k += len(str)
				}
				newStr += string(str[k])
			}
		case "position":
			letter := instr[6]
			index := 0
			for i, r := range str {
				if string(r) == letter {
					index = i
					break
				}
			}
			numberRotations := 0
			numberRotations = 1 + index
			if index > 3 {
				numberRotations++
			}
			newStr = newString(str, []string{"rotate", "right", strconv.Itoa(numberRotations), "step"})
		}
	case "reverse":
		indexA, _ := strconv.Atoi(instr[2])
		indexB, _ := strconv.Atoi(instr[4])
		for i := 0; i < indexA; i++ {
			newStr += string(str[i])
		}
		for i := indexB; i >= indexA; i-- {
			newStr += string(str[i])
		}
		for i := indexB + 1; i < len(str); i++ {
			newStr += string(str[i])
		}
	case "move":
		indexA, _ := strconv.Atoi(instr[2])
		indexB, _ := strconv.Atoi(instr[5])
		for i := 0; i < len(str); i++ {
			if i == indexA {
				continue
			}
			if i == indexB {
				if indexA < indexB {
					newStr += string(str[i])
					newStr += string(str[indexA])
				} else {
					newStr += string(str[indexA])
					newStr += string(str[i])
				}
			}
			if i != indexA && i != indexB {
				newStr += string(str[i])
			}
		}
	}
	return newStr
}

func newStringInv(str string, instr []string) string {
	newStr := ""
	switch instr[0] {
	case "swap":
		switch instr[1] {
		case "position":
			indexA, _ := strconv.Atoi(instr[2])
			indexB, _ := strconv.Atoi(instr[5])
			for i := range str {
				if i == indexA {
					newStr += string(str[indexB])
				}
				if i == indexB {
					newStr += string(str[indexA])
				}
				if i != indexA && i != indexB {
					newStr += string(str[i])
				}
			}
		case "letter":
			for _, r := range str {
				if string(r) == instr[2] {
					newStr += instr[5]
				}
				if string(r) == instr[5] {
					newStr += instr[2]
				}
				if string(r) != instr[2] && string(r) != instr[5] {
					newStr += string(r)
				}
			}
		}
	case "rotate":
		switch instr[3] {
		case "step", "steps":
			var direction int
			steps, _ := strconv.Atoi(instr[2])
			switch instr[1] {
			case "left":
				direction = -1
			case "right":
				direction = 1
			}
			for i := range str {
				k := i + direction*steps
				if k >= len(str) {
					k = k % len(str)
				}
				for k < 0 {
					k += len(str)
				}
				newStr += string(str[k])
			}
		case "position":
			letter := instr[6]
			index := 0
			for i, r := range str {
				if string(r) == letter {
					index = i
					break
				}
			}
			numberRotations := reverse(index)
			newStr = newString(str, []string{"rotate", "right", strconv.Itoa(numberRotations), "step"})
		}
	case "reverse":
		indexA, _ := strconv.Atoi(instr[2])
		indexB, _ := strconv.Atoi(instr[4])
		for i := 0; i < indexA; i++ {
			newStr += string(str[i])
		}
		for i := indexB; i >= indexA; i-- {
			newStr += string(str[i])
		}
		for i := indexB + 1; i < len(str); i++ {
			newStr += string(str[i])
		}
	case "move":
		indexA, _ := strconv.Atoi(instr[5])
		indexB, _ := strconv.Atoi(instr[2])
		for i := 0; i < len(str); i++ {
			if i == indexA {
				continue
			}
			if i == indexB {
				if indexA < indexB {
					newStr += string(str[i])
					newStr += string(str[indexA])
				} else {
					newStr += string(str[indexA])
					newStr += string(str[i])
				}
			}
			if i != indexA && i != indexB {
				newStr += string(str[i])
			}
		}
	}
	return newStr
}

func part1(s string) string {
	instrList := format(s)
	str := ""
	if len(instrList) == 8 {
		str = "abcde"
	} else {
		str = "abcdefgh"
	}
	for _, instr := range instrList {
		str = newString(str, instr)
	}
	return str
}

func part2(s string) string {
	instrList := format(s)
	str := "fbgdceah"
	for i := len(instrList) - 1; i >= 0; i-- {
		str = newStringInv(str, instrList[i])
	}
	return str
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day21/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day21/input.txt")

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
