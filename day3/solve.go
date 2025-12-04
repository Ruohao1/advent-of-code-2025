package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1(line string) int {
	count := make(map[int][]int)

	for i := 0; i < 10; i++ {
		count[i] = []int{}
	}

	for i, r := range line {
		c := int(r - '0')
		count[c] = append(count[c], i)
	}

	first := -1
	firstIndex := -1
	last := -1

	for last < 0 {
		for i := 9; i >= 0; i-- {
			if first == -1 {
				if len(count[i]) > 0 {
					if len(count[i]) == 1 && count[i][0] == len(line)-1 {
						continue
					}
					first = i
					firstIndex = count[i][0]
					count[i] = count[i][1:]
					break
				}
				continue
			}
			if len(count[i]) > 0 {
				for j := len(count[i]) - 1; j >= 0; j-- {
					if count[i][j] > firstIndex {
						last = i
						break
					}
				}
				if last >= 0 {
					break
				}
				count[i] = []int{}
			}
		}
	}

	return first*10 + last
}

func part2(line string) int {
	count := make(map[int][]int)

	for i := 0; i < 10; i++ {
		count[i] = []int{}
	}

	for i, r := range line {
		c := int(r - '0')
		count[c] = append(count[c], i)
	}

	maxDigits := 12
	digits := []int{}
	currentIndex := -1

	for len(digits) < maxDigits {
		found := false
		for i := 9; i >= 0; i-- {
			if len(count[i]) > 0 {
				for j := 0; j < len(count[i]); j++ {

					if len(line)-count[i][j] == maxDigits-len(digits) {
						for k := count[i][j]; k < len(line); k++ {
							digits = append(digits, int(line[k]-'0'))
							currentIndex = k
						}
						found = true
					}
					if currentIndex < count[i][j] && len(line)-count[i][j] > maxDigits-len(digits) {
						digits = append(digits, i)
						currentIndex = count[i][j]
						count[i] = count[i][j+1:]
						found = true
					}

					if found {
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	res := digits[0]
	for i := 1; i < len(digits); i++ {
		res *= 10
		res += digits[i]
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

	for scanner.Scan() {
		line := scanner.Text()
		switch part {
		case 1:
			result += part1(line)
		case 2:
			result += part2(line)
		default:
			fmt.Println("Unknown part:", part)
			return
		}
	}

	fmt.Printf("Part %d: %d\n", part, result)
}
