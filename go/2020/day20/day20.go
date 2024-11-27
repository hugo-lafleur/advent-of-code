package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type tile struct {
	id  int
	tab [][]string
}

type couple struct {
	a, b int
}

func Split(r rune) bool {
	return r == ' ' || r == ':'
}

func format(s string) []tile {
	res := []tile{}
	tiles := strings.Split(s, "\n\n")
	for _, t := range tiles {
		lines := strings.Split(t, "\n")
		id, _ := strconv.Atoi(strings.FieldsFunc(lines[0], Split)[1])
		tab := [][]string{}
		for _, line := range lines[1:] {
			tab = append(tab, strings.Split(line, ""))
		}
		res = append(res, tile{id, tab})
	}
	return res
}

func rotateTab(tab [][]string) [][]string {
	matrix := [][]string{}
	for i := 0; i < len(tab); i++ {
		matrix = append(matrix, make([]string, len(tab[i])))
		for j := 0; j < len(tab); j++ {
			matrix[i][j] = tab[i][j]
		}
	}
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}

func flip(tab [][]string) [][]string {
	matrix := [][]string{}
	for i := 0; i < len(tab); i++ {
		matrix = append(matrix, make([]string, len(tab[i])))
		for j := 0; j < len(tab); j++ {
			matrix[i][j] = tab[len(tab)-1-i][j]
		}
	}
	return matrix
}

func edges(tab [][]string) [][]string {
	var up, left, down, right []string
	l := len(tab)
	for i := 0; i < len(tab); i++ {
		up = append(up, tab[0][i])
		down = append(down, tab[l-1][i])
		left = append(left, tab[i][0])
		right = append(right, tab[i][l-1])
	}
	res := [][]string{}
	res = append(res, up, right, down, left)
	return res
}

func sameEdge(edge1, edge2 []string) bool {
	for i := 0; i < len(edge1); i++ {
		if edge1[i] != edge2[i] {
			return false
		}
	}
	return true
}

func sameEdgeReverse(edge1, edge2 []string) bool {
	for i := 0; i < len(edge1); i++ {
		if edge1[i] != edge2[len(edge2)-1-i] {
			return false
		}
	}
	return true
}

func reverse(edge []string) []string {
	res := []string{}
	for i := len(edge) - 1; i >= 0; i-- {
		res = append(res, edge[i])
	}
	return res
}

func firstTab(tab [][]string, paths map[string]couple) [][]string {
	startTab := tab
	_, ok1 := paths[strings.Join(edges(startTab)[1], "")]
	_, ok3 := paths[strings.Join(edges(startTab)[2], "")]
	for !(ok1 && ok3) {
		startTab = flip(startTab)
		_, ok1 = paths[strings.Join(edges(startTab)[1], "")]
		_, ok3 = paths[strings.Join(edges(startTab)[2], "")]
		if ok1 && ok3 {
			break
		}
		startTab = flip(startTab)
		startTab = rotateTab(startTab)
		_, ok1 = paths[strings.Join(edges(startTab)[1], "")]
		_, ok3 = paths[strings.Join(edges(startTab)[2], "")]
	}
	return startTab
}

func nextTab(edge []string, paths map[string]couple, curr int, side int, tabs map[int][][]string) (int, [][]string) {
	pass := paths[strings.Join(edge, "")]
	var nextId int
	if pass.a == curr {
		nextId = pass.b
	} else {
		nextId = pass.a
	}
	nextTab := tabs[nextId]
	for !sameEdge(edge, edges(nextTab)[side]) {
		nextTab = flip(nextTab)
		if sameEdge(edge, edges(nextTab)[side]) {
			break
		}
		nextTab = flip(nextTab)
		nextTab = rotateTab(nextTab)
	}
	return nextId, nextTab
}

func countMonster(image [][]string) int {
	res := 0
	reg1, _ := regexp.Compile("^(#|.){18}#$")
	reg2, _ := regexp.Compile("^#(#|.){4}##(#|.){4}##(#|.){4}###$")
	reg3, _ := regexp.Compile("^#(#|.){2}#(#|.){2}#(#|.){2}#(#|.){2}#(#|.){2}#$")
	for i := 0; i < len(image)-2; i++ {
		for j := 0; j < len(image)-20; j++ {
			if reg1.MatchString(strings.Join(image[i][j:(j+19)], "")) && reg2.MatchString(strings.Join(image[i+1][j:(j+20)], "")) && reg3.MatchString(strings.Join(image[i+2][(j+1):(j+17)], "")) {
				res++
			}
		}
	}
	return res
}

func part1(s string) int {
	res := 1
	tilesList := format(s)
	for i, t := range tilesList {
		k := 0
		for _, edge1 := range edges(t.tab) {
			for j, t2 := range tilesList {
				if i != j {
					for _, edge2 := range edges(t2.tab) {
						if sameEdge(edge1, edge2) || sameEdgeReverse(edge1, edge2) {
							k++
						}
					}
				}
			}
		}
		if k == 2 {
			res *= t.id
		}
	}
	return res
}

func part2(s string) int {
	path := make(map[string]couple)
	match := make(map[int]int)
	tabs := make(map[int][][]string)
	tilesList := format(s)
	for i, t := range tilesList {
		tabs[t.id] = t.tab
		for _, edge1 := range edges(t.tab) {
			for j, t2 := range tilesList {
				if i != j {
					for _, edge2 := range edges(t2.tab) {
						if sameEdge(edge1, edge2) || sameEdgeReverse(edge1, edge2) {
							path[strings.Join(edge1, "")] = couple{t.id, t2.id}
							path[strings.Join(reverse(edge1), "")] = couple{t.id, t2.id}
							match[t.id]++
						}
					}
				}
			}
		}
	}
	image := [][][][]string{}
	var startID int
	var startTab [][]string
	N := int(math.Sqrt(float64(len(tilesList))))
	for key, value := range match {
		if value == 2 {
			startID = key
			break
		}
	}
	if len(tilesList) == 9 {
		startID = 1951
	}
	startTab = tabs[startID]
	startTab = firstTab(startTab, path)
	for i := 0; i < N; i++ {
		image = append(image, make([][][]string, N))
		image[i][0] = startTab
		currId := startID
		currTab := startTab
		for j := 1; j < N; j++ {
			currId, currTab = nextTab(edges(currTab)[1], path, currId, 3, tabs)
			image[i][j] = currTab
		}
		if i == N-1 {
			break
		}
		startID, startTab = nextTab(edges(startTab)[2], path, startID, 0, tabs)
	}
	imageClean := [][]string{}
	for i := 0; i < len(image); i++ {
		for x := 1; x < len(startTab)-1; x++ {
			imageClean = append(imageClean, []string{})
			for j := 0; j < len(image); j++ {
				for y := 1; y < len(startTab)-1; y++ {
					imageClean[x-1+i*(len(startTab)-2)] = append(imageClean[x-1+i*(len(startTab)-2)], image[i][j][x][y])
				}
			}
		}
	}
	for countMonster(imageClean) == 0 {
		imageClean = flip(imageClean)
		if countMonster(imageClean) != 0 {
			break
		}
		imageClean = flip(imageClean)
		imageClean = rotateTab(imageClean)
	}
	res := 0
	for i := 0; i < len(imageClean); i++ {
		for j := 0; j < len(imageClean[i]); j++ {
			if imageClean[i][j] == "#" {
				res++
			}
		}
	}
	return res - countMonster(imageClean)*15
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day20/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2020/day20/input.txt")

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
