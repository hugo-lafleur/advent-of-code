package main

import (
	"fmt"
	"log"
	"math/bits"
	"os"
	"strconv"
	"strings"
	"time"
)

type image struct {
	n    int
	size int
}

func Split(r rune) bool {
	return r == '/' || r == ' '
}

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.FieldsFunc(line, Split))
	}
	return res
}

func gridToImage(g []string) image {
	b := ""
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			if rune(g[i][j]) == '.' || rune(g[i][j]) == '0' {
				b += "0"
			} else {
				b += "1"
			}
		}
	}
	n, _ := strconv.ParseInt(b, 2, 64)
	return image{int(n), len(g)}
}

func flips(grid []string) [][]string {
	if len(grid) == 2 {
		a, b, c, d := string(grid[0][0]), string(grid[0][1]), string(grid[1][0]), string(grid[1][1])
		return [][]string{grid, {b + a, d + c}}
	}
	a, b, c, d, e, f, g, h, i := string(grid[0][0]), string(grid[0][1]), string(grid[0][2]), string(grid[1][0]), string(grid[1][1]), string(grid[1][2]), string(grid[2][0]), string(grid[2][1]), string(grid[2][2])
	return [][]string{grid, {c + b + a, f + e + d, i + h + g}}
}

func rotates(grid []string) [][]string {
	res := [][]string{}
	res = append(res, grid)
	if len(grid) == 2 {
		a, b, c, d := string(grid[0][0]), string(grid[0][1]), string(grid[1][0]), string(grid[1][1])
		res = append(res, []string{c + a, d + b})
		res = append(res, []string{d + c, b + a})
		res = append(res, []string{b + d, a + c})
	} else {
		a, b, c, d, e, f, g, h, i := string(grid[0][0]), string(grid[0][1]), string(grid[0][2]), string(grid[1][0]), string(grid[1][1]), string(grid[1][2]), string(grid[2][0]), string(grid[2][1]), string(grid[2][2])
		res = append(res, []string{g + d + a, h + e + b, i + f + c})
		res = append(res, []string{i + h + g, f + e + d, c + b + a})
		res = append(res, []string{c + f + i, b + e + h, a + d + g})
	}
	return res
}

func imageToGrid(i image) []string {
	bin := strconv.FormatInt(int64(i.n), 2)
	if i.size == 2 {
		for len(bin) != 4 {
			bin = "0" + bin
		}
		return []string{bin[:2], bin[2:4]}
	}
	if i.size == 3 {
		for len(bin) != 9 {
			bin = "0" + bin
		}
		return []string{bin[:3], bin[3:6], bin[6:9]}
	}
	if i.size == 4 {
		for len(bin) != 16 {
			bin = "0" + bin
		}
		return []string{bin[:4], bin[4:8], bin[8:12], bin[12:16]}
	} else {
		for len(bin) != 36 {
			bin = "0" + bin
		}
		return []string{bin[:6], bin[6:12], bin[12:18], bin[18:24], bin[24:30], bin[30:36]}
	}

}

func partition(i image) []image {
	grid := imageToGrid(i)
	res := []image{}
	if i.size == 4 {
		for j := 0; j < 4; j++ {
			l1 := grid[(j/2)*2]
			l2 := grid[(j/2)*2+1]
			k := (j % 2) * 2
			res = append(res, gridToImage([]string{l1[k : k+2], l2[k : k+2]}))
		}
		return res
	}
	if i.size == 6 {
		for j := 0; j < 9; j++ {
			l1 := grid[(j/3)*2]
			l2 := grid[(j/3)*2+1]
			k := (j % 3) * 2
			res = append(res, gridToImage([]string{l1[k : k+2], l2[k : k+2]}))
		}
		return res

	}
	return []image{}
}

func fusion(l []image) image {
	g1 := imageToGrid(l[0])
	g2 := imageToGrid(l[1])
	g3 := imageToGrid(l[2])
	g4 := imageToGrid(l[3])
	return gridToImage([]string{g1[0] + g2[0], g1[1] + g2[1], g1[2] + g2[2], g3[0] + g4[0], g3[1] + g4[1], g3[2] + g4[2]})
}

func next(i image, rules map[image]image) []image {
	grid := imageToGrid(i)
	if len(grid) == 2 || len(grid) == 3 {
		for key, value := range rules {
			if i == key {
				return []image{value}
			}
		}
	}
	if len(grid) == 4 {
		p := partition(i)
		temp := []image{}
		for _, x := range p {
			temp = append(temp, next(x, rules)...)
		}
		fusion := fusion(temp)
		return []image{fusion}
	}
	if len(grid) == 6 {
		p := partition(i)
		res := []image{}
		for _, x := range p {
			res = append(res, next(x, rules)...)
		}
		return res
	}
	return []image{}
}

func part1(s string) int {
	c := 0
	input := format(s)
	rules := make(map[image]image)
	for _, line := range input {
		if len(line) == 6 {
			for _, grid := range flips(line[:2]) {
				for _, r := range rotates(grid) {
					rules[gridToImage(r)] = gridToImage(line[3:])
				}
			}
		}
		if len(line) == 8 {
			for _, grid := range flips(line[:3]) {
				for _, r := range rotates(grid) {
					rules[gridToImage(r)] = gridToImage(line[4:])
				}
			}
		}
	}
	curr := []image{{143, 3}}
	i := 0
	var limit int
	if len(input) == 2 {
		limit = 2
	} else {
		limit = 5
	}
	for i < limit {
		temp := []image{}
		for _, n := range curr {
			temp = append(temp, next(n, rules)...)
		}
		curr = temp
		i++
	}
	for _, x := range curr {
		c += bits.OnesCount(uint(x.n))
	}
	return c
}

func part2(s string) int {
	c := 0
	input := format(s)
	rules := make(map[image]image)
	for _, line := range input {
		if len(line) == 6 {
			for _, grid := range flips(line[:2]) {
				for _, r := range rotates(grid) {
					rules[gridToImage(r)] = gridToImage(line[3:])
				}
			}
		}
		if len(line) == 8 {
			for _, grid := range flips(line[:3]) {
				for _, r := range rotates(grid) {
					rules[gridToImage(r)] = gridToImage(line[4:])
				}
			}
		}
	}
	curr := make(map[image]int)
	curr[image{143, 3}] = 1
	i := 0
	limit := 18
	for i < limit {
		new := make(map[image]int)
		for key, value := range curr {
			for _, im := range next(key, rules) {
				new[im] += value
			}
		}
		curr = new
		i++
	}
	for key, value := range curr {
		c += (bits.OnesCount(uint(key.n)) * value)
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day21/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day21/input.data")

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
