package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solvePart1(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solvePart2(lines)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 \\\

func parsePart1(lines []string) [][]string {
	data := make([][]string, 0)
	for _, line := range lines {
		var idx int
		for _, row := range strings.Split(line, " ") {
			if row == "" {
				continue
			}
			if len(data) <= idx {
				data = append(data, []string{})
			}
			data[idx] = append(data[idx], row)
			idx++
		}
	}
	return data
}

func solvePart1(lines []string) (int, error) {
	data := parsePart1(lines)
	return solve(data)
}

func solve(data [][]string) (int, error) {
	var total int
	for _, row := range data {
		op := row[len(row)-1]
		var value int
		switch op {
		case "*":
			value = 1
		case "+":
			value = 0
		default:
			return 0, fmt.Errorf("unknown operation: %v", op)
		}

		for j := 0; j < len(row)-1; j++ {
			x, err := strconv.Atoi(row[j])
			if err != nil {
				return 0, err
			}
			switch op {
			case "*":
				value *= x
			case "+":
				value += x
			}
		}
		total += value
	}
	return total, nil
}

// / Part 2 \\\

func parsePart2(lines []string) [][]string {
	data := make([][]string, 0)
	var idx int
	for x := 0; x < len(lines[0]); x++ {
		var numStr string
		for y := 0; y < len(lines)-1; y++ {
			if lines[y][x] == ' ' {
				continue
			}
			numStr += string(lines[y][x])
		}
		if len(data) <= idx {
			data = append(data, []string{})
		}
		if numStr == "" {
			idx++
		} else {
			data[idx] = append(data[idx], numStr)
		}
	}

	idx = 0
	for _, op := range strings.Split(lines[len(lines)-1], " ") {
		if op == "" {
			continue
		}
		data[idx] = append(data[idx], op)
		idx++
	}

	return data
}

func solvePart2(lines []string) (int, error) {
	data := parsePart2(lines)
	return solve(data)
}
