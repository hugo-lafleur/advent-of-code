package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func part1(s string) string {
	res := []string{}
	input := s
	i := 0
	for len(res) != 8 {
		toBeHashed := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(toBeHashed))
		strHash := hex.EncodeToString(hash[:])
		if strHash[0] == '0' && strHash[1] == '0' && strHash[2] == '0' && strHash[3] == '0' && strHash[4] == '0' {
			res = append(res, string(strHash[5]))
		}
		i++
	}
	return strings.Join(res, "")
}

func part2(s string) string {
	res := make([]string, 8)
	input := s
	i := 0
	k := 0
	done := make(map[int]bool)
	for k < 8 {
		toBeHashed := input + strconv.Itoa(i)
		hash := md5.Sum([]byte(toBeHashed))
		strHash := hex.EncodeToString(hash[:])
		if strHash[0] == '0' && strHash[1] == '0' && strHash[2] == '0' && strHash[3] == '0' && strHash[4] == '0' {
			n, err := strconv.Atoi(string(strHash[5]))
			if err == nil && n < 8 && !done[n] {
				res[n] = string(strHash[6])
				done[n] = true
				k++
			}
		}
		i++
	}
	return strings.Join(res, "")
}

func main() {
	content, err := os.ReadFile("test.data")

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

	content, err = os.ReadFile("input.data")

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
