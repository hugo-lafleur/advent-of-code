package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type ressource struct {
	n    int
	ress string
}

type formulas map[ressource][]ressource

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func format(s string) formulas {
	lines := strings.Split(s, "\n")
	res := make(formulas)
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		temp := []ressource{}
		for i := 0; i < len(lineSplit)-3; i += 2 {
			n, _ := strconv.Atoi(lineSplit[i])
			ress := lineSplit[i+1]
			r := ressource{n, ress}
			temp = append(temp, r)
		}
		n, _ := strconv.Atoi(lineSplit[len(lineSplit)-2])
		r := lineSplit[len(lineSplit)-1]
		res[ressource{n, r}] = temp
	}
	return res
}

func calculRessources(n int, ress string, f formulas, excess map[string]int) int {
	res := 0
	var d int
	_, ok := excess[ress]
	if ok {
		for excess[ress] > 0 && n > 0 {
			n--
			excess[ress]--
		}
	}
	if ress == "ORE" {
		return n
	}
	var l []ressource
	for key, value := range f {
		if key.ress == ress {
			l = value
			d = key.n
		}
	}

	for _, x := range l {
		for n%d != 0 {
			n++
			excess[ress]++
		}
		res += calculRessources((n*x.n)/d, x.ress, f, excess)

	}
	return res
}

func part1(s string) int {
	f := format(s)
	return calculRessources(1, "FUEL", f, make(map[string]int))
}

func part2(s string) int {
	f := format(s)
	a := 0
	b := 1000000000000
	for {
		c := (a + b) / 2
		if calculRessources(c, "FUEL", f, make(map[string]int)) > 1000000000000 {
			b = c
		} else {
			a = c
		}
		if b-a < 2 {
			return a
		}
	}
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day14/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2019/day14/input.txt")

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
