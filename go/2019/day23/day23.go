package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) map[int]int {
	res := make(map[int]int)
	lines := strings.Split(s, ",")
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		res[i] = n
	}
	return res
}

func opcodeSolve(n int) (int, int, int, int) {
	s := strconv.Itoa(n)
	for len(s) != 5 {
		s = "0" + s
	}
	var a, b, c, de int
	a, _ = strconv.Atoi(string(s[0]))
	b, _ = strconv.Atoi(string(s[1]))
	c, _ = strconv.Atoi(string(s[2]))
	de, _ = strconv.Atoi(s[3:])
	return de, c, b, a
}

func solve(s string, input, output chan int) {
	p := format(s)
	i := 0
	relativeBase := 0
mainLoop:
	for {
		opcode, mode1, mode2, mode3 := opcodeSolve(p[i])
		var a, b, c, res int
		if opcode == 99 {
			break
		}
		switch mode1 {
		case 0:
			a = p[p[i+1]]
		case 1:
			a = p[i+1]
		case 2:
			a = p[relativeBase+p[i+1]]
		}
		switch mode2 {
		case 0:
			b = p[p[i+2]]
		case 1:
			b = p[i+2]
		case 2:
			b = p[relativeBase+p[i+2]]
		}
		switch mode3 {
		case 0:
			c = p[i+3]
		case 2:
			c = p[i+3] + relativeBase
		}
		switch opcode {
		case 1:
			res = a + b
			i += 3
		case 2:
			res = a * b
			i += 3
		case 3:
			temp := <-input
			switch mode1 {
			case 0:
				p[p[i+1]] = temp
			case 2:
				p[relativeBase+p[i+1]] = temp
			}
			i += 2
			continue mainLoop
		case 4:
			output <- a
			i += 2
			continue mainLoop
		case 5:
			if a != 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 6:
			if a == 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 7:
			if a < b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 8:
			if a == b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 9:
			relativeBase += a
			i += 2
			continue mainLoop
		}
		p[c] = res
		i++
	}
}

func isIdle(m map[int]int) bool {
	for _, x := range m {
		if x < 1 {
			return false
		}
	}
	return true
}

func part1(s string) int {
	inputs := []chan int{}
	outputs := []chan int{}
	res := make(chan int)
	for i := 0; i < 50; i++ {
		inputs = append(inputs, make(chan int))
		outputs = append(outputs, make(chan int))
		go solve(s, inputs[i], outputs[i])
		inputs[i] <- i
	}
	go func() {
		for {
			for i := 0; i < 50; i++ {
				select {
				case address := <-outputs[i]:
					X := <-outputs[i]
					Y := <-outputs[i]
					if address == 255 {
						res <- Y
						return
					}
					inputs[address] <- X
					inputs[address] <- Y
					i = address - 1
				case inputs[i] <- -1:
				}
			}
		}
	}()
	return <-res
}

func part2(s string) int {
	inputs := []chan int{}
	outputs := []chan int{}
	NAT := make(chan int, 2)
	res := make(chan int)
	var mem int
	idle := make(map[int]int)
	for i := 0; i < 50; i++ {
		inputs = append(inputs, make(chan int))
		outputs = append(outputs, make(chan int))
		go solve(s, inputs[i], outputs[i])
		inputs[i] <- i
		idle[i] = -1
	}
	go func() {
		for {
			for i := 0; i < 50; i++ {
				select {
				case address := <-outputs[i]:
					X := <-outputs[i]
					Y := <-outputs[i]
					if address == 255 {
						if len(NAT) == 2 {
							<-NAT
							<-NAT
						}
						NAT <- X
						NAT <- Y
						continue
					}
					inputs[address] <- X
					inputs[address] <- Y
					idle[address] = 0
					i = address - 1
				case inputs[i] <- -1:
					idle[i]++
				}
			}
			if isIdle(idle) {
				X, Y := <-NAT, <-NAT
				if Y == mem {
					res <- mem
				} else {
					mem = Y
				}
				inputs[0] <- X
				inputs[0] <- Y
				idle[0] = 0
			}
		}
	}()
	return <-res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day23/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
