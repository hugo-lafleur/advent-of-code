package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type cube struct {
	x, y, z int
}

type cube4d struct {
	x, y, z, w int
}

func format(s string) map[cube]bool {
	res := make(map[cube]bool)
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ""))
	}
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == "#" {
				res[cube{i, j, 0}] = true
			}
		}
	}
	return res
}

func format4d(s string) map[cube4d]bool {
	res := make(map[cube4d]bool)
	lines := strings.Split(s, "\n")
	tab := [][]string{}
	for _, line := range lines {
		tab = append(tab, strings.Split(line, ""))
	}
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == "#" {
				res[cube4d{i, j, 0, 0}] = true
			}
		}
	}
	return res
}

func addCube(c1, c2 cube) cube {
	return cube{c1.x + c2.x, c1.y + c2.y, c1.z + c2.z}
}

func addCube4d(c1, c2 cube4d) cube4d {
	return cube4d{c1.x + c2.x, c1.y + c2.y, c1.z + c2.z, c1.w + c2.w}
}

func adjacents(c cube) []cube {
	res := []cube{}
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				res = append(res, addCube(c, cube{x, y, z}))
			}
		}
	}
	return res
}

func adjacents4d(c cube4d) []cube4d {
	res := []cube4d{}
	for _, x := range []int{-1, 0, 1} {
		for _, y := range []int{-1, 0, 1} {
			for _, z := range []int{-1, 0, 1} {
				for _, w := range []int{-1, 0, 1} {
					res = append(res, addCube4d(c, cube4d{x, y, z, w}))
				}
			}
		}
	}
	return res
}

func numbersActive(c cube, cubes map[cube]bool) int {
	res := 0
	for _, neighbors := range adjacents(c) {
		if cubes[neighbors] {
			res++
		}
	}
	if cubes[c] {
		return res - 1
	}
	return res
}

func numbersActive4d(c cube4d, cubes map[cube4d]bool) int {
	res := 0
	for _, neighbors := range adjacents4d(c) {
		if cubes[neighbors] {
			res++
		}
	}
	if cubes[c] {
		return res - 1
	}
	return res
}

func part1(s string) int {
	cubes := format(s)
	cycles := 0
	for cycles < 6 {
		newCubes := make(map[cube]bool)
		for cube := range cubes {
			for _, neighbors := range adjacents(cube) {
				active := numbersActive(neighbors, cubes)
				if cubes[neighbors] {
					if active == 2 || active == 3 {
						newCubes[neighbors] = true
					}
					continue
				} else {
					if active == 3 {
						newCubes[neighbors] = true
					}
					continue
				}
			}
		}
		cubes = newCubes
		cycles++
	}
	res := 0
	for _, state := range cubes {
		if state {
			res++
		}
	}
	return res
}

func part2(s string) int {
	cubes := format4d(s)
	cycles := 0
	for cycles < 6 {
		newCubes := make(map[cube4d]bool)
		for cube := range cubes {
			for _, neighbors := range adjacents4d(cube) {
				active := numbersActive4d(neighbors, cubes)
				if cubes[neighbors] {
					if active == 2 || active == 3 {
						newCubes[neighbors] = true
					}
					continue
				} else {
					if active == 3 {
						newCubes[neighbors] = true
					}
					continue
				}
			}
		}
		cubes = newCubes
		cycles++
	}
	res := 0
	for _, state := range cubes {
		if state {
			res++
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day17/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day17/input.data")

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
