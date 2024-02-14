package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type cube struct {
	x1    int
	x2    int
	y1    int
	y2    int
	z1    int
	z2    int
	state int
	inter int
}

func format(s string) [][]int {
	data := [][]int{}
	b := 0
	str := ""
	tab := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = line + ","
		tab = append(tab, strings.Split(line, ""))
	}
	for i, line := range tab {
		b = 0
		if line[1] == "n" {
			b = 1
		}
		data = append(data, []int{})
		data[i] = append(data[i], b)
		for j, c := range line {
			if c == "=" {
				k := j + 1
				for line[k] != "." {
					str = str + line[k]
					k = k + 1
				}
				n, _ := strconv.Atoi(str)
				data[i] = append(data[i], n)
				k = k + 2
				str = ""
				for line[k] != "," {
					str = str + line[k]
					k = k + 1
				}
				n, _ = strconv.Atoi(str)
				data[i] = append(data[i], n)
				str = ""
			}
		}
	}
	return data
}

func cubes_on(cubes [200][200][200]int) int {
	c := 0
	for _, d1 := range cubes {
		for _, d2 := range d1 {
			for _, x := range d2 {
				if x == 1 {
					c += 1
				}
			}
		}
	}
	return c
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func initialization(data [][]int) [200][200][200]int {
	cubes := [200][200][200]int{}

	for _, line := range data {
		if abs(line[1]) > 100 {
			continue
		}
		for x := line[1]; x < line[2]+1; x++ {
			for y := line[3]; y < line[4]+1; y++ {
				for z := line[5]; z < line[6]+1; z++ {
					cubes[x+50][y+50][z+50] = line[0]
				}
			}
		}
	}
	return cubes
}

func cube_struct(data [][]int) []cube {
	l := []cube{}
	for _, line := range data {
		l = append(l, cube{line[1], line[2], line[3], line[4], line[5], line[6], line[0], 0})
	}
	return l
}

func intersect(a cube, b cube) bool {
	return (a.x1 <= b.x2 && a.x2 >= b.x1) && (a.y1 <= b.y2 && a.y2 >= b.y1) && (a.z1 <= b.z2 && a.z2 >= b.z1)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
func intersection(a cube, b cube) cube {
	c := cube{}
	c.x1 = max(a.x1, b.x1)
	c.x2 = min(a.x2, b.x2)
	c.y1 = max(a.y1, b.y1)
	c.y2 = min(a.y2, b.y2)
	c.z1 = max(a.z1, b.z1)
	c.z2 = min(a.z2, b.z2)
	c.state = a.state
	if a.inter == 1 || b.inter == 1 {
		c.inter = 0
	} else {
		c.inter = 1
	}
	return c
}

func volume(a cube) int {
	return (1 + a.x2 - a.x1) * (1 + a.y2 - a.y1) * (1 + a.z2 - a.z1) * a.state
}

func add_volumes(cubes []cube) int {
	seen := []cube{}
	v := 0
	for _, cube := range cubes {
		v += volume(cube)
		for _, kube := range seen {
			if intersect(kube, cube) {
				inters := intersection(kube, cube)
				if inters.inter == 1 {
					v -= volume(intersection(kube, cube))
				} else {
					v += volume(intersection(kube, cube))
				}
				seen = append(seen, inters)
			}
		}
		seen = append(seen, cube)
	}
	return v
}

func part1(s string) int {
	data := format(s)
	return cubes_on(initialization(data))
}

func part2(s string) int {
	data := format(s)
	cubes := cube_struct(data)
	return add_volumes(cubes)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day22/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2021/day22/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2021/day22/input.data")

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
