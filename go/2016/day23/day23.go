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
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func part1(s string) int {
	instr := format(s)
	registers := make(map[string]int)
	for _, str := range []string{"a", "b", "c", "d"} {
		registers[str] = 0
	}
	registers["a"] = 7
	for i := 0; i < len(instr); i++ {
		line := instr[i]
		switch line[0] {
		case "cpy":
			_, notOk := strconv.Atoi(line[2])
			if notOk == nil {
				continue
			}
			n, err := strconv.Atoi(line[1])
			if err == nil {
				registers[line[2]] = n
			} else {
				registers[line[2]] = registers[line[1]]
			}
		case "inc":
			registers[line[1]]++
		case "dec":
			registers[line[1]]--
		case "jnz":
			c, err := strconv.Atoi(line[1])
			if registers[line[1]] != 0 || (err == nil && c != 0) {
				n, isInt := strconv.Atoi(line[2])
				if isInt == nil {
					i += n - 1
				} else {
					i += registers[line[2]] - 1
				}

			}
		case "tgl":
			n := registers[line[1]]
			j := i + n
			if j < len(instr) {
				lineToReplace := instr[j]
				switch lineToReplace[0] {
				case "cpy":
					instr[j][0] = "jnz"
				case "jnz":
					instr[j][0] = "cpy"
				case "inc":
					instr[j][0] = "dec"
				case "dec":
					instr[j][0] = "inc"
				case "tgl":
					instr[j][0] = "inc"
				}
			}
		}

	}
	return registers["a"]
}

func part2(s string) int {
	instr := format(s)
	registers := make(map[string]int)
	for _, str := range []string{"a", "b", "c", "d"} {
		registers[str] = 0
	}
	registers["a"] = 12
	for i := 0; i < len(instr); i++ {
		line := instr[i]
		switch line[0] {
		case "cpy":
			if i+5 < len(instr) && instr[i+1][0] == "inc" && instr[i+2][0] == "dec" && instr[i+3][0] == "jnz" && instr[i+4][0] == "dec" && instr[i+5][0] == "jnz" {
				if instr[i][2] == instr[i+2][1] && instr[i][2] == instr[i+3][1] {
					if instr[i+4][1] == instr[i+5][1] {
						if instr[i+3][2] == "-2" && instr[i+5][2] == "-5" {
							n, isInt := strconv.Atoi(instr[i][1])
							if isInt == nil {
								registers[instr[i+1][1]] += n * registers[instr[i+4][1]]
								registers[instr[i][2]] = 0
								registers[instr[i+4][1]] = 0
								i += 5
								continue
							} else {
								registers[instr[i+1][1]] += registers[instr[i][1]] * registers[instr[i+4][1]]
								registers[instr[i][2]] = 0
								registers[instr[i+4][1]] = 0
								i += 5
								continue
							}
						}
					}
				}
			}
			_, notOk := strconv.Atoi(line[2])
			if notOk == nil {
				continue
			}
			n, err := strconv.Atoi(line[1])
			if err == nil {
				registers[line[2]] = n
			} else {
				registers[line[2]] = registers[line[1]]
			}
		case "inc":
			registers[line[1]]++
		case "dec":
			registers[line[1]]--
		case "jnz":
			c, err := strconv.Atoi(line[1])
			if registers[line[1]] != 0 || (err == nil && c != 0) {
				n, isInt := strconv.Atoi(line[2])
				if isInt == nil {
					i += n - 1
				} else {
					i += registers[line[2]] - 1
				}

			}
		case "tgl":
			n := registers[line[1]]
			j := i + n
			if j < len(instr) {
				lineToReplace := instr[j]
				switch lineToReplace[0] {
				case "cpy":
					instr[j][0] = "jnz"
				case "jnz":
					instr[j][0] = "cpy"
				case "inc":
					instr[j][0] = "dec"
				case "dec":
					instr[j][0] = "inc"
				case "tgl":
					instr[j][0] = "inc"
				}
			}
		}

	}
	return registers["a"]
}

func main() {
	content, err := os.ReadFile("../../../inputs/2016/day23/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2016/day23/input.data")

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
