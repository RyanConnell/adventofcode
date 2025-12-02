package main

import (
	"testing"

	"github.com/RyanConnell/adventocode/2025/pkg/tester"
)

func TestDay02(t *testing.T) {
	tester.TimeAndCheck(2, []tester.TestCase[int]{
		{
			Description: "Part 1 (sample)",
			File:        "sample.txt",
			Solution:    solve,
			Expected:    1227775554,
		},
		{
			Description: "Part 1",
			File:        "input.txt",
			Solution:    solve,
			Expected:    17077011375,
		},
		{
			Description: "Part 2 (sample)",
			File:        "sample.txt",
			Solution:    solvePart2,
			Expected:    4174379265,
		},
		{
			Description: "Part 2",
			File:        "input.txt",
			Solution:    solvePart2,
			Expected:    36037497037,
		},
	})
}
