package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/gonum/mat"
)

type hailstone struct {
	x, y, z    int64
	dx, dy, dz int64
}

type vector struct {
	x, y, z int64
}

func Split(r rune) bool {
	return r == ' ' || r == ','
}

func format(s string) []hailstone {
	lines := strings.Split(s, "\n")
	res := []hailstone{}
	for _, line := range lines {
		lineSplit := strings.FieldsFunc(line, Split)
		x, _ := strconv.Atoi(lineSplit[0])
		y, _ := strconv.Atoi(lineSplit[1])
		z, _ := strconv.Atoi(lineSplit[2])
		dx, _ := strconv.Atoi(lineSplit[4])
		dy, _ := strconv.Atoi(lineSplit[5])
		dz, _ := strconv.Atoi(lineSplit[6])
		res = append(res, hailstone{int64(x), int64(y), int64(z), int64(dx), int64(dy), int64(dz)})
	}
	return res
}

func diffVector(v1, v2 vector) vector {
	return vector{v1.x - v2.x, v1.y - v2.y, v1.z - v2.z}
}

func intToFloat(l []int64) []float64 {
	res := []float64{}
	for _, x := range l {
		res = append(res, float64(x))
	}
	return res
}

func getMatrix(a, va, b, vb vector) (*mat.Dense, *mat.Dense, []int64, []int64) {
	D := diffVector(a, b)
	DV := diffVector(va, vb)

	tabA := []int64{0, -DV.z, DV.y, 0, -D.z, D.y, DV.z, 0, -DV.x, D.z, 0, -D.x, -DV.y, DV.x, 0, -D.y, D.x, 0}
	tabB := []int64{b.y*vb.z - b.z*vb.y - (a.y*va.z - a.z*va.y), b.z*vb.x - b.x*vb.z - (a.z*va.x - a.x*va.z), b.x*vb.y - b.y*vb.x - (a.x*va.y - a.y*va.x)}
	A := mat.NewDense(3, 6, intToFloat(tabA))
	B := mat.NewDense(3, 1, intToFloat(tabB))

	return A, B, tabA, tabB
}

func addRows(A, B *mat.Dense) *mat.Dense {
	tab := []float64{}
	rA, c := A.Dims()
	rB, _ := B.Dims()
	for i := 0; i < rA; i++ {
		tab = append(tab, A.RawRowView(i)...)
	}
	for i := 0; i < rB; i++ {
		tab = append(tab, B.RawRowView(i)...)
	}
	R := mat.NewDense(rA+rB, c, tab)
	return R
}

func part1(s string) int { // may not work with other inputs due to float precision but who knows
	c := 0
	list := format(s)
	min, max := float64(200000000000000), float64(400000000000000)
	if len(list) == 5 {
		min, max = 7, 27
	}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			A, B := list[i], list[j]
			denominator := float64(A.dx*B.dy - A.dy*B.dx)
			if denominator == 0 {
				continue
			}
			numerator := float64(B.dy*(B.x-A.x) - B.dx*(B.y-A.y))
			t := numerator / denominator
			numerator = float64(A.dy*(B.x-A.x) - A.dx*(B.y-A.y))
			u := numerator / denominator
			if t < 0 || u < 0 {
				continue
			}
			px := float64(A.x) + t*float64(A.dx)
			py := float64(A.y) + t*float64(A.dy)
			if px >= min && px <= max && py >= min && py <= max {
				c++
			}

		}
	}
	return c
}

func part2(s string) int {
	list := format(s)
	var marginOfError = int64(10)
	a := vector{list[0].x, list[0].y, list[0].z}
	b := vector{list[1].x, list[1].y, list[1].z}
	c := vector{list[2].x, list[2].y, list[2].z}
	va := vector{list[0].dx, list[0].dy, list[0].dz}
	vb := vector{list[1].dx, list[1].dy, list[1].dz}
	vc := vector{list[2].dx, list[2].dy, list[2].dz}
	A1, B1, tA1, tB1 := getMatrix(a, va, b, vb)
	A2, B2, tA2, tB2 := getMatrix(a, va, c, vc)
	var R mat.Dense
	A := addRows(A1, A2)
	B := addRows(B1, B2)
	R.Solve(A, B)

	//this is very ugly
	for i := -marginOfError; i <= marginOfError; i++ {
		for j := -marginOfError; j <= marginOfError; j++ {
			for k := -marginOfError; k <= marginOfError; k++ {
				rock := hailstone{}
				rock.x = int64(math.Round(R.At(0, 0))) + i
				rock.y = int64(math.Round(R.At(1, 0))) + j
				rock.z = int64(math.Round(R.At(2, 0))) + k
				rock.dx = int64(math.Round(R.At(3, 0)))
				rock.dy = int64(math.Round(R.At(4, 0)))
				rock.dz = int64(math.Round(R.At(5, 0)))
				var ok = true
				for l := 0; l < 3; l++ {
					ok = ok && (tA1[l*6+0]*rock.x+tA1[l*6+1]*rock.y+tA1[l*6+2]*rock.z+tA1[l*6+3]*rock.dx+tA1[l*6+4]*rock.dy+tA1[l*6+5]*rock.dz-tB1[l] == 0)
					ok = ok && (tA2[l*6+0]*rock.x+tA2[l*6+1]*rock.y+tA2[l*6+2]*rock.z+tA2[l*6+3]*rock.dx+tA2[l*6+4]*rock.dy+tA2[l*6+5]*rock.dz-tB2[l] == 0)
				}
				if ok {
					return int(rock.x + rock.z + rock.y)
				}
			}
		}
	}
	return 0
}

func main() {
	content, err := os.ReadFile("../../../inputs/2023/day24/test.data")

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

	content, err = os.ReadFile("../../../inputs/2023/day24/input.data")

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
