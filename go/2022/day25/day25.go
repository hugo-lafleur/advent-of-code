package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) [][]string {
	res := [][]string{}
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		res = append(res, strings.Split(line, ""))
	}
	return res
}

func snafuToDec(n []string) int {
	res := 0
	m := 1
	for i := len(n) - 1; i >= 0; i-- {
		switch n[i] {
		case "1":
			res += m
		case "2":
			res += 2 * m
		case "-":
			res -= m
		case "=":
			res -= 2 * m
		}
		m *= 5
	}
	return res
}

func part1(s string) string {
	dec := 0
	list := format(s)
	for _, number := range list {
		dec += snafuToDec(number)
	}
	base5 := strconv.FormatInt(int64(dec), 5)
	base5Tab := strings.Split(base5, "")
	carry := 0
	res := ""
	for i := len(base5Tab) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(base5Tab[i])
		n = n + carry
		carry = 0
		if n > 2 {
			n = n - 5
			carry = 1
		}
		if n == -1 {
			res = "-" + res
		}
		if n == -2 {
			res = "=" + res
		}
		if n >= 0 && n <= 2 {
			res = strconv.Itoa(n) + res
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day25/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day25/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
