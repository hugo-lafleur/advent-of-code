package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type list []int

func (l *list) reverse(index int, length int) {
	n := len(*l)
	i := index
	j := (index + length - 1) % n
	k := 0
	for k < length/2 {
		(*l)[i], (*l)[j] = (*l)[j], (*l)[i]
		i = (i + 1) % n
		j = (j - 1 + n) % n
		k++
	}
}

func format2(s string) []int {
	res := []int{}
	for _, r := range s {
		res = append(res, int(r))
	}
	res = append(res, 17, 31, 73, 47, 23)
	return res
}

func xor(l []int) int {
	res := l[0]
	for i := 1; i < len(l); i++ {
		res = res ^ l[i]
	}
	return res
}

func knotHash(s string) string {
	lengths := format2(s)
	l := 256
	var list list
	for i := 0; i < l; i++ {
		list = append(list, i)
	}
	skipSize := 0
	curr := 0
	for i := 0; i < 64; i++ {
		for _, length := range lengths {
			list.reverse(curr, length)
			curr = (curr + length + skipSize) % len(list)
			skipSize++
		}
	}
	xors := []int{}
	for i := 0; i < l/16; i++ {
		xors = append(xors, xor(list[16*i:16*(i+1)]))
	}
	res := ""
	for _, x := range xors {
		res += fmt.Sprintf("%02x", x)
	}
	return res
}

func hexToBin(r rune) string {
	switch r {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'a':
		return "1010"
	case 'b':
		return "1011"
	case 'c':
		return "1100"
	case 'd':
		return "1101"
	case 'e':
		return "1110"
	case 'f':
		return "1111"
	}
	return ""
}

func Ones(s string) int {
	c := 0
	for _, r := range s {
		if r == '1' {
			c++
		}
	}
	return c
}

func binToArray(s string) []int {
	res := []int{}
	for _, r := range s {
		if r == '1' {
			res = append(res, 1)
			continue
		}
		res = append(res, 0)
	}
	return res
}

func part1(s string) int {
	c := 0
	for i := 0; i < 128; i++ {
		row := s + "-" + strconv.Itoa(i)
		knot := knotHash(row)
		bin := ""
		for _, r := range knot {
			bin += hexToBin(r)
		}
		c += Ones(bin)
	}
	return c
}

func isValid(i, j int) bool {
	return i >= 0 && j >= 0 && j < 128 && i < 128
}

func visit(i, j int, tab [][]int, visited map[int]bool) map[int]bool {
	visited[128*i+j] = true
	for _, x := range []int{i + 1, i - 1} {
		_, v := visited[128*x+j]
		if !v && isValid(x, j) && tab[x][j] == 1 {
			visited = visit(x, j, tab, visited)
		}
	}
	for _, y := range []int{j + 1, j - 1} {
		_, v := visited[128*i+y]
		if !v && isValid(i, y) && tab[i][y] == 1 {
			visited = visit(i, y, tab, visited)
		}
	}
	return visited
}

func part2(s string) int {
	c := 0
	tab := [][]int{}
	visited := make(map[int]bool)
	for i := 0; i < 128; i++ {
		row := s + "-" + strconv.Itoa(i)
		knot := knotHash(row)
		bin := ""
		for _, r := range knot {
			bin += hexToBin(r)
		}
		arr := binToArray(bin)
		tab = append(tab, arr)
	}
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			_, v := visited[128*i+j]
			if !v && tab[i][j] == 1 {
				visited = visit(i, j, tab, visited)
				c++
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day14/test.data")

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

	content, err = os.ReadFile("../../../inputs/2017/day14/input.data")

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
