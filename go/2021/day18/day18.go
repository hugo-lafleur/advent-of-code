package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

type pair struct {
	tag   int
	left  *pair
	right *pair
	depth int
}

func format(s string) []string {
	return strings.Split(s, "\n")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func explode(s string) string {
	k := 0
	var maxK int
	var res string
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if r == '[' {
			k++
		}
		if r == ']' {
			k--
		}
		maxK = max(maxK, k)
	}
	if maxK < 5 {
		return s
	}
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if r == '[' {
			k++
		}
		if r == ']' {
			k--
		}
		if k == 5 && r >= '0' && r <= '9' {
			start := i - 1
			nStr := ""
			for r = rune(s[i]); r >= '0' && r <= '9'; r = rune(s[i]) {
				nStr += string(r)
				i++
			}
			i++
			left, _ := strconv.Atoi(nStr)
			nStr = ""
			for r = rune(s[i]); r >= '0' && r <= '9'; r = rune(s[i]) {
				nStr += string(r)
				i++
			}
			end := i + 1
			right, _ := strconv.Atoi(nStr)
			i--
			before := ""
			for j := start; j >= 0; j-- {
				r = rune(s[j])
				if r >= '0' && r <= '9' {
					l := j
					for r = rune(s[l]); r >= '0' && r <= '9'; r = rune(s[l]) {
						before = string(r) + before
						l--
					}
					beforeInt, _ := strconv.Atoi(before)
					res = s[:(l+1)] + strconv.Itoa(beforeInt+left) + s[(j+1):start]
					break
				}
			}
			if len(res) == 0 {
				res = s[:start]
			}
			res += "0"
			after := ""
			for j := end; j < len(s); j++ {
				r = rune(s[j])
				if r >= '0' && r <= '9' {
					l := j
					for r = rune(s[l]); r >= '0' && r <= '9'; r = rune(s[l]) {
						after = after + string(r)
						l++
					}
					afterInt, _ := strconv.Atoi(after)
					res += s[end:j] + strconv.Itoa(afterInt+right) + s[l:]
					return res
				}
			}
			return res + s[end:]
		}
	}
	return s
}

func split(s string) string {
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if r >= '0' && r <= '9' {
			j := i
			nStr := ""
			for r = rune(s[j]); r >= '0' && r <= '9'; r = rune(s[j]) {
				nStr += string(r)
				j++
			}
			n, _ := strconv.Atoi(nStr)
			if n > 9 {
				res := s[:i] + "["
				if n%2 == 0 {
					res += strconv.Itoa(n/2) + "," + strconv.Itoa(n/2) + "]" + s[j:]
					return res
				} else {
					res += strconv.Itoa(n/2) + "," + strconv.Itoa((n/2)+1) + "]" + s[j:]
					return res
				}
			}
		}
	}
	return s
}

func add(s1, s2 string) string {
	return "[" + s1 + "," + s2 + "]"
}

func reduce(s string) string {
	exploded := explode(s)
	if exploded != s {
		return reduce(exploded)
	} else {
		splited := split(s)
		if splited != s {
			return reduce(splited)
		} else {
			return splited
		}
	}
}

func tree(s string) *pair {
	lines := strings.Split(s, "\n")
	var q deque.Deque[*pair]
	for _, line := range lines {
		id := 1
		root := new(pair)
		root.depth = 0
		root.tag = -1
		q.PushBack(root)
		l := len(line)
		for i := 0; i < l; i++ {
			r := rune(line[i])
			curr := q.PopBack()
			if i == l-1 {
				return curr
			}
			if r == '[' {
				child := new(pair)
				child.depth = curr.depth + 1
				curr.left = child
				q.PushBack(curr)
				q.PushBack(child)
			}
			if r == ',' {
				curr.tag = -1
				child := new(pair)
				child.depth = curr.depth + 1
				curr.right = child
				q.PushBack(curr)
				q.PushBack(child)
			}
			if r == ']' {
				continue
			}
			if r >= '0' && r <= '9' {
				nStr := ""
				for ; r >= '0' && r <= '9'; r = rune(line[i]) {
					nStr += string(r)
					i++
				}
				i = i - 1
				n, _ := strconv.Atoi(nStr)
				curr.tag = n
				id++
			}
		}
	}
	return new(pair)
}

func (p *pair) magnitude() int {
	if p.left == nil && p.right == nil {
		return p.tag
	}
	return 3*p.left.magnitude() + 2*p.right.magnitude()
}

func part1(s string) int {
	list := format(s)
	curr := list[0]
	for i := 1; i < len(list); i++ {
		curr = reduce(add(curr, list[i]))
	}
	return tree(curr).magnitude()
}

func part2(s string) int {
	list := format(s)
	var res int
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			res = max(res, tree(reduce(add(list[i], list[j]))).magnitude())
			res = max(res, tree(reduce(add(list[j], list[i]))).magnitude())
		}
	}
	return res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2021/day18/test.txt")

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

	content, err = os.ReadFile("../../../inputs/2021/day18/input.txt")

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
