package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector struct {
	x, y, z int
}

type moon struct {
	pos, vel vector
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, list ...int) int {
	res := a * b / GCD(a, b)
	for i := 0; i < len(list); i++ {
		res = LCM(res, list[i])
	}
	return res
}

func Split(r rune) bool {
	return r == ',' || r == ' ' || r == '=' || r == '<' || r == '>'
}

func format(s string) []moon {
	lines := strings.Split(s, "\n")
	res := []moon{}
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		x, _ := strconv.Atoi(lineSplit[1])
		y, _ := strconv.Atoi(lineSplit[3])
		z, _ := strconv.Atoi(lineSplit[5])
		res = append(res, moon{vector{x, y, z}, vector{0, 0, 0}})
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func energy(m moon) int {
	return (abs(m.pos.x) + abs(m.pos.y) + abs(m.pos.z)) * (abs(m.vel.x) + abs(m.vel.y) + abs(m.vel.z))
}

func dimToString(a, b, c, d, e, f, g, h int) string {
	return strconv.Itoa(a) + "," + strconv.Itoa(b) + "," + strconv.Itoa(c) + "," + strconv.Itoa(d) + "," + strconv.Itoa(e) + "," + strconv.Itoa(f) + "," + strconv.Itoa(g) + "," + strconv.Itoa(h)
}

func part1(s string) int {
	c := 0
	moons := format(s)
	steps := 1000
	n := len(moons)
	for k := 0; k < steps; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				m1 := moons[i]
				m2 := moons[j]
				if m1.pos.x > m2.pos.x {
					moons[i].vel.x--
					moons[j].vel.x++
				}
				if m1.pos.x < m2.pos.x {
					moons[i].vel.x++
					moons[j].vel.x--
				}
				if m1.pos.y > m2.pos.y {
					moons[i].vel.y--
					moons[j].vel.y++
				}
				if m1.pos.y < m2.pos.y {
					moons[i].vel.y++
					moons[j].vel.y--
				}
				if m1.pos.z > m2.pos.z {
					moons[i].vel.z--
					moons[j].vel.z++
				}
				if m1.pos.z < m2.pos.z {
					moons[i].vel.z++
					moons[j].vel.z--
				}
			}
		}
		for i := 0; i < n; i++ {
			moons[i].pos.x += moons[i].vel.x
			moons[i].pos.y += moons[i].vel.y
			moons[i].pos.z += moons[i].vel.z
		}
	}
	for _, m := range moons {
		c += energy(m)
	}
	return c
}

func part2(s string) int {
	visited := make(map[string]int)
	res := make(map[int]int)
	moons := format(s)
	n := len(moons)
	steps := 0
	for {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				m1 := moons[i]
				m2 := moons[j]
				if m1.pos.x > m2.pos.x {
					moons[i].vel.x--
					moons[j].vel.x++
				}
				if m1.pos.x < m2.pos.x {
					moons[i].vel.x++
					moons[j].vel.x--
				}
				if m1.pos.y > m2.pos.y {
					moons[i].vel.y--
					moons[j].vel.y++
				}
				if m1.pos.y < m2.pos.y {
					moons[i].vel.y++
					moons[j].vel.y--
				}
				if m1.pos.z > m2.pos.z {
					moons[i].vel.z--
					moons[j].vel.z++
				}
				if m1.pos.z < m2.pos.z {
					moons[i].vel.z++
					moons[j].vel.z--
				}
			}
		}
		for i := 0; i < n; i++ {
			moons[i].pos.x += moons[i].vel.x
			moons[i].pos.y += moons[i].vel.y
			moons[i].pos.z += moons[i].vel.z
		}
		steps++
		var strX, strY, strZ string
		strX = dimToString(moons[0].pos.x, moons[1].pos.x, moons[2].pos.x, moons[3].pos.x, moons[0].vel.x, moons[1].vel.x, moons[2].vel.x, moons[3].vel.x)
		strY = dimToString(moons[0].pos.y, moons[1].pos.y, moons[2].pos.y, moons[3].pos.y, moons[0].vel.y, moons[1].vel.y, moons[2].vel.y, moons[3].vel.y)
		strZ = dimToString(moons[0].pos.z, moons[1].pos.z, moons[2].pos.z, moons[3].pos.z, moons[0].vel.z, moons[1].vel.z, moons[2].vel.z, moons[3].vel.z)
		stepX, okX := visited[strX]
		_, doneX := res[0]
		if okX && !doneX {
			res[0] = steps - stepX
		}
		stepY, okY := visited[strY]
		_, doneY := res[1]
		if okY && !doneY {
			res[1] = steps - stepY
		}
		stepZ, okZ := visited[strZ]
		_, doneZ := res[2]
		if okZ && !doneZ {
			res[2] = steps - stepZ
		}
		if doneX && doneY && doneZ {
			return LCM(res[0], res[1], res[2])
		}
		visited[strX] = steps
		visited[strY] = steps
		visited[strZ] = steps
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day12/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2019/day12/input.txt")

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
