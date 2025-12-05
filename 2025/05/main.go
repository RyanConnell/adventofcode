package main

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
	"golang.org/x/exp/slices"
)

func main() {
	lines := parser.MustReadFile("input/input.txt")

	solutionPart1, err := solve(lines)
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

func parseInput(lines []string) ([][]int, []int, error) {
	ranges := [][]int{}
	var ids []int
	var gatherIDs bool
	for _, line := range lines {
		if line == "" {
			gatherIDs = true
			continue
		}
		if gatherIDs {
			id, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, err
			}
			ids = append(ids, id)
		} else {
			rangeStr := strings.Split(line, "-")
			left, err := strconv.Atoi(rangeStr[0])
			if err != nil {
				return nil, nil, err
			}
			right, err := strconv.Atoi(rangeStr[1])
			if err != nil {
				return nil, nil, err
			}
			ranges = append(ranges, []int{left, right})
		}
	}
	return ranges, ids, nil
}

func solve(lines []string) (int, error) {
	ranges, ids, err := parseInput(lines)
	if err != nil {
		return 0, err
	}

	var count int
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				count++
				break
			}
		}
	}
	return count, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	ranges, _, err := parseInput(lines)
	if err != nil {
		return 0, err
	}
	slices.SortFunc(ranges, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	// Merge ranges
	for i := 1; i < len(ranges); i++ {
		lastEnding := ranges[i-1][1]
		if ranges[i][1] <= lastEnding { // Remove range
			ranges = append(ranges[:i], ranges[i+1:]...)
			i--
		} else if ranges[i][0] <= lastEnding { // Squash range
			ranges[i-1][1] = ranges[i][1]
			ranges = append(ranges[:i], ranges[i+1:]...)
			i--
		}
	}
	var count int
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}
	return count, nil
}
