package main

import (
	"fmt"
	"strconv"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines, 2)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 1 Result: %d\n", solutionPart1)

	solutionPart2, err := solve(lines, 12)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Part 2 Result: %d\n", solutionPart2)
}

/// Part 1 & 2 \\\

func solve(lines []string, activeCount int) (int, error) {
	var voltageSum int
	for _, line := range lines {
		var voltageStr string
		var idx int
		lastIdx := -1
		for i := range activeCount {
			// Search for the largest number that would still leave activeCount-i entries remaining
			idx = lastIdx + 1
			for j := lastIdx + 1; j <= len(line)-activeCount+i; j++ {
				if line[j] > line[idx] {
					idx = j
				}
			}
			voltageStr += string(line[idx])
			lastIdx = idx
		}
		maxVoltage, err := strconv.Atoi(voltageStr)
		if err != nil {
			return 0, err
		}
		voltageSum += maxVoltage
	}
	return voltageSum, nil
}
