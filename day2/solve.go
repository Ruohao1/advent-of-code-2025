package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidPart1(id string) bool {
	n := len(id)
	if n%2 != 0 {
		return false
	}
	return id[:n/2] == id[n/2:]
}

func part1() {
	inputFile, err := os.OpenFile("input", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	_ = scanner.Scan()
	line := scanner.Text()
	sum := 0
	idRanges := strings.Split(line, ",")

	for _, idRange := range idRanges {
		split := strings.Split(strings.TrimSpace(idRange), "-")

		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			if isInvalidPart1(strconv.Itoa(i)) {
				sum += i
			}
		}

	}

	fmt.Println("Part 1:", sum)
}

func isInvalidPart2(id string) bool {
	n := len(id)

	for i := 2; i <= n; i++ {
		sliceLen := n / i
		if n%sliceLen != 0 {
			continue
		}

		isInvalid := true

		for j := 0; j < n-2*sliceLen+1; j += sliceLen {
			if id[j:j+sliceLen] != id[j+sliceLen:j+2*sliceLen] {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			return isInvalid
		}
	}

	return false
}

func part2() {
	inputFile, err := os.OpenFile("input", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	_ = scanner.Scan()
	line := scanner.Text()
	sum := 0
	idRanges := strings.Split(line, ",")

	for _, idRange := range idRanges {
		split := strings.Split(strings.TrimSpace(idRange), "-")

		start, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for i := start; i <= end; i++ {
			if isInvalidPart2(strconv.Itoa(i)) {
				sum += i
			}
		}

	}

	fmt.Println("Part 2:", sum)
}

func main() {
	part1()
	part2()
}
