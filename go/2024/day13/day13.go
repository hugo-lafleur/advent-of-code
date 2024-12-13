package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Button struct {
	X, Y int
}

type Prize struct {
	X, Y int
}

type Machine struct {
	buttonA, buttonB Button
	prize            Prize
}

func parse(s string) []Machine {
	var result []Machine
	var lines = strings.Split(s, "\n")
	for i := 0; i < len(lines); i += 4 {
		var aX, aY, bX, bY, prizeX, prizeY int
		fmt.Sscanf(lines[i+0], "Button A: X+%d, Y+%d", &aX, &aY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)
		result = append(result, Machine{buttonA: Button{aX, aY}, buttonB: Button{bX, bY}, prize: Prize{prizeX, prizeY}})
	}
	return result

}

func solve(m Machine) int {
	var b = (m.buttonA.Y*m.prize.X - m.buttonA.X*m.prize.Y) /
		(m.buttonA.Y*m.buttonB.X - m.buttonA.X*m.buttonB.Y)
	var a = (m.prize.X - m.buttonB.X*b) / m.buttonA.X
	var x = m.buttonA.X*a + m.buttonB.X*b
	var y = m.buttonA.Y*a + m.buttonB.Y*b
	if x == m.prize.X && y == m.prize.Y {
		return 3*a + b
	}
	return 0
}

func part1(s string) int {
	var machines = parse(s)
	var result int
	for _, m := range machines {
		result += solve(m)
	}
	return result
}

func part2(s string) int {
	var machines = parse(s)
	var result int
	for _, m := range machines {
		m.prize.X += 10000000000000
		m.prize.Y += 10000000000000
		result += solve(m)
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2024/day13/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2024/day13/input.txt")

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
