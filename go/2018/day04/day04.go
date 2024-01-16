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

type event struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	action string
}

type eventList []event

func Split(r rune) bool {
	return r == '[' || r == ']' || r == ':' || r == ' ' || r == '#' || r == '-'
}

func (e eventList) Len() int {
	return len(e)
}

func (e eventList) Less(i, j int) bool {
	ev1 := e[i]
	ev2 := e[j]
	if ev1.year < ev2.year {
		return true
	}
	if ev1.year > ev2.year {
		return false
	} else {
		if ev1.month < ev2.month {
			return true
		}
		if ev1.month > ev2.month {
			return false
		} else {
			if ev1.day < ev2.day {
				return true
			}
			if ev1.day > ev2.day {
				return false
			} else {
				if ev1.hour < ev2.hour {
					return true
				}
				if ev1.hour > ev2.hour {
					return false
				} else {
					return ev1.minute < ev2.minute
				}
			}
		}
	}
}

func (e eventList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func format(s string) eventList {
	lines := strings.Split(s, "\n")
	res := eventList{}
	for _, line := range lines {
		strs := strings.FieldsFunc(line, Split)
		year, _ := strconv.Atoi(strs[0])
		month, _ := strconv.Atoi(strs[1])
		day, _ := strconv.Atoi(strs[2])
		hour, _ := strconv.Atoi(strs[3])
		minute, _ := strconv.Atoi(strs[4])
		res = append(res, event{year, month, day, hour, minute, strs[6]})
	}
	sort.Sort(res)
	return res
}

func sum(m map[int]int) int {
	res := 0
	for _, value := range m {
		res += value
	}
	return res
}

func max(m map[int]int) int {
	res := 0
	max := 0
	for key, value := range m {
		if value > max {
			max = value
			res = key
		}
	}
	return res
}

func part1(s string) int {
	events := eventList(format(s))
	minutesAsleep := make(map[int]map[int]int)
	currGuard := -1
	for _, event := range events {
		if event.action != "asleep" && event.action != "up" {
			n, _ := strconv.Atoi(event.action)
			minutesAsleep[n] = make(map[int]int)
		}
	}
	for i, event := range events {
		if event.action != "asleep" && event.action != "up" {
			n, _ := strconv.Atoi(event.action)
			currGuard = n
		}
		if event.action == "up" {
			lastEvent := events[i-1]
			for j := lastEvent.minute; j < event.minute; j++ {
				minutesAsleep[currGuard][j]++
			}
		}

	}
	mostAsleepGuard := -1
	maxMinutes := -1
	for key, value := range minutesAsleep {
		m := sum(value)
		if m > maxMinutes {
			maxMinutes = m
			mostAsleepGuard = key
		}
	}
	return mostAsleepGuard * max(minutesAsleep[mostAsleepGuard])
}

func part2(s string) int {
	events := eventList(format(s))
	minutesAsleep := make(map[int]map[int]int)
	currGuard := -1
	for _, event := range events {
		if event.action != "asleep" && event.action != "up" {
			n, _ := strconv.Atoi(event.action)
			minutesAsleep[n] = make(map[int]int)
		}
	}
	for i, event := range events {
		if event.action != "asleep" && event.action != "up" {
			n, _ := strconv.Atoi(event.action)
			currGuard = n
		}
		if event.action == "up" {
			lastEvent := events[i-1]
			for j := lastEvent.minute; j < event.minute; j++ {
				minutesAsleep[currGuard][j]++
			}
		}

	}
	guardChosen := -1
	minuteChosen := -1
	maxTime := -1
	for key, value := range minutesAsleep {
		for min, time := range value {
			if time > maxTime {
				maxTime = time
				guardChosen = key
				minuteChosen = min
			}
		}
	}
	return guardChosen * minuteChosen
}

func main() {
	content, err := os.ReadFile("../../../inputs/2018/day04/test.data")

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

	content, err = os.ReadFile("../../../inputs/2018/day04/input.data")

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
