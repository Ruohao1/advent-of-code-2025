package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1(scanner *bufio.Scanner) int {
	res := 0
	lines := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Fields(line))
	}

	operators := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	for j, op := range operators {
		columnValue := -1
		for i := range len(lines) {
			value, err := strconv.Atoi(lines[i][j])
			if err != nil {
				panic("error parsing value")
			}
			if columnValue == -1 {
				columnValue = value
			} else {
				switch op[0] {
				case '+':
					columnValue += value
				case '*':
					columnValue *= value
				}
			}
		}
		res += columnValue
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	res := 0
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	operators := lines[len(lines)-1]
	lines = lines[:len(lines)-1]

	var current int
	var columnValue int
	column := 0
	var op rune

	for j, r := range operators {
		current = -1
		if r == '+' || r == '*' {
			op = r
			res += columnValue
			columnValue = -1
			column++
		}

		for i := range len(lines) {
			if lines[i][j] == ' ' {
				continue
			}

			if current == -1 {
				current = int(lines[i][j] - '0')
			} else {
				current = current*10 + int(lines[i][j]-'0')
			}
		}
		if current == -1 {
			continue
		}
		if columnValue == -1 {
			columnValue = current
		} else {
			switch op {
			case '+':
				columnValue += current
			case '*':
				columnValue *= current
			}
		}
	}

	res += columnValue
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
