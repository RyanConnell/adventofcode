package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/RyanConnell/adventocode/2025/pkg/parser"
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

func solve(lines []string) (int, error) {
	idRanges := strings.Split(lines[0], ",")
	var invalid int
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 != 0 {
				continue
			}
			if s[:len(s)/2] == s[len(s)/2:] {
				invalid += i
			}
		}
	}
	return invalid, nil
}

/// Part 2 \\\

func solvePart2(lines []string) (int, error) {
	idRanges := strings.Split(lines[0], ",")
	var invalid int
	for _, idRange := range idRanges {
		parts := strings.Split(idRange, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			var match bool
			for chunkSize := 1; chunkSize <= len(s)/2; chunkSize++ {
				if isInvalid(s, chunkSize) {
					match = true
					break
				}
			}
			if match {
				invalid += i
			}
		}
	}
	return invalid, nil
}

func isInvalid(si string, chunkSize int) bool {
	if len(si)%chunkSize != 0 {
		return false
	}
	chunks := len(si) / chunkSize
	expected := si[:chunkSize]
	invalid := true
	for chunk := 1; chunk < chunks; chunk++ {
		end := chunkSize * (chunk + 1)
		if end > len(si) {
			end -= 1
		}
		if si[chunkSize*chunk:end] != expected {
			invalid = false
			break
		}
	}
	return invalid
}
