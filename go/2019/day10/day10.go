package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

type angle struct {
	ast    complex128
	vector complex128
}

type angleList []angle

func (a angleList) Len() int {
	return len(a)
}
func (a angleList) Less(i, j int) bool {
	a1 := math.Atan2(-imag(a[i].vector), real(a[i].vector))
	a2 := math.Atan2(-imag(a[j].vector), real(a[j].vector))
	if a1 > math.Pi/2 {
		a1 = -2*math.Pi + a1
	}
	if a2 > math.Pi/2 {
		a2 = -2*math.Pi + a2
	}
	return a1 > a2
}
func (a angleList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func format(s string) []complex128 {
	lines := strings.Split(s, "\n")
	res := []complex128{}
	for i, line := range lines {
		lineSplit := strings.Split(line, "")
		for j, s := range lineSplit {
			if s == "#" {
				res = append(res, complex(float64(j), float64(i)))
			}
		}
	}
	return res
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func miniDiff(c complex128) complex128 {
	gcd := abs(GCD(int(real(c)), int(imag(c))))
	r := int(real(c)) / gcd
	i := int(imag(c)) / gcd
	return complex(float64(r), float64(i))
}

func isIn(c complex128, l []complex128) bool {
	for _, x := range l {
		if c == x {
			return true
		}
	}
	return false
}

func maxMapValue(m map[complex128]int) int {
	var max int
	for _, value := range m {
		max = value
	}
	for _, value := range m {
		if value > max {
			max = value
		}
	}
	return max
}

func maxMapKey(m map[complex128]int) complex128 {
	var max int
	var res complex128
	for _, value := range m {
		max = value
	}
	for key, value := range m {
		if value > max {
			max = value
			res = key
		}
	}
	return res
}

func isDetectable(ast, otherAst complex128, asteroids []complex128) bool {
	if ast != otherAst {
		minDiff := miniDiff(otherAst - ast)
		diff := minDiff
		for {
			if ast+diff == otherAst {
				return true
			}
			if isIn(ast+diff, asteroids) && ast+diff != otherAst && ast+diff != ast {
				return false
			}
			diff += minDiff
		}
	}
	return false
}

func part1(s string) int {
	asteroids := format(s)
	detect := make(map[complex128]int)
	for _, ast := range asteroids {
		k := 0
		for _, otherAst := range asteroids {
			if isDetectable(ast, otherAst, asteroids) {
				k++
			}
		}
		detect[ast] = k
	}
	return maxMapValue(detect)
}

func part2(s string) int {
	asteroids := format(s)
	detect := make(map[complex128]int)
	for _, ast := range asteroids {
		k := 0
		for _, otherAst := range asteroids {
			if isDetectable(ast, otherAst, asteroids) {
				k++
			}
		}
		detect[ast] = k
	}
	c := 0
	var detectableByMain angleList
	mainAst := maxMapKey(detect)
	for c < 200 {
		for _, otherAst := range asteroids {
			if isDetectable(mainAst, otherAst, asteroids) {
				detectableByMain = append(detectableByMain, angle{otherAst, otherAst - mainAst})
				c++
			}
		}
	}
	sort.Sort(detectableByMain)
	res := detectableByMain[199]
	return 100*int(real(res.ast)) + int(imag(res.ast))
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day10/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2019/day10/input.txt")

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
