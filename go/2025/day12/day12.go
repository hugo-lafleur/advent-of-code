package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) ([]int, [][]int) {
	lines := strings.Split(s, "\n")
	shapes := []int{}
	regions := [][]int{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if !strings.Contains(line, "x") {
			cnt := 0
			for j := 1; j <= 3; j++ {
				cnt += strings.Count(lines[i+j], "#")
			}
			shapes = append(shapes, cnt)
			i += 4
		} else {
			parts := strings.FieldsFunc(line, func(r rune) bool {
				return r == 'x' || r == ':' || r == ' '
			})
			region := []int{}
			for _, p := range parts {
				n, _ := strconv.Atoi(p)
				region = append(region, n)
			}
			regions = append(regions, region)
		}
	}
	return shapes, regions
}

func part1(s string) int {
	shapes, regions := format(s)
	result := 0
	for j, region := range regions {
		area := region[0] * region[1]
		sum := 0
		tiles := 0
		for i := 2; i < len(region); i++ {
			sum += region[i]
			tiles += region[i] * shapes[i-2]
		}
		if area >= 9*sum {
			result++
		} else if area < tiles {
			continue
		} else {
			if j == 0 && len(regions) == 3 {
				result++
			}
		}
	}
	return result
}

func main() {
	content, err := os.ReadFile("../../../inputs/2025/day12/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2025/day12/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
