package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

func hexa_to_bits(input string) []int {
	s := []int{}
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '0':
			s = append(s, 0, 0, 0, 0)
		case '1':
			s = append(s, 0, 0, 0, 1)
		case '2':
			s = append(s, 0, 0, 1, 0)
		case '3':
			s = append(s, 0, 0, 1, 1)
		case '4':
			s = append(s, 0, 1, 0, 0)
		case '5':
			s = append(s, 0, 1, 0, 1)
		case '6':
			s = append(s, 0, 1, 1, 0)
		case '7':
			s = append(s, 0, 1, 1, 1)
		case '8':
			s = append(s, 1, 0, 0, 0)
		case '9':
			s = append(s, 1, 0, 0, 1)
		case 'A', 'a':
			s = append(s, 1, 0, 1, 0)
		case 'B', 'b':
			s = append(s, 1, 0, 1, 1)
		case 'C', 'c':
			s = append(s, 1, 1, 0, 0)
		case 'D', 'd':
			s = append(s, 1, 1, 0, 1)
		case 'E', 'e':
			s = append(s, 1, 1, 1, 0)
		case 'F', 'f':
			s = append(s, 1, 1, 1, 1)
		}
	}
	return s
}

func litteral(input []int) bool {
	if input[5]+2*input[4]+4*input[3] == 4 {
		return true
	} else {
		return false
	}

}

func version(input []int) int {
	return input[2] + 2*input[1] + 4*input[0]
}

func length(input []int) int {
	return input[6]
}

func pow(x int, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func type0(input []int) int {
	s := 0
	for i := 7; i < 22; i++ {
		s += input[i] * pow(2, 21-i)
	}
	return s
}

func type1(input []int) int {
	s := 0
	for i := 7; i < 18; i++ {
		s += input[i] * pow(2, 17-i)
	}
	return s
}

func len_type1(input []int) int {
	l := type1(input)
	c := 0
	i := 0
	current := input[18:]
	for c < l {
		if litteral(current[i:]) {
			i += 6
			for current[i] == 1 {
				i += 5
			}
			i += 5
			c += 1
		} else {
			if length(current[i:]) == 0 {
				i += 22 + type0(current[i:])
				c += 1
			} else {
				i += 18 + len_type1(current[i:])
				c += 1
			}
		}
	}
	return i
}
func packets(input []int) [][]int {
	s := [][]int{}
	i := 0
	j := 0
	l := 0
	var current []int
	if length(input) == 0 {
		l = type0(input)
		current = input[22:]
	} else {
		l = len_type1(input)
		current = input[18:]
	}
	for i < l {
		if litteral(current[j:]) {
			i += 6
			for current[i] == 1 {
				i += 5
			}
			i += 5
			s = append(s, current[j:i])
			j = i
		} else {
			if length(current[j:]) == 0 {
				s = append(s, current[j:j+22+type0(current[j:])])
				i = j + 22 + type0(current[j:])
				j = i
			} else {
				s = append(s, current[j:j+18+len_type1(current[j:])])
				i = j + 18 + len_type1(current[j:])
				j = i
			}
		}
	}
	return s
}

func version_sum(input []int) int {
	s := 0
	s += version(input)
	if litteral(input) {
		return s
	} else {
		slice := packets(input)
		for i := 0; i < len(slice); i++ {
			s += version_sum(slice[i])
		}
		return s
	}
}

func values(input []int) int {
	b := []int{}
	current := input[6:]
	i := 0
	v := 0
	for current[i] == 1 {
		b = append(b, current[i+1:i+5]...)
		i = i + 5
	}
	b = append(b, current[i+1:i+5]...)
	for j := 0; j < len(b); j++ {
		v += b[j] * pow(2, len(b)-j-1)
	}
	return v
}

func id(input []int) int {
	s := 4*input[3] + 2*input[4] + input[5]
	return s
}

func calculate(input []int) int {
	s := 0
	if litteral(input) {
		s = values(input)
	} else {
		if id(input) == 0 {
			s = 0
			p := packets(input)
			for i := 0; i < len(p); i++ {
				s = s + calculate(p[i])
			}
		}
		if id(input) == 1 {
			s = 1
			p := packets(input)
			for i := 0; i < len(p); i++ {
				s = s * calculate(p[i])
			}
		}
		if id(input) == 2 {
			p := packets(input)
			s = -1
			for i := 0; i < len(p); i++ {
				if s == -1 || calculate(p[i]) < s {
					s = calculate(p[i])
				}
			}
			return s
		}
		if id(input) == 3 {
			p := packets(input)
			s = -1
			for i := 0; i < len(p); i++ {
				if s == -1 || calculate(p[i]) > s {
					s = calculate(p[i])
				}
			}
			return s
		}
		if id(input) == 5 {
			p := packets(input)
			if calculate(p[0]) > calculate(p[1]) {
				s = 1
			} else {
				s = 0
			}
		}
		if id(input) == 6 {
			p := packets(input)
			if calculate(p[0]) >= calculate(p[1]) {
				s = 0
			} else {
				s = 1
			}
		}
		if id(input) == 7 {
			p := packets(input)
			if calculate(p[0]) == calculate(p[1]) {
				s = 1
			} else {
				s = 0
			}
		}
	}
	return s
}

func part1(s string) int {
	hex := hexa_to_bits(s)
	return version_sum(hex)
}

func part2(s string) int {
	hex := hexa_to_bits(s)
	return calculate(hex)
}
func main() {
	content, err := os.ReadFile("../../../inputs/2021/day16/test1.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2021/day16/test2.txt")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2021/day16/input.txt")

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
