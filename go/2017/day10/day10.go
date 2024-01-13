package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type list []int

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func format(s string) []int {
	strs := strings.FieldsFunc(s, Split)
	res := []int{}
	for _, str := range strs {
		n, _ := strconv.Atoi(str)
		res = append(res, n)
	}
	return res
}

func format2(s string) []int {
	res := []int{}
	for _, r := range s {
		res = append(res, int(r))
	}
	res = append(res, 17, 31, 73, 47, 23)
	return res
}

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

func xor(l []int) int {
	res := l[0]
	for i := 1; i < len(l); i++ {
		res = res ^ l[i]
	}
	return res
}

func sumString(l []string) string {
	res := ""
	for _, x := range l {
		res += x
	}
	return res
}

func part1(s string) int {
	lengths := format(s)
	var l int
	if len(lengths) == 4 {
		l = 5
	} else {
		l = 256
	}
	var list list
	for i := 0; i < l; i++ {
		list = append(list, i)
	}
	skipSize := 0
	curr := 0
	for _, length := range lengths {
		list.reverse(curr, length)
		curr = (curr + length + skipSize) % len(list)
		skipSize++
	}
	return list[0] * list[1]
}

func part2(s string) string {
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
	res := []string{}
	for _, x := range xors {
		res = append(res, fmt.Sprintf("%02x", x))
	}
	return sumString(res)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2017/day10/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day10/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2017/day10/input.data")

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
