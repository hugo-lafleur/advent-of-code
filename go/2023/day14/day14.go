package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type platform struct {
	rocks [][]string
}

var cache = make(map[string]int)

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func (p platform) tilt(direction string) {
	tab := p.rocks
	if direction == "N" {
		for r := 0; r < len(tab); r++ {
			for i := 1; i < len(tab); i++ {
				for j := 0; j < len(tab[i]); j++ {
					if tab[i][j] == "O" && tab[i-1][j] == "." {
						tab[i][j] = "."
						tab[i-1][j] = "O"
					}
				}
			}
		}
	}
	if direction == "S" {
		for r := 0; r < len(tab); r++ {
			for i := 0; i < len(tab)-1; i++ {
				for j := 0; j < len(tab[i]); j++ {
					if tab[i][j] == "O" && tab[i+1][j] == "." {
						tab[i][j] = "."
						tab[i+1][j] = "O"
					}
				}
			}
		}
	}
	if direction == "W" {
		for r := 0; r < len(tab); r++ {
			for i := 0; i < len(tab); i++ {
				for j := 1; j < len(tab[i]); j++ {
					if tab[i][j] == "O" && tab[i][j-1] == "." {
						tab[i][j] = "."
						tab[i][j-1] = "O"
					}
				}
			}
		}
	}
	if direction == "E" {
		for r := 0; r < len(tab); r++ {
			for i := 0; i < len(tab); i++ {
				for j := 0; j < len(tab[i])-1; j++ {
					if tab[i][j] == "O" && tab[i][j+1] == "." {
						tab[i][j] = "."
						tab[i][j+1] = "O"
					}
				}
			}
		}
	}
}

func (p platform) count() int {
	res := 0
	tab := p.rocks
	for k, line := range tab {
		for _, char := range line {
			if char == "O" {
				res += len(tab) - k
			}
		}
	}
	return res
}

func (p platform) cycle() {
	p.tilt("N")
	p.tilt("W")
	p.tilt("S")
	p.tilt("E")
}

func sliceToString(tab [][]string) string {
	res := ""
	for _, line := range tab {
		for _, x := range line {
			res += x
		}
	}
	return res
}

func part1(s string) int {
	tab := format(s)
	p := platform{tab}
	p.tilt("N")
	return p.count()
}

func part2(s string) int {
	d := 1000000000
	tab := format(s)
	p := platform{tab}
	for i := 0; i < d; i++ {
		p.cycle()
		l, ok := cache[sliceToString(p.rocks)]
		if ok {
			k := i - l
			for i < d-k {
				i += k
			}
			i++
			for i < d {
				p.cycle()
				i++
			}
			return p.count()
		}
		cache[sliceToString(p.rocks)] = i
	}
	return 0
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
