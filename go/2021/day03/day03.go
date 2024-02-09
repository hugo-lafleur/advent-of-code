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
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func mostColumn(tab [][]string, index int) int {
	ones := 0
	zeros := 0
	for j := 0; j < len(tab); j++ {
		if tab[j][index] == "1" {
			ones++
		} else {
			zeros++
		}
	}
	if ones > zeros {
		return 1
	}
	if zeros > ones {
		return 0
	}
	return 2
}

func part1(s string) int {
	tab := format(s)
	gamma := ""
	epsilon := ""
	for i := 0; i < len(tab[0]); i++ {
		most := mostColumn(tab, i)
		gamma += strconv.Itoa(most)
		epsilon += strconv.Itoa(1 - most)
	}
	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(g * e)
}

func part2(s string) int {
	tab := format(s)
	generator := ""
	scrubber := ""
	currTab := [][]string{}
	for k := 0; k < len(tab); k++ {
		currTab = append(currTab, tab[k])
	}
	for j := 0; j < len(tab[0]); j++ {
		most := mostColumn(currTab, j)
		if most == 2 {
			most = 1
		}
		newTab := [][]string{}
		for k := 0; k < len(currTab); k++ {
			if currTab[k][j] == strconv.Itoa(most) {
				newTab = append(newTab, currTab[k])
			}
		}
		currTab = newTab
		if len(currTab) == 1 {
			generator = strings.Join(currTab[0], "")
			break
		}
	}
	currTab = [][]string{}
	for k := 0; k < len(tab); k++ {
		currTab = append(currTab, tab[k])
	}
	for j := 0; j < len(tab[0]); j++ {
		most := mostColumn(currTab, j)
		least := 1 - most
		if most == 2 {
			least = 0
		}
		newTab := [][]string{}
		for k := 0; k < len(currTab); k++ {
			if currTab[k][j] == strconv.Itoa(least) {
				newTab = append(newTab, currTab[k])
			}
		}
		currTab = newTab
		if len(currTab) == 1 {
			scrubber = strings.Join(currTab[0], "")
			break
		}
	}
	g, _ := strconv.ParseInt(generator, 2, 64)
	scr, _ := strconv.ParseInt(scrubber, 2, 64)
	return int(g * scr)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day03/test.data")

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

	content, err = os.ReadFile("../../../inputs/2021/day03/input.data")

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
