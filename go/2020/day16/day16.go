package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type interval struct {
	a, b int
}

type doubleInterval struct {
	first, second interval
}

type ticket []int

func Split(r rune) bool {
	return r == ' ' || r == '-'
}

func format(s string) ([]doubleInterval, ticket, []ticket) {
	listIntervals := []doubleInterval{}
	var personalTicket ticket
	nearbyTickets := []ticket{}
	lines := strings.Split(s, "\n")
	start := 0
	for i := start; i < len(lines); i++ {
		line := lines[i]
		if len(line) < 2 {
			start = i + 2
			break
		}
		linesSplit := strings.FieldsFunc(line, Split)
		var a, b, c, d int
		if strings.ContainsAny(linesSplit[1], ":") {
			a, _ = strconv.Atoi(linesSplit[2])
			b, _ = strconv.Atoi(linesSplit[3])
			c, _ = strconv.Atoi(linesSplit[5])
			d, _ = strconv.Atoi(linesSplit[6])
		} else {
			a, _ = strconv.Atoi(linesSplit[1])
			b, _ = strconv.Atoi(linesSplit[2])
			c, _ = strconv.Atoi(linesSplit[4])
			d, _ = strconv.Atoi(linesSplit[5])
		}
		listIntervals = append(listIntervals, doubleInterval{interval{a, b}, interval{c, d}})
	}
	for i := start; i < len(lines); i++ {
		line := lines[i]
		if len(line) < 2 {
			start = i + 2
			break
		}
		lineSplit := strings.Split(line, ",")
		for _, s := range lineSplit {
			n, _ := strconv.Atoi(s)
			personalTicket = append(personalTicket, n)
		}
	}
	for i := start; i < len(lines); i++ {
		line := lines[i]
		var currTicket ticket
		lineSplit := strings.Split(line, ",")
		for _, s := range lineSplit {
			n, _ := strconv.Atoi(s)
			currTicket = append(currTicket, n)
		}
		nearbyTickets = append(nearbyTickets, currTicket)
	}
	return listIntervals, personalTicket, nearbyTickets

}

func isInInterval(n int, i interval) bool {
	return n >= i.a && n <= i.b
}

func isInIntervals(n int, l []doubleInterval) bool {
	for _, inter := range l {
		if isInInterval(n, inter.first) || isInInterval(n, inter.second) {
			return true
		}
	}
	return false
}

func try(i, j int, intervals []doubleInterval, tickets []ticket) bool {
	for _, ticket := range tickets {
		if !(isInInterval(ticket[j], intervals[i].first) || isInInterval(ticket[j], intervals[i].second)) {
			return false
		}
	}
	return true
}

func removeElement(n int, l []int) []int {
	res := []int{}
	for _, x := range l {
		if x != n {
			res = append(res, x)
		}
	}
	return res
}

func solve(m map[int][]int) {
	for i := 0; i < len(m); i++ {
		for key, value := range m {
			if len(value) == 1 {
				for otherKey, otherValue := range m {
					if key != otherKey {
						m[otherKey] = removeElement(value[0], otherValue)
					}
				}
			}
		}
	}
}

func part1(s string) int {
	c := 0
	intervals, _, nearbyTickets := format(s)
	for _, tick := range nearbyTickets {
		for _, n := range tick {
			if !isInIntervals(n, intervals) {
				c += n
			}
		}
	}
	return c
}

func part2(s string) int {
	intervals, myTicket, nearbyTickets := format(s)
	validTickets := []ticket{}
nextTicket:
	for _, tick := range nearbyTickets {
		for _, n := range tick {
			if !isInIntervals(n, intervals) {
				continue nextTicket
			}
		}
		validTickets = append(validTickets, tick)
	}
	possibilites := make(map[int][]int)
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals); j++ {
			if try(i, j, intervals, validTickets) {
				possibilites[i] = append(possibilites[i], j)
			}
		}
	}
	solve(possibilites)
	if len(intervals) == 3 {
		res := 1
		for i := 0; i < 3; i++ {
			res *= myTicket[possibilites[i][0]]
		}
		return res
	}
	res := 1
	for i := 0; i < 6; i++ {
		res *= myTicket[possibilites[i][0]]
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day16/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day16/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day16/input.txt")

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
