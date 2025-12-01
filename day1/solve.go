package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	file, err := os.OpenFile("input", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	current := 50
	lineNumber := 1
	var temp int
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing rotation:", err)
			return
		}

		switch direction {
		case "L":
			temp = current - rotation
		case "R":
			temp = current + rotation
		default:
			fmt.Printf("Unknown direction '%s' at line %d\n", line, lineNumber)
			return
		}

		current = (100 + temp) % 100
		if current == 0 {
			count++
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", count)
}

func part2() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	current := 50 // position on [0..99]
	lineNumber := 1
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			lineNumber++
			continue
		}

		direction := line[0]
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing rotation at line", lineNumber, ":", err)
			return
		}

		full := rotation / 100
		rem := rotation % 100
		count += full

		switch direction {
		case 'R':
			temp := current + rem
			if current > 0 && temp >= 100 {
				count++
			}
			current = (current + rem) % 100

		case 'L':
			temp := current - rem
			if current > 0 && temp <= 0 {
				count++
			}
			current = (100 + temp) % 100

		default:
			fmt.Printf("Unknown direction '%c' at line %d\n", direction, lineNumber)
			return
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 2:", count)
}

func main() {
	part1()
	part2()
}
