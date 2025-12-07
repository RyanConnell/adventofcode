package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay07(t *testing.T) {
	tester.TimeAndCheck(7, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solvePart1,
			Expected:    21,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solvePart1,
			Expected:    1570,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    40,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    15118009521693,
		},
	})
}
