package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	tab := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		tab = append(tab, strings.Split(line, " "))
	}
	return tab

}

func part1(s string) int {
	tab := format(s)
	signal := map[string]int{}
	i := 0
	j := 0
	n := 65536
	on := map[string]bool{}
	for j < n {
		on[strconv.Itoa(j)] = true
		signal[strconv.Itoa(j)] = j
		j++
	}
	for !on["a"] {
		for _, cmd := range tab {
			if cmd[0] == "NOT" && on[cmd[1]] {
				signal[cmd[3]] = 65535 ^ signal[cmd[1]]
				on[cmd[3]] = true
			}
			if cmd[1] == "AND" && on[cmd[0]] && on[cmd[2]] {
				signal[cmd[4]] = signal[cmd[0]] & signal[cmd[2]]
				on[cmd[4]] = true
			}
			if cmd[1] == "OR" && on[cmd[0]] && on[cmd[2]] {
				signal[cmd[4]] = signal[cmd[0]] | signal[cmd[2]]
				on[cmd[4]] = true
			}
			if cmd[1] == "LSHIFT" && on[cmd[0]] {
				n, _ := strconv.Atoi(cmd[2])
				signal[cmd[4]] = signal[cmd[0]] << n
				on[cmd[4]] = true
			}
			if cmd[1] == "RSHIFT" && on[cmd[0]] {
				n, _ := strconv.Atoi(cmd[2])
				signal[cmd[4]] = signal[cmd[0]] >> n
				on[cmd[4]] = true
			}
			if cmd[1] == "->" {
				n, _ := strconv.Atoi(cmd[0])
				if n != 0 || cmd[0] == "0" {
					signal[cmd[2]] = n
					on[cmd[2]] = true
				} else {
					if on[cmd[0]] {
						signal[cmd[2]] = signal[cmd[0]]
						on[cmd[2]] = true
					}
				}
			}
		}
		i++
	}
	return signal["a"]
}

func part2(s string) int {
	tab := format(s)
	signal := map[string]int{}
	i := 0
	j := 0
	n := 1674
	on := map[string]bool{}
	for j < n {
		on[strconv.Itoa(j)] = true
		signal[strconv.Itoa(j)] = j
		j++
	}
	a := part1(s)
	for !on["a"] {
		for _, cmd := range tab {
			if cmd[0] == "NOT" && on[cmd[1]] {
				signal[cmd[3]] = 65535 ^ signal[cmd[1]]
				on[cmd[3]] = true
			}
			if cmd[1] == "AND" && on[cmd[0]] && on[cmd[2]] {
				signal[cmd[4]] = signal[cmd[0]] & signal[cmd[2]]
				on[cmd[4]] = true
			}
			if cmd[1] == "OR" && on[cmd[0]] && on[cmd[2]] {
				signal[cmd[4]] = signal[cmd[0]] | signal[cmd[2]]
				on[cmd[4]] = true
			}
			if cmd[1] == "LSHIFT" && on[cmd[0]] {
				n, _ := strconv.Atoi(cmd[2])
				signal[cmd[4]] = signal[cmd[0]] << n
				on[cmd[4]] = true
			}
			if cmd[1] == "RSHIFT" && on[cmd[0]] {
				n, _ := strconv.Atoi(cmd[2])
				signal[cmd[4]] = signal[cmd[0]] >> n
				on[cmd[4]] = true
			}
			if cmd[1] == "->" {
				n, _ := strconv.Atoi(cmd[0])
				if n != 0 || cmd[0] == "0" {
					signal[cmd[2]] = n
					on[cmd[2]] = true
				} else {
					if on[cmd[0]] {
						signal[cmd[2]] = signal[cmd[0]]
						on[cmd[2]] = true
					}
				}
				if cmd[2] == "b" {
					signal["b"] = a
					on["b"] = true
				}
			}
		}
		i++
	}
	return signal["a"]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2015/day07/test.data")

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

	content, err = os.ReadFile("../../../inputs/2015/day07/input.data")

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
