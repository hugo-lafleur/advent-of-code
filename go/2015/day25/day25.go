package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(s string) int {
	line := strings.Split(s, " ")
	row := 0
	column := 0
	for i, w := range line {
		if w == "row" {
			n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
			row = n
		}
		if w == "column" {
			n, _ := strconv.Atoi(line[i+1][:len(line[i+1])-1])
			column = n
		}

	}
	tab := [7000][7000]int{}
	n := row + column - 1
	i := 1
	tab[0][0] = 20151125
	for i < n {
		j := 0
		for j <= i {
			if j == 0 {
				tab[i-j][j] = (252533 * tab[0][i-1]) % 33554393
			} else {
				tab[i-j][j] = (252533 * tab[i-j+1][j-1]) % 33554393
			}
			j++
		}
		i++
	}
	return tab[2980][3074]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day25/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
