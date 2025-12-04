package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay04(t *testing.T) {
	tester.TimeAndCheck(4, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    13,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, false) },
			Expected:    1527,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    43,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, true) },
			Expected:    8690,
		},
	})
}
