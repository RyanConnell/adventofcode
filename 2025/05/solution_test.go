package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay05(t *testing.T) {
	tester.TimeAndCheck(5, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    3,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solve,
			Expected:    811,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    14,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    338189277144473,
		},
		{
			Description: "Part 2 (sample2)",
			File:        "sample2.txt",
			Solution:    solvePart2,
			Expected:    13,
		},
	})
}
