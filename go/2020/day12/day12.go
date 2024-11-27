package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type instruction struct {
	name  string
	value int
}

type object struct {
	position  complex128
	direction complex128
}

func angleToComplex(value int) complex128 {
	switch value {
	case 90:
		return 1i
	case 180:
		return -1
	case 270:
		return -1i
	}
	return 0
}

func format(s string) []instruction {
	lines := strings.Split(s, "\n")
	res := []instruction{}
	for _, line := range lines {
		name := line[0:1]
		n, _ := strconv.Atoi(line[1:])
		res = append(res, instruction{name: name, value: n})
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(s string) int {
	list := format(s)
	ship := object{complex(0, 0), 1}
	for _, instruction := range list {
		switch instruction.name {
		case "N":
			ship.position = ship.position + complex(0, float64(instruction.value))
		case "S":
			ship.position = ship.position + complex(0, -float64(instruction.value))
		case "E":
			ship.position = ship.position + complex(float64(instruction.value), 0)
		case "W":
			ship.position = ship.position + complex(-float64(instruction.value), 0)
		case "L":
			ship.direction = ship.direction * angleToComplex(instruction.value)
		case "R":
			ship.direction = ship.direction * (complex(real(angleToComplex(instruction.value)), -imag(angleToComplex(instruction.value))))
		case "F":
			ship.position = ship.position + complex(float64(instruction.value), 0)*ship.direction
		}
	}
	return abs(int(real(ship.position))) + abs(int(imag(ship.position)))
}

func part2(s string) int {
	list := format(s)
	ship := object{complex(0, 0), 1}
	waypoint := object{complex(10, 1), 1}
	for _, instruction := range list {
		switch instruction.name {
		case "N":
			waypoint.position = waypoint.position + complex(0, float64(instruction.value))
		case "S":
			waypoint.position = waypoint.position + complex(0, -float64(instruction.value))
		case "E":
			waypoint.position = waypoint.position + complex(float64(instruction.value), 0)
		case "W":
			waypoint.position = waypoint.position + complex(-float64(instruction.value), 0)
		case "L":
			waypoint.position = waypoint.position * angleToComplex(instruction.value)
		case "R":
			waypoint.position = waypoint.position * (complex(real(angleToComplex(instruction.value)), -imag(angleToComplex(instruction.value))))
		case "F":
			ship.position = ship.position + complex(float64(instruction.value), 0)*waypoint.position
		}
	}
	return abs(int(real(ship.position))) + abs(int(imag(ship.position)))
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day12/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day12/input.txt")

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
