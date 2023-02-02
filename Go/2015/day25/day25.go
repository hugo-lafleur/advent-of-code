package main

import (
	"fmt"
	"time"
)

func part1() int {
	tab := [7000][7000]int{}
	n := 2980 + 3074 + 1
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
	//fmt.Println(tab)
	return tab[2980][3074]
}

func part2() int {
	return 0
}

func main() {
	fmt.Printf("Input :\n")
	start := time.Now()
	fmt.Printf("Part 1 : %d\n", part1())
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("Part 2 : %d\n", part2())
	fmt.Println(time.Since(start))
}
