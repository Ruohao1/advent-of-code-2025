package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	parseRanges := true
	ranges := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseRanges = false
			continue
		}
		if parseRanges {
			split := strings.Split(line, "-")
			start, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}
			end, err := strconv.Atoi(split[1])
			if err != nil {
				panic(err)
			}
			ranges = append(ranges, []int{start, end})
		} else {
			id, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			for _, rangePair := range ranges {
				if id >= rangePair[0] && id <= rangePair[1] {
					res++
					break
				}
			}
		}

	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	ranges := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		split := strings.Split(line, "-")
		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, []int{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	mergedRanges := [][]int{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		if ranges[i][0] <= mergedRanges[len(mergedRanges)-1][1] {
			mergedRanges[len(mergedRanges)-1][1] = max(mergedRanges[len(mergedRanges)-1][1], ranges[i][1])
		} else {
			mergedRanges = append(mergedRanges, ranges[i])
		}
	}

	for _, rangePair := range mergedRanges {
		res += rangePair[1] - rangePair[0] + 1
	}

	return res
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

	scanner := bufio.NewScanner(inputFile)

	result := 0

	switch part {
	case 1:
		result = part1(scanner)
	case 2:
		result = part2(scanner)
	default:
		fmt.Println("Unknown part:", part)
		return
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
