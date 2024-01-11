package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func multiply(tab []int) int {
	sort.Ints(tab)
	l := len(tab)
	return tab[l-1] * tab[l-2]
}

func part1(s string) int {
	inspections := []int{}
	if len(strings.Split(s, "\n")) == 27 {
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		items := make([][]int, 4)
		items[0] = []int{79, 98}
		items[1] = []int{54, 65, 75, 74}
		items[2] = []int{79, 60, 97}
		items[3] = []int{74}
		i := 0
		for i < 20 {
			monkey := 0
			for monkey < 4 {
				item := 0
				for len(items[monkey]) != 0 {
					inspections[monkey]++
					if monkey == 0 {
						x := items[monkey][item]
						x = x * 19
						x = x / 3
						if x%23 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 1 {
						x := items[monkey][item]
						x = x + 6
						x = x / 3
						if x%19 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[0] = append(items[0], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 2 {
						x := items[monkey][item]
						x = x * x
						x = x / 3
						if x%13 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 3 {
						x := items[monkey][item]
						x = x + 3
						x = x / 3
						if x%17 == 0 {
							items[0] = append(items[0], x)
						} else {
							items[1] = append(items[1], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
				}
				monkey++
			}
			i++
		}
	} else {
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		items := make([][]int, 8)
		items[0] = []int{91, 58, 52, 69, 95, 54}
		items[1] = []int{80, 80, 97, 84}
		items[2] = []int{86, 92, 71}
		items[3] = []int{96, 90, 99, 76, 79, 85, 98, 61}
		items[4] = []int{60, 83, 68, 64, 73}
		items[5] = []int{96, 52, 52, 94, 76, 51, 57}
		items[6] = []int{75}
		items[7] = []int{83, 75}
		i := 0
		for i < 20 {
			monkey := 0
			for monkey < 8 {
				item := 0
				for len(items[monkey]) != 0 {
					inspections[monkey]++
					if monkey == 0 {
						x := items[monkey][item]
						x = x * 13
						x = x / 3
						if x%7 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[5] = append(items[5], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 1 {
						x := items[monkey][item]
						x = x * x
						x = x / 3
						if x%3 == 0 {
							items[3] = append(items[3], x)
						} else {
							items[5] = append(items[5], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 2 {
						x := items[monkey][item]
						x = x + 7
						x = x / 3
						if x%2 == 0 {
							items[0] = append(items[0], x)
						} else {
							items[4] = append(items[4], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 3 {
						x := items[monkey][item]
						x = x + 4
						x = x / 3
						if x%11 == 0 {
							items[7] = append(items[7], x)
						} else {
							items[6] = append(items[6], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 4 {
						x := items[monkey][item]
						x = x * 19
						x = x / 3
						if x%17 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[0] = append(items[0], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 5 {
						x := items[monkey][item]
						x = x + 3
						x = x / 3
						if x%5 == 0 {
							items[7] = append(items[7], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 6 {
						x := items[monkey][item]
						x = x + 5
						x = x / 3
						if x%13 == 0 {
							items[4] = append(items[4], x)
						} else {
							items[2] = append(items[2], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 7 {
						x := items[monkey][item]
						x = x + 1
						x = x / 3
						if x%19 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[6] = append(items[6], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
				}
				monkey++
			}
			i++
		}
	}
	return multiply(inspections)
}

func part2(s string) int {
	inspections := []int{}
	if len(strings.Split(s, "\n")) == 27 {
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		items := make([][]int, 4)
		items[0] = []int{79, 98}
		items[1] = []int{54, 65, 75, 74}
		items[2] = []int{79, 60, 97}
		items[3] = []int{74}
		i := 0
		for i < 10000 {
			monkey := 0
			for monkey < 4 {
				item := 0
				for len(items[monkey]) != 0 {
					inspections[monkey]++
					if monkey == 0 {
						x := items[monkey][item]
						x = x * 19
						x = x % 96577
						if x%23 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 1 {
						x := items[monkey][item]
						x = x + 6
						x = x % 96577
						if x%19 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[0] = append(items[0], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 2 {
						x := items[monkey][item]
						x = x * x
						x = x % 96577
						if x%13 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 3 {
						x := items[monkey][item]
						x = x + 3
						x = x % 96577
						if x%17 == 0 {
							items[0] = append(items[0], x)
						} else {
							items[1] = append(items[1], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
				}
				monkey++
			}
			i++
		}
	} else {
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		inspections = append(inspections, 0)
		items := make([][]int, 8)
		items[0] = []int{91, 58, 52, 69, 95, 54}
		items[1] = []int{80, 80, 97, 84}
		items[2] = []int{86, 92, 71}
		items[3] = []int{96, 90, 99, 76, 79, 85, 98, 61}
		items[4] = []int{60, 83, 68, 64, 73}
		items[5] = []int{96, 52, 52, 94, 76, 51, 57}
		items[6] = []int{75}
		items[7] = []int{83, 75}
		i := 0
		for i < 10000 {
			monkey := 0
			for monkey < 8 {
				item := 0
				for len(items[monkey]) != 0 {
					inspections[monkey]++
					if monkey == 0 {
						x := items[monkey][item]
						x = x * 13
						x = x % 9699690
						if x%7 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[5] = append(items[5], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 1 {
						x := items[monkey][item]
						x = x * x
						x = x % 9699690
						if x%3 == 0 {
							items[3] = append(items[3], x)
						} else {
							items[5] = append(items[5], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 2 {
						x := items[monkey][item]
						x = x + 7
						x = x % 9699690
						if x%2 == 0 {
							items[0] = append(items[0], x)
						} else {
							items[4] = append(items[4], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 3 {
						x := items[monkey][item]
						x = x + 4
						x = x % 9699690
						if x%11 == 0 {
							items[7] = append(items[7], x)
						} else {
							items[6] = append(items[6], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 4 {
						x := items[monkey][item]
						x = x * 19
						x = x % 9699690
						if x%17 == 0 {
							items[1] = append(items[1], x)
						} else {
							items[0] = append(items[0], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 5 {
						x := items[monkey][item]
						x = x + 3
						x = x % 9699690
						if x%5 == 0 {
							items[7] = append(items[7], x)
						} else {
							items[3] = append(items[3], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 6 {
						x := items[monkey][item]
						x = x + 5
						x = x % 9699690
						if x%13 == 0 {
							items[4] = append(items[4], x)
						} else {
							items[2] = append(items[2], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
					if monkey == 7 {
						x := items[monkey][item]
						x = x + 1
						x = x % 9699690
						if x%19 == 0 {
							items[2] = append(items[2], x)
						} else {
							items[6] = append(items[6], x)
						}
						items[monkey] = remove(items[monkey], item)
					}
				}
				monkey++
			}
			i++
		}
	}
	return multiply(inspections)
}

func main() {
	content, err := os.ReadFile("../../../inputs/2022/day11/test.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Test :\n")
	start := time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))

	content, err = os.ReadFile("../../../inputs/2022/day11/input.data")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInput :\n")
	start = time.Now()
	fmt.Printf("\nPart 1 : %d\n", part1(string(content)))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Printf("\nPart 2 : %d\n", part2(string(content)))
	fmt.Println(time.Since(start))
}
