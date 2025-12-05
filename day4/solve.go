package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(lines []string) int {
	res := 0
	dirs := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] != '@' {
				continue
			}

			count := 0
			for _, d := range dirs {
				ni := i + d[0]
				nj := j + d[1]

				if ni < 0 || ni >= len(lines) {
					continue
				}
				if nj < 0 || nj >= len(lines[ni]) {
					continue
				}

				if lines[ni][nj] == '@' {
					count++
				}
			}

			if count < 4 {
				res++
			}
		}
	}
	return res
}

func part2(lines []string) int {
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}

	res := 0
	dirs := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	changed := true
	remove := [][]int{}

	for changed {
		changed = false
		for _, index := range remove {
			grid[index[0]][index[1]] = '.'
		}
		remove = [][]int{}

		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] != '@' {
					continue
				}

				count := 0
				for _, d := range dirs {
					ni := i + d[0]
					nj := j + d[1]

					if ni < 0 || ni >= len(grid) {
						continue
					}
					if nj < 0 || nj >= len(grid[ni]) {
						continue
					}

					if grid[ni][nj] == '@' {
						count++
					}
				}

				if count < 4 {
					changed = true
					remove = append(remove, []int{i, j})
					res++
				}
			}
		}
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
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	switch part {
	case 1:
		result += part1(lines)
	case 2:
		result += part2(lines)
	default:
		fmt.Println("Unknown part:", part)
		return
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
