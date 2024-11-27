package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

const MaxInt = 163904084080993271554974776000

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func newStack(dq deque.Deque[int]) deque.Deque[int] {
	var res deque.Deque[int]
	for dq.Len() != 0 {
		n := dq.PopFront()
		res.PushFront(n)
	}
	return res
}

func increment(dq deque.Deque[int], incr int) deque.Deque[int] {
	var res deque.Deque[int]
	for i := 0; i < dq.Len(); i++ {
		res.PushBack(i)
	}
	for dq.Len() != 0 {
		n := dq.PopFront()
		res.Set(0, n)
		res.Rotate(incr)
	}
	return res
}

func cut(dq deque.Deque[int], n int) deque.Deque[int] {
	var res deque.Deque[int]
	for i := 0; i < dq.Len(); i++ {
		res.PushBack(dq.At(i))
	}
	if n > 0 {
		for n > 0 {
			x := res.PopFront()
			res.PushBack(x)
			n--
		}
	}
	if n < 0 {
		for n < 0 {
			x := res.PopBack()
			res.PushFront(x)
			n++
		}
	}
	return res
}

func modinv(n, L int) int {
	z := big.NewInt(0)
	z.ModInverse(big.NewInt(int64(n)), big.NewInt(int64(L)))
	return int(z.Int64())
}

func mulmod(x, y, L int) int {
	z := big.NewInt(0)
	z.Mul(big.NewInt(int64(x)), big.NewInt(int64(y)))
	z.Mod(z, big.NewInt(int64(L)))
	return int(z.Int64())
}

func pow(x, n, L int) int {
	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		return pow(mulmod(x, x, L), n/2, L)
	}
	return mulmod(x, pow(x, n-1, L), L)
}

func part1(s string) int {
	instr := format(s)
	var dq deque.Deque[int]
	var L int
	var card int
	if len(instr) < 15 {
		L = 10
		card = 0
	} else {
		L = 10007
		card = 2019
	}
	for i := 0; i < L; i++ {
		dq.PushBack(i)
	}
	for _, line := range instr {
		switch line[1] {
		case "into":
			dq = newStack(dq)
		case "with":
			n, _ := strconv.Atoi(line[3])
			dq = increment(dq, n)
		default:
			n, _ := strconv.Atoi(line[1])
			dq = cut(dq, n)
		}
	}
	return dq.Index(func(x int) bool { return x == card })
}

func part2(s string) int {
	instr := format(s)
	slices.Reverse(instr)
	L := 119315717514047
	n := 101741582076661
	pos := 2020
	a, b := 1, 0
	for _, line := range instr {
		switch line[1] {
		case "into":
			a = -a
			b = L - b - 1
		case "with":
			n, _ := strconv.Atoi(line[3])
			z := modinv(n, L)
			a = mulmod(a, z, L)
			b = mulmod(b, z, L)
		default:
			n, _ := strconv.Atoi(line[1])
			b = (b + n) % L
		}
	}
	A := pow(a, n, L)
	B := mulmod(A-1, modinv(a-1, L), L)
	B = mulmod(B, b, L)
	return (mulmod(A, pos, L) + B) % L
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day22/test.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2019/day22/input.txt")

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
