package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2024/pkg/tester"
)

func TestDay10(t *testing.T) {
	tester.TimeAndCheck(10, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    36,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solve,
			Expected:    430,
		},
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    81,
		},
		{
			Description: "Part 1 (input)",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    928,
		},
	})
}
