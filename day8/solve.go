package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const N_PAIRS = 1000

type box struct {
	x, y, z int
}

func dist(a, b box) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2) + math.Pow(float64(a.z-b.z), 2))
}

type boxPair struct {
	d float64
	a box
	b box
}

func part1(lines []string) int {
	boxPairs := []boxPair{}

	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		a := box{x, y, z}
		for j, line2 := range lines {
			if i <= j {
				continue
			}
			split := strings.Split(line2, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			z, _ := strconv.Atoi(split[2])
			b := box{x, y, z}
			d := dist(a, b)
			boxPairs = append(boxPairs, boxPair{d: d, a: a, b: b})
		}
	}

	slices.SortFunc(boxPairs, func(a, b boxPair) int {
		if a.d < b.d {
			return -1
		} else if a.d == b.d {
			return 0
		} else {
			return 1
		}
	})

	circuits := map[box]int{}
	circuitsLen := []int{}
	n := 0

	for i := range N_PAIRS {
		boxPair := boxPairs[i]
		c1, ok1 := circuits[boxPair.a]
		c2, ok2 := circuits[boxPair.b]
		// fmt.Println(boxPair, c1, ok1, c2, ok2)
		if ok1 && ok2 {
			if c1 != c2 {
				circuitsLen[c1] += circuitsLen[c2] - 1
				for k, v := range circuits {
					if v == c2 {
						circuits[k] = c1
					}
				}
				circuitsLen[c2] = 1
			} else {
				continue
			}
		}
		if ok1 {
			circuits[boxPair.b] = c1
			circuitsLen[c1]++
		} else if ok2 {
			circuits[boxPair.a] = c2
			circuitsLen[c2]++
		} else {
			circuitsLen = append(circuitsLen, 2)
			circuits[boxPair.a] = n
			circuits[boxPair.b] = n
			n++
		}
		// fmt.Println(circuits)
		// fmt.Println(circuitsLen)
	}

	// fmt.Println(circuits)
	// fmt.Println(circuitsLen)

	slices.SortFunc(circuitsLen, func(a, b int) int {
		if a > b {
			return -1
		} else if a == b {
			return 0
		} else {
			return 1
		}
	})

	res := 1
	for k := range 3 {
		res *= circuitsLen[k]
	}
	return res
}

func part2(lines []string) int {
	boxPairs := []boxPair{}

	for i, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])
		a := box{x, y, z}
		for j, line2 := range lines {
			if i <= j {
				continue
			}
			split := strings.Split(line2, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])
			z, _ := strconv.Atoi(split[2])
			b := box{x, y, z}
			d := dist(a, b)
			boxPairs = append(boxPairs, boxPair{d: d, a: a, b: b})
		}
	}

	slices.SortFunc(boxPairs, func(a, b boxPair) int {
		if a.d < b.d {
			return -1
		} else if a.d == b.d {
			return 0
		} else {
			return 1
		}
	})
	circuits := map[box]int{}
	circuitsLen := []int{}
	n := 0

	for _, boxPair := range boxPairs {
		c1, ok1 := circuits[boxPair.a]
		c2, ok2 := circuits[boxPair.b]
		// fmt.Println(boxPair, c1, ok1, c2, ok2)
		if ok1 && ok2 {
			if c1 != c2 {
				circuitsLen[c1] += circuitsLen[c2] - 1
				// fmt.Println(circuitsLen[c1])
				if circuitsLen[c1] == len(lines) {
					return boxPair.a.x * boxPair.b.x
				}

				for k, v := range circuits {
					if v == c2 {
						circuits[k] = c1
					}
				}
				circuitsLen[c2] = 1
			} else {
				continue
			}
		}
		if ok1 {
			circuits[boxPair.b] = c1
			circuitsLen[c1]++
			if circuitsLen[c1] == len(lines) {
				return boxPair.a.x * boxPair.b.x
			}
		} else if ok2 {
			circuits[boxPair.a] = c2
			circuitsLen[c2]++
			if circuitsLen[c2] == len(lines) {
				return boxPair.a.x * boxPair.b.x
			}
		} else {
			circuitsLen = append(circuitsLen, 2)
			circuits[boxPair.a] = n
			circuits[boxPair.b] = n
			n++
		}
		// fmt.Println(circuits)
		// fmt.Println(circuitsLen)
	}
	return 0
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run solve.go <inputName> <part>")
		return
	}
	inputName := os.Args[1]
	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Error parsing part:", err)
		return
	}

	inputFile, err := os.OpenFile(inputName, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	lines := []string{}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	result := 0

	switch part {
	case 1:
		result = part1(lines)
	case 2:
		result = part2(lines)
	default:
		fmt.Println("Unknown part:", part)
		return
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
