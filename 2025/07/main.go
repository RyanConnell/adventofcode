package main

import (
	"fmt"

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

func solvePart1(lines []string) (int, error) {
	splits, _ := solve(lines)
	return splits, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	_, timelines := solve(lines)
	return timelines, nil
}

func solve(lines []string) (int, int) {
	state := make([]int, len(lines[0]))
	for i, char := range lines[0] {
		if char == 'S' {
			state[i] = 1
		}
	}

	var splits int
	for _, line := range lines {
		for x, s := range state {
			if s == 0 {
				continue
			}
			if line[x] == '.' {
				continue
			}
			if x+1 < len(state) {
				state[x+1] += state[x]
			}
			if x-1 >= 0 {
				state[x-1] += state[x]
			}
			state[x] = 0
			splits++
		}
	}

	var timelines int
	for _, i := range state {
		timelines += i
	}
	return splits, timelines
}
