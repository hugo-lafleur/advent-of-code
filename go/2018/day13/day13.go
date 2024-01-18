package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type cart struct {
	index int
	x, y  int
	d     string
	k     int
}

type carts []cart

func (c carts) Len() int {
	return len(c)
}

func (c carts) Less(i, j int) bool {
	if c[i].x == c[j].x {
		return c[i].y < c[j].y
	}
	return c[i].x < c[j].x
}

func (c carts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func collide(carts carts) (int, int, bool) {
	l := len(carts)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if carts[i].x == carts[j].x && carts[i].y == carts[j].y {
				return carts[i].x, carts[i].y, true
			}
		}
	}
	return 0, 0, false
}

func remove(l carts, i, j int) carts {
	var res carts
	for k, c := range l {
		if k == i || k == j {
			continue
		}
		res = append(res, c)
	}
	return res
}
func part1(s string) string {
	tab := format(s)
	var carts carts
	k := 0
	for i, line := range tab {
		for j, s := range line {
			switch s {
			case "v":
				carts = append(carts, cart{k, i, j, "d", 0})
				tab[i][j] = "|"
				k++
			case "^":
				carts = append(carts, cart{k, i, j, "u", 0})
				tab[i][j] = "|"
				k++
			case "<":
				carts = append(carts, cart{k, i, j, "l", 0})
				tab[i][j] = "-"
				k++
			case ">":
				carts = append(carts, cart{k, i, j, "r", 0})
				tab[i][j] = "-"
				k++
			}
		}
	}
	for {
		sort.Sort(carts)
		for i, cart := range carts {
			x, y := cart.x, cart.y
			d := cart.d
			switch tab[x][y] {
			case "-":
				if d == "l" {
					carts[i].y--
				}
				if d == "r" {
					carts[i].y++
				}
			case "|":
				if d == "u" {
					carts[i].x--
				}
				if d == "d" {
					carts[i].x++
				}
			case "/":
				switch d {
				case "u":
					carts[i].y++
					carts[i].d = "r"
				case "l":
					carts[i].x++
					carts[i].d = "d"
				case "r":
					carts[i].x--
					carts[i].d = "u"
				case "d":
					carts[i].y--
					carts[i].d = "l"
				}
			case "\\":
				switch d {
				case "u":
					carts[i].y--
					carts[i].d = "l"
				case "l":
					carts[i].x--
					carts[i].d = "u"
				case "r":
					carts[i].x++
					carts[i].d = "d"
				case "d":
					carts[i].y++
					carts[i].d = "r"
				}
			case "+":
				switch d {
				case "u":
					switch cart.k {
					case 0:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k++
					case 1:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k++
					case 2:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k = 0
					}
				case "d":
					switch cart.k {
					case 0:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k++
					case 1:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k++
					case 2:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k = 0
					}
				case "l":
					switch cart.k {
					case 0:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k++
					case 1:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k++
					case 2:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k = 0
					}
				case "r":
					switch cart.k {
					case 0:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k++
					case 1:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k++
					case 2:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k = 0
					}
				}
			}
			a, b, isCollision := collide(carts)
			if isCollision {
				return strconv.Itoa(b) + "," + strconv.Itoa(a)
			}
		}
	}
}

func part2(s string) string {
	tab := format(s)
	var carts carts
	k := 0
	for i, line := range tab {
		for j, s := range line {
			switch s {
			case "v":
				carts = append(carts, cart{k, i, j, "d", 0})
				tab[i][j] = "|"
				k++
			case "^":
				carts = append(carts, cart{k, i, j, "u", 0})
				tab[i][j] = "|"
				k++
			case "<":
				carts = append(carts, cart{k, i, j, "l", 0})
				tab[i][j] = "-"
				k++
			case ">":
				carts = append(carts, cart{k, i, j, "r", 0})
				tab[i][j] = "-"
				k++
			}
		}
	}
	for {
		sort.Sort(carts)
		for i := 0; i < len(carts); i++ {
			if i < 0 || i >= len(carts) {
				continue
			}
			cart := carts[i]
			x, y := cart.x, cart.y
			d := cart.d
			switch tab[x][y] {
			case "-":
				if d == "l" {
					carts[i].y--
				}
				if d == "r" {
					carts[i].y++
				}
			case "|":
				if d == "u" {
					carts[i].x--
				}
				if d == "d" {
					carts[i].x++
				}
			case "/":
				switch d {
				case "u":
					carts[i].y++
					carts[i].d = "r"
				case "l":
					carts[i].x++
					carts[i].d = "d"
				case "r":
					carts[i].x--
					carts[i].d = "u"
				case "d":
					carts[i].y--
					carts[i].d = "l"
				}
			case "\\":
				switch d {
				case "u":
					carts[i].y--
					carts[i].d = "l"
				case "l":
					carts[i].x--
					carts[i].d = "u"
				case "r":
					carts[i].x++
					carts[i].d = "d"
				case "d":
					carts[i].y++
					carts[i].d = "r"
				}
			case "+":
				switch d {
				case "u":
					switch cart.k {
					case 0:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k++
					case 1:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k++
					case 2:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k = 0
					}
				case "d":
					switch cart.k {
					case 0:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k++
					case 1:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k++
					case 2:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k = 0
					}
				case "l":
					switch cart.k {
					case 0:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k++
					case 1:
						carts[i].y--
						carts[i].d = "l"
						carts[i].k++
					case 2:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k = 0
					}
				case "r":
					switch cart.k {
					case 0:
						carts[i].x--
						carts[i].d = "u"
						carts[i].k++
					case 1:
						carts[i].y++
						carts[i].d = "r"
						carts[i].k++
					case 2:
						carts[i].x++
						carts[i].d = "d"
						carts[i].k = 0
					}
				}
			}
			for j := 0; j < len(carts); j++ {
				if i != j && carts[i].x == carts[j].x && carts[i].y == carts[j].y {
					carts = remove(carts, i, j)
					i--
					if j < i+1 {
						i--
					}
					break
				}
			}
		}
		if len(carts) == 1 {
			a, b := carts[0].x, carts[0].y
			return strconv.Itoa(b) + "," + strconv.Itoa(a)
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day13/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))

	content, err = os.ReadFile("../../../inputs/2018/day13/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2018/day13/input.data")

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
