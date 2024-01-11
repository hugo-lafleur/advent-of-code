package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type flipFlops map[string]string

type conjuctions map[string]flipFlops

type queue []string

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *queue) Push(s string) {
	*q = append(*q, s)
}

func (q *queue) Pop() (string, bool) {
	if q.IsEmpty() {
		return "", false
	}
	s := (*q)[0]
	*q = (*q)[1:]
	return s, true
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func isIn(s string, l []string) bool {
	for _, x := range l {
		if x == s {
			return true
		}
	}
	return false
}

func isDone(toRX map[string]int) bool {
	for _, value := range toRX {
		if value == 0 {
			return false
		}
	}
	return true
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

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func part1(s string) int {
	config := format(s)
	flipFlop := flipFlops(make(map[string]string))
	conjuction := conjuctions(make(map[string]flipFlops))
	nature := make(map[string]string)
	lows := 0
	highs := 0
	for _, line := range config {
		if line[0][0] == '%' {
			flipFlop[line[0][1:]] = "off"
			nature[line[0][1:]] = "%"
		}
		if line[0][0] == '&' {
			configConjuction := flipFlops(make(map[string]string))
			conj := line[0][1:]
			for _, line2 := range config {
				if isIn(conj, line2) {
					configConjuction[line2[0][1:]] = "low"
				}
			}
			conjuction[conj] = configConjuction
			nature[line[0][1:]] = "&"
		}
	}
	var queue queue
	for k := 0; k < 1000; k++ {
		lows++
		for _, line := range config {
			if line[0] == "broadcaster" {
				for i := 2; i < len(line); i++ {
					next := line[i]
					if flipFlop[next] == "off" {
						flipFlop[next] = "on"
					} else {
						flipFlop[next] = "off"
					}
					lows++
					queue.Push(next)
				}
			}
		}
		for !queue.IsEmpty() {
			curr, _ := queue.Pop()
			for _, line := range config {
				if line[0][1:] == curr {
					for i := 2; i < len(line); i++ {
						next := line[i]
						if nature[curr] == "%" && nature[next] == "%" {
							if flipFlop[curr] == "off" {
								if flipFlop[next] == "off" {
									flipFlop[next] = "on"
									queue.Push(next)
								} else {
									flipFlop[next] = "off"
									queue.Push(next)
								}
								lows++
							} else {
								highs++
							}
						}
						if nature[curr] == "%" && nature[next] == "&" {
							if flipFlop[curr] != "on" {
								conjuction[next][curr] = "low"
								lows++
							} else {
								conjuction[next][curr] = "high"
								highs++
							}
							queue.Push(next)
						}
						if nature[curr] == "&" && nature[next] == "%" {
							high := true
							for _, value := range conjuction[curr] {
								if value != "high" {
									high = false
								}
							}
							if high {
								if flipFlop[next] == "on" {
									flipFlop[next] = "off"
									queue.Push(next)
								} else {
									flipFlop[next] = "on"
									queue.Push(next)
								}
								lows++
							} else {
								highs++
							}
						}
						if nature[curr] == "&" && nature[next] == "&" {
							high := true
							for _, value := range conjuction[curr] {
								if value != "high" {
									high = false
								}
							}
							if high {
								conjuction[next][curr] = "low"
								lows++
							} else {
								conjuction[next][curr] = "high"
								highs++
							}
							queue.Push(next)
						}
						if nature[curr] == "&" && (next == "output" || next == "rx") {
							high := true
							for _, value := range conjuction[curr] {
								if value != "high" {
									high = false
								}
							}
							if high {
								lows++
							} else {
								highs++
							}
						}
					}
				}
			}
		}
	}
	return lows * highs
}

func part2(s string) int {
	config := format(s)
	flipFlop := flipFlops(make(map[string]string))
	conjuction := conjuctions(make(map[string]flipFlops))
	nature := make(map[string]string)
	sendToRX := ""
	toRX := make(map[string]int)
	for _, line := range config {
		if line[0][0] == '%' {
			flipFlop[line[0][1:]] = "off"
			nature[line[0][1:]] = "%"
		}
		if line[0][0] == '&' {
			configConjuction := flipFlops(make(map[string]string))
			conj := line[0][1:]
			for _, line2 := range config {
				if isIn(conj, line2) {
					configConjuction[line2[0][1:]] = "low"
				}
			}
			conjuction[conj] = configConjuction
			nature[line[0][1:]] = "&"
			if line[2] == "rx" {
				sendToRX = conj
				for key := range configConjuction {
					toRX[key] = 0
				}
			}
		}
	}
	var queue queue
	for k := 0; k >= 0; k++ {
		if isDone(toRX) {
			res := []int{}
			for _, value := range toRX {
				res = append(res, value)
			}
			return LCM(1, 1, res...)
		}
		for _, line := range config {
			if line[0] == "broadcaster" {
				for i := 2; i < len(line); i++ {
					next := line[i]
					if flipFlop[next] == "off" {
						flipFlop[next] = "on"
					} else {
						flipFlop[next] = "off"
					}
					queue.Push(next)
				}
			}
		}
		for !queue.IsEmpty() {
			curr, _ := queue.Pop()
			for _, line := range config {
				if line[0][1:] == curr {
					for i := 2; i < len(line); i++ {
						next := line[i]
						if nature[curr] == "%" && nature[next] == "%" {
							if flipFlop[curr] == "off" {
								if flipFlop[next] == "off" {
									flipFlop[next] = "on"
									queue.Push(next)
								} else {
									flipFlop[next] = "off"
									queue.Push(next)
								}

							}
						}
						if nature[curr] == "%" && nature[next] == "&" {
							if flipFlop[curr] == "off" {
								conjuction[next][curr] = "low"
							} else {
								conjuction[next][curr] = "high"
							}
							queue.Push(next)
						}
						if nature[curr] == "&" && nature[next] == "%" {
							high := true
							for _, value := range conjuction[curr] {
								if value != "high" {
									high = false
								}
							}
							if high {
								if flipFlop[next] == "on" {
									flipFlop[next] = "off"
									queue.Push(next)
								} else {
									flipFlop[next] = "on"
									queue.Push(next)
								}
							}
						}
						if nature[curr] == "&" && nature[next] == "&" {
							high := true
							for _, value := range conjuction[curr] {
								if value != "high" {
									high = false
								}
							}
							if high {
								conjuction[next][curr] = "low"
							} else {
								conjuction[next][curr] = "high"
							}
							queue.Push(next)
						}
						if curr == sendToRX && next == "rx" {
							for key, value := range toRX {
								if value == 0 {
									if conjuction[sendToRX][key] == "high" {
										toRX[key] = k + 1
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day20/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2023/day20/input.data")

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
