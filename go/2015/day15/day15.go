package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]int {
	lines := strings.Split(s, "\n")
	l := len(lines)
	tab := [][]string{}
	for _, x := range lines {
		tab = append(tab, strings.Split(x, " "))
	}
	graph := make([][]int, l)
	for i := range graph {
		graph[i] = make([]int, 5)
	}
	for i, x := range tab {
		a, _ := strconv.Atoi(x[2][:len(x[2])-1])
		b, _ := strconv.Atoi(x[4][:len(x[4])-1])
		c, _ := strconv.Atoi(x[6][:len(x[6])-1])
		d, _ := strconv.Atoi(x[8][:len(x[8])-1])
		e, _ := strconv.Atoi(x[10])
		graph[i][0] = a
		graph[i][1] = b
		graph[i][2] = c
		graph[i][3] = d
		graph[i][4] = e
	}
	return graph
}

func part1(s string) int {
	graph := format(s)
	n := len(graph)
	res := 0
	tmp := 0
	if n == 2 {
		i := 0
		for i < 100 {
			j := 0
			for j < 100 {
				if i+j == 100 {
					c := i*graph[0][0] + j*graph[1][0]
					d := i*graph[0][1] + j*graph[1][1]
					f := i*graph[0][2] + j*graph[1][2]
					t := i*graph[0][3] + j*graph[1][3]
					if c < 0 || d < 0 || f < 0 || t < 0 {
						tmp = 0
					} else {
						tmp = c * d * f * t
					}
					if tmp > res {
						res = tmp
					}
				}
				j++
			}
			i++
		}
	}
	if n == 4 {
		i := 0
		for i < 100 {
			j := 0
			for j < 100 {
				k := 0
				for k < 100 {
					l := 0
					for l < 100 {
						if i+j+k+l == 100 {
							c := i*graph[0][0] + j*graph[1][0] + k*graph[2][0] + l*graph[3][0]
							d := i*graph[0][1] + j*graph[1][1] + k*graph[2][1] + l*graph[3][1]
							f := i*graph[0][2] + j*graph[1][2] + k*graph[2][2] + l*graph[3][2]
							t := i*graph[0][3] + j*graph[1][3] + k*graph[2][3] + l*graph[3][3]
							if c < 0 || d < 0 || f < 0 || t < 0 {
								tmp = 0
							} else {
								tmp = c * d * f * t
							}
							if tmp > res {
								res = tmp
							}
						}
						l++
					}
					k++
				}
				j++
			}
			i++
		}
	}
	return res
}

func part2(s string) int {
	graph := format(s)
	n := len(graph)
	res := 0
	tmp := 0
	if n == 2 {
		i := 0
		for i < 100 {
			j := 0
			for j < 100 {
				if i+j == 100 && i*graph[0][4]+j*graph[1][4] == 500 {
					c := i*graph[0][0] + j*graph[1][0]
					d := i*graph[0][1] + j*graph[1][1]
					f := i*graph[0][2] + j*graph[1][2]
					t := i*graph[0][3] + j*graph[1][3]
					if c < 0 || d < 0 || f < 0 || t < 0 {
						tmp = 0
					} else {
						tmp = c * d * f * t
					}
					if tmp > res {
						res = tmp
					}
				}
				j++
			}
			i++
		}
	}
	if n == 4 {
		i := 0
		for i < 100 {
			j := 0
			for j < 100 {
				k := 0
				for k < 100 {
					l := 0
					for l < 100 {
						if i+j+k+l == 100 && i*graph[0][4]+j*graph[1][4]+k*graph[2][4]+l*graph[3][4] == 500 {
							c := i*graph[0][0] + j*graph[1][0] + k*graph[2][0] + l*graph[3][0]
							d := i*graph[0][1] + j*graph[1][1] + k*graph[2][1] + l*graph[3][1]
							f := i*graph[0][2] + j*graph[1][2] + k*graph[2][2] + l*graph[3][2]
							t := i*graph[0][3] + j*graph[1][3] + k*graph[2][3] + l*graph[3][3]
							if c < 0 || d < 0 || f < 0 || t < 0 {
								tmp = 0
							} else {
								tmp = c * d * f * t
							}
							if tmp > res {
								res = tmp
							}
						}
						l++
					}
					k++
				}
				j++
			}
			i++
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day15/test.data")

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

	content, err = os.ReadFile("../../../inputs/2015/day15/input.data")

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
