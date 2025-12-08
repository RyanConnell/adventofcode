package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay08(t *testing.T) {
	tester.TimeAndCheck(8, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 10) },
			Expected:    40,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 1000) },
			Expected:    57970,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    25272,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    8520040659,
		},
	})
}
