package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay03(t *testing.T) {
	tester.TimeAndCheck(3, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 2) },
			Expected:    357,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 2) },
			Expected:    17324,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 12) },
			Expected:    3121910778619,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    func(lines []string) (int, error) { return solve(lines, 12) },
			Expected:    171846613143331,
		},
	})
}
