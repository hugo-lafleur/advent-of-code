package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type cube struct {
	x, y, z int
}

func Split(r rune) bool {
	return r == ',' || r == '~'
}

func format(s string) (map[cube]int, [][]cube) {
	res := make(map[cube]int)
	bricks := [][]cube{}
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		br := []cube{}
		lineSplit := strings.FieldsFunc(line, Split)
		a, _ := strconv.Atoi(lineSplit[0])
		b, _ := strconv.Atoi(lineSplit[1])
		c, _ := strconv.Atoi(lineSplit[2])
		d, _ := strconv.Atoi(lineSplit[3])
		e, _ := strconv.Atoi(lineSplit[4])
		f, _ := strconv.Atoi(lineSplit[5])
		for x := a; x <= d; x++ {
			for y := b; y <= e; y++ {
				for z := c; z <= f; z++ {
					res[cube{x, y, z}] = i
					br = append(br, cube{x, y, z})
				}
			}
		}
		bricks = append(bricks, br)
	}
	return res, bricks
}

func remove[T comparable](l []T, n T) []T {
	res := []T{}
	for _, x := range l {
		if x != n {
			res = append(res, x)
		}
	}
	return res
}

func fullMapping(s string) (map[cube]int, [][]cube) {
	mapping, bricks := format(s)
	L := len(strings.Split(s, "\n"))
	movement := true
	for movement {
		movement = false
		newMapping := make(map[cube]int)
		for i := 0; i < L; i++ {
			brick := bricks[i]
			fall := true
			for _, c := range brick {
				n, ok := mapping[cube{c.x, c.y, c.z - 1}]
				if (ok && n != i) || c.z == 1 {
					fall = false
					break
				}
			}
			if fall {
				movement = true
				for j, c := range brick {
					newMapping[cube{c.x, c.y, c.z - 1}] = i
					bricks[i][j] = cube{c.x, c.y, c.z - 1}
				}
			} else {
				for _, c := range brick {
					newMapping[cube{c.x, c.y, c.z}] = i
				}
			}
		}
		mapping = newMapping
	}
	return mapping, bricks
}

func part1(s string) int {
	mapping, bricks := fullMapping(s)
	L := len(strings.Split(s, "\n"))
	supporting := make([][]int, L)
	for i := 0; i < L; i++ {
		brick := bricks[i]
		supporting[i] = []int{}
		for _, c := range brick {
			n, ok := mapping[cube{c.x, c.y, c.z + 1}]
			if ok && n != i {
				supporting[i] = append(supporting[i], n)
			}
		}
	}
	c := 0
	for key, value := range supporting {
	loop:
		for _, b := range value {
			for key2, value2 := range supporting {
				if key != key2 && slices.Contains(value2, b) {
					continue loop
				}
			}
			c++
			break
		}
	}
	return L - c
}

func part2(s string) int {
	mapping, bricks := fullMapping(s)
	L := len(strings.Split(s, "\n"))
	supported := make([][]int, L)
	supporting := make([][]int, L)
	toDisintigrate := []int{}
	for i := 0; i < L; i++ {
		brick := bricks[i]
		supported[i] = []int{}
		for _, c := range brick {
			n, ok := mapping[cube{c.x, c.y, c.z - 1}]
			if ok && n != i {
				supported[i] = append(supported[i], n)
			}
			if c.z == 1 {
				supported[i] = []int{-1}
			}
		}
	}
	for i := 0; i < L; i++ {
		brick := bricks[i]
		supporting[i] = []int{}
		for _, c := range brick {
			n, ok := mapping[cube{c.x, c.y, c.z + 1}]
			if ok && n != i {
				supporting[i] = append(supporting[i], n)
			}
		}
	}
	for key, value := range supporting {
	loop:
		for _, b := range value {
			for key2, value2 := range supporting {
				if key != key2 && slices.Contains(value2, b) {
					continue loop
				}
			}
			toDisintigrate = append(toDisintigrate, key)
			break
		}
	}
	c := 0
	for _, br := range toDisintigrate {
		var q deque.Deque[int]
		copy := slices.Clone(supported)
		temp := 0
		q.PushBack(br)
		for q.Len() != 0 {
			curr := q.PopBack()
			temp++
			for key := range copy {
				for _, x := range copy[key] {
					if x == curr {
						copy[key] = remove(copy[key], x)
						break
					}
				}
				if len(copy[key]) == 0 {
					copy[key] = append(copy[key], -1)
					q.PushBack(key)
				}
			}
		}
		c += temp - 1
	}

	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day22/test.data")

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

	content, err = os.ReadFile("../../../inputs/2023/day22/input.data")

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
