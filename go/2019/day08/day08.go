package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) []string {
	return strings.Split(s, "")
}

func createMap(l [][]string) []int {
	res := make([]int, 3)
	for _, line := range l {
		for _, s := range line {
			n, _ := strconv.Atoi(s)
			if n < 3 {
				res[n]++
			}
		}
	}
	return res
}

func minZero(m map[int][]int) int {
	var min, index int
	for key, value := range m {
		min = value[0]
		index = key
	}
	for key, value := range m {
		if value[0] < min {
			min = value[0]
			index = key
		}
	}
	return index
}

func part1(s string) int {
	input := format(s)
	var n, m int
	if len(input) == 12 {
		n, m = 3, 2
	} else {
		n, m = 25, 6
	}
	layers := [][][]string{}
	for l := 0; l < len(input); l += n * m {
		layer := [][]string{}
		for r := l; r < l+n*m; r += n {
			row := []string{}
			for s := r; s < r+n; s++ {
				row = append(row, input[s])
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}
	mapOfLayers := make(map[int][]int)
	for i, layer := range layers {
		mapOfLayers[i] = createMap(layer)
	}
	index := minZero(mapOfLayers)
	return mapOfLayers[index][1] * mapOfLayers[index][2]
}

func part2(s string) string {
	input := format(s)
	var n, m int
	if len(input) == 16 {
		n, m = 2, 2
	} else {
		n, m = 25, 6
	}
	layers := [][][]string{}
	for l := 0; l < len(input); l += n * m {
		layer := [][]string{}
		for r := l; r < l+n*m; r += n {
			row := []string{}
			for s := r; s < r+n; s++ {
				row = append(row, input[s])
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}
	res := [][]string{}
	for i := 0; i < m; i++ {
		res = append(res, make([]string, n))
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			l := 0
			for layers[l][i][j] == "2" {
				l++
			}
			res[i][j] = layers[l][i][j]
		}
	}
	toPrint := "\n"
	for _, line := range res {
		for _, s := range line {
			switch s {
			case "1":
				toPrint += "#"
			case "0":
				toPrint += "."
			}
			toPrint += " "
		}
		toPrint += "\n"
	}
	return toPrint
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day08/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))

	content, err = os.ReadFile("../../../inputs/2019/day08/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day08/input.txt")

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
