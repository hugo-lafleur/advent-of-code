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

func printTab(t [][]string) {
	for _, line := range t {
		fmt.Println(line)
	}
}

func part1(s string) int {
	c := 0
	list := format(s)
	var wide int
	var length int
	if len(list) == 4 {
		wide = 7
		length = 3
	} else {
		wide = 50
		length = 6
	}
	tab := make([][]bool, length)
	for i := 0; i < length; i++ {
		tab[i] = make([]bool, wide)
	}
	for _, line := range list {
		if line[0] == "rect" {
			nbrs := strings.Split(line[1], "x")
			n1, _ := strconv.Atoi(nbrs[0])
			n2, _ := strconv.Atoi(nbrs[1])
			for i := 0; i < n1; i++ {
				for j := 0; j < n2; j++ {
					tab[j][i] = true
				}
			}
		}
		if line[0] == "rotate" {
			if line[1] == "row" {
				eq := strings.Split(line[2], "=")
				row, _ := strconv.Atoi(eq[1])
				n, _ := strconv.Atoi(line[4])
				temp := []bool{}
				for j := 0; j < wide; j++ {
					temp = append(temp, tab[row][j])
				}
				for j := 0; j < wide; j++ {
					tab[row][j] = temp[((j-n)+wide)%wide]
				}
			}
			if line[1] == "column" {
				eq := strings.Split(line[2], "=")
				column, _ := strconv.Atoi(eq[1])
				n, _ := strconv.Atoi(line[4])
				temp := []bool{}
				for j := 0; j < length; j++ {
					temp = append(temp, tab[j][column])
				}
				for j := 0; j < length; j++ {
					tab[j][column] = temp[((j-n)+length)%length]
				}
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < wide; j++ {
			if tab[i][j] {
				c++
			}
		}
	}
	return c
}

func part2(s string) {
	c := 0
	list := format(s)
	var wide int
	var length int
	if len(list) == 4 {
		wide = 7
		length = 3
	} else {
		wide = 50
		length = 6
	}
	tab := make([][]string, length)
	for i := 0; i < length; i++ {
		tab[i] = make([]string, wide)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < wide; j++ {
			tab[i][j] = "."
		}
	}
	for _, line := range list {
		if line[0] == "rect" {
			nbrs := strings.Split(line[1], "x")
			n1, _ := strconv.Atoi(nbrs[0])
			n2, _ := strconv.Atoi(nbrs[1])
			for i := 0; i < n1; i++ {
				for j := 0; j < n2; j++ {
					tab[j][i] = "#"
				}
			}
		}
		if line[0] == "rotate" {
			if line[1] == "row" {
				eq := strings.Split(line[2], "=")
				row, _ := strconv.Atoi(eq[1])
				n, _ := strconv.Atoi(line[4])
				temp := []string{}
				for j := 0; j < wide; j++ {
					temp = append(temp, tab[row][j])
				}
				for j := 0; j < wide; j++ {
					tab[row][j] = temp[((j-n)+wide)%wide]
				}
			}
			if line[1] == "column" {
				eq := strings.Split(line[2], "=")
				column, _ := strconv.Atoi(eq[1])
				n, _ := strconv.Atoi(line[4])
				temp := []string{}
				for j := 0; j < length; j++ {
					temp = append(temp, tab[j][column])
				}
				for j := 0; j < length; j++ {
					tab[j][column] = temp[((j-n)+length)%length]
				}
			}
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < wide; j++ {
			if tab[i][j] == "#" {
				c++
			}
		}
	}
	printTab(tab)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day08/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : \n")
	part2(string(content))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day08/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : \n")
	part2(string(content))
	fmt.Println(time.Since(start))
}
