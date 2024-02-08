package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == ':' || r == ' ' || r == '"'
}

func format(s string) (map[int]string, []string) {
	max := 5
	res := make(map[int]string)
	parts := strings.Split(s, "\n\n")
	lines := strings.Split(parts[0], "\n")
	for len(res) != len(lines) {
		for _, line := range lines {
			lineSplit := strings.FieldsFunc(line, Split)
			n, _ := strconv.Atoi(lineSplit[0])
			_, done := res[n]
			if !done {
				if strings.ContainsAny(line, "ab") {
					res[n] = lineSplit[1]
				} else {
					if strings.Contains(line, "|") {
						if len(lineSplit) == 6 {
							a, _ := strconv.Atoi(lineSplit[1])
							b, _ := strconv.Atoi(lineSplit[2])
							c, _ := strconv.Atoi(lineSplit[4])
							d, _ := strconv.Atoi(lineSplit[5])
							_, okA := res[a]
							_, okB := res[b]
							_, okC := res[c]
							_, okD := res[d]
							if okA && okB && okC && okD {
								res[n] = "(" + res[a] + res[b] + "|" + res[c] + res[d] + ")"
							}
						}
						if len(lineSplit) == 4 {
							a, _ := strconv.Atoi(lineSplit[1])
							b, _ := strconv.Atoi(lineSplit[3])
							_, okA := res[a]
							_, okB := res[b]

							if okA && okB {
								res[n] = "(" + res[a] + "|" + res[b] + ")"
							}
						}
						if len(lineSplit) == 5 && n == 8 {
							a, _ := strconv.Atoi(lineSplit[1])
							_, okA := res[a]
							if okA {
								res[n] = "(" + res[a] + ")+"
							}
						}
						if len(lineSplit) == 7 && n == 11 {
							a, _ := strconv.Atoi(lineSplit[1])
							b, _ := strconv.Atoi(lineSplit[2])
							_, okA := res[a]
							_, okB := res[b]
							if okA && okB {
								res[n] = "("
								for i := 1; i < max; i++ {
									res[n] += res[a] + "{" + strconv.Itoa(i) + "}" + res[b] + "{" + strconv.Itoa(i) + "}" + "|"
								}
								res[n] += res[a] + "{" + strconv.Itoa(max) + "}" + res[b] + "{" + strconv.Itoa(max) + "}" + ")"
							}
						}
					}
					if len(lineSplit) == 2 {
						a, _ := strconv.Atoi(lineSplit[1])
						_, okA := res[a]
						if okA {
							res[n] = res[a]
						}
					}
					if len(lineSplit) == 3 {
						a, _ := strconv.Atoi(lineSplit[1])
						b, _ := strconv.Atoi(lineSplit[2])
						_, okA := res[a]
						_, okB := res[b]
						if okA && okB {
							res[n] = res[a] + res[b]
						}
					}
					if len(lineSplit) == 4 {
						a, _ := strconv.Atoi(lineSplit[1])
						b, _ := strconv.Atoi(lineSplit[2])
						c, _ := strconv.Atoi(lineSplit[3])
						_, okA := res[a]
						_, okB := res[b]
						_, okC := res[c]
						if okA && okB && okC {
							res[n] = res[a] + res[b] + res[c]
						}
					}
				}
			}
		}
	}
	return res, strings.Split(parts[1], "\n")
}

func part1(s string) int {
	c := 0
	rules, msgs := format(s)
	for _, msg := range msgs {
		if ok, _ := regexp.MatchString("^"+rules[0]+"$", msg); ok {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	rules, msgs := format(s)
	for _, msg := range msgs {
		if ok, _ := regexp.MatchString("^"+rules[0]+"$", msg); ok {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day19/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day19/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day19/input1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day19/input2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
