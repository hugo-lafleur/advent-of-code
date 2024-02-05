package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Split(r rune) bool {
	return r == ' ' || r == ':' || r == '\n'
}

func format(s string) []map[string]string {
	res := []map[string]string{}
	passports := strings.Split(s, "\n\n")
	for _, passport := range passports {
		passportSplit := strings.FieldsFunc(passport, Split)
		pass := make(map[string]string)
		for i := 0; i < len(passportSplit); i += 2 {
			pass[passportSplit[i]] = passportSplit[i+1]
		}
		res = append(res, pass)
	}
	return res
}

func hasAllRequiredFields(m map[string]string) bool {
	for _, s := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, ok := m[s]
		if !ok {
			return false
		}
	}
	return true
}

func validBYR(s string) bool {
	n, err := strconv.Atoi(s)
	return (err == nil) && n >= 1920 && n <= 2002
}

func validIYR(s string) bool {
	n, err := strconv.Atoi(s)
	return (err == nil) && n >= 2010 && n <= 2020
}

func validEYR(s string) bool {
	n, err := strconv.Atoi(s)
	return (err == nil) && n >= 2020 && n <= 2030
}

func validHGT(s string) bool {
	if len(s) == 5 {
		hgt := s[0:3]
		msr := s[3:]
		n, err := strconv.Atoi(hgt)
		return (err == nil) && msr == "cm" && n >= 150 && n <= 193
	}
	if len(s) == 4 {
		hgt := s[0:2]
		msr := s[2:]
		n, err := strconv.Atoi(hgt)
		return (err == nil) && msr == "in" && n >= 59 && n <= 76
	}
	return false
}

func isHex(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f')
}

func validHCL(s string) bool {
	if len(s) != 7 {
		return false
	}
	for i, r := range s {
		if i == 0 && r != '#' {
			return false
		}
		if i > 0 && !isHex(r) {
			return false
		}
	}
	return true
}

func validECL(s string) bool {
	return s == "amb" || s == "blu" || s == "brn" || s == "gry" || s == "grn" || s == "hzl" || s == "oth"
}

func validPID(s string) bool {
	_, err := strconv.Atoi(s)
	return len(s) == 9 && (err == nil)
}

func part1(s string) int {
	c := 0
	passports := format(s)
	for _, pass := range passports {
		if hasAllRequiredFields(pass) {
			c++
		}
	}
	return c
}

func part2(s string) int {
	c := 0
	passports := format(s)
	for _, pass := range passports {
		//fmt.Println(hasAllRequiredFields(pass), validBYR(pass["byr"]), validIYR(pass["iyr"]), validEYR(pass["eyr"]), validHGT(pass["hgt"]), validHCL(pass["hcl"]), validECL(pass["ecl"]), validPID(pass["pid"]))
		if hasAllRequiredFields(pass) && validBYR(pass["byr"]) && validIYR(pass["iyr"]) && validEYR(pass["eyr"]) && validHGT(pass["hgt"]) && validHCL(pass["hcl"]) && validECL(pass["ecl"]) && validPID(pass["pid"]) {
			c++
		}
	}
	return c
}

func main() {
	content, err := os.ReadFile("../../../inputs/2020/day04/test1.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day04/test2.data")

	if err != nil {
		log.Fatal(err)
	}
	start = time.Now()
	fmt.Printf("\nPart 2 : %v\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2020/day04/input.data")

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
