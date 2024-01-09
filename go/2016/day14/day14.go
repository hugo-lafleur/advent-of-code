package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func MoreHashing(s string) string {
	res := GetMD5Hash(s)
	for i := 0; i < 2016; i++ {
		res = GetMD5Hash(res)
	}
	return res
}

func threepeat(s string) (byte, bool) {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i+1] == s[i+2] {
			return s[i], true
		}
	}
	return byte(0), false
}

func fivepeat(s string, char byte) bool {
	for i := 0; i < len(s)-4; i++ {
		if s[i] == char && s[i+1] == char && s[i+2] == char && s[i+3] == char && s[i+4] == char {
			return true
		}
	}
	return false
}

func part1(s string) int {
	key := 0
	index := 0
	limit := 100000
	hashDb := make(map[string]string)
	for i := 0; i < limit; i++ {
		hashDb[s+strconv.Itoa(i)] = GetMD5Hash(s + strconv.Itoa(i))
	}
	for key != 64 {
		hash := hashDb[s+strconv.Itoa(index)]
		char, ok := threepeat(hash)
		if ok {
			for j := index + 1; j < index+1001; j++ {
				if fivepeat(hashDb[(s+strconv.Itoa(j))], char) {
					key++
					break
				}
			}
		}
		index++
	}
	return index - 1
}

func part2(s string) int {
	key := 0
	index := 0
	limit := 50000
	hashDb := make(map[string]string)
	for i := 0; i < limit; i++ {
		hashDb[s+strconv.Itoa(i)] = MoreHashing(s + strconv.Itoa(i))
		if i == 10 {
			fmt.Println(hashDb[s+strconv.Itoa(i)])
		}
	}
	for key != 64 {
		hash := hashDb[s+strconv.Itoa(index)]
		char, ok := threepeat(hash)
		if ok {
			for j := index + 1; j < index+1001; j++ {
				if fivepeat(hashDb[(s+strconv.Itoa(j))], char) {
					key++
					break
				}
			}
		}
		index++
	}
	return index - 1
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
