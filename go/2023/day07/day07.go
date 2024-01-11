package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type hands []string

type hands2 []string

func format(s string) [][]string {
	lines := strings.Split(s, "\n")
	res := [][]string{}
	for _, line := range lines {
		res = append(res, strings.Split(line, " "))
	}
	return res
}

func handsNature(s string) int {
	dict := make(map[rune]int)
	for _, r := range s {
		dict[r]++
	}
	fours := []rune{}
	threes := []rune{}
	twos := []rune{}
	for card, occ := range dict {
		if occ == 5 {
			return 7
		}
		if occ == 4 {
			fours = append(fours, card)
		}
		if occ == 3 {
			threes = append(threes, card)
		}
		if occ == 2 {
			twos = append(twos, card)
		}
	}
	if len(fours) == 1 {
		return 6
	}
	if len(threes) == 1 && len(twos) == 1 {
		return 5
	}
	if len(threes) == 1 && len(twos) == 0 {
		return 4
	}
	if len(twos) == 2 {
		return 3
	}
	if len(twos) == 1 {
		return 2
	} else {
		return 1
	}
}

func compareCards(r1, r2 rune) int {
	if int(r1) < 58 && int(r2) < 58 {
		return int(r1) - int(r2)
	}
	if int(r1) > 64 && int(r2) < 58 {
		return 1
	}
	if int(r1) < 58 && int(r2) > 64 {
		return -1
	}
	if int(r1) > 64 && int(r2) > 64 {
		if r1 == 'J' {
			if r2 == 'A' || r2 == 'Q' || r2 == 'K' {
				return -1
			}
			if r2 == 'J' {
				return 0
			} else {
				return 1
			}
		}
		if r2 == 'J' {
			if r1 == 'A' || r1 == 'Q' || r1 == 'K' {
				return 1
			} else {
				return -1
			}
		} else {
			return int(r2) - int(r1)
		}
	}
	return 0
}

func (h hands) Len() int {
	return len(h)
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hands) Less(i, j int) bool {
	h1 := h[i]
	h2 := h[j]
	n1 := handsNature(h1)
	n2 := handsNature(h2)
	if n1 > n2 {
		return false
	}
	if n1 < n2 {
		return true
	} else {
		i := 0
		for i < 5 {
			a := h1[i]
			b := h2[i]
			c := compareCards(rune(a), rune(b))
			if c < 0 {
				return true
			}
			if c > 0 {
				return false
			}
			i++
		}
	}
	return false
}

func part1(s string) int {
	c := 0
	list := format(s)
	handsList := []string{}
	for _, line := range list {
		handsList = append(handsList, line[0])
	}
	sort.Sort(hands(handsList))
	for i, hands := range handsList {
		for _, line := range list {
			if hands == line[0] {
				n, _ := strconv.Atoi(line[1])
				c += (i + 1) * n
			}
		}
	}
	return c
}

func compareCards2(r1, r2 rune) int {
	if r1 == r2 {
		return 0
	}
	if r1 == 'J' {
		return -1
	}
	if r2 == 'J' {
		return 1
	}
	if int(r1) < 58 && int(r2) < 58 {
		return int(r1) - int(r2)
	}
	if int(r1) > 64 && int(r2) < 58 {
		return 1
	}
	if int(r1) < 58 && int(r2) > 64 {
		return -1
	}
	if int(r1) > 64 && int(r2) > 64 {
		return int(r2) - int(r1)
	}
	return 0
}

func handsNature2(s string) int {
	dict := make(map[rune]int)
	for _, r := range s {
		dict[r]++
	}
	fours := []rune{}
	threes := []rune{}
	twos := []rune{}
	j := 0
	for card, occ := range dict {
		if card != 'J' {
			if occ == 5 {
				return 7
			}
			if occ == 4 {
				fours = append(fours, card)
			}
			if occ == 3 {
				threes = append(threes, card)
			}
			if occ == 2 {
				twos = append(twos, card)
			}
		} else {
			j = occ
		}
	}
	if j == 5 {
		return 7
	}
	if len(fours) == 1 {
		if j == 1 {
			return 7
		} else {
			return 6
		}
	}
	if len(threes) == 1 && len(twos) == 1 {
		return 5
	}
	if len(threes) == 1 && len(twos) == 0 {
		if j == 1 {
			return 6
		}
		if j == 2 {
			return 7
		} else {
			return 4
		}
	}
	if len(twos) == 2 {
		if j == 1 {
			return 5
		} else {
			return 3
		}
	}
	if len(twos) == 1 {
		if j == 1 {
			return 4
		}
		if j == 2 {
			return 6
		}
		if j == 3 {
			return 7
		} else {
			return 2
		}
	} else {
		if j == 1 {
			return 2
		}
		if j == 2 {
			return 4
		}
		if j == 3 {
			return 6
		}
		if j == 4 {
			return 7
		} else {
			return 1
		}
	}
}
func (h hands2) Len() int {
	return len(h)
}

func (h hands2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hands2) Less(i, j int) bool {
	h1 := h[i]
	h2 := h[j]
	n1 := handsNature2(h1)
	n2 := handsNature2(h2)
	if n1 > n2 {
		return false
	}
	if n1 < n2 {
		return true
	} else {
		i := 0
		for i < 5 {
			a := h1[i]
			b := h2[i]
			c := compareCards2(rune(a), rune(b))
			if c < 0 {
				return true
			}
			if c > 0 {
				return false
			}
			i++
		}
	}
	return false
}

func part2(s string) int {
	c := 0
	list := format(s)
	handsList := []string{}
	for _, line := range list {
		handsList = append(handsList, line[0])
	}
	sort.Sort(hands2(handsList))
	for i, hands := range handsList {
		for _, line := range list {
			if hands == line[0] {
				n, _ := strconv.Atoi(line[1])
				c += (i + 1) * n
			}
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day07/test.data")

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

	content, err = os.ReadFile("../../../inputs/2023/day07/input.data")

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
