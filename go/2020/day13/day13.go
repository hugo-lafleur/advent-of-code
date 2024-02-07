package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func format(s string) (int, []string) {
	lines := strings.Split(s, "\n")
	n, _ := strconv.Atoi(lines[0])
	return n, strings.Split(lines[1], ",")
}

func chineseRemainder(num, rem []*big.Int) *big.Int {
	p := big.NewInt(1)
	for _, x := range num {
		p.Mul(p, x)
	}
	var pp, x, y, gcd, res, z big.Int
	for i := range num {
		pp.Div(p, num[i])
		gcd.GCD(&x, &y, &pp, num[i])
		res.Add(&res, z.Mul(rem[i], z.Mul(&x, &pp)))
	}
	return res.Mod(&res, p)
}

func part1(s string) int {
	timestamp, IDS := format(s)
	min := 100000
	minID := 0
	for _, ID := range IDS {
		n, err := strconv.Atoi(ID)
		if err == nil {
			wait := n - (timestamp % n)
			if wait < min {
				min = wait
				minID = n
			}
		}
	}
	return min * minID
}

func part2(s string) int {
	_, IDS := format(s)
	num := []*big.Int{}
	rem := []*big.Int{}
	for i, ID := range IDS {
		n, err := strconv.Atoi(ID)
		if err == nil {
			num = append(num, big.NewInt(int64(n)))
			rem = append(rem, big.NewInt(int64(n-i)))
		}
	}
	return int(chineseRemainder(num, rem).Int64())
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day13/test.data")

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

	content, err = os.ReadFile("../../../inputs/2020/day13/input.data")

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
