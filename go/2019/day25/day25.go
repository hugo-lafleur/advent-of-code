package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gammazero/deque"
)

type Room struct {
	name  string
	path  []string
	items []string
}

type subsets [][]string

func (a subsets) Len() int           { return len(a) }
func (a subsets) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a subsets) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func format(s string) map[int]int {
	res := make(map[int]int)
	lines := strings.Split(s, ",")
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		res[i] = n
	}
	return res
}

func opcodeSolve(n int) (int, int, int, int) {
	s := strconv.Itoa(n)
	for len(s) != 5 {
		s = "0" + s
	}
	var a, b, c, de int
	a, _ = strconv.Atoi(string(s[0]))
	b, _ = strconv.Atoi(string(s[1]))
	c, _ = strconv.Atoi(string(s[2]))
	de, _ = strconv.Atoi(s[3:])
	return de, c, b, a
}

func stringToASCII(s string) []int {
	res := []int{}
	for _, r := range s {
		res = append(res, int(r))
	}
	return res
}

func ASCIIToString(l []int) string {
	res := ""
	for _, x := range l {
		res += string(rune(x))
	}
	return res
}

func stringToRoom(s string) Room {
	var name string
	var path, items []string
	lines := strings.Split(s, "\n")
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "==") {
			name = line[3 : len(line)-3]
			path = []string{}
			items = []string{}
		}
		if strings.Contains(line, "Doors here lead:") {
			i++
			line = lines[i]
			for strings.Contains(line, "-") {
				path = append(path, line[2:])
				i++
				line = lines[i]
			}
		}
		if strings.Contains(line, "Items here:") {
			i++
			line = lines[i]
			for strings.Contains(line, "-") {
				items = append(items, line[2:])
				i++
				line = lines[i]
			}
		}
	}
	return Room{name, path, items}
}

func solve(s string, in, out chan string) {
	var output []int
	var inputList []int
	p := format(s)
	i := 0
	relativeBase := 0
mainLoop:
	for {
		opcode, mode1, mode2, mode3 := opcodeSolve(p[i])
		var a, b, c, res int
		if opcode == 99 {
			break
		}
		switch mode1 {
		case 0:
			a = p[p[i+1]]
		case 1:
			a = p[i+1]
		case 2:
			a = p[relativeBase+p[i+1]]
		}
		switch mode2 {
		case 0:
			b = p[p[i+2]]
		case 1:
			b = p[i+2]
		case 2:
			b = p[relativeBase+p[i+2]]
		}
		switch mode3 {
		case 0:
			c = p[i+3]
		case 2:
			c = p[i+3] + relativeBase
		}
		switch opcode {
		case 1:
			res = a + b
			i += 3
		case 2:
			res = a * b
			i += 3
		case 3:
			if len(inputList) == 0 {
				str := ASCIIToString(output)
				output = []int{}
				out <- str
				str = <-in
				inputList = stringToASCII(str)
			}
			input := inputList[0]
			inputList = inputList[1:]
			switch mode1 {
			case 0:
				p[p[i+1]] = input
			case 2:
				p[relativeBase+p[i+1]] = input
			}
			i += 2
			continue mainLoop
		case 4:
			output = append(output, a)
			i += 2
			continue mainLoop
		case 5:
			if a != 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 6:
			if a == 0 {
				i = b
				continue mainLoop
			}
			i += 3
			continue mainLoop
		case 7:
			if a < b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 8:
			if a == b {
				res = 1
			} else {
				res = 0
			}
			i += 3
		case 9:
			relativeBase += a
			i += 2
			continue mainLoop
		}
		p[c] = res
		i++
	}
	out <- ASCIIToString(output)
}

func opposite(d string) string {
	switch d {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	}
	return ""
}

func allSubsets(l []string) subsets {
	res := [][]string{{}}
	for _, s := range l {
		temp := [][]string{}
		for _, subset := range res {
			var new []string
			new = append(new, subset...)
			new = append(new, s)
			temp = append(temp, new)
		}
		res = append(res, temp...)
	}
	return subsets(res)
}

func part1(s string) int {
	in, out := make(chan string), make(chan string)
	res := make(chan int)
	visitedRoom := make(map[string]bool)
	var dq deque.Deque[string]
	var wg sync.WaitGroup
	var path []string
	var pathToSecurity []string
	var goCheck string
	wg.Add(1)
	inventory := []string{}
	go solve(s, in, out)
	go func() {
		var room Room
		for {
			str := <-out
			if strings.Contains(str, "You take the") {
				room.items = room.items[1:]
			} else {
				room = stringToRoom(str)
			}
			_, roomVisited := visitedRoom[room.name]
			if len(room.items) != 0 && room.items[0] != "molten lava" && room.items[0] != "photons" && room.items[0] != "giant electromagnet" && room.items[0] != "infinite loop" && room.items[0] != "escape pod" {
				str = "take " + room.items[0] + "\n"
				inventory = append(inventory, room.items[0])
				in <- str
				continue
			}
			if !roomVisited {
				if room.name != "Security Checkpoint" {
					var last string
					if dq.Len() != 0 {
						last = dq.Back()
					}
					for _, direction := range room.path {
						if dq.Len() != 0 && direction == last {
							continue
						}
						dq.PushBack(direction)
					}
				} else {
					pathToSecurity = path
					last := dq.Back()
					for _, direction := range room.path {
						if direction != last {
							goCheck = direction
						}
					}
				}
				visitedRoom[room.name] = true
			} else {
				dq.PopBack()
			}
			if dq.Len() == 0 {
				defer wg.Done()
				break
			}
			dir := dq.PopBack()
			path = append(path, dir)
			str = dir + "\n"
			dq.PushBack(opposite(dir))
			in <- str
		}
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		for {
			in <- pathToSecurity[0] + "\n"
			pathToSecurity = pathToSecurity[1:]
			<-out
			if len(pathToSecurity) == 0 {
				defer wg.Done()
				break
			}
		}

	}()
	wg.Wait()
	wg.Add(1)
	allSubsetsInventory := allSubsets(inventory)
	go func() {
		for {
			item := inventory[0]
			inventory = inventory[1:]
			str := "drop " + item + "\n"
			in <- str
			<-out
			if len(inventory) == 0 {
				defer wg.Done()
				break
			}

		}
	}()
	wg.Wait()
	sort.Sort(allSubsetsInventory)
	go func() {
		for _, subset := range allSubsetsInventory {
			if len(subset) != 0 {
				for _, item := range subset {
					str := "take " + item + "\n"
					in <- str
					<-out
				}
				in <- (goCheck + "\n")
				str := <-out
				if strings.Contains(str, "Santa") {
					split := strings.Split(str, " ")
					for _, word := range split {
						n, ok := strconv.Atoi(word)
						if ok == nil {
							res <- n
						}
					}
				} else {
					for _, item := range subset {
						str := "drop " + item + "\n"
						in <- str
						<-out
					}
				}
			}
		}
	}()
	return <-res
}

func main() {
	content, err := os.ReadFile("../../../inputs/2019/day25/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Input :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %v\n", part1(string(content)))
	fmt.Println(time.Since(start))
}
