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
	instrs := format(s)
	registers := make(map[string]int)
	var sound int
	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]
		switch instr[0] {
		case "snd":
			sound = registers[instr[1]]
		case "set":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] = n
			} else {
				registers[instr[1]] = registers[instr[2]]
			}
		case "add":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] += n
			} else {
				registers[instr[1]] += registers[instr[2]]
			}
		case "mul":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] *= n
			} else {
				registers[instr[1]] *= registers[instr[2]]
			}
		case "mod":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] = registers[instr[1]] % n
			} else {
				registers[instr[1]] = registers[instr[1]] % registers[instr[2]]
			}
		case "rcv":
			if registers[instr[1]] != 0 {
				return sound
			}
		case "jgz":
			n, err1 := strconv.Atoi(instr[1])
			m, err2 := strconv.Atoi(instr[2])
			if err1 != nil {
				n = registers[instr[1]]
			}
			if err2 != nil {
				m = registers[instr[2]]
			}
			if n > 0 {
				i = i + m - 1
			}
		}
	}
	return 0
}

func part2aux(s string, id int, out, in, res chan int) {
	instrs := format(s)
	registers := make(map[string]int)
	registers["p"] = id
	send := 0
	for i := 0; i < len(instrs); i++ {
		instr := instrs[i]
		switch instr[0] {
		case "snd":
			n, err := strconv.Atoi(instr[1])
			send++
			if err == nil {
				out <- n
			} else {
				out <- registers[instr[1]]
			}
		case "set":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] = n
			} else {
				registers[instr[1]] = registers[instr[2]]
			}
		case "add":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] += n
			} else {
				registers[instr[1]] += registers[instr[2]]
			}
		case "mul":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] *= n
			} else {
				registers[instr[1]] *= registers[instr[2]]
			}
		case "mod":
			n, err := strconv.Atoi(instr[2])
			if err == nil {
				registers[instr[1]] = registers[instr[1]] % n
			} else {
				registers[instr[1]] = registers[instr[1]] % registers[instr[2]]
			}
		case "rcv":
			select {
			case received := <-in:
				registers[instr[1]] = received
			case <-time.After(1 * time.Second):
				if id == 1 {
					res <- send
					return
				}
			}

		case "jgz":
			n, err1 := strconv.Atoi(instr[1])
			m, err2 := strconv.Atoi(instr[2])
			if err1 != nil {
				n = registers[instr[1]]
			}
			if err2 != nil {
				m = registers[instr[2]]
			}
			if n > 0 {
				i = i + m - 1
			}
		}
	}
	if id == 1 {
		res <- send
	}
}

func part2(s string) int {
	ch0 := make(chan int, 1000)
	ch1 := make(chan int, 1000)
	res := make(chan int)

	go part2aux(s, 0, ch0, ch1, res)
	go part2aux(s, 1, ch1, ch0, res)

	time.Sleep(2 * time.Second)
	return <-res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day18/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day18/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day18/input.data")

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
