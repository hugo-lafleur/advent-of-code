package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type vector struct {
	x, y, z int
}

type particle struct {
	position     vector
	velocity     vector
	acceleration vector
}

func Split(r rune) bool {
	return r == '>' || r == '<' || r == ',' || r == ' '
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func get(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(p particle) int {
	return abs(p.position.x) + abs(p.position.y) + abs(p.position.z)
}

func collide(p1, p2 particle) bool {
	return p1.position.x == p2.position.x && p1.position.y == p2.position.y && p1.position.z == p2.position.z
}

func collision(particles []particle) []particle {
	res := []particle{}
loop:
	for _, p1 := range particles {
		c := 0
		for _, p2 := range particles {
			if collide(p1, p2) {
				c++
			}
			if c > 1 {
				continue loop
			}
		}
		res = append(res, p1)
	}
	return res
}

func part1(s string) int {
	list := format(s)
	particles := []particle{}
	for _, line := range list {
		p := vector{get(line[1]), get(line[2]), get(line[3])}
		v := vector{get(line[5]), get(line[6]), get(line[7])}
		a := vector{get(line[9]), get(line[10]), get(line[11])}
		particles = append(particles, particle{p, v, a})
	}
	tick := 0
	for tick < 500 {
		for i, particle := range particles {
			particle.velocity.x += particle.acceleration.x
			particle.velocity.y += particle.acceleration.y
			particle.velocity.z += particle.acceleration.z
			particle.position.x += particle.velocity.x
			particle.position.y += particle.velocity.y
			particle.position.z += particle.velocity.z
			particles[i] = particle
		}
		tick++
	}
	min := distance(particles[0])
	res := 0
	for i, particle := range particles {
		if distance(particle) < min {
			min = distance(particle)
			res = i
		}
	}
	return res
}

func part2(s string) int {
	list := format(s)
	particles := []particle{}
	for _, line := range list {
		p := vector{get(line[1]), get(line[2]), get(line[3])}
		v := vector{get(line[5]), get(line[6]), get(line[7])}
		a := vector{get(line[9]), get(line[10]), get(line[11])}
		particles = append(particles, particle{p, v, a})
	}
	tick := 0
	for tick < 500 {
		for i, particle := range particles {
			particle.velocity.x += particle.acceleration.x
			particle.velocity.y += particle.acceleration.y
			particle.velocity.z += particle.acceleration.z
			particle.position.x += particle.velocity.x
			particle.position.y += particle.velocity.y
			particle.position.z += particle.velocity.z
			particles[i] = particle
		}
		particles = collision(particles)
		tick++
	}
	return len(particles)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day20/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day20/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day20/input.data")

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
